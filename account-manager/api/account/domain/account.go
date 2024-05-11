package domain

import (
	"account-manager/config/models"

	"github.com/gin-gonic/gin"
)

type AccRepositories interface {
	RegisterUser(ctx *gin.Context, req models.ReqRegister) (err error)
	GetUserByEmail(ctx *gin.Context, email string) (respEmail string, err error)
	GetUserIdByEmail(ctx *gin.Context, email string) (ID string, err error)
	UpdateUserProfile(ctx *gin.Context, req models.ReqUpdateUserProfile, userId string) error
	GetCurrentUser(ctx *gin.Context, userId string) (resp models.RespCurrentUser, err error)

	AddAccountBankUser(ctx *gin.Context, req models.ReqAddAccBank, userId string) error
	GetListAccBank(ctx *gin.Context, userId string) (resp []models.RespListAccBank, err error)
	GetDetailAccBank(ctx *gin.Context, accNumber string) (resp models.RespDetailAccBank, err error)

	// TRANS
	UpdateAccBank(ctx *gin.Context, finalBalance int64, accNumber string) error
}

type AccUsecase interface {
	RegisterUser(ctx *gin.Context, req models.ReqRegister) (err error)
	GetUserIdByEmail(ctx *gin.Context, email string) (ID string, err error)
	UpdateUserProfile(ctx *gin.Context, req models.ReqUpdateUserProfile, userId string) error
	GetCurrentUser(ctx *gin.Context, userId string) (resp models.RespCurrentUser, err error)

	AddAccountBankUser(ctx *gin.Context, req models.ReqAddAccBank, userId string) error
	GetListAccBank(ctx *gin.Context, userId string) (resp []models.RespListAccBank, err error)
	GetDetailAccBank(ctx *gin.Context, accNumber string) (resp models.RespDetailAccBank, err error)

	// TRANS
	UpdateAccBankSum(ctx *gin.Context, amount int64, accNumber string) error
	UpdateAccBankReduc(ctx *gin.Context, amount int64, accNumber string) error
}
