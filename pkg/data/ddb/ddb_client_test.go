package ddb

import (
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDdbClient(t *testing.T) {
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
					DynamoDb: &commons_conf.Data_DynamoDb{
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
			_, _, err := NewDdbClient(tt.args.confData, logger)
			if tt.wantErr {
				assert.Error(t, err) // Check that an error occurred
			} else {
				assert.NoError(t, err) // Check that no error occurred
			}
		})
	}
}
