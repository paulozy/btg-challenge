package server

var Routes = []Handler{}

// func PopulateRoutes(mongoClient *mongo.Client) []Handler {
// 	addOrderRoutes(mongoClient)
// 	return Routes
// }

// func addOrderRoutes(mongoClient *mongo.Client) {
// 	orderRepository := repositories.NewOrderRepository(mongoClient, "btg_challenge", "orders")
// 	saveOrderUseCase := usecases.NewSaveOrderUseCase(orderRepository)

// 	orderUseCases := controllers.OrderUseCasesInput{
// 		SaveOrderUseCase: saveOrderUseCase,
// 	}

// 	orderController := controllers.NewOrderController(orderRepository, orderUseCases)

// 	orderControllerRoutes := []Handler{}

// 	Routes = append(Routes, orderControllerRoutes...)
// }
