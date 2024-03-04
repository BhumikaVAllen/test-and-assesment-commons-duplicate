package publisher

import (
	"fmt"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/mocks"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/data/entity"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/publisher/event"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type TestSuitePublisher struct {
	suite.Suite
	logger    log.Logger
	publisher mocks.PubSubPublisher
}

func TestRunTestSuitePublisher(t *testing.T) {
	suite.Run(t, &TestSuitePublisher{})
}

func (ts *TestSuitePublisher) SetupSuite() {
	logger := log.NewStdLogger(os.Stdout)
	ts.logger = logger
}

func (ts *TestSuitePublisher) BeforeTest(suiteName, testName string) {
	ts.publisher = mocks.PubSubPublisher{}
}

func (ts *TestSuitePublisher) TestNewTestAndAssessmentPublisherClient_Success() {
	conf := createConf()
	publisher, err := NewTestAndAssessmentPublisherClient(conf, ts.logger)
	assert.Nil(ts.T(), err)
	assert.NotNil(ts.T(), publisher)

	taaPublisher := publisher.(*TestAndAssessmentPublisher)
	assert.Equal(ts.T(), taaPublisher.topicName, "test")
	assert.NotNil(ts.T(), taaPublisher.log)
	assert.NotNil(ts.T(), taaPublisher.publisher)
}

func (ts *TestSuitePublisher) TestTestAndAssessmentPublisher_Success() {
	event := createEvent()
	data, err := event.GetStringMessage()
	assert.Nil(ts.T(), err)
	ts.publisher.On("Publish", mock.Anything, "test", data, mock.Anything).Return(nil)
	taaPublisher := &TestAndAssessmentPublisher{
		log:       log.NewHelper(ts.logger),
		publisher: &ts.publisher,
		topicName: "test",
	}
	err = taaPublisher.PublishEvent(event)
	assert.Nil(ts.T(), err)

}

func (ts *TestSuitePublisher) TestTestAndAssessmentPublisher_Failed() {
	event := createEvent()
	data, err := event.GetStringMessage()
	assert.Nil(ts.T(), err)
	ts.publisher.On("Publish", mock.Anything, "test", string(data), mock.Anything).Return(pbTypes.ErrorEventPublishFailed(fmt.Sprintf("Error in publishing to SNS topic: %s, error: %v", "test", err)))

	taaPublisher := &TestAndAssessmentPublisher{
		log:       log.NewHelper(ts.logger),
		publisher: &ts.publisher,
		topicName: "test",
	}
	err = taaPublisher.PublishEvent(event)
	assert.NotNil(ts.T(), err)
	assert.ErrorContains(ts.T(), err, "Error in publishing to SNS topic")
}

func createConf() *commons_conf.Publisher {

	return &commons_conf.Publisher{
		TestAndAssessment: &commons_conf.Publisher_Topic{
			TopicName:    "test",
			AwsRegion:    "ap-south-1",
			MinBatchSize: 1,
			MaxBatchSize: 2,
			MaxHandlers:  2,
		},
	}
}

func createEvent() event.TestAndAssessmentEvent {

	eb := entity.TestInfoEntity{
		Assignment: entity.Assignment{
			Batches: []*entity.Batch{
				{
					BatchCode: "bacth",
					BatchID:   "232",
				}},
			Centers:  nil,
			Students: nil,
		},
		Syllabus: []*entity.Syllabus{{
			GroupNodeID: "aaaaaaa",
			NodeIds:     []string{"cccccc", "ddddd"},
			TaxonomyID:  "yyyyyy",
			Metadata:    &entity.SyllabusMetadata{Notes: "nnnnnnn"},
		}}}

	return event.TestAndAssessmentEvent{
		BaseEvent: event.BaseEvent{
			Event:      "test",
			EntityID:   "1",
			EntityName: "test1",
			Timestamp:  0,
		},
		Data: map[string]interface{}{
			"test_id":    "abc",
			"assignment": eb.Assignment,
			"syllabus":   eb.Syllabus,
		},
	}
}
