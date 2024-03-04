package publisher

import (
	ps "github.com/Allen-Career-Institute/go-kratos-commons/pubsub/v1"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/publisher/event"
)

type Publisher interface {
	PublishEvent(message event.Event) error
}

type PubSubPublisher interface {
	ps.Publisher
}
