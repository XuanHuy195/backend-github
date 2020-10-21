package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"backend-github-trending/repository/repo_impl"
	"backend-github-trending/router"
	"github.com/labstack/echo/v4"
)

func main() {
	sql:= &db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "Huy",
		Password: "huydoi1955",
		DbName: "Golang",
	}
	sql.Connect()

	defer sql.Close()
	e := echo.New()

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.Api{
		Echo: e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":5500"))
}
