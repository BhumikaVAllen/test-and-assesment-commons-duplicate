package mongodb

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func setupMongo() (context.Context, MongoRepository) {
	ctx := context.Background()
	d := &MongoClient{
		client: mongoClient,
	}
	repo := NewMongoRepositoryImpl(d, log.DefaultLogger)
	return ctx, repo
}

func TestNewMongoRepositoryImpl(t *testing.T) {
	type args struct {
		data   *MongoClient
		logger log.Logger
	}
	tests := []struct {
		name string
		args args
		want MongoRepository
	}{
		{
			name: "Test new mongo impl",
			args: args{
				data: &MongoClient{
					client: mongoClient,
				},
				logger: log.DefaultLogger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mongoStore := NewMongoRepositoryImpl(tt.args.data, tt.args.logger)
			assert.NotNil(t, mongoStore)
		})
	}
}

type data struct {
	name string
}

func TestMongoRepositoryImpl_InsertDocument(t *testing.T) {
	ctx, repo := setupMongo()
	type fields struct {
		logger *log.Helper
		client *mongo.Client
	}

	type args struct {
		ctx            context.Context
		data           data
		dbName         string
		collectionName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test InsertDocument impl",
			args: args{
				ctx: ctx,
				data: data{
					name: "random",
				},
				dbName:         "testdb",
				collectionName: "sample-collection",
			},
			fields: fields{
				logger: log.NewHelper(log.DefaultLogger),
				client: mongoClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply, err := repo.InsertDocument(tt.args.ctx, tt.args.data, tt.args.collectionName)
			assert.Nil(t, err)
			assert.NotNil(t, reply)
		})
	}
}

func TestMongoRepositoryImpl_Get(t *testing.T) {
	ctx, repo := setupMongo()

	type fields struct {
		logger *log.Helper
		client *mongo.Client
	}
	type args struct {
		ctx            context.Context
		filter         bson.M
		dbName         string
		data           data
		collectionName string
		setup          func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		setup  func()
	}{
		{
			name: "Test get impl",
			args: args{
				ctx:            ctx,
				filter:         bson.M{"name": "mohit"},
				dbName:         "testdb",
				collectionName: "sample-collection",
			},
			fields: fields{
				logger: log.NewHelper(log.DefaultLogger),
				client: mongoClient,
			},
			setup: func() {
				dd := data{
					name: "mohit",
				}
				id, err := repo.InsertDocument(ctx, dd, "sample-collection")
				if err != nil {
					log.Fatalf("setup failed")
				}
				log.Infof("id: %v", id)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := repo.Get(tt.args.ctx, tt.args.filter, tt.args.collectionName)
			assert.Nil(t, got)
			assert.NotNil(t, err)
		})
	}
}

func TestMongoRepositoryImpl_InsertDocuments(t *testing.T) {
	ctx, repo := setupMongo()
	type fields struct {
		logger *log.Helper
		client *mongo.Client
	}

	type args struct {
		ctx            context.Context
		data           data
		dbName         string
		collectionName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test InsertDocument impl",
			args: args{
				ctx: ctx,
				data: data{
					name: "random",
				},
				dbName:         "testdb",
				collectionName: "sample-collection",
			},
			fields: fields{
				logger: log.NewHelper(log.DefaultLogger),
				client: mongoClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var entities []DocumentCollectionMap
			entities = append(entities, DocumentCollectionMap{
				Document:   tt.args.data,
				Collection: tt.args.collectionName,
			})
			reply, err := repo.InsertDocuments(tt.args.ctx, entities)
			assert.NotNil(t, err)
			assert.Nil(t, reply)
		})
	}
}

func TestMongoRepositoryImpl_Update(t *testing.T) {
	ctx, repo := setupMongo()

	type fields struct {
		logger *log.Helper
		client *mongo.Client
	}
	type args struct {
		ctx            context.Context
		filter         bson.M
		dbName         string
		data           data
		collectionName string
		setup          func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		setup  func()
	}{
		{
			name: "Test get impl",
			args: args{
				ctx:            ctx,
				filter:         bson.M{"name": "mohit"},
				dbName:         "testdb",
				collectionName: "sample-collection",
			},
			fields: fields{
				logger: log.NewHelper(log.DefaultLogger),
				client: mongoClient,
			},
			setup: func() {
				dd := data{
					name: "mohit",
				}
				id, err := repo.InsertDocument(ctx, dd, "sample-collection")
				if err != nil {
					log.Fatalf("setup failed")
				}
				log.Infof("id: %v", id)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			dd := data{
				name: "mohitverma",
			}
			got, err := repo.Update(tt.args.ctx, tt.args.filter, dd, tt.args.collectionName)
			assert.Nil(t, got)
			assert.NotNil(t, err)
		})
	}
}
