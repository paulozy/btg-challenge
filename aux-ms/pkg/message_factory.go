package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"math/rand"
)

type Item struct {
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Order struct {
	OrderCode  int    `json:"orderCode"`
	ClientCode int    `json:"clientCode"`
	Items      []Item `json:"items"`
}

func CreateMessages(clientQty, ordersQtyPerClient int) []string {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator

	products := []Item{
		{Product: "pencil", Quantity: 100, Price: 1.10},
		{Product: "notebook", Quantity: 10, Price: 1.00},
		{Product: "eraser", Quantity: 50, Price: 0.50},
		{Product: "pen", Quantity: 30, Price: 1.20},
		{Product: "marker", Quantity: 20, Price: 1.30},
	}

	messages := make([]string, 0, clientQty*ordersQtyPerClient)

	for i := 1; i <= clientQty; i++ {
		for j := 1; j <= ordersQtyPerClient; j++ {
			order := Order{
				OrderCode:  time.Now().Nanosecond() + j,
				ClientCode: i,
			}
			// Randomly pick 1 to 3 products for each order
			numItems := rand.Intn(3) + 1
			var orderItems []Item
			for k := 0; k < numItems; k++ {
				product := products[rand.Intn(len(products))]
				// Randomly adjust quantity and ensure it is not zero
				product.Quantity = rand.Intn(100) + 1
				orderItems = append(orderItems, product)
			}
			order.Items = orderItems

			// Convert order struct to JSON
			jsonOrder, err := json.Marshal(order)
			if err != nil {
				fmt.Println("Error marshalling order:", err)
				continue
			}
			messages = append(messages, string(jsonOrder))
		}
	}

	return messages
}
