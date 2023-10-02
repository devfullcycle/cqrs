package event

import "time"

type ProductCreated struct {
	Name    string
	Payload interface{}
}

func NewProductCreated() *ProductCreated {
	return &ProductCreated{
		Name: "ProductCreated",
	}
}

func (e *ProductCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ProductCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *ProductCreated) GetName() string {
	return e.Name
}

func (e *ProductCreated) GetDateTime() time.Time {
	return time.Now()
}
