package controllers

import (
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database"
	"github.com/paulozy/btg-challenge/order-ms/internal/usecases"
)

type OrderController struct {
	orderRepository  database.OrderRepositoryInterface
	saveOrderUseCase *usecases.SaveOrderUseCase
}

type OrderUseCasesInput struct {
	SaveOrderUseCase *usecases.SaveOrderUseCase
}

func NewOrderController(or database.OrderRepositoryInterface, usecases OrderUseCasesInput) *OrderController {
	return &OrderController{
		orderRepository:  or,
		saveOrderUseCase: usecases.SaveOrderUseCase,
	}
}
