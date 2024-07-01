package main

import (
	"github.com/paulozy/btg-challenge/order-ms/configs"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/server"
	services "github.com/paulozy/btg-challenge/order-ms/internal/services/aws"
)

func main() {
	configs, err := configs.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	awsService, err := services.NewAWS(
		configs.AWSAccessKey,
		configs.AWSSecretKey,
	)
	if err != nil {
		panic(err)
	}

	mongoClient, cancel := database.NewDatabaseConnection(configs.DatabaseURL)
	defer cancel()

	sqsService := services.NewSQSService(&awsService.Session, &configs.SQSCreatedOrdersQueueUrl)
	go sqsService.ReadMessagesAndSaveOrder(mongoClient)

	server.PopulateRoutes(mongoClient)

	server := server.NewServer(configs.WebHost, configs.WebPort, configs.Env)
	server.AddHandlers()
	server.Start()
}
