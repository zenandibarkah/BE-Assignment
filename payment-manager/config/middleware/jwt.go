package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"payment-manager/config/db"
	"payment-manager/helper"
	"payment-manager/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func MiddlewareJWTOverride() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := MiddlewareJWT(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func MiddlewareJWT(ctx *gin.Context) error {

	tokenString, err := ExtractToken(ctx)
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(helper.MyConfig.JWTSecret), nil
	})
	//
	if err != nil {
		return errors.New("invalid token")
	}

	_, err = service.HttpReqTokenCheck(tokenString, helper.MyConfig.TokenValidate)
	if err != nil {
		return err
	}

	// ctx.Next()
	return nil
}

func ExtractToken(ctx *gin.Context) (string, error) {
	var tokenString string
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		return "", errors.New("autorization not found")
	}

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	}

	if tokenString == "" {
		return "", errors.New("token is empty")
	}
	return tokenString, nil
}

func GetUid(ctx *gin.Context) (string, error) {
	tokenString, err := ExtractToken(ctx)
	if err != nil {
		return "", err
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(helper.MyConfig.JWTSecret), nil
	})
	//
	if err != nil {
		return "", errors.New("invalid token")
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return "", errors.New("invalid token")
	}
	// fmt.Println(claims["id_user"])

	db, err := db.GetConnectionDB()
	if err != nil {
		helper.Log.Errorln("Error in GetConnectionDB", err.Error())
		return "", err
	}

	var id string
	if err := db.QueryRow("select id from users where id=$1", claims["id_user"]).Scan(&id); err != nil {
		return "", err
	}

	if id != claims["id_user"] {
		return "", errors.New("the user belonging to this token no logger exists")
	}

	return id, nil
}
