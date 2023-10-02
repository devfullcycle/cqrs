package command

import (
	"github.com/devfullcycle/cqrs/internal/command/domain"
	"github.com/devfullcycle/cqrs/pkg/events"
)

type PlaceOrderInput struct {
	OrderItems []string
}

type PlaceOrderCommand struct {
	Repository        domain.OrderRepository
	ProductRepository domain.ProductRepository
	EventDispatcher   events.EventDispatcherInterface
	OrderCreatedEvent events.EventInterface
}

func (c *PlaceOrderCommand) Handle(input *PlaceOrderInput) {
	order := domain.NewOrder()
	for _, itemID := range input.OrderItems {
		product := c.ProductRepository.FindByID(itemID)
		orderItem := domain.NewOrderItem(product, 1, product.Price)
		order.AddOrderItem(orderItem)
	}
	c.Repository.Save(order)
	c.OrderCreatedEvent.SetPayload(order)
	c.EventDispatcher.Dispatch(c.OrderCreatedEvent)
}
