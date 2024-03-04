package filestore

import (
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/filestore/request"
)

type S3Repository interface {
	IFilestoreRepository[request.S3Request]
}
