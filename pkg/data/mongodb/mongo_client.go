package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"os"
)

type MongoClient struct {
	client *mongo.Client
	dbName string
}

type DatabaseCred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DatabaseCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewMongoClient(c *commons_conf.Data_MongoDb, logger log.Logger) (*MongoClient, error) {
	l := log.NewHelper(logger)

	credentials := ReadCredentials(c.CredFileLocation)
	if credentials == nil {
		err := errors.New(401, "connection error", "Error fetching mongo credentials")
		panic(err)
	}
	log.Info("Started connecting with mongo server")
	credential := options.Credential{
		Username: credentials.Username,
		Password: credentials.Password,
	}
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().SetMonitor(otelmongo.NewMonitor()).ApplyURI(c.Addr).SetAuth(credential))
	if err != nil {
		l.Errorf("Error while connecting to mongo ddb")
		return nil, errors.New(500, "connection error", fmt.Sprintf("Error on connecting Mongo server: %v", err))
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Errorf("Error while connecting to mongo db")
		return nil, fmt.Errorf("connection error: Error pinging Mongo server: %v", err)
	}
	l.Info("Completed  connecting with mongo server successfully")
	return &MongoClient{client: client, dbName: c.DbName}, nil
}

func ReadCredentials(fileName string) *DatabaseCreds {
	// read our opened jsonFile as a byte array.
	byteValue, err := os.ReadFile(fileName)

	if err != nil {
		log.Errorf("Error while reading the file for mongo creds")
		return nil
	}

	// we initialize our Users array
	var creds DatabaseCreds

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &creds)

	if err != nil {
		log.Errorf("Error while marshaling  the file for mongo creds")
		return nil
	}

	return &creds
}
