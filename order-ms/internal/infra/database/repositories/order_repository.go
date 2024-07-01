package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/paulozy/btg-challenge/order-ms/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(mongoClient *mongo.Client, database, collectionName string) *OrderRepository {
	collection := mongoClient.Database(database).Collection(collectionName)

	return &OrderRepository{
		collection,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(
		ctx,
		bson.D{
			{Key: "clientCode", Value: order.ClientCode},
			{Key: "orderCode", Value: order.OrderCode},
			{Key: "items", Value: order.Items},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetByClientCode(clientCode int) ([]entity.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	filter := bson.D{{Key: "clientCode", Value: clientCode}}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("porra", err)
		return nil, err
	}

	defer cur.Close(ctx)
	var orders []entity.Order

	for cur.Next(ctx) {
		var order entity.Order

		err := cur.Decode(&order)
		if err != nil {
			fmt.Println("porra2", err)
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) FindByOrderCode(clientCode, orderCode int) (*entity.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := r.collection.FindOne(
		ctx,
		bson.D{
			{Key: "clientCode", Value: clientCode},
			{Key: "orderCode", Value: orderCode},
		},
	)

	var order *entity.Order

	err := res.Decode(&order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
