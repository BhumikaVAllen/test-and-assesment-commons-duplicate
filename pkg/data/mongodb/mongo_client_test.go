package mongodb

import (
	"context"
	"fmt"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"testing"
)

// MockClient is a mock implementation of *mongo.Client
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Connect(_ context.Context, _ ...*options.ClientOptions) error {
	args := m.Called()
	return args.Error(0)
}

func TestNewData_Success(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test-creds.json")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()

	// Define the contents of the JSON file
	credsJSON := `{"username":"testuser","password":"testpass"}`

	// Write the JSON contents to the temporary file
	_, err = tmpfile.Write([]byte(credsJSON))
	if err != nil {
		t.Fatal(err)
	}

	mockClient := new(MockClient)

	c := &commons_conf.Data_MongoDb{
		CredFileLocation: tmpfile.Name(),
	}

	// Set up expectations for Connect and Ping methods
	mockClient.On("Connect").Return(mockClient)
	mockClient.On("Ping").Return(mockClient)

	data, err := NewMongoClient(c, log.DefaultLogger)
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestReadCredentials_Success(t *testing.T) {
	// Create a temporary JSON file for testing
	tmpfile, err := ioutil.TempFile("", "test-creds.json")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()

	// Define the contents of the JSON file
	credsJSON := `{"username":"testuser","password":"testpass"}`

	// Write the JSON contents to the temporary file
	_, err = tmpfile.Write([]byte(credsJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Call the function being tested
	credentials := ReadCredentials(tmpfile.Name())
	fmt.Println(credentials)

	// Assert the results
	expectedCreds := &DatabaseCreds{Username: "testuser", Password: "testpass"}
	assert.NotNil(t, credentials)
	assert.Equal(t, expectedCreds, credentials)

}

func TestReadCredentials_ErrorFileRead(t *testing.T) {

	// Call the function being tested with a non-existent file
	credentials := ReadCredentials("nonexistent.json")

	// Assert the results
	assert.Nil(t, credentials)

}

func TestReadCredentials_ErrorUnmarshal(t *testing.T) {
	// Create a temporary JSON file for testing
	tmpfile, err := ioutil.TempFile("", "test-creds.json")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()

	// Define invalid JSON content
	invalidJSON := `invalid_json`

	// Write the invalid JSON content to the temporary file
	_, err = tmpfile.Write([]byte(invalidJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Call the function being tested
	credentials := ReadCredentials(tmpfile.Name())

	// Assert the results
	assert.Nil(t, credentials)

}
