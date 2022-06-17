package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"wallet/models"
)

type dbQE interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

const (
	//getBalanceQuery        = "SELECT current_balance FROM user_transactions WHERE user_id = $1 ORDER BY id DESC LIMIT 1"
	getBalanceQuery2       = "SELECT current_balance FROM user_balance WHERE user_id = $1 LIMIT 1"
	insertBalanceQuery     = "INSERT INTO user_balance (user_id, current_balance) VALUES ($1, $2)"
	updateBalanceQuery     = "UPDATE user_balance SET current_balance = $1 WHERE user_id = $2"
	insertTransactionQuery = "INSERT INTO user_transactions (user_id, amount, description) VALUES ($1, $2, $3)"
)

var NoRowsAffectedError = errors.New("zero row affected")
var CouldNotFindUserBalance = errors.New("could not find user balance")

type r1WalletRepository struct {
	db  *sql.DB
	dbq dbQE
	tx  *sql.Tx
}

func NewR1WalletRepository(db *sql.DB) *r1WalletRepository {
	return &r1WalletRepository{
		db:  db,
		dbq: db,
		tx:  nil,
	}
}

func (r1 *r1WalletRepository) GetBalanceByUserID(ID int) (int, error) {
	r, err := r1.dbq.Query(getBalanceQuery2, ID)
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

func (r1 *r1WalletRepository) InsertBalance(userID, currentBalance int) error {
	println(userID, currentBalance)
	r, err := r1.dbq.Exec(insertBalanceQuery, userID, currentBalance)
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

func (r1 *r1WalletRepository) UpdateBalance(currentBalance, userID int) error {
	r, err := r1.dbq.Exec(updateBalanceQuery, currentBalance, userID)
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

func (r1 *r1WalletRepository) Insert(ut models.UserTransactionModel) error {
	r, err := r1.dbq.Exec(insertTransactionQuery, ut.UserID, ut.Amount, ut.Description)
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

func (r1 *r1WalletRepository) InsertTransaction(ut models.UserTransactionModel) error {
	trx, err := r1.beginTransaction()
	if err != nil {
		return err
	}

	bu, err := trx.GetBalanceByUserID(ut.UserID)
	if err != nil {
		ib := trx.InsertBalance(ut.UserID, ut.Amount)
		if ib != nil {
			println(ib)
			er := trx.rollbackTransaction()
			if er != nil {
				return er
			}
			return ib
		}
	} else {
		ub := trx.UpdateBalance(bu+ut.Amount, ut.UserID)

		if ub != nil {
			er := trx.rollbackTransaction()
			if er != nil {
				return er
			}
			return ub
		}
	}

	r := trx.Insert(ut)
	if r != nil {
		er := trx.rollbackTransaction()
		if er != nil {
			return er
		}
		return r
	}

	return trx.commitTransaction()
}

func (r1 *r1WalletRepository) beginTransaction() (*r1WalletRepository, error) {
	tx, err := r1.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return &r1WalletRepository{}, err
	}

	return &r1WalletRepository{tx: tx, dbq: tx}, nil
}

func (r1 *r1WalletRepository) commitTransaction() error {
	if r1.tx == nil {
		return fmt.Errorf("you cant commit tansaction befor start it")
	}

	return r1.tx.Commit()
}

func (r1 *r1WalletRepository) rollbackTransaction() error {
	if r1.tx == nil {
		return fmt.Errorf("you cant rollback tansaction befor start it")
	}

	return r1.tx.Rollback()
}
