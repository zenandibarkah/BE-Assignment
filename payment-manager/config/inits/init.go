package inits

import (
	"net/http"
	"payment-manager/config/db"
	"payment-manager/config/models"
	"payment-manager/helper"

	"github.com/gin-gonic/gin"
)

func init() {
	if helper.MyConfig.Environment == "PROD" {
		gin.SetMode(gin.ReleaseMode)
		helper.Log.Printf("Starting %s on PRODUCTION Environment", helper.MyConfig.AppName)
	} else {
		helper.Log.Printf("Starting %s on DEVELOPMENT Environment", helper.MyConfig.Environment)
	}

	if err := db.InitDBConnection(); err != nil {
		helper.Log.Fatal(err.Error())
	}
}

func HandlerHealthCheck(ctx *gin.Context) {
	resp := models.Response{StatusCode: http.StatusOK, Status: "success", Message: "Health Check Success", DisplayMsg: "warning.healthCheckSuccess", Response: nil}
	ctx.JSON(http.StatusOK, resp)
}

func HandlerPanic(ctx *gin.Context, err interface{}) {
	helper.Log.Error(err.(error).Error())
	resp := models.Response{StatusCode: http.StatusInternalServerError, Status: "failed", Message: "internal server error, Please come back later", DisplayMsg: "warning.internalError", Response: nil}
	ctx.JSON(http.StatusInternalServerError, resp)
}
