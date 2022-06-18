package redis

import (
	"context"
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

func (r *redisConsumer) Consume(ctx context.Context, channelName string, bufSize int) chan string {
	eventCh := make(chan string, bufSize)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(eventCh)
				return
			default:
				ch, err := r.repository.Redis.Dequeue(channelName)
				if err != nil {
					fmt.Println("could not dequeue from redis", err)
					continue
				}
				eventCh <- ch
			}
		}
	}()

	return eventCh
}
