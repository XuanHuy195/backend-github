package router

import (
	"backend-github-trending/handler"
	"backend-github-trending/middleware"
	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *Api) SetupRouter() {
	//user
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	//profile
	profile := api.Echo.Group("/user", middleware.JWTMiddleWare())
	profile.GET("/profile", api.UserHandler.Profile)
	profile.PUT("/profile/update", api.UserHandler.UpdateProfile)

}
