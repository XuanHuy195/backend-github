package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//logic
			req := req.RequestSignUp{}
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			if req.Email != "admin@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "Bạn không có quyền truy cập",
					Data:       nil,
				})
			}
			return next(c)
		}
	}
}
