package wallet

type Wallet interface {
	GetBalance(userID int) (int, error)
	Decrease(userID int, amount int, description string) error
	Increase(userID int, amount int, description string) error
}
