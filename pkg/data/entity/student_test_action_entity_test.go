package entity

import (
	pbEnums "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/enums"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestStudentTestActionEntity_FromPB(t *testing.T) {
	type args struct {
		studentTestActionInfo *pbTypes.StudentTestActionInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *StudentTestActionEntity
		wantErr bool
	}{
		{
			name: "Successful",
			args: args{
				studentTestActionInfo: getMockStudentTestActionInfoWithoutTime(),
			},
			want:    getMockStudentTestActionEntityWithoutTime(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &StudentTestActionEntity{}
			entity.FromPB(tt.args.studentTestActionInfo)
			if !tt.wantErr {
				assert.Equal(t, tt.want, entity)
			} else {
				assert.NotEqual(t, tt.want, entity)
			}

		})
	}
}

func TestStudentTestActionEntity_ToPB(t *testing.T) {
	timeNow := time.Now()
	type args struct {
		entity *StudentTestActionEntity
	}
	tests := []struct {
		name    string
		args    args
		want    *pbTypes.StudentTestActionInfo
		wantErr bool
	}{
		{
			name: "Successful",
			args: args{
				entity: getMockStudentTestActionEntity(timeNow),
			},
			want:    getMockStudentTestActionInfo(timeNow),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := &pbTypes.StudentTestActionInfo{}
			entity := tt.args.entity
			entity.ToPB(info)
			if !tt.wantErr {
				assert.Equal(t, tt.want, info)
			} else {
				assert.NotEqual(t, tt.want, info)
			}
		})
	}
}

func getMockStudentTestActionInfoWithoutTime() *pbTypes.StudentTestActionInfo {
	info := &pbTypes.StudentTestActionInfo{
		TestId:         "123",
		StudentId:      "456",
		SetId:          "Set-A",
		QuestionId:     "789",
		SectionName:    "Math",
		MarkedResponse: "A",
		ActionType:     pbEnums.QuestionAction_ANSWERED,
	}
	return info
}

func getMockStudentTestActionEntityWithoutTime() *StudentTestActionEntity {
	entity := &StudentTestActionEntity{
		TestIDStudentID: "123#456",
		StudentID:       "456",
		SetID:           "Set-A",
		QuestionID:      "789",
		SectionName:     "Math",
		MarkedResponse:  "A",
		ActionType:      "ANSWERED",
	}
	return entity
}

func getMockStudentTestActionEntity(time time.Time) *StudentTestActionEntity {
	entity := &StudentTestActionEntity{
		TestIDStudentID: "123#456",
		StudentID:       "456",
		SetID:           "Set-A",
		QuestionID:      "789",
		SectionName:     "Math",
		MarkedResponse:  "A",
		ActionType:      "ANSWERED",
		CreatedAt:       time,
		UpdatedAt:       time,
	}
	return entity
}

func getMockStudentTestActionInfo(time time.Time) *pbTypes.StudentTestActionInfo {
	info := &pbTypes.StudentTestActionInfo{
		TestId:         "123",
		StudentId:      "456",
		SetId:          "Set-A",
		QuestionId:     "789",
		SectionName:    "Math",
		MarkedResponse: "A",
		ActionType:     pbEnums.QuestionAction_ANSWERED,
		CreatedAt:      timestamppb.New(time),
		UpdatedAt:      timestamppb.New(time),
	}
	return info
}
