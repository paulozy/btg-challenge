package database

import "github.com/paulozy/btg-challenge/order-ms/internal/entity"

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	GetByClientCode(clientCode int) ([]entity.Order, error)
	FindByOrderCode(clientCode, orderCode int) (*entity.Order, error)
}
