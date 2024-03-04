package mongodb

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
)

var mongoClient *mongo.Client

func TestMain(m *testing.M) {

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		panic(err)
	}
	mappedPort, err := container.MappedPort(ctx, "27017")
	if err != nil {
		panic(err)
	}

	uri := fmt.Sprintf("mongodb://%s:%s", hostIP, mappedPort.Port())
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
