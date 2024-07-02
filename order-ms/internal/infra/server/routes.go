package server

import (
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database/repositories"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/server/controllers"
	"github.com/paulozy/btg-challenge/order-ms/internal/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME string = "orders"

var Routes = []Handler{}

func PopulateRoutes(mongoClient *mongo.Client, databaseName string) []Handler {
	addOrderRoutes(mongoClient, databaseName)
	return Routes
}

func addOrderRoutes(mongoClient *mongo.Client, databaseName string) {
	orderRepository := repositories.NewOrderRepository(mongoClient, databaseName, COLLECTION_NAME)
	listOrdersByClientCodeUseCase := usecases.NewListOrdersByClientCodeUseCase(orderRepository)
	showOrderByOrderCodeUseCase := usecases.NewShowOrderByOrderCodeUseCase(orderRepository)

	orderUseCases := controllers.OrderUseCasesInput{
		ListOrdersByClientCodeUseCase: listOrdersByClientCodeUseCase,
		ShowOrderByOrderCodeUseCase:   showOrderByOrderCodeUseCase,
	}

	orderController := controllers.NewOrderController(orderUseCases)

	orderControllerRoutes := []Handler{
		{
			Path:   "/orders",
			Method: "GET",
			Func:   orderController.ListByClientCode,
		},
		{
			Path:   "/orders/:orderCode",
			Method: "GET",
			Func:   orderController.GetOrderByOrderCode,
		},
	}

	Routes = append(Routes, orderControllerRoutes...)
}
