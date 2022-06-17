package repositories

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *redisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Dequeue(channelName string) (string, error) {
	res, err := r.client.BLPop(0*time.Second, channelName).Result()
	if err != nil {
		return "", err
	}

	return res[1], nil
}

func (r *redisRepository) Enqueue(message []byte, channelName string) error {
	_, err := r.client.RPush(channelName, message).Result()
	return err
}

func (r *redisRepository) Increase(key string) error {
	_, err := r.client.Incr(key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) SetValue(key string, value interface{}) error {
	_, err := r.client.Set(key, value, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) GetValue(key string) (string, error) {
	v, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return v, nil
}
