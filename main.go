package main

import (
	"github.com/fernandormoraes/stonksdayapi/app/handlers"
	"github.com/fernandormoraes/stonksdayapi/app/repositories"
	"github.com/fernandormoraes/stonksdayapi/app/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	stockHandler := *handlers.NewStockHandler(services.NewStockService(repositories.NewStockRepository()))

	e.GET("/:stock", stockHandler.GetStock)

	e.Logger.Fatal(e.Start(":1323"))
}
