package redis

import (
	"fmt"
	"wallet/repositories"
)

type redisConsumer struct {
	repository *repositories.Repository
}

func NewRedisConsumer(repository *repositories.Repository) *redisConsumer {
	return &redisConsumer{
		repository: repository,
	}
}

func (r *redisConsumer) Consume(messages chan<- string, channelName string) {
	for {
		ch, err := r.repository.Redis.Dequeue(channelName)
		if err != nil {
			fmt.Println("could not dequeue from redis", err)
			continue
		}

		messages <- ch
	}
}
