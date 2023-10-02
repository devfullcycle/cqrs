package query

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindProductQuery struct {
	mongoClient *mongo.Client
}

type Product struct {
	ID    string  `bson:"_id" json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewFindProductQuery(mongoClient *mongo.Client) *FindProductQuery {
	return &FindProductQuery{mongoClient: mongoClient}
}

// find all products
func (q *FindProductQuery) FindAll() ([]*Product, error) {
	products := []*Product{}

	collection := q.mongoClient.Database("ecommerce").Collection("products")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		product := &Product{}
		if err := cursor.Decode(product); err != nil {
			fmt.Println(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// find product by id
func (q *FindProductQuery) FindByID(id string) (*Product, error) {
	product := &Product{}
	collection := q.mongoClient.Database("ecommerce").Collection("products")
	f := collection.FindOne(context.TODO(), bson.M{"_id": id})
	if err := f.Decode(product); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return product, nil
}
