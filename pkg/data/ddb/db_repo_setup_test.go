package ddb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
)

var ddb *dynamodb.DynamoDB

func TestMain(m *testing.M) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "amazon/dynamodb-local:latest",
		Cmd:          []string{"-jar", "DynamoDBLocal.jar", "-inMemory"},
		ExposedPorts: []string{"8000/tcp"},
		WaitingFor:   wait.ForListeningPort("8000"),
	}

	dynamodbContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to create container: %s", err))
	}

	defer func(dynamodbContainer testcontainers.Container, ctx context.Context) {
		err := dynamodbContainer.Terminate(ctx)
		if err != nil {
			log.Fatalf("Failed to terminate container: %s", err)
		}
	}(dynamodbContainer, ctx)

	ip, err := dynamodbContainer.Host(ctx)

	if err != nil {
		panic(err)
	}

	port, err := dynamodbContainer.MappedPort(ctx, "8000")

	if err != nil {
		panic(err)
	}

	log.Infof("DynamoDB running on %s:%s", ip, port)

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("test", "test", ""),
		Endpoint:    aws.String(fmt.Sprintf("http://%s:%s", ip, port)),
		Region:      aws.String("ap-south-1"),
	}))

	ddb = dynamodb.New(sess)

	//createTestTable()
	createStudentTestActionTable()

	os.Exit(m.Run())
}
