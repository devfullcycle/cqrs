// Handle is a method that implements EventHandlerInterface
// persist data using mongodb
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

type ProductCreatedEventHandler struct {
	MongoDBConnection *mongo.Client
}

type ProductCreatedEventData struct {
	ID    string
	Name  string
	Price float64
}

func (h *ProductCreatedEventHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	// Unmarshal the event data into a struct
	var eventData ProductCreatedEventData
	payload, ok := event.GetPayload().([]byte)
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
	collection := h.MongoDBConnection.Database("ecommerce").Collection("products")

	// use id event as id mongodb
	_, err = collection.InsertOne(context.Background(),
		bson.M{
			"_id":   eventData.ID,
			"name":  eventData.Name,
			"price": eventData.Price,
		})
	if err != nil {
		log.Printf("Error inserting event data into MongoDB: %v", err)
		return
	}
}
