package repository

import (
	"account-manager/api/account/domain"
	"account-manager/config/models"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type AccRepositories struct {
	db *sql.DB
}

func NewAccRepo(db *sql.DB) domain.AccRepositories {
	return &AccRepositories{db: db}
}

func (r *AccRepositories) RegisterUser(ctx *gin.Context, req models.ReqRegister) (err error) {

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) returning id"
	if _, err = r.db.ExecContext(ctx, query, req.Username, req.Email, req.Password); err != nil {
		return err
	}

	return
}

func (r *AccRepositories) GetUserByEmail(ctx *gin.Context, email string) (respEmail string, err error) {
	query := "SELECT email FROM users WHERE email = $1"
	if err = r.db.QueryRowContext(ctx, query, email).Scan(&respEmail); err != nil {
		if err == sql.ErrNoRows {
			return respEmail, nil
		}
		return
	}

	return
}

func (r *AccRepositories) GetUserIdByEmail(ctx *gin.Context, email string) (ID string, err error) {
	query := "SELECT id FROM users WHERE email = $1"

	if err = r.db.QueryRowContext(ctx, query, email).Scan(&ID); err != nil {
		return
	}

	return
}

func (r *AccRepositories) GetCurrentUser(ctx *gin.Context, userId string) (resp models.RespCurrentUser, err error) {
	query := "SELECT id, username, email, coalesce(phone, '') from users where id = $1"
	if err = r.db.QueryRowContext(ctx, query, userId).Scan(&resp.ID, &resp.Username, &resp.Email, &resp.Phone); err != nil {
		return resp, err
	}

	return
}

func (r *AccRepositories) UpdateUserProfile(ctx *gin.Context, req models.ReqUpdateUserProfile, userId string) error {
	query := "UPDATE users set username = $1, phone = $2 where id = $3"
	if _, err := r.db.ExecContext(ctx, query, req.Username, req.Phone, userId); err != nil {
		return err
	}

	return nil
}

func (r *AccRepositories) AddAccountBankUser(ctx *gin.Context, req models.ReqAddAccBank, userId string) error {
	query := "INSERT INTO account_bank(user_id, acc_name, bank_name, acc_number) VALUES ($1, $2, $3, $4)"

	if _, err := r.db.ExecContext(ctx, query, userId, req.AccName, req.BankName, req.AccNumber); err != nil {
		return err
	}
	return nil
}

func (r *AccRepositories) GetListAccBank(ctx *gin.Context, userId string) (resp []models.RespListAccBank, err error) {
	query := "SELECT id, bank_name, acc_number from account_bank where user_id=$1"
	rows, err := r.db.QueryContext(ctx, query, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	defer rows.Close()

	for rows.Next() {
		var data models.RespListAccBank
		if err = rows.Scan(&data.ID, &data.BankName, &data.AccNumber); err != nil {
			return nil, err
		}

		resp = append(resp, data)
	}

	return
}

func (r *AccRepositories) GetDetailAccBank(ctx *gin.Context, accNumber string) (resp models.RespDetailAccBank, err error) {
	query := "SELECT id, acc_name, bank_name, acc_number, balance from account_bank where acc_number=$1"

	if err = r.db.QueryRowContext(ctx, query, accNumber).Scan(&resp.ID, &resp.AccName, &resp.BankName, &resp.AccNumber, &resp.Saldo); err != nil {
		return resp, err
	}

	return
}

func (r *AccRepositories) UpdateAccBank(ctx *gin.Context, finalBalance int64, accNumber string) error {
	query := "UPDATE account_bank set balance = $1 where acc_number = $2"
	if _, err := r.db.ExecContext(ctx, query, finalBalance, accNumber); err != nil {
		return err
	}

	return nil
}
