package ddb

import (
	"context"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
)

type StudentTestActionRepository interface {
	BaseRepository[entity.StudentTestActionEntity, string, uint]
	FindAllByTestIDStudentID(ctx context.Context, testID string, studentID string) ([]*entity.StudentTestActionEntity, error)
	FindAllByTestIDStudentIDNamespace(_ context.Context, testID, studentID, namespace string) ([]*entity.StudentTestActionEntity, error)
}
