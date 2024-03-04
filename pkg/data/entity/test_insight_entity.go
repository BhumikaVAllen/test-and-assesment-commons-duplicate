package entity

import (
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"time"
)

// TestInsightEntity for storing test level insights.
type TestInsightEntity struct {
	TestID            string               `bson:"testId"`
	StudentCount      int64                `bson:"studentCount"`
	AverageMarks      float32              `bson:"averageMarks"`
	TopperMarks       float32              `bson:"topperMarks"`
	SectionalInsights []TestSectionInsight `bson:"sectionalInsights"`
	CreatedAt         time.Time            `bson:"createdAt"`
	UpdatedAt         time.Time            `bson:"updatedAt"`
}

type TestSectionInsight struct {
	SectionName  string  `bson:"sectionName"`
	AverageMarks float32 `bson:"averageMarks"`
}

func (tsi *TestSectionInsight) ToPB(insight *pbTypes.TestSectionInsight) {
	insight.SectionName = tsi.SectionName
	insight.AverageMarks = tsi.AverageMarks
}

func (entity *TestInsightEntity) ToPB(testInsight *pbTypes.TestInsight) {
	testInsight.TestId = entity.TestID
	testInsight.StudentCount = entity.StudentCount
	testInsight.AverageMarks = entity.AverageMarks
	testInsight.TopperMarks = entity.TopperMarks
	testInsight.CreatedAt = entity.CreatedAt.UnixMilli()
	testInsight.UpdatedAt = entity.UpdatedAt.UnixMilli()
	for _, sectionInsight := range entity.SectionalInsights {
		pbTestSectionInsight := new(pbTypes.TestSectionInsight)
		sectionInsight.ToPB(pbTestSectionInsight)
		testInsight.SectionalInsights = append(testInsight.SectionalInsights, pbTestSectionInsight)
	}
}
