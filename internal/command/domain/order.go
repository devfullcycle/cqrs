package domain

import "github.com/google/uuid"

type OrderRepository interface {
	Save(order *Order)
}

type ProductRepository interface {
	Save(product Product)
	FindByID(id string) Product
}

type Order struct {
	ID         string
	Total      float64
	OrderItems []OrderItem
}

func NewOrder() *Order {
	return &Order{
		ID: uuid.New().String(),
	}
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) Product {
	return Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

type OrderItem struct {
	ID       string
	Product  Product
	Quantity int
	Price    float64
}

func NewOrderItem(product Product, quantity int, price float64) OrderItem {
	return OrderItem{
		ID:       uuid.New().String(),
		Product:  product,
		Quantity: quantity,
		Price:    price,
	}
}

func (o *Order) AddOrderItem(orderItem OrderItem) {
	o.OrderItems = append(o.OrderItems, orderItem)
}

func (o *Order) RemoveOrderItem(orderItem OrderItem) {
	for i, item := range o.OrderItems {
		if item.ID == orderItem.ID {
			o.OrderItems = append(o.OrderItems[:i], o.OrderItems[i+1:]...)
		}
	}
}

func (o *Order) GetOrderItems() []OrderItem {
	return o.OrderItems
}

func (o *Order) GetTotal() float64 {
	var total float64
	for _, item := range o.OrderItems {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
