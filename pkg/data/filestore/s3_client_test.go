package filestore

import (
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewS3Client(t *testing.T) {
	logger := log.DefaultLogger
	type args struct {
		confData *commons_conf.Data
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		errResponse error
	}{
		{
			name: "New data successful",
			args: args{
				confData: &commons_conf.Data{
					S3: &commons_conf.Data_S3{
						Region: "ap-south-1",
					},
				},
			},
			wantErr:     false,
			errResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := NewS3Client(tt.args.confData.S3, logger)
			if tt.wantErr {
				assert.Error(t, err) // Check that an error occurred
			} else {
				assert.NoError(t, err) // Check that no error occurred
			}
		})
	}
}
