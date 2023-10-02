package event

import "time"

type OrderCreatedEvent struct {
	Name    string
	Payload interface{}
}

func NewOrderCreatedEvent() *OrderCreatedEvent {
	return &OrderCreatedEvent{
		Name: "OrderCreatedEvent",
	}
}

func (e *OrderCreatedEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderCreatedEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderCreatedEvent) GetName() string {
	return e.Name
}

func (e *OrderCreatedEvent) GetDateTime() time.Time {
	return time.Now()
}
