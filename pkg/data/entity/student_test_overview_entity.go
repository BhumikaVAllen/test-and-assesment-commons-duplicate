package entity

import (
	enums "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/enums"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// StudentTestOverviewEntity TODO : Add TestMode
type StudentTestOverviewEntity struct {
	StudentID    string    `bson:"studentId"`
	TestID       string    `bson:"testId"`
	TestCategory string    `bson:"testCategory"`
	TestStatus   string    `bson:"testStatus"`
	SetID        string    `bson:"setId"`
	StartTime    time.Time `bson:"startTime,unixtime"`
	FinishTime   time.Time `bson:"finishTime,unixtime"`
	CreatedAt    time.Time `bson:"createdAt,unixtime"`
	UpdatedAt    time.Time `bson:"updatedAt,unixtime"`
}

func (entity *StudentTestOverviewEntity) FromPB(info *pbTypes.StudentTestOverviewInfo) {
	entity.StudentID = info.GetStudentId()
	entity.TestID = info.GetTestId()
	entity.TestStatus = info.GetTestStatus().String()
	entity.SetID = info.GetSetId()
	if info.StartTime != nil {
		entity.StartTime = info.StartTime.AsTime()
	}
	if info.FinishTime != nil {
		entity.FinishTime = info.FinishTime.AsTime()
	}
}

func (entity *StudentTestOverviewEntity) ToPB(info *pbTypes.StudentTestOverviewInfo) {
	info.StudentId = entity.StudentID
	info.TestId = entity.TestID
	info.TestStatus = enums.StudentTestStatus(enums.StudentTestStatus_value[entity.TestStatus])
	info.StartTime = timestamppb.New(entity.StartTime)
	info.FinishTime = timestamppb.New(entity.FinishTime)
	info.SetId = entity.SetID
}
