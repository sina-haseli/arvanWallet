package models

type UserTransactionModel struct {
	UserID         int
	Amount         int
	CurrentBalance int
	Description    string
}
