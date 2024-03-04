package entity

import (
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"time"
)

// StudentTestInsightEntity for storing test X student level insights
type StudentTestInsightEntity struct {
	StudentID         string           `bson:"studentId"`
	TestID            string           `bson:"testId"`
	CreatedAt         time.Time        `bson:"createdAt,unixtime"`
	UpdatedAt         time.Time        `bson:"updatedAt,unixtime"`
	SectionalInsights []SectionInsight `bson:"sectionalInsights"`
	PeerInsights      []PeerInsight    `bson:"peerInsights"`
	OverallInsight    OverallInsight   `bson:"overallInsight"`
}

type SectionInsight struct {
	SectionName        string  `bson:"sectionName"`
	Attempted          int64   `bson:"attempted"`
	TotalQuestions     int64   `bson:"totalQuestions"`
	CorrectQuestions   int64   `bson:"correctQuestions"`
	InCorrectQuestions int64   `bson:"inCorrectQuestions"`
	MarksScored        float32 `bson:"marksScored"`
	MaximumMarks       float32 `bson:"maximumMarks"`
	Rank               int64   `bson:"rank"`
	AttemptPercentage  float32 `bson:"attemptPercentage"`
	Accuracy           float32 `bson:"accuracy"`
	Percentage         float32 `bson:"percentage"`
	Percentile         float32 `bson:"percentile"`
}

type PeerInsight struct {
	SectionName                       string  `bson:"sectionName"`
	PercentileSectionMarksAvg         float32 `bson:"percentileSectionMarksAvg"`
	QuestionsForNextPercentileSection int64   `bson:"questionsForNextPercentileSection"`
}

type OverallInsight struct {
	Attempted          int64   `bson:"attempted"`
	TotalQuestions     int64   `bson:"totalQuestions"`
	CorrectQuestions   int64   `bson:"correctQuestions"`
	InCorrectQuestions int64   `bson:"inCorrectQuestions"`
	MarksScored        float32 `bson:"marksScored"`
	MaximumMarks       float32 `bson:"maximumMarks"`
	Rank               int64   `bson:"rank"`
	AttemptPercentage  float32 `bson:"attemptPercentage"`
	Accuracy           float32 `bson:"accuracy"`
	Percentage         float32 `bson:"percentage"`
	Percentile         float32 `bson:"percentile"`
}

func (si *SectionInsight) ToPB(insight *pbTypes.SectionInsight) {
	insight.SectionName = si.SectionName
	insight.Attempted = si.Attempted
	insight.TotalQuestions = si.TotalQuestions
	insight.CorrectQuestions = si.CorrectQuestions
	insight.IncorrectQuestions = si.InCorrectQuestions
	insight.MarksScored = si.MarksScored
	insight.MaximumMarks = si.MaximumMarks
	insight.Rank = si.Rank
	insight.AttemptPercentage = si.AttemptPercentage
	insight.Accuracy = si.Accuracy
	insight.Percentage = si.Percentage
	insight.Percentile = si.Percentile

}

func (pi *PeerInsight) ToPB(insight *pbTypes.PeerInsight) {
	insight.SectionName = pi.SectionName
	insight.PercentileSectionMarksAvg = pi.PercentileSectionMarksAvg
	insight.QuestionsForNextPercentileSection = pi.QuestionsForNextPercentileSection
}

func (oi *OverallInsight) ToPB(insight *pbTypes.OverallInsight) {
	insight.Attempted = oi.Attempted
	insight.TotalQuestions = oi.TotalQuestions
	insight.CorrectQuestions = oi.CorrectQuestions
	insight.IncorrectQuestions = oi.InCorrectQuestions
	insight.MarksScored = oi.MarksScored
	insight.MaximumMarks = oi.MaximumMarks
	insight.Rank = oi.Rank
	insight.AttemptPercentage = oi.AttemptPercentage
	insight.Accuracy = oi.Accuracy
	insight.Percentage = oi.Percentage
	insight.Percentile = oi.Percentile
}

func (entity *StudentTestInsightEntity) ToPB(info *pbTypes.StudentTestInsight) {
	info.TestId = entity.TestID
	info.StudentId = entity.StudentID
	info.CreatedAt = entity.CreatedAt.UnixMilli()
	info.UpdatedAt = entity.UpdatedAt.UnixMilli()
	for _, sectionInsight := range entity.SectionalInsights {
		pbSectionInsight := new(pbTypes.SectionInsight)
		sectionInsight.ToPB(pbSectionInsight)
		info.SectionalInsights = append(info.SectionalInsights, pbSectionInsight)
	}
	for _, peerInsights := range entity.PeerInsights {
		pbPeerInsights := new(pbTypes.PeerInsight)
		peerInsights.ToPB(pbPeerInsights)
		info.PeerInsights = append(info.PeerInsights, pbPeerInsights)
	}
	overallInsight := new(pbTypes.OverallInsight)
	entity.OverallInsight.ToPB(overallInsight)
	info.OverallInsight = overallInsight
}
