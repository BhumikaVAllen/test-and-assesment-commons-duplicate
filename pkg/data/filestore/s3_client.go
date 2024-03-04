package filestore

import (
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-kratos/kratos/v2/log"
)

type S3Client struct {
	S3Uploader   *s3manager.Uploader
	S3Downloader *s3manager.Downloader
	s3Session    *session.Session
}

func NewS3Client(c *commons_conf.Data_S3, logger log.Logger) (*S3Client, func(), error) {
	l := log.NewHelper(logger)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(c.GetRegion()),
	})
	if err != nil {
		l.Errorf("Unable to connect to db : %v", err)
		return nil, nil, err
	}
	s3Uploader := s3manager.NewUploader(sess)
	s3Downloader := s3manager.NewDownloader(sess)
	d := &S3Client{
		S3Uploader:   s3Uploader,
		S3Downloader: s3Downloader,
		s3Session:    sess,
	}
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return d, cleanup, nil
}
