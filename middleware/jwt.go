package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleWare() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JWTCustomClaims{},
		SigningKey: []byte(security.SECRECT_KEY),
	}
	return middleware.JWTWithConfig(config)
}
