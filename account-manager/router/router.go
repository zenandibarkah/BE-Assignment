package router

import (
	acchandler "account-manager/api/account/handler"
	accrepo "account-manager/api/account/repository"
	accusecase "account-manager/api/account/usecase"
	"account-manager/config/db"
	"account-manager/config/inits"
	"account-manager/config/middleware"

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

func InitRouter(router *gin.Engine) {
	public := router.Group("/public")
	private := router.Group("/private")
	private.Use(middleware.MiddlewareJWTOverride())

	db, err := db.GetConnectionDB()
	if err != nil {
		return
	}

	accRepo := accrepo.NewAccRepo(db)
	accUsecase := accusecase.NewAccUsecase(accRepo)
	acchandler.NewAccHandler(public, private, accUsecase)
}
