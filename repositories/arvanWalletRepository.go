package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"wallet/models"
)

const (
	getBalanceQuery        = "SELECT current_balance FROM user_balance WHERE user_id = $1 LIMIT 1"
	insertBalanceQuery     = "INSERT INTO user_balance (user_id, current_balance) VALUES ($1, $2)"
	updateBalanceQuery     = "UPDATE user_balance SET current_balance = $1 WHERE user_id = $2"
	insertTransactionQuery = "INSERT INTO user_transactions (user_id, amount, description) VALUES ($1, $2, $3)"
)

var NoRowsAffectedError = errors.New("zero row affected")
var CouldNotFindUserBalance = errors.New("could not find user balance")

type r1WalletRepository struct {
	db *sql.DB
}

func NewR1WalletRepository(db *sql.DB) *r1WalletRepository {
	return &r1WalletRepository{
		db: db,
	}
}

func (r1 *r1WalletRepository) getBalanceByUserID(ctx context.Context, tx *sql.Tx, ID string) (int, error) {
	r, err := tx.QueryContext(ctx, getBalanceQuery, ID)
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
		return 0, CouldNotFindUserBalance
	}

	var b int
	err = r.Scan(&b)
	if err != nil {
		return 0, err
	}

	return b, nil
}

func (r1 *r1WalletRepository) GetBalanceByUserID(ctx context.Context, ID string) (int, error) {
	r, err := r1.db.QueryContext(ctx, getBalanceQuery, ID)
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
		return 0, CouldNotFindUserBalance
	}

	var b int
	err = r.Scan(&b)
	if err != nil {
		return 0, err
	}

	return b, nil
}

func (r1 *r1WalletRepository) InsertBalance(ctx context.Context, trx *sql.Tx, userID string, currentBalance int) error {
	r, err := trx.ExecContext(ctx, insertBalanceQuery, userID, currentBalance)
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

func (r1 *r1WalletRepository) UpdateBalance(ctx context.Context, trx *sql.Tx, currentBalance int, userID string) error {
	r, err := trx.ExecContext(ctx, updateBalanceQuery, currentBalance, userID)
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

func (r1 *r1WalletRepository) Insert(ctx context.Context, trx *sql.Tx, ut models.UserTransactionModel) error {
	r, err := trx.ExecContext(ctx, insertTransactionQuery, ut.UserID, ut.Amount, ut.Description)
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

func (r1 *r1WalletRepository) InsertTransaction(ctx context.Context, ut models.UserTransactionModel) error {
	trx, err := r1.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	bu, err := r1.getBalanceByUserID(ctx, trx, ut.UserID)
	if err != nil {
		if errors.Is(err, CouldNotFindUserBalance) {
			if ibErr := r1.InsertBalance(ctx, trx, ut.UserID, ut.Amount); ibErr != nil {
				if txErr := trx.Rollback(); txErr != nil {
					return txErr
				}
				return ibErr
			}
		}

		if txErr := trx.Rollback(); txErr != nil {
			return txErr
		}

		return err
	}

	if ubErr := r1.UpdateBalance(ctx, trx, bu+ut.Amount, ut.UserID); ubErr != nil {
		if txErr := trx.Rollback(); txErr != nil {
			return txErr
		}
		return ubErr
	}

	if insertErr := r1.Insert(ctx, trx, ut); insertErr != nil {
		if txErr := trx.Rollback(); txErr != nil {
			return txErr
		}
		return insertErr
	}

	return trx.Commit()
}
