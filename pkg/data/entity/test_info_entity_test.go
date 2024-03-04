package entity

// Commenting temporarily
//import (
//	"github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/enums"
//	"github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
//	"github.com/Allen-Career-Institute/common-protos/test_and_assessment_creator/v1/request"
//	v1 "github.com/Allen-Career-Institute/common-protos/test_and_assessment_creator/v1/response"
//	"github.com/go-kratos/kratos/v2/errors"
//	"github.com/golang/protobuf/ptypes/timestamp"
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)
//
//// nolint: funlen
//func TestTaaCreatorEntity_FromPB(t *testing.T) {
//
//	type args struct {
//		req *request.CreateTaaCreatorRequest
//	}
//
//	tests := []struct {
//		name          string
//		args          args
//		want          *TestInfoEntity
//		wantErr       bool
//		errorResponse *errors.Error
//	}{
//		{
//			name: "Test TAA Creator Entity From PB",
//			args: args{
//				req: &request.CreateTaaCreatorRequest{
//					QuestionPaperId: "6a7eb986-181f-11ee-be56-0242ac120002",
//					DisplayName:     "NEET_UG_2020",
//					Assignment: &types.TestAssignment{
//						Batches:  []string{"BATCH_0001"},
//						Centers:  []string{"CENTER_001"},
//						Students: []string{},
//					},
//					Mode:      enums.TestMode_ONLINE,
//					Stream:    enums.TestStream_MD,
//					Category:  enums.TestCategory_CLASSROOM,
//					CreatedBy: "Mohit",
//					Schedule: &types.TestSchedule{
//						StartTime: &timestamp.Timestamp{
//							Seconds: 1692192600,
//						},
//						EndTime: &timestamp.Timestamp{
//							Seconds: 1692193500,
//						},
//						Duration: &types.Duration{
//							Value: 2,
//							Unit:  enums.DurationUnit_HOUR,
//						},
//						Type: enums.TestScheduleType_FIXED,
//					},
//					MetaData: &types.TestMetaData{
//						SubCategory: "MAJOR",
//					},
//				},
//			},
//			want: &TestInfoEntity{
//				QuestionPaperID: "6a7eb986-181f-11ee-be56-0242ac120002",
//				DisplayName:     "NEET_UG_2020",
//				Assignment: Assignment{
//					Batches:  []string{"BATCH_0001"},
//					Centers:  []string{"CENTER_001"},
//					Students: []string{},
//				},
//				Mode:      enums.TestMode_ONLINE.String(),
//				Stream:    enums.TestStream_MD.String(),
//				Category:  enums.TestCategory_CLASSROOM.String(),
//				CreatedBy: "Mohit",
//				Schedule: Schedule{
//					StartTime: time.Unix(1692192600, 0),
//					EndTime:   time.Unix(1692193500, 0),
//					Duration: Duration{
//						Value: 2,
//						Unit:  enums.DurationUnit_HOUR.String(),
//					},
//					Type: enums.TestScheduleType_FIXED.String(),
//				},
//				MetaData: MetaData{
//					SubCategory: "MAJOR",
//				},
//				CreatedAt: time.Now(),
//				UpdatedAt: time.Now(),
//			},
//			wantErr:       false,
//			errorResponse: nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			entity := &TestInfoEntity{}
//			entity.FromPB(tt.args.req)
//			assert.Equal(t, entity.QuestionPaperID, tt.want.QuestionPaperID)
//			assert.Equal(t, entity.DisplayName, tt.want.DisplayName)
//			assert.Equal(t, entity.Assignment, tt.want.Assignment)
//			assert.Equal(t, entity.Mode, tt.want.Mode)
//			assert.Equal(t, entity.Category, tt.want.Category)
//			assert.Equal(t, entity.Stream, tt.want.Stream)
//			assert.Equal(t, entity.Schedule, tt.want.Schedule)
//			assert.Equal(t, entity.CreatedBy, tt.want.CreatedBy)
//			assert.Equal(t, entity.MetaData, tt.want.MetaData)
//			assert.Equal(t, entity.Schedule.Type, tt.want.Schedule.Type)
//			assert.Equal(t, entity.Schedule.StartTime, tt.want.Schedule.StartTime)
//			assert.Equal(t, entity.Schedule.EndTime, tt.want.Schedule.EndTime)
//			assert.Equal(t, entity.Schedule.Duration.Value, tt.want.Schedule.Duration.Value)
//			assert.Equal(t, entity.Schedule.Duration.Unit, tt.want.Schedule.Duration.Unit)
//		})
//	}
//}
//
//// nolint: funlen
//func TestTaaCreatorEntity_ToPB(t *testing.T) {
//	type args struct {
//		entity *TestInfoEntity
//	}
//
//	currentTime := time.Now()
//
//	tests := []struct {
//		name          string
//		args          args
//		want          *v1.GetTaaCreatorResponse
//		wantErr       bool
//		errorResponse *errors.Error
//	}{
//		{
//			name: "Test TaaCreatorResponse From Entity",
//			args: args{
//				entity: &TestInfoEntity{
//					ID:          "6a7eb986-181f-11ee-be56-0242ac120002",
//					QuestionPaperID: "6a7eb986-181f-11ee-be56-0242ac120002",
//					DisplayName:     "NEET_UG_2020",
//					Assignment: Assignment{
//						Batches:  []string{"BATCH_0001"},
//						Centers:  []string{"CENTER_001"},
//						Students: []string{},
//					},
//					Mode:      enums.TestMode_ONLINE.String(),
//					Stream:    enums.TestStream_MD.String(),
//					Category:  enums.TestCategory_CLASSROOM.String(),
//					CreatedBy: "Mohit",
//					Schedule: Schedule{
//						StartTime: time.Unix(1692192600, 0),
//						EndTime:   time.Unix(1692193500, 0),
//						Duration: Duration{
//							Value: 2,
//							Unit:  enums.DurationUnit_HOUR.String(),
//						},
//						Type: enums.TestScheduleType_FIXED.String(),
//					},
//					MetaData: MetaData{
//						SubCategory: "MAJOR",
//					},
//					QuestionSets: QuestionSets{
//						NoOfSets: 1,
//						Sets:     []string{"Set-A"},
//					},
//					QuestionPaperStats: QuestionPaperStats{
//						TotalQuestions: 3,
//						Sections:       getEntitySections(),
//					},
//					CreatedAt: currentTime,
//					UpdatedAt: currentTime,
//					UpdatedBy: "",
//				},
//			},
//			want: &v1.GetTaaCreatorResponse{
//				TestId:          "6a7eb986-181f-11ee-be56-0242ac120002",
//				QuestionPaperId: "6a7eb986-181f-11ee-be56-0242ac120002",
//				DisplayName:     "NEET_UG_2020",
//				Assignment: &types.TestAssignment{
//					Batches:  []string{"BATCH_0001"},
//					Centers:  []string{"CENTER_001"},
//					Students: []string{},
//				},
//				Mode:     enums.TestMode_ONLINE,
//				Stream:   enums.TestStream_MD,
//				Category: enums.TestCategory_CLASSROOM,
//
//				Schedule: &types.TestSchedule{
//					StartTime: &timestamp.Timestamp{
//						Seconds: 1692192600000,
//					},
//					EndTime: &timestamp.Timestamp{
//						Seconds: 1692193500000,
//					},
//					Duration: &types.Duration{
//						Value: 2,
//						Unit:  enums.DurationUnit_HOUR,
//					},
//					Type: enums.TestScheduleType_FIXED,
//				},
//				MetaData: &types.TestMetaData{
//					SubCategory: "MAJOR",
//				},
//				QuestionSets: &v1.QuestionSets{
//					NoOfSets: 1,
//					Sets:     []string{"Set-A"},
//				},
//				QuestionPaperStats: &v1.QuestionPaperStats{
//					TotalQuestions: 3,
//					Sections:       getPBSections(),
//				},
//				CreatedBy: "Mohit",
//				CreatedAt: &timestamp.Timestamp{
//					Seconds: currentTime.UnixMilli(),
//				},
//				UpdatedAt: &timestamp.Timestamp{
//					Seconds: currentTime.UnixMilli(),
//				},
//				UpdatedBy: "",
//			},
//			wantErr:       false,
//			errorResponse: nil,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			info := &v1.GetTaaCreatorResponse{}
//			entity := tt.args.entity
//			entity.ToPB(info)
//			assert.Equal(t, tt.want, info)
//		})
//	}
//}
//
//func TestTaaCreatorEntity_FromUpdatePB(t *testing.T) {
//
//	type args struct {
//		req *request.UpdateTaaCreatorRequest
//	}
//
//	tests := []struct {
//		name          string
//		args          args
//		want          *TestInfoEntity
//		wantErr       bool
//		errorResponse *errors.Error
//	}{
//		{
//			name: "Test TAA Creator Entity From PB",
//			args: args{
//				req: &request.UpdateTaaCreatorRequest{
//					DisplayName: "NEET_UG_2020",
//					Mode:        enums.TestMode_ONLINE,
//					Stream:      enums.TestStream_MD,
//					Category:    enums.TestCategory_CLASSROOM,
//					ScheduleId:  "schedule_123",
//					Assignment: &request.Assignment{
//						Batches: &request.Batches{
//							Add:    []string{"BATCH_0001"},
//							Remove: []string{"X"},
//						},
//						Centers: &request.Centers{
//							Add:    []string{"CENTER_001"},
//							Remove: []string{"Y"},
//						},
//						Students: &request.Students{
//							Add:    []string{"ST_001"},
//							Remove: []string{"Z"},
//						},
//					},
//					Schedule: &types.TestSchedule{
//						StartTime: &timestamp.Timestamp{
//							Seconds: 1692192600,
//						},
//						EndTime: &timestamp.Timestamp{
//							Seconds: 1692193500,
//						},
//						Duration: &types.Duration{
//							Value: 1,
//							Unit:  enums.DurationUnit_HOUR,
//						},
//						Type: enums.TestScheduleType_FIXED,
//					},
//					MetaData: &types.TestMetaData{
//						SubCategory: "MAJOR",
//					},
//				},
//			},
//			want: &TestInfoEntity{
//				QuestionPaperID: "6a7eb986-181f-11ee-be56-0242ac120002",
//				DisplayName:     "NEET_UG_2020",
//				Mode:            enums.TestMode_ONLINE.String(),
//				Stream:          enums.TestStream_MD.String(),
//				Category:        enums.TestCategory_CLASSROOM.String(),
//				ScheduleID:      "schedule_123",
//				Assignment: Assignment{
//					Batches:  []string{"BATCH_0001"},
//					Centers:  []string{"CENTER_001"},
//					Students: []string{"ST_001"},
//				},
//				Schedule: Schedule{
//					StartTime: time.Unix(1692192600, 0),
//					EndTime:   time.Unix(1692193500, 0),
//					Duration: Duration{
//						Value: 1,
//						Unit:  enums.DurationUnit_HOUR.String(),
//					},
//					Type: enums.TestScheduleType_FIXED.String(),
//				},
//				MetaData: MetaData{
//					SubCategory: "MAJOR",
//				},
//				CreatedAt: time.Now(),
//				UpdatedAt: time.Now(),
//			},
//			wantErr:       false,
//			errorResponse: nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			entity := &TestInfoEntity{
//				Assignment: Assignment{
//					Batches:  []string{"X"},
//					Centers:  []string{"Y"},
//					Students: []string{"Z"},
//				},
//			}
//			entity.FromUpdatePB(tt.args.req)
//			assert.Equal(t, entity.DisplayName, tt.want.DisplayName)
//			assert.Equal(t, entity.Mode, tt.want.Mode)
//			assert.Equal(t, entity.Category, tt.want.Category)
//			assert.Equal(t, entity.Stream, tt.want.Stream)
//			assert.Equal(t, entity.Assignment.Batches, tt.want.Assignment.Batches)
//			assert.Equal(t, entity.Assignment.Centers, tt.want.Assignment.Centers)
//			assert.Equal(t, entity.Assignment.Students, tt.want.Assignment.Students)
//			assert.Equal(t, entity.Schedule, tt.want.Schedule)
//			assert.Equal(t, entity.CreatedBy, tt.want.CreatedBy)
//			assert.Equal(t, entity.MetaData, tt.want.MetaData)
//			assert.Equal(t, entity.Schedule.Type, tt.want.Schedule.Type)
//			assert.Equal(t, entity.Schedule.StartTime, tt.want.Schedule.StartTime)
//			assert.Equal(t, entity.Schedule.EndTime, tt.want.Schedule.EndTime)
//			assert.Equal(t, entity.Schedule.Duration.Value, tt.want.Schedule.Duration.Value)
//			assert.Equal(t, entity.Schedule.Duration.Unit, tt.want.Schedule.Duration.Unit)
//		})
//	}
//}
//
//func getPBSections() []*v1.Sections {
//	sectionA := &v1.Sections{
//		SectionName:    "BIOLOGY",
//		TotalQuestions: 1,
//	}
//
//	sectionB := &v1.Sections{
//		SectionName:    "CHEMISTRY",
//		TotalQuestions: 1,
//	}
//
//	sectionC := &v1.Sections{
//		SectionName:    "PHYSICS",
//		TotalQuestions: 1,
//	}
//	var sectionList []*v1.Sections
//	sectionList = append(sectionList, sectionA)
//	sectionList = append(sectionList, sectionB)
//	sectionList = append(sectionList, sectionC)
//
//	return sectionList
//}
//
//func getEntitySections() []Sections {
//	sectionA := Sections{
//		SectionName:    "BIOLOGY",
//		TotalQuestions: 1,
//	}
//
//	sectionB := Sections{
//		SectionName:    "CHEMISTRY",
//		TotalQuestions: 1,
//	}
//
//	sectionC := Sections{
//		SectionName:    "PHYSICS",
//		TotalQuestions: 1,
//	}
//	var sectionList []Sections
//	sectionList = append(sectionList, sectionA)
//	sectionList = append(sectionList, sectionB)
//	sectionList = append(sectionList, sectionC)
//
//	return sectionList
//}
