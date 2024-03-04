package ddb

import (
	"context"
	"fmt"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createStudentTestActionTable() {
	studentTestActionTable := studentTestActionTableName

	tableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(studentTestActionTablePartitionKey),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(studentTestActionTablePartitionKey),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(studentTestActionTable),
	}
	_, err := ddb.CreateTable(tableInput)

	if err != nil {
		log.Fatalf("Failed to create student test action table: %s, err: %s", studentTestActionTable, err.Error())
	}
	log.Info("Student test action table created successfully!")
}

func setupStudentTestAction() (context.Context, StudentTestActionRepository) {
	ctx := context.Background()
	d := &Client{
		DynamoDB: ddb,
	}
	repo := NewStudentTestActionRepositoryImpl(d, log.DefaultLogger)
	return ctx, repo
}

func TestStudentTestActionRepositoryImpl_Create(t *testing.T) {
	c, repo := setupStudentTestAction()

	type args struct {
		ctx context.Context
		en  *entity.StudentTestActionEntity
	}

	tests := []struct {
		name          string
		args          args
		wantErr       bool
		errorResponse *errors.Error
		setup         func()
		tearDown      func()
	}{
		{
			name: "Create Student Response success",
			args: args{
				ctx: c,
				en: &entity.StudentTestActionEntity{
					TestIDStudentID: "test_123#Student_ID_1",
					SetID:           "SID_01",
					StudentID:       "Student_ID_1",
					QuestionID:      "QID_01",
					SectionName:     "Physics",
					MarkedResponse:  "A",
					ActionType:      "QUESTION_ACTION_ANSWERED",
					//CreatedAt:       time.Now(),
					//UpdatedAt:       time.Now(),
				},
			},
			wantErr:       false,
			errorResponse: nil,
			tearDown: func() {
				t.Log("teardown")
				input := &dynamodb.DeleteItemInput{
					TableName: aws.String(studentTestActionTableName),
					Key: map[string]*dynamodb.AttributeValue{
						studentTestActionTablePartitionKey: {
							S: aws.String("test_123#Student_ID_1"),
						},
					},
				}
				t.Log("input")
				_, err := ddb.DeleteItem(input)
				if err != nil {
					t.Errorf("error while deleting the studentTestAction record: %v", err.Error())
				}
			},
			setup: func() {

			},
		},
		{
			name: "Create Student Response failure",
			args: args{
				ctx: c,
				en: &entity.StudentTestActionEntity{
					//TestIDStudentID: "test_123#Student_ID_1",
					SetID:          "SID_01",
					StudentID:      "Student_ID_1",
					QuestionID:     "QID_01",
					SectionName:    "Physics",
					MarkedResponse: "A",
					ActionType:     "QUESTION_ACTION_ANSWERED",
					//CreatedAt:       time.Now(),
					//UpdatedAt:       time.Now(),
				},
			},
			wantErr:       true,
			errorResponse: pbTypes.ErrorStudentTestActionCreateFailed(""),
			tearDown: func() {

			},
			setup: func() {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			err := repo.Create(tt.args.ctx, tt.args.en)
			if tt.wantErr {
				assert.Error(t, err)
				Err := errors.FromError(err)
				if Err != nil {
					assert.Equal(t, tt.errorResponse.Is(Err), true)
				}
			} else {
				assert.NoError(t, err)
			}
			tt.tearDown()
		})
	}
}

func TestStudentTestActionRepositoryImpl_FindAllByTestIDStudentID(t *testing.T) {
	c, repo := setupStudentTestAction()

	type args struct {
		ctx       context.Context
		testId    string
		StudentId string
	}

	tests := []struct {
		name          string
		args          args
		want          []entity.StudentTestActionEntity
		wantErr       bool
		errorResponse *errors.Error
		setup         func()
		teardown      func()
	}{
		{
			name: "Find all by testIdStudent response successful",
			args: args{
				ctx:       c,
				testId:    "test_123",
				StudentId: "Student_ID_1",
			},
			want: []entity.StudentTestActionEntity{
				{
					TestIDStudentID: "test_123#Student_ID_1",
					SetID:           "SID_01",
					StudentID:       "Student_ID_1",
					QuestionID:      "QID_01",
					SectionName:     "Physics",
					MarkedResponse:  "A",
					ActionType:      "QUESTION_ACTION_ANSWERED",
					//CreatedAt:       time.Now(),
					//UpdatedAt:       time.Now(),
				},
			},
			wantErr:       false,
			errorResponse: nil,
			setup: func() {
				t.Log("setup student test action")
				err := repo.Create(c, &entity.StudentTestActionEntity{
					TestIDStudentID: "test_123#Student_ID_1",
					SetID:           "SID_01",
					StudentID:       "Student_ID_1",
					QuestionID:      "QID_01",
					SectionName:     "Physics",
					MarkedResponse:  "A",
					ActionType:      "QUESTION_ACTION_ANSWERED",
					//CreatedAt:       time.Now(),
					//UpdatedAt:       time.Now(),
				})
				if err != nil {
					log.Fatalf("setup student test action failed")
				}
			},
			teardown: func() {
				t.Log("teardown")
				input := &dynamodb.DeleteItemInput{
					TableName: aws.String(studentTestActionTableName),
					Key: map[string]*dynamodb.AttributeValue{
						studentTestActionTablePartitionKey: {
							S: aws.String("test_123#Student_ID_1"),
						},
					},
				}
				t.Log("delete")
				_, err := ddb.DeleteItem(input)
				if err != nil {
					t.Errorf("error while deleting the record: %v", err.Error())
				}
			},
		},
		{
			name: "Failed to find StudentTestActions",
			args: args{
				ctx:       c,
				testId:    "test_1234",
				StudentId: "Student_ID_1",
			},
			want: []entity.StudentTestActionEntity{
				{
					TestIDStudentID: "test_1234#Student_ID_1",
					SetID:           "SID_01",
					StudentID:       "Student_ID_1",
					QuestionID:      "QID_01",
					SectionName:     "Physics",
					MarkedResponse:  "A",
					ActionType:      "QUESTION_ACTION_ANSWERED",
					//CreatedAt:       time.Now(),
					//UpdatedAt:       time.Now(),
				},
			},
			wantErr:       true,
			errorResponse: nil,
			setup: func() {
				t.Log("setup student test action")
				err := repo.Create(c, &entity.StudentTestActionEntity{
					TestIDStudentID: "test_123#Student_ID_1",
					SetID:           "SID_01",
					StudentID:       "Student_ID_1",
					QuestionID:      "QID_01",
					SectionName:     "Physics",
					MarkedResponse:  "A",
					ActionType:      "QUESTION_ACTION_ANSWERED",
					//CreatedAt:       time.Now(),
					//UpdatedAt:       time.Now(),
				})
				if err != nil {
					log.Fatalf("setup student test action failed")
				}
			},
			teardown: func() {
				t.Log("teardown")
				input := &dynamodb.DeleteItemInput{
					TableName: aws.String(studentTestActionTableName),
					Key: map[string]*dynamodb.AttributeValue{
						studentTestActionTablePartitionKey: {
							S: aws.String("test_123#Student_ID_1"),
						},
					},
				}
				t.Log("delete")
				_, err := ddb.DeleteItem(input)
				if err != nil {
					t.Errorf("error while deleting the record: %v", err.Error())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := repo.FindAllByTestIDStudentID(tt.args.ctx, tt.args.testId, tt.args.StudentId)
			if !tt.wantErr {
				assert.Equal(t, tt.wantErr, err != nil, "FindById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, len(tt.want), len(got), "FindAllByTestIDStudentID length got = %v, want %v", len(got), len(tt.want))
				if got != nil && len(tt.want) > 0 {
					assert.NotEmpty(t, got)
					if tt.want[0].TestIDStudentID != "" {
						assert.Equal(t, tt.want[0].TestIDStudentID, got[0].TestIDStudentID, "FindAllByTestIDStudentID() got = %v, want %v", got[0].TestIDStudentID, tt.want[0].TestIDStudentID)
					}
					assert.Equal(t, tt.want[0].SetID, got[0].SetID, "FindAllByTestIDStudentID() got = %v, want %v", got[0].SetID, tt.want[0].SetID)
					assert.Equal(t, tt.want[0].StudentID, got[0].StudentID, "FindAllByTestIDStudentID() got = %v, want %v", got[0].StudentID, tt.want[0].StudentID)
					assert.Equal(t, tt.want[0].MarkedResponse, got[0].MarkedResponse, "FindAllByTestIDStudentID() got = %v, want %v", got[0].MarkedResponse, tt.want[0].MarkedResponse)
					assert.Equal(t, tt.want[0].QuestionID, got[0].QuestionID, "FindAllByTestIDStudentID() got = %v, want %v", got[0].QuestionID, tt.want[0].QuestionID)
					assert.Equal(t, tt.want[0].SectionName, got[0].SectionName, "FindAllByTestIDStudentID() got = %v, want %v", got[0].SectionName, tt.want[0].SectionName)
					assert.Equal(t, tt.want[0].ActionType, got[0].ActionType, "FindAllByTestIDStudentID() got = %v, want %v", got[0].ActionType, tt.want[0].ActionType)
				}
			} else {
				assert.NotEqual(t, tt.wantErr, err != nil, "FindById() error = %v, wantErr %v", err, tt.wantErr)
				assert.NotEqual(t, len(tt.want), len(got), "FindAllByTestIDStudentID length got = %v, want %v", len(got), len(tt.want))
			}

			tt.teardown()
		})
	}
}

func TestStudentTestActionRepositoryImpl_List(t *testing.T) {
	c, repo := setupStudentTestAction()

	type args struct {
		ctx    context.Context
		offset uint
		limit  int
	}
	tests := []struct {
		name        string
		args        args
		want        []*entity.StudentTestActionEntity
		wantErr     bool
		errResponse error
	}{
		{
			name: "list operation not implemented",
			args: args{
				ctx:    c,
				offset: 1,
				limit:  1,
			},
			want:        nil,
			wantErr:     true,
			errResponse: fmt.Errorf("list operation not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.List(tt.args.ctx, tt.args.offset, tt.args.limit)
			if tt.wantErr {
				assert.Error(t, err)
				if err != nil {
					assert.Equal(t, tt.errResponse, err, "List() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got, "List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentTestActionRepositoryImpl_Update(t *testing.T) {
	c, repo := setupStudentTestAction()
	type args struct {
		ctx context.Context
		en  *entity.StudentTestActionEntity
	}
	tests := []struct {
		name        string
		args        args
		want        *entity.StudentTestActionEntity
		wantErr     bool
		errResponse error
	}{
		{
			name: "Update studentTestAction not implemented",
			args: args{
				ctx: c,
				en:  &entity.StudentTestActionEntity{},
			},
			want:        nil,
			wantErr:     true,
			errResponse: fmt.Errorf("update operation not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Update(tt.args.ctx, tt.args.en)
			if tt.wantErr {
				assert.Error(t, err)
				if err != nil {
					assert.Equal(t, tt.errResponse, err, "Update() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStudentTestActionRepositoryImpl_FindByID(t *testing.T) {
	c, repo := setupStudentTestAction()

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name        string
		args        args
		want        *entity.StudentTestActionEntity
		wantErr     bool
		errResponse error
	}{
		{
			name: "FindbyId Not Implemented",
			args: args{
				ctx: c,
				id:  "",
			},
			want:        nil,
			wantErr:     true,
			errResponse: fmt.Errorf("findByID operation not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindByID(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				if err != nil {
					assert.Equal(t, tt.errResponse, err, "FindByID() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got, "FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
