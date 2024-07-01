package entity

type Item struct {
	Product  string  `bson:"product" json:"product"`
	Quantity int     `bson:"quantity" json:"quantity"`
	Price    float64 `bson:"price" json:"price"`
}

type Order struct {
	OrderCode  int    `bson:"orderCode" json:"orderCode"`
	ClientCode int    `bson:"clientCode" json:"clientCode"`
	Items      []Item `bson:"items" json:"items"`
}

func NewOrder(orderCode, clientCode int, items []Item) *Order {
	order := &Order{
		OrderCode:  orderCode,
		ClientCode: clientCode,
		Items:      items,
	}

	return order
}
