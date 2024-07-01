package services

import (
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/paulozy/btg-challenge/order-ms/internal/entity"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/database/repositories"
	"github.com/paulozy/btg-challenge/order-ms/internal/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type SQSService struct {
	queueURL       *string
	timeout        int64
	maxNumMessages int64
	svc            *sqs.SQS
}

func NewSQSService(sess *session.Session, queueURL *string) *SQSService {
	sqsSession := sqs.New(sess)

	timeout := 10 * time.Second.Milliseconds()
	maxNumMessages := int64(10)

	return &SQSService{
		queueURL,
		timeout,
		maxNumMessages,
		sqsSession,
	}
}

func (ssvc *SQSService) GetMessages() (*sqs.ReceiveMessageOutput, error) {
	msgRes, err := ssvc.svc.ReceiveMessage(
		&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            ssvc.queueURL,
			MaxNumberOfMessages: aws.Int64(ssvc.maxNumMessages),
			VisibilityTimeout:   aws.Int64(ssvc.timeout),
		},
	)
	if err != nil {
		return nil, err
	}

	return msgRes, nil
}

func (ssvc *SQSService) DeleteMessage(msg sqs.Message) error {
	_, err := ssvc.svc.DeleteMessage(
		&sqs.DeleteMessageInput{
			QueueUrl:      ssvc.queueURL,
			ReceiptHandle: msg.ReceiptHandle,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (ssvc *SQSService) ReadMessagesAndSaveOrder(mongoClient *mongo.Client) {
	orderRepository := repositories.NewOrderRepository(
		mongoClient,
		"btg_challenges",
		"orders",
	)
	saveOrderUseCase := usecases.NewSaveOrderUseCase(orderRepository)

	for {
		messages, err := ssvc.GetMessages()
		if err != nil {
			panic(err)
		}

		for _, msg := range messages.Messages {
			var order entity.Order

			err := json.Unmarshal([]byte(*msg.Body), &order)
			if err != nil {
				log.Fatalf("Error unmarshaling message to order: %s", err)
			}

			_, ucError := saveOrderUseCase.Execute(usecases.SaveOrderInput(order))
			if ucError.Message != "" {
				panic(ucError.Message)
			}

			err = ssvc.DeleteMessage(*msg)
			if err != nil {
				log.Fatalf("error to delete message from queue: %s", err)
			}
		}
	}
}
