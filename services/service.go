package services

import (
	"wallet/repositories"
	"wallet/services/consumer"
	"wallet/services/consumer/redis"
	"wallet/services/wallet"
	"wallet/services/wallet/arvanWallet"
)

type Services struct {
	Consumer consumer.Consumer
	Wallet   wallet.Wallet
}

func NewServices(repository *repositories.Repository) *Services {
	return &Services{
		Consumer: redis.NewRedisConsumer(repository),
		Wallet:   arvanWallet.NewR1Wallet(repository),
	}
}
