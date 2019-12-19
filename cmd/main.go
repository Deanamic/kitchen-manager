package main

import (
	// "context"
	"fmt"
	// "github.com/Deanamic/kitchen-helper/pkg/domain"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Deanamic/kitchen-helper/pkg/storage"
	web "github.com/Deanamic/kitchen-helper/pkg/http"
	"log"
	"net/http"
)

func main() {
	repo := storage.New()
	s := web.New(repo)
	fmt.Println("Listening on port 9999")
	log.Fatal(http.ListenAndServe("127.0.0.1:9999", s.Router))

	/*
	ingredient := domainj.Ingredient{
		Name:     "potato",
		Quantity: 10,
	}
	recipe := domain.Recipe{
		Name:         "potato salad",
		Ingredients:  []domain.Ingredient{ingredient},
		Instructions: []string{"peel the potato", "eat the potato"},
	}
	fmt.Println(recipe)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("test").Collection("test")
	insertResult, err := collection.InsertOne(context.TODO(), recipe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	var result domain.Recipe
	filter := bson.D{{"name", "potato salad"}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	fmt.Printf("Found a single document: %+v\n", result.Ingredients)
*/
}
