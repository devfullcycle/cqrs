// BEGIN: abpxx6d04wxr
package handler_test

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"testing"

	"github.com/devfullcycle/cqrs/internal/command/event"
	"github.com/devfullcycle/cqrs/internal/command/event/handler"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestProductCreatedEventHandler_Handle(t *testing.T) {
	// Set up MongoDB connection
	// use login and password
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions.Auth = &options.Credential{
		Username: "root",
		Password: "example",
	}
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	defer mongoClient.Disconnect(context.Background())
	// clear collection
	// defer mongoClient.Database("ecommerce").Collection("products").Drop(context.Background())

	// Create test event data
	eventData := handler.ProductCreatedEventData{
		ID:    "123456789012",
		Name:  "Test Product 2",
		Price: 9.99,
	}
	payload, err := json.Marshal(eventData)
	if err != nil {
		log.Fatalf("Error marshalling test event data: %v", err)
	}

	// Create test event
	event := event.NewProductCreated()
	event.SetPayload(payload)

	// Create test handler
	eventHandler := handler.ProductCreatedEventHandler{
		MongoDBConnection: mongoClient,
	}

	// Create wait group for handling the event
	var wg sync.WaitGroup
	wg.Add(1)

	// Handle the event
	eventHandler.Handle(event, &wg)

	// Wait for the handler to finish
	wg.Wait()

	// Check that the event data was inserted into MongoDB
	collection := mongoClient.Database("ecommerce").Collection("products")
	result := collection.FindOne(context.Background(), bson.M{"_id": eventData.ID})
	assert.NoError(t, result.Err())
	var insertedData map[string]interface{}
	err = result.Decode(&insertedData)
	assert.NoError(t, err)
	assert.Equal(t, eventData.ID, insertedData["_id"])
	assert.Equal(t, eventData.Name, insertedData["name"])
	assert.Equal(t, eventData.Price, insertedData["price"])
}
