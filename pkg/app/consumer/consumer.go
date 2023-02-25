package consumer

import "github.com/go-redis/redis"

type (
	Consumer interface {
		Consume() (string, string, error)
	}

	RedisConsumer struct {
		cli   *redis.PubSub
		topic string
	}
)

const (
	RATE = "RATE"
	REF  = "REF"
)

func InitRedisConsumer(cli *redis.PubSub, topic string) *RedisConsumer {
	return &RedisConsumer{cli, topic}
}

func (f *RedisConsumer) Consume() (string, string, error) {
	ch, err := f.cli.ReceiveMessage()
	if err != nil {
		return "", "", err
	}

	switch f.topic {
	case RATE:
	case REF:
	default:
	}

	return ch.Payload, f.topic, nil
}
