package mongodb

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database = "test-and-assessment"
)

type mongoRepositoryImpl struct {
	log    *log.Helper
	client *mongo.Client
}

type DocumentCollectionMap struct {
	Document   interface{}
	Collection string
}

func NewMongoRepositoryImpl(mongoClient *MongoClient, logger log.Logger) MongoRepository {
	return &mongoRepositoryImpl{
		client: mongoClient.client,
		log:    log.NewHelper(logger),
	}
}

func (mongoRepositoryImpl *mongoRepositoryImpl) InsertDocument(ctx context.Context, document interface{}, collectionName string) (interface{}, error) {
	mongoRepositoryImpl.log.WithContext(ctx).Infof("insert document: +v", document)

	db := mongoRepositoryImpl.client.Database(database)
	collection := db.Collection(collectionName)

	data, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	mongoRepositoryImpl.log.WithContext(ctx).Infof("Insert document in mongoDB = %+v", data)
	return data.InsertedID, nil
}

// InsertDocuments Documents in documentsCollectionMap should be in order of insertion
func (mongoRepositoryImpl *mongoRepositoryImpl) InsertDocuments(ctx context.Context, documentsCollectionList []DocumentCollectionMap) ([]interface{}, error) {
	mongoRepositoryImpl.log.WithContext(ctx).Infof(" insert documents in transactional")

	var collections []*mongo.Collection

	for _, value := range documentsCollectionList {
		collections = append(collections, mongoRepositoryImpl.client.Database(database).Collection(value.Collection))
	}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		for i, value := range documentsCollectionList {
			if _, err := collections[i].InsertOne(sessCtx, value.Document); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}

	session, err := mongoRepositoryImpl.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return nil, err
	}

	mongoRepositoryImpl.log.WithContext(ctx).Infof("Inserted Documents in mongo ddb = %+v", result)
	return nil, nil
}

func (mongoRepositoryImpl *mongoRepositoryImpl) Get(ctx context.Context, filter interface{}, collectionName string) (*mongo.SingleResult, error) {
	mongoRepositoryImpl.log.WithContext(ctx).Infof("Get TestInfo from mongoDb: %+v", filter)

	db := mongoRepositoryImpl.client.Database(database)
	collection := db.Collection(collectionName)

	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			mongoRepositoryImpl.log.Info("No document found")
		}
		return nil, result.Err()
	}
	return result, nil
}

func (mongoRepositoryImpl *mongoRepositoryImpl) Update(ctx context.Context, filter interface{}, document interface{}, collectionName string) (*mongo.SingleResult, error) {
	mongoRepositoryImpl.log.WithContext(ctx).Infof("Update TestInfo from mongoDb: %+v", filter)

	db := mongoRepositoryImpl.client.Database(database)
	collection := db.Collection(collectionName)

	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	result := collection.FindOneAndReplace(ctx, filter, document, opts)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			mongoRepositoryImpl.log.Info("No document found")
		}
		return nil, result.Err()
	}
	return result, nil
}
