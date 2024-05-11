package usecase

import (
	"payment-manager/api/payment/domain"
	"payment-manager/config/models"
	"payment-manager/service"

	"github.com/gin-gonic/gin"
)

type PaymentUsecase struct {
	paymentRepo domain.PaymentRepositories
}

func NewPaymentUsecase(paymentRepo domain.PaymentRepositories) domain.PaymentUsecase {
	return &PaymentUsecase{paymentRepo: paymentRepo}
}

func (u *PaymentUsecase) InsertHistoryTrans(ctx *gin.Context, token string, urlSum string, urlReduc string, req models.ReqTransaction, typeTrans string, userId string) error {
	if typeTrans == "CREDIT" {
		if err := service.HttpReqBalanceReduc(models.ReqUpdateAccBalance{AccNumber: req.SourceAccNum, Amount: req.Amount}, token, urlReduc); err != nil {
			return err
		}

		if err := service.HttpReqBalanceSum(models.ReqUpdateAccBalance{AccNumber: req.DestAccNum, Amount: req.Amount}, token, urlSum); err != nil {
			return err
		}
	} else if typeTrans == "DEBIT" {
		if err := service.HttpReqBalanceSum(models.ReqUpdateAccBalance{AccNumber: req.DestAccNum, Amount: req.Amount}, token, urlSum); err != nil {
			return err
		}
	}

	if err := u.paymentRepo.InsertHistoryTrans(ctx, req, typeTrans, userId); err != nil {
		return err
	}

	return nil
}

func (u *PaymentUsecase) GetListTrans(ctx *gin.Context, userId string) (resp []models.RespListTrans, err error) {
	resp, err = u.paymentRepo.GetListTrans(ctx, userId)
	if err != nil {
		return nil, err
	}

	return
}

func (u *PaymentUsecase) GetDetailTrans(ctx *gin.Context, idTrans string) (resp models.RespDetailTrans, err error) {
	resp, err = u.paymentRepo.GetDetailTrans(ctx, idTrans)
	if err != nil {
		return resp, err
	}

	return
}
