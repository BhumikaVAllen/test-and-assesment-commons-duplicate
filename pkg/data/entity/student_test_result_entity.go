package entity

import "time"

type StudentTestResultEntity struct {
	StudentID          string             `bson:"studentId"`
	TestID             string             `bson:"testId"`
	ProvisionalResult  ProvisionalResult  `bson:"provisionalResult"`
	ConsolidatedResult ConsolidatedResult `bson:"consolidatedResult"`
	CreatedAt          time.Time          `bson:"createdAt,unixtime"`
	UpdatedAt          time.Time          `bson:"updatedAt,unixtime"`
}

type ProvisionalResult struct {
	SectionalResult []SectionalResult `bson:"sectionalResult"`
}

type ConsolidatedResult struct {
	MarksScored  float32 `bson:"marksScored"`
	MaximumMarks float32 `bson:"maximumMarks"`
	Rank         int     `bson:"rank"`
}

type SectionalResult struct {
	SectionName        string  `bson:"sectionName"`
	Attempted          int     `bson:"attempted"`
	TotalQuestions     int     `bson:"totalQuestions"`
	CorrectQuestions   int     `bson:"correctQuestions"`
	InCorrectQuestions int     `bson:"inCorrectQuestions"`
	MarksScored        float32 `bson:"marksScored"`
	MaximumMarks       float32 `bson:"maximumMarks"`
	Rank               int     `bson:"rank"`
}
