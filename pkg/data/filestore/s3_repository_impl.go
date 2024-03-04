package filestore

import (
	"bytes"
	"fmt"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/filestore/request"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type s3RepositoryImpl struct {
	s3Uploader   *s3manager.Uploader
	s3Downloader *s3manager.Downloader
	log          *log.Helper
	s3Client     *s3.S3
}

func NewS3RepositoryImpl(data *S3Client, logger log.Logger) S3Repository {
	return &s3RepositoryImpl{
		s3Uploader:   data.S3Uploader,
		s3Downloader: data.S3Downloader,
		log:          log.NewHelper(logger),
		s3Client:     s3.New(data.s3Session),
	}
}

func (s3RepositoryImpl *s3RepositoryImpl) UploadFile(s3Request *request.S3Request) (*string, error) {
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(s3Request.Bucket),
		Key:    aws.String(s3Request.Key),
		Body:   bytes.NewReader(s3Request.ByteData),
	}
	uploadOutput, err := s3RepositoryImpl.s3Uploader.Upload(upParams)
	if err != nil {
		errorMsg := fmt.Sprintf("Error occurred when uploading to S3 : %v", err)
		s3RepositoryImpl.log.Errorf(errorMsg)
		return nil, pbTypes.ErrorFileStoreUploadFailed(errorMsg)
	}
	return &uploadOutput.Location, nil
}

func (s3RepositoryImpl *s3RepositoryImpl) DownloadFile(s3Request *request.S3Request) ([]byte, error) {
	dataBuff := &aws.WriteAtBuffer{}
	_, err := s3RepositoryImpl.s3Downloader.Download(dataBuff, &s3.GetObjectInput{
		Bucket: aws.String(s3Request.Bucket),
		Key:    aws.String(s3Request.Key),
	})
	if err != nil {
		errorMsg := fmt.Sprintf("Error occurred when downloading from S3 : %v", err)
		s3RepositoryImpl.log.Errorf(errorMsg)
		return nil, pbTypes.ErrorFileStoreDownloadFailed(errorMsg)
	}
	return dataBuff.Bytes(), nil
}
func (s3RepositoryImpl *s3RepositoryImpl) GeneratePreSignedURl(s3Request *request.S3Request) (*string, error) {

	r, _ := s3RepositoryImpl.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Request.Bucket),
		Key:    aws.String(s3Request.Key),
	})
	urlStr, err := r.Presign(15 * time.Minute)
	if err != nil {
		errorMsg := fmt.Sprintf("error while getting presigned url  for test solution : %v", err)
		s3RepositoryImpl.log.Errorf(errorMsg)
		return nil, err
	}
	return &urlStr, nil
}
