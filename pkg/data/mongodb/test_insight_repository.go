package mongodb

import (
	"context"
	"errors"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/constants"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const testInsightCollection = "testInsight"

type ITestInsightRepository interface {
	CreateTestInsight(ctx context.Context, TestInsight entity.TestInsightEntity) error
	FindByTestID(ctx context.Context, testID string) (*entity.TestInsightEntity, error)
}

type TestInsightRepository struct {
	log *log.Helper
	db  *mongo.Collection
}

func NewTestInsightRepository(logger log.Logger, mongoClient *MongoClient) *TestInsightRepository {
	db := mongoClient.client.Database(mongoClient.dbName)
	collection := db.Collection(testInsightCollection)
	return &TestInsightRepository{
		log: log.NewHelper(logger),
		db:  collection,
	}
}

func (s *TestInsightRepository) CreateTestInsight(ctx context.Context, testInsight entity.TestInsightEntity) error {
	filter := bson.D{
		{
			testIdKey, testInsight.TestID,
		},
	}
	testInsight.CreatedAt = time.Now()
	testInsight.UpdatedAt = time.Now()
	opts := options.FindOneAndReplace().SetUpsert(true).SetReturnDocument(options.After)
	dbRes := s.db.FindOneAndReplace(ctx, filter, testInsight, opts)
	if dbRes.Err() != nil {
		s.log.WithContext(ctx).Errorf("CreateTestInsight failed, err:%+v", dbRes.Err())
		return dbRes.Err()
	}
	return nil
}

func (s *TestInsightRepository) FindByTestID(ctx context.Context, testID string) (*entity.TestInsightEntity, error) {
	filter := bson.D{
		{Key: testIdKey, Value: testID},
	}
	result := s.db.FindOne(ctx, filter)
	if result.Err() != nil {
		s.log.WithContext(ctx).Errorf("FindByTestID failed for TestInsightEntity, err:%+v", result.Err())
		return nil, result.Err()
	}
	var TestInsight entity.TestInsightEntity
	err := result.Decode(&TestInsight)
	if err != nil {
		s.log.WithContext(ctx).Errorf("FindByTestID decoding failed, err:%+v", err)
		return nil, errors.New(constants.DBErrorMessage)
	}
	return &TestInsight, nil
}
