package consumer

import "context"

type Consumer interface {
	Consume(ctx context.Context, channelName string, bufSize int) chan string
}
