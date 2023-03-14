package services

import (
	"github.com/fernandormoraes/stonksdayapi/app/models"
	"github.com/fernandormoraes/stonksdayapi/app/repositories"
)

type IStockService interface {
	GetStock(stockName string) (models.Stock, error)
}

type StockService struct {
	stockRepository repositories.IStockRepository
}

func NewStockService(stockRepository repositories.IStockRepository) *StockService {
	return &StockService{stockRepository: stockRepository}
}

func (s StockService) GetStock(stockName string) (models.Stock, error) {
	return s.stockRepository.GetStock(stockName)
}
