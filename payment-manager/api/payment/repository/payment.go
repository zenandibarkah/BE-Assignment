package repository

import (
	"database/sql"
	"payment-manager/api/payment/domain"
	"payment-manager/config/models"

	"github.com/gin-gonic/gin"
)

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) domain.PaymentRepositories {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) InsertHistoryTrans(ctx *gin.Context, req models.ReqTransaction, typeTrans string, userId string) error {
	query := `INSERT INTO history_trans(user_id, source_bank, source_account_number, destination_bank, destination_account_number, amount, trans_type)
			VALUES($1, $2, $3, $4, $5, $6, $7)`

	if _, err := r.db.ExecContext(ctx, query, userId, req.SourceBank, req.SourceAccNum, req.DestBank, req.DestAccNum, req.Amount, typeTrans); err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepo) GetListTrans(ctx *gin.Context, userId string) (resp []models.RespListTrans, err error) {
	query := `SELECT id, source_bank, destination_bank, trans_type from history_trans
			WHERE user_id = $1 order by created_at desc`

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	defer rows.Close()

	for rows.Next() {
		var data models.RespListTrans
		if err = rows.Scan(&data.ID, &data.SourceBank, &data.DestBank, &data.TransType); err != nil {
			return nil, err
		}

		resp = append(resp, data)
	}

	return
}

func (r *PaymentRepo) GetDetailTrans(ctx *gin.Context, idTrans string) (resp models.RespDetailTrans, err error) {
	query := `SELECT id, source_bank, source_account_number, destination_bank, destination_account_number, amount, trans_type, created_at
				FROM history_trans WHERE id = $1`

	if err = r.db.QueryRowContext(ctx, query, idTrans).Scan(&resp.ID, &resp.SourceBank, &resp.SourceAccNum,
		&resp.DestBank, &resp.DestAccNum, &resp.Amount, &resp.TransType, &resp.TransDate); err != nil {
		return resp, err
	}

	return
}
