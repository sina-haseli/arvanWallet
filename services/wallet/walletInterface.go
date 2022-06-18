package wallet

import "context"

type Wallet interface {
	GetBalance(ctx context.Context, userID string) (int, error)
	Increase(ctx context.Context, userID string, amount int, description string) error
}
