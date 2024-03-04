package data_filters

import (
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"time"
)

type StudentTestFilter struct {
	TestID       string
	TestCategory string
	Status       string
	StudentID    string
}

type TestFilter struct {
	ID            string
	Category      string
	Status        string
	Stream        string
	Class         string
	PageSize      int64
	PageNo        int64
	SortField     string
	OrderBy       string
	FromDate      time.Time
	ToDate        time.Time
	Centers       string
	TestType      string
	UserId        string
	SearchKeyword string
}

type TestFilterResponse struct {
	TotalResults int64
	PageNo       int64
	PageSize     int64
	TestInfo     []*entity.TestInfoEntity
}
type SearchFilter struct {
	SearchKeyword string
	PageSize      int64
	PageNo        int64
	SortField     string
	OrderBy       string
}
