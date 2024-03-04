package mongodb

import (
	"context"
	"errors"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/constants"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/mongodb/data_filters"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	studentTestOverviewCollection = "studentTestOverview"
	studentIdKey                  = "studentId"
	testIdKey                     = "testId"
	testCategoryKey               = "testCategory"
	testStatusKey                 = "testStatus"
)

type IStudentTestOverviewRepository interface {
	CreateStudentTest(ctx context.Context, studentTestInput entity.StudentTestOverviewEntity) error
	FindByStudentIDTestID(ctx context.Context, studentID string, testID string) (*entity.StudentTestOverviewEntity, error)
	FetchTestsByFilters(ctx context.Context, studentTestFilter data_filters.StudentTestFilter) ([]*entity.StudentTestOverviewEntity, error)
	UpdateStudentTestStatus(ctx context.Context, testID string, studentID string, status string) error
	UpdateStudentTest(ctx context.Context, studentTestOverviewEntity *entity.StudentTestOverviewEntity) (*entity.StudentTestOverviewEntity, error)
	FetchAllTestsByFilters(ctx context.Context, studentTestFilter data_filters.StudentTestFilter) ([]*entity.StudentTestOverviewEntity, error)
	FindByTestStudentIDTestID(ctx context.Context, studentID string, testID string) (*entity.StudentTestOverviewEntity, error)
}

type StudentTestOverviewRepository struct {
	log *log.Helper
	db  *mongo.Collection
}

func NewStudentTestOverviewRepository(mongoClient *MongoClient, logger log.Logger) *StudentTestOverviewRepository {
	db := mongoClient.client.Database(mongoClient.dbName)
	collection := db.Collection(studentTestOverviewCollection)
	return &StudentTestOverviewRepository{
		log: log.NewHelper(logger),
		db:  collection,
	}
}

func (s *StudentTestOverviewRepository) CreateStudentTest(ctx context.Context, studentTestInput entity.StudentTestOverviewEntity) error {
	_, err := s.db.InsertOne(ctx, studentTestInput)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Error occurred while insert StudentTestOverview err: %+v", err)
		return err
	}
	return nil
}

func (s *StudentTestOverviewRepository) UpdateStudentTestStatus(ctx context.Context, testID string, studentID string, status string) error {
	filter := bson.D{
		{Key: studentIdKey, Value: studentID},
		{Key: testIdKey, Value: testID},
	}

	update := bson.M{"$set": bson.M{testStatusKey: status}}
	result, err := s.db.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		s.log.WithContext(ctx).Errorf("Update Student Status Test failed err:%+v", err)
		return err
	}

	if result.MatchedCount == 0 {
		s.log.WithContext(ctx).Error(constants.NoDocumentFoundMessage)
		return mongo.ErrNoDocuments
	}
	return nil
}

func (s *StudentTestOverviewRepository) FindByStudentIDTestID(ctx context.Context, studentID string, testID string) (*entity.StudentTestOverviewEntity, error) {
	filter := bson.D{
		{Key: studentIdKey, Value: studentID},
		{Key: testIdKey, Value: testID},
	}
	result := s.db.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return &entity.StudentTestOverviewEntity{}, nil
		}
		return nil, errors.New("DB error")
	}

	var studentTestOverview entity.StudentTestOverviewEntity
	err := result.Decode(&studentTestOverview)
	if err != nil {
		return nil, errors.New("DB error")
	}
	return &studentTestOverview, nil
}

func (s *StudentTestOverviewRepository) FindByTestStudentIDTestID(ctx context.Context, studentID string, testID string) (*entity.StudentTestOverviewEntity, error) {
	filter := bson.D{
		{Key: studentIdKey, Value: studentID},
		{Key: testIdKey, Value: testID},
	}
	result := s.db.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, result.Err()
		}
		return nil, errors.New("DB error")
	}

	var studentTestOverview entity.StudentTestOverviewEntity
	err := result.Decode(&studentTestOverview)
	if err != nil {
		return nil, errors.New("DB error")
	}
	return &studentTestOverview, nil
}

func (s *StudentTestOverviewRepository) FetchTestsByFilters(ctx context.Context, studentTestFilter data_filters.StudentTestFilter) ([]*entity.StudentTestOverviewEntity, error) {
	if studentTestFilter.StudentID == "" {
		return nil, errors.New("student id field is mandatory")
	}

	filter := bson.M{}
	filter[studentIdKey] = studentTestFilter.StudentID

	if studentTestFilter.TestID != "" {
		filter[testIdKey] = studentTestFilter.TestID
	}

	if studentTestFilter.TestCategory != "" {
		filter[testCategoryKey] = studentTestFilter.TestCategory
	}

	if studentTestFilter.Status != "" {
		filter[testStatusKey] = studentTestFilter.Status
	}
	cursor, err := s.db.Find(ctx, filter)
	if err != nil {
		s.log.WithContext(ctx).Errorf("FetchTestsByFilters failed, err:%+v", err)
		return nil, err
	}

	studentTestOverviewResults := make([]*entity.StudentTestOverviewEntity, 0)
	for cursor.Next(ctx) {
		var studentTestOverview entity.StudentTestOverviewEntity
		err := cursor.Decode(&studentTestOverview)
		if err != nil {
			s.log.WithContext(ctx).Errorf("FetchTestsByFilters decoding failed, err:%+v", err)
			return nil, err
		}
		studentTestOverviewResults = append(studentTestOverviewResults, &studentTestOverview)
	}
	return studentTestOverviewResults, nil
}

func (s *StudentTestOverviewRepository) FetchAllTestsByFilters(ctx context.Context, studentTestFilter data_filters.StudentTestFilter) ([]*entity.StudentTestOverviewEntity, error) {
	filter := bson.M{}
	if studentTestFilter.TestID != "" {
		filter[testIdKey] = studentTestFilter.TestID
	}

	if studentTestFilter.Status != "" {
		filter[testStatusKey] = studentTestFilter.Status
	}
	cursor, err := s.db.Find(ctx, filter)
	if err != nil {
		s.log.WithContext(ctx).Errorf("FetchTestsByFilters failed, err:%+v", err)
		return nil, err
	}

	studentTestOverviewResults := make([]*entity.StudentTestOverviewEntity, 0)
	for cursor.Next(ctx) {
		var studentTestOverview entity.StudentTestOverviewEntity
		err := cursor.Decode(&studentTestOverview)
		if err != nil {
			s.log.WithContext(ctx).Errorf("FetchTestsByFilters decoding failed, err:%+v", err)
			return nil, err
		}
		studentTestOverviewResults = append(studentTestOverviewResults, &studentTestOverview)
	}
	return studentTestOverviewResults, nil
}

func (s *StudentTestOverviewRepository) FetchStudentCountByFilters(ctx context.Context, studentTestFilter data_filters.StudentTestFilter) (int64, error) {
	filter := bson.M{}
	if studentTestFilter.TestID != "" {
		filter[testIdKey] = studentTestFilter.TestID
	}

	if studentTestFilter.Status != "" {
		filter[testStatusKey] = studentTestFilter.Status
	}
	count, err := s.db.CountDocuments(ctx, filter)
	if err != nil {
		s.log.WithContext(ctx).Errorf("fetchStudentCountByFilters count failed with err: %+v", err)
		return 0, err
	}
	return count, nil
}

func (s *StudentTestOverviewRepository) UpdateStudentTest(ctx context.Context, studentTestOverviewEntity *entity.StudentTestOverviewEntity) (*entity.StudentTestOverviewEntity, error) {
	filter := bson.D{
		{Key: studentIdKey, Value: studentTestOverviewEntity.StudentID},
		{Key: testIdKey, Value: studentTestOverviewEntity.TestID},
	}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	result := s.db.FindOneAndReplace(ctx, filter, studentTestOverviewEntity, opts)
	if result.Err() != nil {
		s.log.WithContext(ctx).Errorf("UpdateStudentTestOverview failed, err:%+v", result.Err())
		return nil, result.Err()
	}
	var updatedTestOverviewEntity entity.StudentTestOverviewEntity
	err := result.Decode(&updatedTestOverviewEntity)
	if err != nil {
		s.log.WithContext(ctx).Errorf("UpdateStudentTestOverview decoding failed, err:%+v", err)
		return nil, err
	}
	return &updatedTestOverviewEntity, nil
}
