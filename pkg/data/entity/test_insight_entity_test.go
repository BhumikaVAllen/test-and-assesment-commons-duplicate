package entity

import (
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTestSectionInsight_ToPB(t *testing.T) {
	tests := []struct {
		name     string
		expected pbTypes.TestSectionInsight
		arg      TestSectionInsight
	}{
		{
			name: "Returns updated response from given input",
			expected: pbTypes.TestSectionInsight{
				SectionName:  "PHYSICS SECTION-A",
				AverageMarks: 12.2,
			},
			arg: TestSectionInsight{
				SectionName:  "PHYSICS SECTION-A",
				AverageMarks: 12.2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pbTypes.TestSectionInsight{}
			tt.arg.ToPB(&res)
			assert.Equal(t, tt.expected, res)
		})
	}

}

func TestTestInsightEntity_ToPB(t *testing.T) {
	tests := []struct {
		name     string
		expected pbTypes.TestInsight
		arg      TestInsightEntity
	}{
		{
			name: "Returns updated response from given input",
			expected: pbTypes.TestInsight{
				TestId:       "test_3Nn8AI8JAbzx",
				StudentCount: 2,
				AverageMarks: 23,
				TopperMarks:  23,
				CreatedAt:    1707299305000,
				UpdatedAt:    1707299305000,
				SectionalInsights: []*pbTypes.TestSectionInsight{
					{
						SectionName:  "physics",
						AverageMarks: 20,
					},
					{
						SectionName:  "chemistry",
						AverageMarks: 12,
					},
				},
			},
			arg: TestInsightEntity{
				TestID:       "test_3Nn8AI8JAbzx",
				StudentCount: 2,
				AverageMarks: 23,
				TopperMarks:  23,
				SectionalInsights: []TestSectionInsight{
					{
						SectionName:  "physics",
						AverageMarks: 20,
					},
					{
						SectionName:  "chemistry",
						AverageMarks: 12,
					},
				},
				CreatedAt: time.UnixMilli(1707299305000),
				UpdatedAt: time.UnixMilli(1707299305000),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pbTypes.TestInsight{}
			tt.arg.ToPB(&res)
			assert.Equal(t, tt.expected, res)
		})
	}

}
