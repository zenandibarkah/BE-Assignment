package usecase

import (
	"account-manager/api/account/domain"
	"account-manager/config/models"
	"account-manager/helper"
	"errors"

	"github.com/gin-gonic/gin"
)

type AccUsecase struct {
	AccRepo domain.AccRepositories
}

func NewAccUsecase(accRepo domain.AccRepositories) domain.AccUsecase {
	return &AccUsecase{AccRepo: accRepo}
}

func (u *AccUsecase) RegisterUser(ctx *gin.Context, req models.ReqRegister) (err error) {
	hashPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashPassword

	email, err := u.AccRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	if email == req.Email {
		return errors.New("email already registered")
	}

	err = u.AccRepo.RegisterUser(ctx, req)
	if err != nil {
		return err
	}

	return
}

func (u *AccUsecase) GetUserIdByEmail(ctx *gin.Context, email string) (ID string, err error) {
	ID, err = u.AccRepo.GetUserIdByEmail(ctx, email)
	if err != nil {
		return
	}

	return
}

func (u *AccUsecase) GetCurrentUser(ctx *gin.Context, userId string) (resp models.RespCurrentUser, err error) {
	resp, err = u.AccRepo.GetCurrentUser(ctx, userId)
	if err != nil {
		return resp, err
	}

	return
}

func (u *AccUsecase) UpdateUserProfile(ctx *gin.Context, req models.ReqUpdateUserProfile, userId string) error {
	if err := u.AccRepo.UpdateUserProfile(ctx, req, userId); err != nil {
		return err
	}

	return nil
}

func (u *AccUsecase) AddAccountBankUser(ctx *gin.Context, req models.ReqAddAccBank, userId string) error {
	if err := u.AccRepo.AddAccountBankUser(ctx, req, userId); err != nil {
		return err
	}

	return nil
}

func (u *AccUsecase) GetListAccBank(ctx *gin.Context, userId string) (resp []models.RespListAccBank, err error) {
	resp, err = u.AccRepo.GetListAccBank(ctx, userId)
	if err != nil {
		return nil, err
	}

	return
}

func (u *AccUsecase) GetDetailAccBank(ctx *gin.Context, accNumber string) (resp models.RespDetailAccBank, err error) {
	resp, err = u.AccRepo.GetDetailAccBank(ctx, accNumber)
	if err != nil {
		return resp, err
	}

	return
}

func (u *AccUsecase) UpdateAccBankSum(ctx *gin.Context, amount int64, accNumber string) error {
	resp, err := u.AccRepo.GetDetailAccBank(ctx, accNumber)
	if err != nil {
		return err
	}

	finalBalance := resp.Saldo + amount

	if err := u.AccRepo.UpdateAccBank(ctx, finalBalance, accNumber); err != nil {
		return err
	}

	return nil
}

func (u *AccUsecase) UpdateAccBankReduc(ctx *gin.Context, amount int64, accNumber string) error {
	resp, err := u.AccRepo.GetDetailAccBank(ctx, accNumber)
	if err != nil {
		return err
	}

	finalBalance := resp.Saldo - amount
	if finalBalance < 0 {
		return errors.New("balance is 0")
	}

	if err := u.AccRepo.UpdateAccBank(ctx, finalBalance, accNumber); err != nil {
		return err
	}

	return nil
}
