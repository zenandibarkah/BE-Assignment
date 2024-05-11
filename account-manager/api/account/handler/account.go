package handler

import (
	"account-manager/api/account/domain"
	"account-manager/config/middleware"
	"account-manager/config/models"
	"account-manager/helper"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type AccHandler struct {
	AccUsecase domain.AccUsecase
}

func NewAccHandler(gpublic *gin.RouterGroup, gprivate *gin.RouterGroup, accUsecase domain.AccUsecase) {
	handler := &AccHandler{AccUsecase: accUsecase}

	gpublic.POST("/register", handler.RegisterUser)
	gpublic.POST("/login", handler.Login)

	gprivate.GET("/logout", Logout)
	gprivate.GET("/token/validate", TokenValidate)
	gprivate.GET("/myuser", handler.GetCurrentUser)
	gprivate.POST("/user/update", handler.UpdateUserProfile)

	gprivate.POST("/user/bank/add", handler.AddAccountBankUser)
	gprivate.GET("/user/banks", handler.GetListAccBank)
	gprivate.GET("/user/bank/:accNumber", handler.GetDetailAccBank)
	gprivate.POST("/bank/balance/sum", handler.UpdateAccBankSum)
	gprivate.POST("/bank/balance/reduc", handler.UpdateAccBankReduc)
}

func TokenValidate(ctx *gin.Context) {
	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "User is valid", DisplayMsg: "warning.sucess", Response: userId})
}

func (h *AccHandler) RegisterUser(ctx *gin.Context) {
	var req models.ReqRegister

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	err := h.AccUsecase.RegisterUser(ctx, req)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: err.Error(), DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Register Successfull", DisplayMsg: "warning.sucess", Response: req.Username})
}

func (h *AccHandler) Login(ctx *gin.Context) {
	var req models.ReqLogin

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	userId, err := h.AccUsecase.GetUserIdByEmail(ctx, req.Email)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	respLogin, err := middleware.GenerateToken(userId)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Login Successfull", DisplayMsg: "warning.sucess", Response: respLogin})
}

func Logout(ctx *gin.Context) {
	ctx.Request.Header.Del("Authorization")

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Logout Successfull", DisplayMsg: "warning.sucess", Response: nil})
}

func (h *AccHandler) GetCurrentUser(ctx *gin.Context) {
	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	resp, err := h.AccUsecase.GetCurrentUser(ctx, userId)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Get Current User Successfull", DisplayMsg: "warning.sucess", Response: resp})

}

func (h *AccHandler) UpdateUserProfile(ctx *gin.Context) {
	var req models.ReqUpdateUserProfile

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	if err := h.AccUsecase.UpdateUserProfile(ctx, req, userId); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Update User Successfull", DisplayMsg: "warning.sucess", Response: nil})
}

func (h *AccHandler) AddAccountBankUser(ctx *gin.Context) {
	var req models.ReqAddAccBank

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	if err := h.AccUsecase.AddAccountBankUser(ctx, req, userId); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Add Account Bank Successfull", DisplayMsg: "warning.sucess", Response: nil})
}

func (h *AccHandler) GetListAccBank(ctx *gin.Context) {
	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	resp, err := h.AccUsecase.GetListAccBank(ctx, userId)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Get List Account Bank Successfull", DisplayMsg: "warning.sucess", Response: resp})
}

func (h *AccHandler) GetDetailAccBank(ctx *gin.Context) {
	accNumber := ctx.Param("accNumber")

	resp, err := h.AccUsecase.GetDetailAccBank(ctx, accNumber)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Get Detail Account Bank Successfull", DisplayMsg: "warning.sucess", Response: resp})
}

func (h *AccHandler) UpdateAccBankSum(ctx *gin.Context) {
	var req models.ReqUpdateAccBalance

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	if err := h.AccUsecase.UpdateAccBankSum(ctx, req.Amount, req.AccNumber); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Update Balance Successfull", DisplayMsg: "warning.sucess", Response: nil})
}

func (h *AccHandler) UpdateAccBankReduc(ctx *gin.Context) {
	var req models.ReqUpdateAccBalance

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	if err := h.AccUsecase.UpdateAccBankReduc(ctx, req.Amount, req.AccNumber); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Update Balance Successfull", DisplayMsg: "warning.sucess", Response: nil})
}
