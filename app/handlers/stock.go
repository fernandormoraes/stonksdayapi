package handlers

import (
	"net/http"

	"github.com/fernandormoraes/stonksdayapi/app/services"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type IStockHandler interface {
	getStock(c echo.Context) error
}

type StockHandler struct {
	stockService services.IStockService
}

func NewStockHandler(stockService services.IStockService) *StockHandler {
	return &StockHandler{stockService: stockService}
}

func (s StockHandler) GetStock(c echo.Context) error {
	stockName := c.Param("stock")

	stock, err := s.stockService.GetStock(stockName)

	if err != nil {
		logrus.Print(err)
	}

	return c.JSON(http.StatusOK, stock)
}
