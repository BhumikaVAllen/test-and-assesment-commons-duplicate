package mongodb

import (
	"context"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

var testInsightEntry = entity.TestInsightEntity{
	TestID:       "test_h2poJlz9WZAg",
	StudentCount: 2,
	AverageMarks: 23,
	TopperMarks:  23,
	SectionalInsights: []entity.TestSectionInsight{
		{
			SectionName:  "physics",
			AverageMarks: 20,
		},
		{
			SectionName:  "chemistry",
			AverageMarks: 12,
		},
	},
}

func TestTestInsightRepository_CreateTestInsight(t *testing.T) {
	d := &MongoClient{
		client: mongoClient,
		dbName: "test-and-assessments",
	}
	ctx := context.Background()
	testInsightRepo := NewTestInsightRepository(log.NewStdLogger(os.Stdout), d)
	testInsightEntry := entity.TestInsightEntity{
		TestID:       "test_h2poJlz9WZAg",
		StudentCount: 2,
		AverageMarks: 23,
		TopperMarks:  23,
		SectionalInsights: []entity.TestSectionInsight{
			{
				SectionName:  "physics",
				AverageMarks: 20,
			},
			{
				SectionName:  "chemistry",
				AverageMarks: 12,
			},
		},
	}
	type assertions struct {
		wantErr    bool
		assertFunc func(t *testing.T, ctx context.Context, testInsight entity.TestInsightEntity)
	}

	tests := []struct {
		name        string
		assertions  assertions
		prepareFunc func(*testing.T)
		arg         entity.TestInsightEntity
	}{
		{
			name: "Inserts document in DB",
			assertions: assertions{
				wantErr: false,
				assertFunc: func(t *testing.T, ctx context.Context, testInsightArg entity.TestInsightEntity) {
					filter := bson.D{
						{Key: testIdKey, Value: testInsightArg.TestID},
					}
					res := testInsightRepo.db.FindOne(ctx, filter)
					err := res.Err()
					assert.Nil(t, err)
					var testInsight entity.TestInsightEntity
					err = res.Decode(&testInsight)
					assert.Nil(t, err)
					assert.Equal(t, testInsightArg.TestID, testInsight.TestID)
					assert.Equal(t, testInsightArg.StudentCount, testInsight.StudentCount)
					assert.Equal(t, testInsightArg.AverageMarks, testInsight.AverageMarks)
					assert.Equal(t, testInsightArg.TopperMarks, testInsight.TopperMarks)
					assert.Equal(t, testInsightArg.SectionalInsights, testInsight.SectionalInsights)
				},
			},
			arg: entity.TestInsightEntity{
				TestID:       "test_3Nn8AI8JAbzx",
				StudentCount: 2,
				AverageMarks: 23,
				TopperMarks:  23,
				SectionalInsights: []entity.TestSectionInsight{
					{
						SectionName:  "physics",
						AverageMarks: 20,
					},
					{
						SectionName:  "chemistry",
						AverageMarks: 12,
					},
				},
			},
			prepareFunc: func(t *testing.T) {},
		},
		{
			name: "Upsert document in DB",
			assertions: assertions{
				wantErr: false,
				assertFunc: func(t *testing.T, ctx context.Context, testInsightArg entity.TestInsightEntity) {
					filter := bson.D{
						{Key: testIdKey, Value: testInsightArg.TestID},
					}
					res := testInsightRepo.db.FindOne(ctx, filter)
					var testInsight entity.TestInsightEntity
					err := res.Decode(&testInsight)
					assert.Nil(t, err)
					assert.Equal(t, testInsightArg.TestID, testInsight.TestID)
					assert.Equal(t, testInsightArg.StudentCount, testInsight.StudentCount)
					assert.Equal(t, testInsightArg.AverageMarks, testInsight.AverageMarks)
					assert.Equal(t, testInsightArg.TopperMarks, testInsight.TopperMarks)
					assert.Equal(t, testInsightArg.SectionalInsights, testInsight.SectionalInsights)
					assert.NotEqual(t, testInsightArg.StudentCount, testInsightEntry.StudentCount)
					assert.NotEqual(t, testInsightArg.AverageMarks, testInsightEntry.AverageMarks)
					assert.NotEqual(t, testInsightArg.TopperMarks, testInsightEntry.TopperMarks)
					assert.NotEqual(t, testInsightArg.SectionalInsights, testInsightEntry.SectionalInsights)
				},
			},
			arg: entity.TestInsightEntity{
				TestID:       "test_h2poJlz9WZAg",
				StudentCount: 4,
				AverageMarks: 13,
				TopperMarks:  51,
				SectionalInsights: []entity.TestSectionInsight{
					{
						SectionName:  "physics",
						AverageMarks: 12,
					},
					{
						SectionName:  "chemistry",
						AverageMarks: 24,
					},
				},
			},
			prepareFunc: func(t *testing.T) {
				err, _ := testInsightRepo.db.InsertOne(ctx, testInsightEntry)
				assert.NotNil(t, err)
			},
		},
	}
	t.Cleanup(func() {
		_, _ = testInsightRepo.db.DeleteMany(ctx, bson.D{})
	})
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.prepareFunc(t)
			err := testInsightRepo.CreateTestInsight(ctx, testCase.arg)
			if testCase.assertions.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			testCase.assertions.assertFunc(t, ctx, testCase.arg)
		})
	}
}

func TestTestInsightRepository_FindByTestID(t *testing.T) {
	d := &MongoClient{
		client: mongoClient,
		dbName: "test-and-assessments",
	}
	ctx := context.Background()
	testInsightRepo := NewTestInsightRepository(log.NewStdLogger(os.Stdout), d)
	type assertions struct {
		wantErr    bool
		assertFunc func(t *testing.T, ctx context.Context, testInsight *entity.TestInsightEntity)
	}
	tests := []struct {
		name        string
		assertions  assertions
		prepareFunc func(*testing.T)
		arg         string
	}{
		{
			name: "Return document from DB",
			assertions: assertions{
				wantErr: false,
				assertFunc: func(t *testing.T, ctx context.Context, testInsight *entity.TestInsightEntity) {
					assert.Equal(t, testInsightEntry.TestID, testInsight.TestID)
					assert.Equal(t, testInsightEntry.StudentCount, testInsight.StudentCount)
					assert.Equal(t, testInsightEntry.AverageMarks, testInsight.AverageMarks)
					assert.Equal(t, testInsightEntry.TopperMarks, testInsight.TopperMarks)
					assert.Equal(t, testInsightEntry.SectionalInsights, testInsight.SectionalInsights)
				},
			},
			prepareFunc: func(t *testing.T) {
				err, _ := testInsightRepo.db.InsertOne(ctx, testInsightEntry)
				assert.NotNil(t, err)
			},
			arg: testInsightEntry.TestID,
		},
		{
			name: "Return error when no document is found",
			assertions: assertions{
				wantErr: true,
				assertFunc: func(t *testing.T, ctx context.Context, testInsight *entity.TestInsightEntity) {
					assert.Nil(t, testInsight)
				},
			},
			prepareFunc: func(t *testing.T) {},
			arg:         "random_test_id",
		},
		{
			name: "Return error when decoding document fails",
			assertions: assertions{
				wantErr: true,
				assertFunc: func(t *testing.T, ctx context.Context, testInsight *entity.TestInsightEntity) {
					assert.Nil(t, testInsight)
				},
			},
			prepareFunc: func(t *testing.T) {
				err, _ := testInsightRepo.db.InsertOne(ctx, struct {
					TestID       string `bson:"testId"`
					StudentCount string `bson:"studentCount"`
				}{
					TestID:       "corrupted_document_id",
					StudentCount: "stringCount",
				})
				assert.NotNil(t, err)
			},
			arg: "corrupted_document_id",
		},
	}
	t.Cleanup(func() {
		_, _ = testInsightRepo.db.DeleteMany(ctx, bson.D{})
	})
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.prepareFunc(t)
			res, err := testInsightRepo.FindByTestID(ctx, testCase.arg)
			if testCase.assertions.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			testCase.assertions.assertFunc(t, ctx, res)
		})
	}
}
