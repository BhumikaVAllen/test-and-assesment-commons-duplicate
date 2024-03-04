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

const studentTestInsightCollection = "studentTestInsight"

type IStudentTestInsightRepository interface {
	CreateStudentTestInsight(ctx context.Context, studentTestInsight entity.StudentTestInsightEntity) error
	FindByStudentIDTestID(ctx context.Context, studentID, testID string) (*entity.StudentTestInsightEntity, error)
}

type StudentTestInsightRepository struct {
	log *log.Helper
	db  *mongo.Collection
}

func NewStudentTestInsightRepository(logger log.Logger, mongoClient *MongoClient) *StudentTestInsightRepository {
	db := mongoClient.client.Database(mongoClient.dbName)
	collection := db.Collection(studentTestInsightCollection)
	return &StudentTestInsightRepository{
		log: log.NewHelper(logger),
		db:  collection,
	}
}

func (s *StudentTestInsightRepository) CreateStudentTestInsight(ctx context.Context, studentTestInsight entity.StudentTestInsightEntity) error {
	filter := bson.D{
		{Key: testIdKey, Value: studentTestInsight.TestID},
		{Key: studentIdKey, Value: studentTestInsight.StudentID},
	}
	studentTestInsight.CreatedAt = time.Now()
	studentTestInsight.UpdatedAt = time.Now()
	opts := options.FindOneAndReplace().SetUpsert(true).SetReturnDocument(options.After)
	dbRes := s.db.FindOneAndReplace(ctx, filter, studentTestInsight, opts)
	if dbRes.Err() != nil {
		s.log.WithContext(ctx).Errorf("CreateStudentTestInsight failed, err:%+v", dbRes.Err())
		return dbRes.Err()
	}
	return nil
}

func (s *StudentTestInsightRepository) FindByStudentIDTestID(ctx context.Context, studentID string, testID string) (*entity.StudentTestInsightEntity, error) {
	filter := bson.D{
		{Key: studentIdKey, Value: studentID},
		{Key: testIdKey, Value: testID},
	}
	result := s.db.FindOne(ctx, filter)
	if result.Err() != nil {
		s.log.WithContext(ctx).Errorf("FindByStudentIDTestID failed for StudentTestInsightEntity , err:%+v", result.Err())
		return nil, result.Err()
	}
	var studentTestInsight entity.StudentTestInsightEntity
	err := result.Decode(&studentTestInsight)
	if err != nil {
		s.log.WithContext(ctx).Errorf("StudentTestInsightEntity: FindByStudentIDTestID decoding failed, err:%+v", err)
		return nil, errors.New(constants.DBErrorMessage)
	}
	return &studentTestInsight, nil
}

func (s *StudentTestInsightRepository) FetchStudentCountByTestID(ctx context.Context, testID string) (int64, error) {
	filter := bson.M{
		testIdKey: testID,
	}
	count, err := s.db.CountDocuments(ctx, filter)
	if err != nil {
		s.log.WithContext(ctx).Errorf("fetchStudentCountByTestID count failed with err: %+v", err)
		return 0, err
	}
	return count, nil
}
