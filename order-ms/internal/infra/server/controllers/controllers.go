package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database"
	"github.com/paulozy/btg-challenge/order-ms/internal/usecases"
)

type OrderController struct {
	orderRepository               database.OrderRepositoryInterface
	listOrdersByClientCodeUseCase *usecases.ListOrdersByClientCodeUseCase
}

type OrderUseCasesInput struct {
	ListOrdersByClientCodeUseCase *usecases.ListOrdersByClientCodeUseCase
}

func NewOrderController(or database.OrderRepositoryInterface, usecases OrderUseCasesInput) *OrderController {
	return &OrderController{
		orderRepository:               or,
		listOrdersByClientCodeUseCase: usecases.ListOrdersByClientCodeUseCase,
	}
}

func (oc *OrderController) ListByClientCode(c *gin.Context) {
	clientCode := c.Query("clientCode")
	if clientCode == "" {
		c.JSON(400, gin.H{"error": "missing param", "reason": "missing clientCode param"})
		return
	}

	normalizedClientCode, err := strconv.Atoi(clientCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "reason": err})
		return
	}

	input := usecases.ListOrdersByClientCodeInput{
		ClientCode: normalizedClientCode,
	}

	orders, ucError := oc.listOrdersByClientCodeUseCase.Execute(input)
	if ucError.Message != "" {
		c.JSON(ucError.Status, gin.H{"error": ucError.Message, "reason": ucError.Error})
		return
	}

	c.JSON(200, gin.H{"count": len(orders), "rows": orders})
}
