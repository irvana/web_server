package publisher

import (
	"web_server/pkg/app/consumer"
	"web_server/pkg/wamp"

	"github.com/gammazero/nexus/v3/client"
)

type (
	Publisher interface {
		Publish()
	}

	CustomPublisher struct {
		consumer consumer.Consumer
		wampCli  *client.Client
	}
)

func InitPublisher(consumer consumer.Consumer, wamp *wamp.Wamp) Publisher {
	return &CustomPublisher{consumer, wamp.Client}
}

func (c *CustomPublisher) Publish() {
	for {

		msg, _, err := c.consumer.Consume()
		if err != nil {
			// return err
			return
		}

		if msg != "" {
			// json.Unmarshal([]byte(msg))
			// c.wampCli.Publish(topic)

		}
		// return nil
	}
}
