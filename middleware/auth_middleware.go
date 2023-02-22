package middleware

import (
	"context"
	"ecourse-app/config"
	"ecourse-app/exception"
	"ecourse-app/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		exception.ErrorHandler(writer, request, err)
	}()
	if request.URL.Path != "/admin/login" && request.URL.Path != "/admin/register" {

		authorization := request.Header.Get("Authorization")
		if authorization != "" {
			if !strings.Contains(authorization, "Bearer") {
				panic(exception.NewUnauthorizedError("invalid token"))
			}

			tokenString := strings.Replace(authorization, "Bearer", "", -1)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				method, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok || method != config.JWT_SIGNING_METHOD {
					return nil, fmt.Errorf("signin method invalid")
				}

				return config.JWT_KEY, nil
			})
			helper.PanicError(err)

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				panic(exception.NewUnauthorizedError("invalid token"))
			}

			ctx := context.WithValue(context.Background(), "adminInfo", claims)
			request = request.WithContext(ctx)

		} else {
			panic(exception.NewUnauthorizedError("token not found"))
		}

	}

	middleware.Handler.ServeHTTP(writer, request)
}
