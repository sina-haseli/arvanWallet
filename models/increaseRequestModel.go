package models

type IncreaseRequestModel struct {
	UserID string `json:"user_id"`
	Amount int    `json:"amount"`
}
