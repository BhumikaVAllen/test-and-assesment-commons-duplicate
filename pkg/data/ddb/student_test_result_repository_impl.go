package ddb

import (
	"context"
	"errors"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-kratos/kratos/v2/log"
)

const (
// studentTestResultTableName         = "StudentTestResults"
// studentTestResultTablePartitionKey = "StudentId"
// studentTestResultTableSortKey      = "TestId"
)

type studentTestResultRepositoryImpl struct {
	ddb *dynamodb.DynamoDB
	log *log.Helper
}

func NewStudentTestResultRepositoryImpl(data *Client, logger log.Logger) StudentTestResultRepository {
	return &studentTestResultRepositoryImpl{
		ddb: data.DynamoDB,
		log: log.NewHelper(logger),
	}
}

func (*studentTestResultRepositoryImpl) Create(_ context.Context, _ *entity.StudentTestResultEntity) error {
	//not implemented
	err := errors.New("update operation not implemented")
	return err
}

func (*studentTestResultRepositoryImpl) Update(_ context.Context, _ *entity.StudentTestResultEntity) (*entity.StudentTestResultEntity, error) {
	//not implemented
	err := errors.New("update operation not implemented")
	return nil, err
}

func (*studentTestResultRepositoryImpl) FindByID(_ context.Context, _ string) (*entity.StudentTestResultEntity, error) {
	//not implemented
	//s.log.Infof()
	err := errors.New("findByID operation not implemented")
	return nil, err
}

func (*studentTestResultRepositoryImpl) List(_ context.Context, _ uint, _ int) ([]*entity.StudentTestResultEntity, error) {
	//not implemented
	err := errors.New("list operation not implemented")
	return nil, err
}
