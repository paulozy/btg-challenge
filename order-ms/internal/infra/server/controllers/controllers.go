package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paulozy/btg-challenge/order-ms/internal/usecases"
)

type OrderController struct {
	listOrdersByClientCodeUseCase *usecases.ListOrdersByClientCodeUseCase
	showOrderByOrderCodeUseCase   *usecases.ShowOrderByOrderCodeUseCase
}

type OrderUseCasesInput struct {
	ListOrdersByClientCodeUseCase *usecases.ListOrdersByClientCodeUseCase
	ShowOrderByOrderCodeUseCase   *usecases.ShowOrderByOrderCodeUseCase
}

func NewOrderController(usecases OrderUseCasesInput) *OrderController {
	return &OrderController{
		listOrdersByClientCodeUseCase: usecases.ListOrdersByClientCodeUseCase,
		showOrderByOrderCodeUseCase:   usecases.ShowOrderByOrderCodeUseCase,
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

func (oc *OrderController) GetOrderByOrderCode(c *gin.Context) {
	orderCode := c.Param("orderCode")
	if orderCode == "" {
		c.JSON(400, gin.H{"error": "missing param", "reason": "missing orderCode param"})
		return
	}

	normalizedOrderCode, err := strconv.Atoi(orderCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "reason": err})
		return
	}

	input := usecases.ShowOrderByOrderCodeInput{
		OrderCode: normalizedOrderCode,
	}

	order, ucError := oc.showOrderByOrderCodeUseCase.Execute(input)
	if ucError.Message != "" {
		c.JSON(ucError.Status, gin.H{"error": ucError.Message, "reason": ucError.Error})
		return
	}

	c.JSON(200, gin.H{"data": order})
}
