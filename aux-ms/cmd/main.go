package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/paulozy/btg-challenge/aux-ms/configs"
	"github.com/paulozy/btg-challenge/aux-ms/internal/services"
	"github.com/paulozy/btg-challenge/aux-ms/pkg"

	"math/rand"
)

func main() {
	configs, err := configs.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	awsService, err := services.NewAWSConfig(
		configs.AWSAccessKey,
		configs.AWSSecretKey,
	)
	if err != nil {
		panic(err)
	}

	sqsService := sqs.New(&awsService.Session)

	// clientQty := flag.Int("clientQty", 1, "quantity of clients")
	// ordersPerClient := flag.Int("ordersPerClient", 1, "quantity of orders per client")

	// flag.Parse()

	rand.Seed(time.Now().UnixNano())

	for {
		clientQty := rand.Intn(4) + 1
		ordersPerClient := 1

		fmt.Printf("Creating Order for %d clients \n", clientQty)

		messages := pkg.CreateMessages(clientQty, ordersPerClient)
		for _, msg := range messages {
			sqsService.SendMessage(
				&sqs.SendMessageInput{
					QueueUrl:    &configs.SQSCreatedOrdersQueueUrl,
					MessageBody: aws.String(msg),
				},
			)
		}

		time.Sleep(30 * time.Second)
	}

	// fmt.Println("messages sended")
}
