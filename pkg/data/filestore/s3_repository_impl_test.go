package filestore

import (
	"context"
	"fmt"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/filestore/request"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setupS3Test() (context.Context, S3Repository) {
	ctx := context.Background()
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)
	d := &S3Client{
		S3Uploader:   s3UploaderTest,
		S3Downloader: s3DownloaderTest,
		s3Session:    sess,
	}
	repo := NewS3RepositoryImpl(d, log.NewStdLogger(os.Stdout))
	return ctx, repo
}

func TestNewS3RepositoryImpl_UploadFile(t *testing.T) {
	_, repo := setupS3Test()
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s/%s", "test-and-assessment-question-papers", "testing", "key")
	type args struct {
		qPId string
		qP   []byte
		key  string
	}
	tests := []struct {
		name          string
		args          args
		want          *string
		wantErr       bool
		errorResponse *errors.Error
		setup         func()
		teardown      func()
	}{
		{
			name: "s3 upload failed",
			args: args{
				qPId: "testing",
				qP:   []byte("mohit"),
				key:  "key",
			},
			want:          nil,
			wantErr:       true,
			errorResponse: errors.New(500, "", ""),
			setup:         func() {},
			teardown:      func() {},
		},
		{
			name: "s3 upload test",
			args: args{
				qPId: "test-and-assessment-question-papers",
				qP:   []byte("mohit"),
				key:  "key",
			},
			want:          &objectURL,
			wantErr:       false,
			errorResponse: nil,
			setup:         func() {},
			teardown:      func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := repo.UploadFile(&request.S3Request{
				Bucket:   tt.args.qPId,
				Key:      tt.args.key,
				ByteData: tt.args.qP,
			})
			if tt.wantErr {
				assert.Error(t, err) // Check that an error occurred
				assert.Equal(t, tt.wantErr, err != nil, "Upload Failed() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != nil {
				assert.NotEmpty(t, got)
			}
			tt.teardown()
		})
	}
}

func TestNewS3RepositoryImpl_DownloadFile(t *testing.T) {
	_, repo := setupS3Test()
	objectURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s/%s", "test-and-assessment-question-papers", "testing", "key")
	type args struct {
		qPId string
		key  string
	}
	tests := []struct {
		name          string
		args          args
		want          *string
		wantErr       bool
		errorResponse *errors.Error
		setup         func()
		teardown      func()
	}{
		{
			name: "s3 download failed",
			args: args{
				qPId: "testing",
				key:  "key",
			},
			want:          nil,
			wantErr:       true,
			errorResponse: errors.New(500, "", ""),
			setup:         func() {},
			teardown:      func() {},
		},
		{
			name: "s3 download success",
			args: args{
				qPId: "test-and-assessment-question-papers",
				key:  "key",
			},
			want:          &objectURL,
			wantErr:       false,
			errorResponse: nil,
			setup:         func() {},
			teardown:      func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := repo.DownloadFile(&request.S3Request{
				Bucket: tt.args.qPId,
				Key:    tt.args.key,
			})
			if tt.wantErr {
				assert.Error(t, err) // Check that an error occurred
				assert.Equal(t, tt.wantErr, err != nil, "Download Failed() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != nil {
				assert.NotEmpty(t, got)
			}
			tt.teardown()
		})
	}
}
