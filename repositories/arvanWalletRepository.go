package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"wallet/models"
)

const (
	getBalanceQuery        = "SELECT current_balance FROM user_transactions WHERE user_id = $1 ORDER BY id DESC LIMIT 1"
	insertTransactionQuery = "INSERT INTO user_transactions (user_id, amount, current_balance, description) VALUES ($1, $2, $3, $4)"
)

var NoRowsAffectedError = errors.New("zero row affected")

type r1WalletRepository struct {
	db *sql.DB
}

func NewR1WalletRepository(db *sql.DB) *r1WalletRepository {
	return &r1WalletRepository{
		db: db,
	}
}

func (r1 *r1WalletRepository) GetBalanceByUserID(ID int) (int, error) {
	r, err := r1.db.Query(getBalanceQuery, ID)
	if err != nil {
		return 0, err
	}

	defer func(r *sql.Rows) {
		err := r.Close()
		if err != nil {
			fmt.Println("could not close rows")
		}
	}(r)

	if r.Next() == false {
		return 0, nil
	}

	var b int
	err = r.Scan(&b)
	if err != nil {
		return 0, err
	}

	return b, nil
}

func (r1 *r1WalletRepository) InsertTransaction(ut models.UserTransactionModel) error {
	r, err := r1.db.Exec(insertTransactionQuery, ut.UserID, ut.Amount, ut.CurrentBalance, ut.Description)
	if err != nil {
		return err
	}

	ra, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if ra == 0 {
		return NoRowsAffectedError
	}

	return nil
}
