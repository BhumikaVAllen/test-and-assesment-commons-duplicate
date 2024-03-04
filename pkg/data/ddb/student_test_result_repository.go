package ddb

import (
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
)

type StudentTestResultRepository interface {
	BaseRepository[entity.StudentTestResultEntity, string, uint]
}
