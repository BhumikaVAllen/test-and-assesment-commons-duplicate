package ddb

import (
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-kratos/kratos/v2/log"
)

type Client struct {
	DynamoDB *dynamodb.DynamoDB
}

func NewDdbClient(c *commons_conf.Data, logger log.Logger) (*Client, func(), error) {
	l := log.NewHelper(logger)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(c.DynamoDb.GetRegion()),
	})
	if err != nil {
		l.Errorf("Unable to connect to ddb : %v", err)
		return nil, nil, err
	}
	dynamoDB := dynamodb.New(sess)
	d := &Client{
		DynamoDB: dynamoDB,
	}
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return d, cleanup, nil
}
