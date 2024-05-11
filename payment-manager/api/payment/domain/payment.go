package domain

import (
	"payment-manager/config/models"

	"github.com/gin-gonic/gin"
)

type PaymentRepositories interface {
	InsertHistoryTrans(ctx *gin.Context, req models.ReqTransaction, typeTrans string, userId string) error
	GetListTrans(ctx *gin.Context, userId string) (resp []models.RespListTrans, err error)
	GetDetailTrans(ctx *gin.Context, idTrans string) (resp models.RespDetailTrans, err error)
}

type PaymentUsecase interface {
	InsertHistoryTrans(ctx *gin.Context, token string, urlSum string, urlReduc string, req models.ReqTransaction, typeTrans string, userId string) error
	GetListTrans(ctx *gin.Context, userId string) (resp []models.RespListTrans, err error)
	GetDetailTrans(ctx *gin.Context, idTrans string) (resp models.RespDetailTrans, err error)
}
