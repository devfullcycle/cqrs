package handler

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/devfullcycle/cqrs/pkg/events"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderCreatedEventHandler struct {
	MongoDBConnection *mongo.Client
}

type OrderCreatedEventData struct {
	ID         string
	Total      float64
	OrderItems []OrderItem
}

type OrderItem struct {
	ID       string
	Product  Product
	Quantity int
	Price    float64
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

func (h *OrderCreatedEventHandler) Handle(e events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	// Unmarshal the event data into a struct
	var eventData OrderCreatedEventData
	payload, ok := e.GetPayload().([]byte)
	if !ok {
		log.Printf("Error getting event payload as []byte")
		return
	}
	err := json.Unmarshal(payload, &eventData)
	if err != nil {
		log.Printf("Error unmarshalling event data: %v", err)
		return
	}

	// Insert the event data into MongoDB
	collection := h.MongoDBConnection.Database("ecommerce").Collection("orders")

	// use id event as id mongodb
	// iterate over order items and insert into mongodb
	_, err = collection.InsertOne(context.Background(),
		bson.M{
			"_id":         eventData.ID,
			"total":       eventData.Total,
			"order_items": eventData.OrderItems,
		})
	if err != nil {
		log.Printf("Error inserting event data into MongoDB: %v", err)
		return
	}
}
