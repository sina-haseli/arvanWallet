package consumer

type Consumer interface {
	Consume(messages chan<- string, channelName string)
}
