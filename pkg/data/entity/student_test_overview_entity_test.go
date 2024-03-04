package entity

import (
	pbEnums "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/enums"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestStudentTestOverviewEntity_FromPB(t *testing.T) {
	var timeInt int64 = 1405544146
	timeToTest := time.Unix(timeInt, 0).UTC()

	type args struct {
		studentTestOverviewInfo *pbTypes.StudentTestOverviewInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *StudentTestOverviewEntity
		wantErr bool
	}{
		{
			name: "Test Successful Parsing without time",
			args: args{
				studentTestOverviewInfo: getMockStudentTestOverviewInfoWithoutTime(),
			},
			want:    getMockStudentTestOverviewEntityWithoutTime(),
			wantErr: false,
		},
		{
			name: "Test Time Parsing",
			args: args{
				studentTestOverviewInfo: getMockStudentTestOverviewInfo(timeToTest),
			},
			want:    getMockStudentTestOverviewEntity(timeToTest),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &StudentTestOverviewEntity{}
			entity.FromPB(tt.args.studentTestOverviewInfo)
			if !tt.wantErr {
				assert.Equal(t, tt.want, entity)
			} else {
				assert.NotEqual(t, tt.want, entity)
			}

		})
	}
}

func getMockStudentTestOverviewInfoWithoutTime() *pbTypes.StudentTestOverviewInfo {
	info := &pbTypes.StudentTestOverviewInfo{
		StudentId:  "123",
		TestId:     "456",
		TestStatus: pbEnums.StudentTestStatus_STS_INITIATED,
		SetId:      "Set-A",
	}
	return info
}

func getMockStudentTestOverviewEntityWithoutTime() *StudentTestOverviewEntity {
	entity := &StudentTestOverviewEntity{
		StudentID:  "123",
		TestID:     "456",
		TestStatus: "STS_INITIATED",
		SetID:      "Set-A",
	}
	return entity
}

func getMockStudentTestOverviewInfo(t time.Time) *pbTypes.StudentTestOverviewInfo {
	info := &pbTypes.StudentTestOverviewInfo{
		StudentId:  "123",
		TestId:     "456",
		TestStatus: pbEnums.StudentTestStatus_STS_IN_PROGRESS,
		SetId:      "Set-A",
		StartTime:  timestamppb.New(t),
		FinishTime: timestamppb.New(t),
	}
	return info
}

func getMockStudentTestOverviewEntity(time time.Time) *StudentTestOverviewEntity {
	entity := &StudentTestOverviewEntity{
		StudentID:  "123",
		TestID:     "456",
		TestStatus: "STS_IN_PROGRESS",
		SetID:      "Set-A",
		StartTime:  time,
		FinishTime: time,
	}
	return entity
}
