package handler_test

import (
	"context"
	"sync"
	"testing"

	"github.com/devfullcycle/cqrs/internal/command/event"
	"github.com/devfullcycle/cqrs/internal/command/event/handler"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestOrderCreatedEventHandler_Handle(t *testing.T) {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions.Auth = &options.Credential{
		Username: "root",
		Password: "example",
	}
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}
	mongoClient.Database("ecommerce").Collection("orders").Drop(context.Background())
	defer mongoClient.Disconnect(context.Background())

	// Set up test data
	orderID := "123"
	total := 100.0
	orderItems := []handler.OrderItem{
		{
			ID:       "456",
			Product:  handler.Product{ID: "789", Name: "Test Product", Price: 50.0},
			Quantity: 2,
			Price:    100.0,
		},
	}

	// Insert test data into MongoDB
	collection := mongoClient.Database("ecommerce").Collection("orders")
	_, err = collection.InsertOne(context.Background(),
		bson.M{
			"_id":         orderID,
			"total":       total,
			"order_items": orderItems,
		})
	if err != nil {
		t.Fatalf("Error inserting test data into MongoDB: %v", err)
	}

	// Set up event and handler
	event := event.NewOrderCreatedEvent()
	eventHandler := &handler.OrderCreatedEventHandler{
		MongoDBConnection: mongoClient,
	}

	// Call Handle method
	var wg sync.WaitGroup
	wg.Add(1)
	eventHandler.Handle(event, &wg)
	wg.Wait()

	// Check that test data was inserted into MongoDB
	var result bson.M
	err = collection.FindOne(context.Background(), bson.M{"_id": orderID}).Decode(&result)
	if err != nil {
		t.Fatalf("Error finding test data in MongoDB: %v", err)
	}
	assert.Equal(t, orderID, result["_id"])
	assert.Equal(t, total, result["total"])
}
