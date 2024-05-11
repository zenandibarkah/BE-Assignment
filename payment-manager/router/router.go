package router

import (
	paymenthandler "payment-manager/api/payment/handler"
	paymentrepo "payment-manager/api/payment/repository"
	paymentusecase "payment-manager/api/payment/usecase"
	"payment-manager/config/db"
	"payment-manager/config/inits"
	"payment-manager/config/middleware"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func CreateRouter(isDev bool) *gin.Engine {
	router := gin.New()

	// use Middlerware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.Secure(isDev))
	router.Use(requestid.New())
	router.Use(gin.CustomRecovery(func(ctx *gin.Context, err interface{}) {
		inits.HandlerPanic(ctx, err)
	}))
	router.Use(healthcheck.Default())

	return router
}

func InitRoute(router *gin.Engine) {
	private := router.Group("/private")
	private.Use(middleware.MiddlewareJWTOverride())

	db, err := db.GetConnectionDB()
	if err != nil {
		return
	}

	paymentRepo := paymentrepo.NewPaymentRepo(db)
	paymentUsecase := paymentusecase.NewPaymentUsecase(paymentRepo)
	paymenthandler.NewPaymentHandler(private, paymentUsecase)
}
