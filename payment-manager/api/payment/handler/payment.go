package handler

import (
	"net/http"
	"payment-manager/api/payment/domain"
	"payment-manager/config/middleware"
	"payment-manager/config/models"
	"payment-manager/helper"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUsecace domain.PaymentUsecase
}

func NewPaymentHandler(gprivate *gin.RouterGroup, paymentUsecase domain.PaymentUsecase) {
	handler := &PaymentHandler{paymentUsecace: paymentUsecase}

	gprivate.GET("/history/trans", handler.GetListTrans)
	gprivate.GET("/history/trans/:idTrans", handler.GetDetailTrans)
	gprivate.POST("/trans/send", handler.SendTransaction)
	gprivate.POST("/trans/withdraw", handler.WithdrawTransaction)
}

func (h *PaymentHandler) SendTransaction(ctx *gin.Context) {
	var req models.ReqTransaction

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	token, err := middleware.ExtractToken(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	if err := h.paymentUsecace.InsertHistoryTrans(ctx, token, helper.MyConfig.BalanceSum, helper.MyConfig.BalanceReduc, req, "CREDIT", userId); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Send Transaction Successfull", DisplayMsg: "warning.sucess", Response: nil})
}

func (h *PaymentHandler) WithdrawTransaction(ctx *gin.Context) {
	var req models.ReqTransaction

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "Failed binding from JSON", DisplayMsg: "", Response: nil})
		return
	}

	token, err := middleware.ExtractToken(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	if err := h.paymentUsecace.InsertHistoryTrans(ctx, token, helper.MyConfig.BalanceSum, helper.MyConfig.BalanceReduc, req, "DEBIT", userId); err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Withdraw Transaction Successfull", DisplayMsg: "warning.sucess", Response: nil})
}

func (h *PaymentHandler) GetListTrans(ctx *gin.Context) {
	userId, err := middleware.GetUid(ctx)
	if err != nil {
		helper.Log.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Status: "Unauthorized", Message: "Invalid Token", DisplayMsg: "warning.invalidToken"})
	}

	resp, err := h.paymentUsecace.GetListTrans(ctx, userId)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Get List History Trans Successfull", DisplayMsg: "warning.sucess", Response: resp})
}

func (h *PaymentHandler) GetDetailTrans(ctx *gin.Context) {
	idTrans := ctx.Param("idTrans")

	resp, err := h.paymentUsecace.GetDetailTrans(ctx, idTrans)
	if err != nil {
		helper.Log.Errorf("Error %v for requestId %v", err.Error(), requestid.Get(ctx))
		ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request", Message: "", DisplayMsg: "", Response: nil})
		return
	}

	helper.Log.Infof("Success for requestId %v", requestid.Get(ctx))
	ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Status: "Success", Message: "Get Detail History Trans Successfull", DisplayMsg: "warning.sucess", Response: resp})
}
