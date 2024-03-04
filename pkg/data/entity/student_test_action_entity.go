package entity

import (
	pbEnums "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/enums"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type StudentTestActionEntity struct {
	TestIDStudentID            string    `dynamodbav:"TestIdStudentId"`            // Partition Key
	SectionNamespaceQuestionID string    `dynamodbav:"SectionNamespaceQuestionId"` //sort key
	StudentID                  string    `dynamodbav:"StudentId"`
	SetID                      string    `dynamodbav:"SetId"`
	SectionName                string    `dynamodbav:"SectionName"`
	SectionNamespace           string    `dynamodbav:"SectionNamespace"`
	QuestionID                 string    `dynamodbav:"QuestionId"`
	QuestionSequenceNo         int32     `dynamodbav:"QuestionSequenceNo"`
	ContentLanguage            string    `dynamodbav:"ContentLanguage"`
	ContentId                  string    `dynamodbav:"ContentId"`
	MarkedResponse             string    `dynamodbav:"MarkedResponse"`
	ActionType                 string    `dynamodbav:"ActionType"`
	CreatedAt                  time.Time `dynamodbav:"CreatedAt,unixtime"`
	UpdatedAt                  time.Time `dynamodbav:"UpdatedAt,unixtime"`
}

func (entity *StudentTestActionEntity) FromPB(info *pbTypes.StudentTestActionInfo) {
	if len(info.GetTestId()) > 0 && len(info.GetStudentId()) > 0 {
		entity.TestIDStudentID = info.GetTestId() + "#" + info.GetStudentId()
	}
	entity.SetID = info.GetSetId()
	entity.StudentID = info.GetStudentId()
	entity.QuestionID = info.GetQuestionId()
	entity.SectionName = info.GetSectionName()
	entity.MarkedResponse = info.GetMarkedResponse()
	entity.ActionType = info.ActionType.String()
}

func (entity *StudentTestActionEntity) ToPB(info *pbTypes.StudentTestActionInfo) {
	info.TestId = strings.Split(entity.TestIDStudentID, "#")[0]
	info.SetId = entity.SetID
	info.StudentId = strings.Split(entity.TestIDStudentID, "#")[1]
	info.QuestionId = entity.QuestionID
	info.SectionName = entity.SectionName
	info.MarkedResponse = entity.MarkedResponse
	info.ActionType = pbEnums.QuestionAction(pbEnums.QuestionAction_value[entity.ActionType])
	info.CreatedAt = timestamppb.New(entity.CreatedAt)
	info.UpdatedAt = timestamppb.New(entity.UpdatedAt)
}
