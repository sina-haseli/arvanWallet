package wallet

type Wallet interface {
	GetBalance(userID int) (int, error)
	Increase(userID int, amount int, description string) error
}
