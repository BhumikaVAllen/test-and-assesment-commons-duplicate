package publisher

import (
	"context"
	"fmt"
	pbTypes "github.com/Allen-Career-Institute/common-protos/test_and_assessment_commons/v1/types"
	ps "github.com/Allen-Career-Institute/go-kratos-commons/pubsub/v1"
	"github.com/Allen-Career-Institute/go-kratos-commons/pubsub/v1/factory"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/publisher/event"
	"github.com/go-kratos/kratos/v2/log"
)

type TestAndAssessmentPublisher struct {
	log       *log.Helper
	publisher PubSubPublisher
	topicName string
}

func NewTestAndAssessmentPublisherClient(publisherConfig *commons_conf.Publisher, logger log.Logger) (Publisher, error) {
	pubSub := &factory.PubSubFactoryHandler{}
	var err error
	ctx := context.Background()
	// create a singleton instance for AWSPublisherManager
	publisher, err := pubSub.GetPublisher(ctx, ps.AWS_SNS_SQS)
	if err != nil {
		log.Errorf("Error in Getting IProducer, %v", err)
		return nil, err
	}

	topicConfig := publisherConfig.TestAndAssessment
	config := map[string]interface{}{
		"AWSRegion": topicConfig.AwsRegion,
		"topicID":   topicConfig.TopicName,
		//optional custom configs
		"options": map[string]interface{}{
			"BatcherOptions": map[string]interface{}{
				"MinBatchSize": topicConfig.MinBatchSize,
				"MaxBatchSize": topicConfig.MaxBatchSize,
				"MaxHandlers":  topicConfig.MaxHandlers,
			},
		},
	}
	err = publisher.SetupPublisher(ctx, config)
	if err != nil {
		log.Errorf("Error in SetupPublisher for topic : %s, error : %v", topicConfig.TopicName, err)
		return nil, err
	}
	return &TestAndAssessmentPublisher{
		log:       log.NewHelper(logger),
		publisher: publisher,
		topicName: publisherConfig.TestAndAssessment.GetTopicName(),
	}, nil
}

// PublishEvent GroupId has to be decided for student + test combination
func (p *TestAndAssessmentPublisher) PublishEvent(message event.Event) error {
	if _, ok := message.(event.TestAndAssessmentEvent); !ok {
		p.log.Errorf("Event message is not of type TestAndAssessmentEvent")
		return pbTypes.ErrorEventPublishFailed("Event message is not of type TestAndAssessmentEvent")
	}

	strMsg, err := message.GetStringMessage()
	if err != nil {
		errorMsg := fmt.Sprintf("Error while marshalling: %v", err)
		p.log.Errorf(errorMsg)
		return pbTypes.ErrorEventPublishFailed(errorMsg)
	}
	metaData := map[string]string{
		ps.MESSAGE_GROUP_ID: message.GetEntityID(),
		//ps.DE_DUPLICATION:   uuid.New().String(),
	}

	return p.publishToTopic(p.topicName, strMsg, metaData)
}

func (p *TestAndAssessmentPublisher) publishToTopic(topicName string, message string, metaData map[string]string) error {
	ctx := context.Background()
	err := p.publisher.Publish(ctx, topicName, message, metaData)
	if err != nil {
		errorMsg := fmt.Sprintf("Error in publishing to SNS topic: %s, error: %v", topicName, err)
		p.log.Errorf(errorMsg)
		return pbTypes.ErrorEventPublishFailed(errorMsg)
	}
	p.log.Infof("Sent Message to AWS : %s ", message)
	return nil
}
