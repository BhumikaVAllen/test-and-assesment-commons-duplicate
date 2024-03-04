package entity

import (
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStudentTestInsightEntity_ToPB(t *testing.T) {
	tests := []struct {
		name     string
		expected pbTypes.StudentTestInsight
		arg      StudentTestInsightEntity
	}{
		{
			name: "Returns Proto response for complete response",
			expected: pbTypes.StudentTestInsight{
				TestId:    "test_hoswSlCRlH2s",
				StudentId: "5RBeTjJkErRVvVdqSeuMt",
				CreatedAt: 1707299958000,
				UpdatedAt: 1707299958000,
				SectionalInsights: []*pbTypes.SectionInsight{
					{
						SectionName:        "PHYSICS_SECTION-(A)",
						Attempted:          31,
						TotalQuestions:     35,
						CorrectQuestions:   22,
						IncorrectQuestions: 9,
						MarksScored:        14,
						MaximumMarks:       140,
						Rank:               2,
						AttemptPercentage:  88.57142639160156,
						Accuracy:           29.032258987426758,
						Percentage:         10,
						Percentile:         80,
					},
				},
				PeerInsights: []*pbTypes.PeerInsight{
					{
						SectionName:                       "PHYSICS_SECTION-(A)",
						PercentileSectionMarksAvg:         12,
						QuestionsForNextPercentileSection: 21,
					},
				},
				OverallInsight: &pbTypes.OverallInsight{
					Attempted:          143,
					TotalQuestions:     200,
					CorrectQuestions:   37,
					IncorrectQuestions: 106,
					MarksScored:        42,
					MaximumMarks:       800,
					Rank:               2,
					AttemptPercentage:  71.5,
					Accuracy:           25.874126434326172,
					Percentage:         5.25,
					Percentile:         100,
				},
			},
			arg: StudentTestInsightEntity{
				StudentID: "5RBeTjJkErRVvVdqSeuMt",
				TestID:    "test_hoswSlCRlH2s",
				CreatedAt: time.UnixMilli(1707299958000),
				UpdatedAt: time.UnixMilli(1707299958000),
				SectionalInsights: []SectionInsight{
					{
						SectionName:        "PHYSICS_SECTION-(A)",
						Attempted:          31,
						TotalQuestions:     35,
						CorrectQuestions:   22,
						InCorrectQuestions: 9,
						MarksScored:        14,
						MaximumMarks:       140,
						Rank:               2,
						AttemptPercentage:  88.57142639160156,
						Accuracy:           29.032258987426758,
						Percentage:         10,
						Percentile:         80,
					},
				},
				PeerInsights: []PeerInsight{
					{
						SectionName:                       "PHYSICS_SECTION-(A)",
						PercentileSectionMarksAvg:         12,
						QuestionsForNextPercentileSection: 21,
					},
				},
				OverallInsight: OverallInsight{
					Attempted:          143,
					TotalQuestions:     200,
					CorrectQuestions:   37,
					InCorrectQuestions: 106,
					MarksScored:        42,
					MaximumMarks:       800,
					Rank:               2,
					AttemptPercentage:  71.5,
					Accuracy:           25.874126434326172,
					Percentage:         5.25,
					Percentile:         100,
				},
			},
		},
		{
			name: "Returns Proto response with null sectional and peer insight",
			expected: pbTypes.StudentTestInsight{
				TestId:            "test_hoswSlCRlH2s",
				StudentId:         "5RBeTjJkErRVvVdqSeuMt",
				CreatedAt:         1707299958000,
				UpdatedAt:         1707299958000,
				SectionalInsights: nil,
				PeerInsights:      nil,
				OverallInsight: &pbTypes.OverallInsight{
					Attempted:          143,
					TotalQuestions:     200,
					CorrectQuestions:   37,
					IncorrectQuestions: 106,
					MarksScored:        42,
					MaximumMarks:       800,
					Rank:               2,
					AttemptPercentage:  71.5,
					Accuracy:           25.874126434326172,
					Percentage:         5.25,
					Percentile:         100,
				},
			},
			arg: StudentTestInsightEntity{
				StudentID:         "5RBeTjJkErRVvVdqSeuMt",
				TestID:            "test_hoswSlCRlH2s",
				CreatedAt:         time.UnixMilli(1707299958000),
				UpdatedAt:         time.UnixMilli(1707299958000),
				SectionalInsights: nil,
				PeerInsights:      nil,
				OverallInsight: OverallInsight{
					Attempted:          143,
					TotalQuestions:     200,
					CorrectQuestions:   37,
					InCorrectQuestions: 106,
					MarksScored:        42,
					MaximumMarks:       800,
					Rank:               2,
					AttemptPercentage:  71.5,
					Accuracy:           25.874126434326172,
					Percentage:         5.25,
					Percentile:         100,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pbTypes.StudentTestInsight{}
			tt.arg.ToPB(&res)
			assert.Equal(t, tt.expected, res)
		})
	}

}
