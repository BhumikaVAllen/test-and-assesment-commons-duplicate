package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository interface {
	InsertDocument(ctx context.Context, document interface{}, collectionName string) (interface{}, error)
	InsertDocuments(ctx context.Context, documentsCollectionMap []DocumentCollectionMap) ([]interface{}, error)
	Get(ctx context.Context, filter interface{}, collectionName string) (*mongo.SingleResult, error)
	Update(ctx context.Context, filter interface{}, document interface{}, collectionName string) (*mongo.SingleResult, error)
}
