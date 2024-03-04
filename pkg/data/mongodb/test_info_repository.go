package mongodb

import (
	"context"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/constants"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/mongodb/data_filters"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

const (
	testInfoCollection = "testInfo"
	idKey              = "testId"
	statusKey          = "status"
	sortField          = "createdAt"
	OrderBy            = "desc"
	pageNo             = 1
	pageSize           = 20
	status             = "status"
	stream             = "stream"
	class              = "class"
	category           = "category"
	createdBy          = "createdBy"
	testType           = "type"
)

type ITestInfoRepository interface {
	CreateTest(ctx context.Context, testInfo *entity.TestInfoEntity) error
	GetTest(ctx context.Context, testID string) (*entity.TestInfoEntity, error)
	FilterTests(ctx context.Context, dataFilter data_filters.TestFilter) (*data_filters.TestFilterResponse, error)
	UpdateTestStatus(ctx context.Context, testID string, testStatus string) error
	UpdateTest(ctx context.Context, testInfoEntity *entity.TestInfoEntity) error
}

type TestInfoRepository struct {
	log *log.Helper
	db  *mongo.Collection
}

func NewTestInfoRepository(mongoClient *MongoClient, logger log.Logger) *TestInfoRepository {
	db := mongoClient.client.Database(mongoClient.dbName)
	collection := db.Collection(testInfoCollection)
	return &TestInfoRepository{
		log: log.NewHelper(logger),
		db:  collection,
	}
}

func (t *TestInfoRepository) CreateTest(ctx context.Context, testInfo *entity.TestInfoEntity) error {
	_, err := t.db.InsertOne(ctx, testInfo)
	if err != nil {
		t.log.WithContext(ctx).Errorf("CreateTest failed, err:%+v", err)
		return err
	}
	return nil
}

func (t *TestInfoRepository) GetTest(ctx context.Context, testID string) (*entity.TestInfoEntity, error) {
	filter := bson.D{
		{Key: idKey, Value: testID},
	}
	result := t.db.FindOne(ctx, filter)
	if result.Err() != nil {
		t.log.WithContext(ctx).Errorf("Get Test failed, err:%+v", result.Err())
		return nil, result.Err()
	}
	var testInfo entity.TestInfoEntity
	err := result.Decode(&testInfo)
	if err != nil {
		t.log.WithContext(ctx).Errorf("GetTest decoding failed, err:%+v", err)
		return nil, err
	}
	return &testInfo, nil
}

// UpdateTest TODO - to update the value if the previous get value matches the old testInfo Entity
func (t *TestInfoRepository) UpdateTest(ctx context.Context, testInfoEntity *entity.TestInfoEntity) error {
	filter := bson.D{
		{Key: idKey, Value: testInfoEntity.TestID},
	}
	update := bson.M{"$set": testInfoEntity}
	result, err := t.db.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.log.WithContext(ctx).Errorf("Update  Test failed, err:%+v", err)
		return err
	}
	if result.MatchedCount == 0 {
		t.log.WithContext(ctx).Error(constants.DBErrorMessage)
		return mongo.ErrNoDocuments
	}
	return nil
}
func (t *TestInfoRepository) UpdateTestStatus(ctx context.Context, testID string, testStatus string) error {
	filter := bson.D{
		{Key: idKey, Value: testID},
	}

	update := bson.M{"$set": bson.M{statusKey: testStatus}}
	result, err := t.db.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		t.log.WithContext(ctx).Errorf("Update Status Test failed, err:%+v", err)
		return err
	}

	if result.MatchedCount == 0 {
		t.log.WithContext(ctx).Error(constants.DBErrorMessage)
		return mongo.ErrNoDocuments
	}
	return nil
}

func (t *TestInfoRepository) FilterTests(ctx context.Context, dataFilter data_filters.TestFilter) (*data_filters.TestFilterResponse, error) {
	filter := bson.M{}
	if dataFilter.PageNo <= 0 {
		dataFilter.PageNo = pageNo
	}
	if dataFilter.PageSize <= 0 || dataFilter.PageSize > 20 {
		dataFilter.PageSize = pageSize
	}
	if dataFilter.SortField == "" {
		dataFilter.SortField = sortField
	}
	if dataFilter.OrderBy != "asc" && dataFilter.OrderBy != "desc" {
		dataFilter.OrderBy = OrderBy
	}
	sortDirection := 1 // Ascending
	if dataFilter.OrderBy == "desc" {
		sortDirection = -1 // Descending
	}
	if dataFilter.Category != "" {
		filter[category] = bson.M{"$in": strings.Split(dataFilter.Category, ",")}
	}
	if dataFilter.SearchKeyword == "" {
		if dataFilter.ID != "" {
			filter[idKey] = dataFilter.ID
		}
		if dataFilter.Status != "" {
			filter[status] = bson.M{"$in": strings.Split(dataFilter.Status, ",")}
		}
		if dataFilter.Stream != "" {
			filter[stream] = bson.M{"$in": strings.Split(dataFilter.Stream, ",")}
		}
		if dataFilter.Class != "" {
			filter[class] = bson.M{"$in": strings.Split(dataFilter.Class, ",")}
		}
		if dataFilter.UserId != "" {
			filter[createdBy] = dataFilter.UserId
		}
		if (!dataFilter.FromDate.IsZero()) && (!dataFilter.ToDate.IsZero()) {
			dateRangeFilter := bson.M{}
			dateRangeFilter["$gte"] = dataFilter.FromDate
			dateRangeFilter["$lte"] = dataFilter.ToDate
			filter["schedule.startTime"] = dateRangeFilter
		}
		if dataFilter.Centers != "" {
			filter["assignment.centers"] = bson.M{"$in": strings.Split(dataFilter.Centers, ",")}
		}
		if dataFilter.TestType != "" {
			filter[testType] = bson.M{"$in": strings.Split(dataFilter.TestType, ",")}
		}

	} else {
		filter["$text"] = bson.M{"$search": dataFilter.SearchKeyword}
	}
	totalCount, err := t.db.CountDocuments(ctx, filter)
	maxPage := (totalCount + dataFilter.PageSize - 1) / dataFilter.PageSize
	if maxPage <= 0 {
		maxPage = 1
	}
	if dataFilter.PageNo > maxPage {
		dataFilter.PageNo = maxPage
	}
	skip := (dataFilter.PageNo - 1) * dataFilter.PageSize
	if skip < 0 {
		skip = 0
	}
	findOptions := options.Find().SetSort(bson.D{{dataFilter.SortField, sortDirection}}).SetSkip(int64(skip)).SetLimit(int64(dataFilter.PageSize))
	cursor, err := t.db.Find(ctx, filter, findOptions)
	if err != nil {
		t.log.WithContext(ctx).Errorf("FilterTests failed, err:%+v", err)
		return nil, err
	}
	testInfoResults := make([]*entity.TestInfoEntity, 0)
	for cursor.Next(ctx) {
		var testInfo entity.TestInfoEntity
		err := cursor.Decode(&testInfo)
		if err != nil {
			t.log.WithContext(ctx).Errorf("FilterTests decoding failed, err:%+v", err)
			return nil, err
		}
		testInfoResults = append(testInfoResults, &testInfo)
	}
	return &data_filters.TestFilterResponse{TotalResults: totalCount, PageNo: dataFilter.PageNo, PageSize: dataFilter.PageSize, TestInfo: testInfoResults}, nil
}
