package command

import (
	"github.com/devfullcycle/cqrs/internal/command/domain"
	"github.com/devfullcycle/cqrs/pkg/events"
)

type CreateProductInput struct {
	Name  string
	Price float64
}

type CreateProductEventPayload struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductCommand struct {
	Repository          domain.ProductRepository
	EventDispatcher     events.EventDispatcherInterface
	ProductCreatedEvent events.EventInterface
}

func (c *CreateProductCommand) Handle(input *CreateProductInput) {
	product := domain.NewProduct(input.Name, input.Price)
	c.Repository.Save(product)
	c.ProductCreatedEvent.SetPayload(
		CreateProductEventPayload{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		},
	)
	c.EventDispatcher.Dispatch(c.ProductCreatedEvent)
}
