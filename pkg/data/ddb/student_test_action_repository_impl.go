package ddb

import (
	"context"
	"errors"
	"fmt"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

const (
	studentTestActionTableName         = "StudentTestActions"
	studentTestActionTablePartitionKey = "TestIdStudentId"
	studentTestActionTableSortKey      = "SectionNamespaceQuestionId"
)

type studentTestActionRepositoryImpl struct {
	ddb *dynamodb.DynamoDB
	log *log.Helper
}

func NewStudentTestActionRepositoryImpl(data *Client, logger log.Logger) StudentTestActionRepository {
	return &studentTestActionRepositoryImpl{
		ddb: data.DynamoDB,
		log: log.NewHelper(logger),
	}
}

func (studentTestActionRepositoryImpl *studentTestActionRepositoryImpl) Create(_ context.Context, en *entity.StudentTestActionEntity) error {
	studentTestActionRepositoryImpl.log.Infof("Create StudentTestAction Entry: %v", en)

	en.CreatedAt = time.Now()
	en.UpdatedAt = time.Now()

	item, err := dynamodbattribute.MarshalMap(en)
	if err != nil {
		errorMsg := fmt.Sprintf("Error occurred while marshalling the record for StudentTestAction error: %v", err)
		studentTestActionRepositoryImpl.log.Errorf(errorMsg)
		return pbTypes.ErrorStudentTestActionCreateFailed(errorMsg)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(studentTestActionTableName),
		Item:      item,
	}

	_, err = studentTestActionRepositoryImpl.ddb.PutItem(input)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to create StudentTestAction entry, err: %v", err)
		studentTestActionRepositoryImpl.log.Errorf(errorMsg)
		return pbTypes.ErrorStudentTestActionCreateFailed(errorMsg)
	}
	return nil
}

func (*studentTestActionRepositoryImpl) Update(_ context.Context, _ *entity.StudentTestActionEntity) (*entity.StudentTestActionEntity, error) {
	//not implemented
	err := errors.New("update operation not implemented")
	return nil, err
}

func (*studentTestActionRepositoryImpl) FindByID(_ context.Context, _ string) (*entity.StudentTestActionEntity, error) {
	//not implemented

	err := errors.New("findByID operation not implemented")
	return nil, err
}

func (*studentTestActionRepositoryImpl) List(_ context.Context, _ uint, _ int) ([]*entity.StudentTestActionEntity, error) {
	//not implemented
	err := errors.New("list operation not implemented")
	return nil, err
}

func (studentTestActionRepositoryImpl *studentTestActionRepositoryImpl) FindAllByTestIDStudentID(ctx context.Context, testID string, studentID string) ([]*entity.StudentTestActionEntity, error) {
	studentTestActionRepositoryImpl.log.Infof("FindAllResponsesForStudentTest repo: testId : %s & studentId : %s", testID, studentID)

	testIDStudentID := testID + "#" + studentID
	keyCondition := expression.Key(studentTestActionTablePartitionKey).Equal(expression.Value(testIDStudentID))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		errorMsg := fmt.Sprintf("Failed to build dynamo expression: %v", err)
		studentTestActionRepositoryImpl.log.Errorf(errorMsg)
		return nil, pbTypes.ErrorStudentTestActionGetFailed(errorMsg)
	}

	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(studentTestActionTableName),
	}

	result, err := studentTestActionRepositoryImpl.ddb.QueryWithContext(ctx, params)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to query in dynamodb: %v", err)
		studentTestActionRepositoryImpl.log.Errorf(errorMsg)
		return nil, pbTypes.ErrorStudentTestActionGetFailed(errorMsg)
	}

	items := make([]*entity.StudentTestActionEntity, len(result.Items))

	for index, value := range result.Items {
		responseEntity := entity.StudentTestActionEntity{}

		err = dynamodbattribute.UnmarshalMap(value, &responseEntity)

		if err != nil {
			errorMsg := fmt.Sprintf("Failed to unmarshall dynamodb attributes to Event: %v", err)
			studentTestActionRepositoryImpl.log.Errorf(errorMsg)
			return nil, pbTypes.ErrorStudentTestActionGetFailed(errorMsg)
		}
		items[index] = &responseEntity
	}
	studentTestActionRepositoryImpl.log.Infof("FindAllResponsesForStudentTest repo: testId : %s & studentId : %s, response : %v", testID, studentID, items)
	return items, err
}

func (studentTestActionRepositoryImpl *studentTestActionRepositoryImpl) FindAllByTestIDStudentIDNamespace(ctx context.Context, testID, studentID, namespace string) ([]*entity.StudentTestActionEntity, error) {
	studentTestActionRepositoryImpl.log.Infof("FindAllByTestIDStudentIDNamespace repo: testId : %s & studentId : %s", testID, studentID)

	testIDStudentID := testID + "#" + studentID
	keyCondition := expression.Key(studentTestActionTablePartitionKey).Equal(expression.Value(testIDStudentID)).And(
		expression.Key(studentTestActionTableSortKey).BeginsWith(namespace))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		errorMsg := fmt.Sprintf("Failed to build dynamo expression: %v", err)
		studentTestActionRepositoryImpl.log.WithContext(ctx).Errorf(errorMsg)
		return nil, pbTypes.ErrorStudentTestActionGetFailed(errorMsg)
	}

	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(studentTestActionTableName),
	}

	result, err := studentTestActionRepositoryImpl.ddb.QueryWithContext(ctx, params)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to scan items in dynamodb: %v", err)
		studentTestActionRepositoryImpl.log.WithContext(ctx).Errorf(errorMsg)
		return nil, pbTypes.ErrorStudentTestActionGetFailed(errorMsg)
	}

	items := make([]*entity.StudentTestActionEntity, len(result.Items))

	for index, value := range result.Items {
		responseEntity := entity.StudentTestActionEntity{}

		err = dynamodbattribute.UnmarshalMap(value, &responseEntity)

		if err != nil {
			errorMsg := fmt.Sprintf("Failed to unmarshall dynamodb attributes to Event: %v", err)
			studentTestActionRepositoryImpl.log.WithContext(ctx).Errorf(errorMsg)
			return nil, pbTypes.ErrorStudentTestActionGetFailed(errorMsg)
		}
		items[index] = &responseEntity
	}
	return items, err
}
