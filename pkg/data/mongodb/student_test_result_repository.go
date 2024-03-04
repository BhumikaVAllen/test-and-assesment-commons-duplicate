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
)

const (
	studentTestResultCollection = "studentTestResult"
)

type IStudentTestResultRepository interface {
	CreateStudentTestResult(ctx context.Context, studentTestResult entity.StudentTestResultEntity) error
	UpdateStudentTestResult(ctx context.Context, studentTestResultInput *entity.StudentTestResultEntity) (*entity.StudentTestResultEntity, error)
	FindByStudentIDTestID(ctx context.Context, studentID string, testID string) (*entity.StudentTestResultEntity, error)
}

type StudentTestResultRepository struct {
	log *log.Helper
	db  *mongo.Collection
}

var _ IStudentTestResultRepository = (*StudentTestResultRepository)(nil)

func NewStudentTestResultRepository(mongoClient *MongoClient, logger log.Logger) *StudentTestResultRepository {
	db := mongoClient.client.Database(mongoClient.dbName)
	collection := db.Collection(studentTestResultCollection)
	return &StudentTestResultRepository{
		log: log.NewHelper(logger),
		db:  collection,
	}
}

func (s *StudentTestResultRepository) CreateStudentTestResult(ctx context.Context, studentTestResultInput entity.StudentTestResultEntity) error {
	filter := bson.D{
		{Key: studentIdKey, Value: studentTestResultInput.StudentID},
		{Key: testIdKey, Value: studentTestResultInput.TestID},
	}
	opts := options.FindOneAndReplace().SetUpsert(true).SetReturnDocument(options.After)
	result := s.db.FindOneAndReplace(ctx, filter, studentTestResultInput, opts)
	if result.Err() != nil {
		s.log.WithContext(ctx).Errorf("CreateStudentTestResult failed, err:%+v", result.Err())
		return result.Err()
	}
	return nil
}

func (s *StudentTestResultRepository) UpdateStudentTestResult(ctx context.Context, studentTestResultInput *entity.StudentTestResultEntity) (*entity.StudentTestResultEntity, error) {
	filter := bson.D{
		{Key: studentIdKey, Value: studentTestResultInput.StudentID},
		{Key: testIdKey, Value: studentTestResultInput.TestID},
	}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	result := s.db.FindOneAndReplace(ctx, filter, studentTestResultInput, opts)
	if result.Err() != nil {
		s.log.WithContext(ctx).Errorf("UpdateStudentTestResult failed, err:%+v", result.Err())
		return nil, result.Err()
	}
	var studentTestResult entity.StudentTestResultEntity
	err := result.Decode(&studentTestResult)
	if err != nil {
		s.log.WithContext(ctx).Errorf("UpdateStudentTestResult decoding failed, err:%+v", err)
		return nil, errors.New(constants.DBErrorMessage)
	}
	return &studentTestResult, nil
}

func (s *StudentTestResultRepository) FindByStudentIDTestID(ctx context.Context, studentID string, testID string) (*entity.StudentTestResultEntity, error) {
	filter := bson.D{
		{Key: studentIdKey, Value: studentID},
		{Key: testIdKey, Value: testID},
	}
	result := s.db.FindOne(ctx, filter)
	if result.Err() != nil {
		s.log.WithContext(ctx).Errorf("FindByStudentIDTestID failed for StudentTestResult , err:%+v", result.Err())
		return nil, result.Err()
	}
	var studentTestResult entity.StudentTestResultEntity
	err := result.Decode(&studentTestResult)
	if err != nil {
		s.log.WithContext(ctx).Errorf("FindByStudentIDTestID decoding failed, err:%+v", err)
		return nil, errors.New(constants.DBErrorMessage)
	}
	return &studentTestResult, nil
}
