package router

import (
	"github.com/labstack/echo/v4"
	"github.com/s-yuk/go-bank-api/controller"
)

func Router() {

	e := echo.New()
	e.GET("/api/v1/bank/user", controller.FindAllUser)
	e.GET("/api/v1/bank/user/:id", controller.FindByUserID)
	e.POST("/api/v1/bank/user", controller.CreateUser)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
