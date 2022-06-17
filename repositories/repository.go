package repositories

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	"wallet/models"
)

type Wallet interface {
	GetBalanceByUserID(ID int) (int, error)
	InsertTransaction(ut models.UserTransactionModel) error
}

type Redis interface {
	Dequeue(queueName string) (string, error)
	Enqueue(message []byte, queueName string) error
	Increase(key string) error
	SetValue(key string, value interface{}) error
	GetValue(key string) (string, error)
}

type Repository struct {
	Wallet Wallet
	Redis  Redis
}

func NewRepository(db *sql.DB, re *redis.Client) *Repository {
	return &Repository{
		Wallet: NewR1WalletRepository(db),
		Redis:  NewRedisRepository(re),
	}
}
