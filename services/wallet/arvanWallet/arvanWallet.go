package arvanWallet

import (
	"context"
	"fmt"
	"wallet/models"
	"wallet/repositories"
)

type R1Wallet struct {
	repository *repositories.Repository
}

func NewR1Wallet(repository *repositories.Repository) *R1Wallet {
	return &R1Wallet{
		repository: repository,
	}
}

func (r *R1Wallet) GetBalance(ctx context.Context, userID string) (int, error) {
	b, err := r.repository.Wallet.GetBalanceByUserID(ctx, userID)
	if err != nil {
		return 0, err
	}

	return b, nil
}

func (r *R1Wallet) Increase(ctx context.Context, userID string, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("negative amount")
	}

	err := r.repository.Wallet.InsertTransaction(ctx, models.UserTransactionModel{
		UserID:      userID,
		Amount:      amount,
		Description: description,
	})

	return err
}
