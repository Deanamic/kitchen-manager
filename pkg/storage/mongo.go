package storage

import (
	"context"
	"github.com/Deanamic/kitchen-helper/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"fmt"
)

/*
MongoRepository TODO
*/
type MongoRepository struct {
	client *mongo.Client
}

const (
	databaseName      = "Kitchen-helper"
	usersCollection   = "Users"
	recipesCollection = "Recipes"
)

/*
New TODO
*/
func New() MongoRepository {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return MongoRepository{client: client}
}

/*
FindUser TODO
*/
func (r MongoRepository) FindUser(id string) (*domain.User, error) {
	collection := r.client.Database(databaseName).Collection(usersCollection)
	var user domain.User
	err := collection.FindOne(context.Background(), bson.M{"username": id}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return &user, nil
}

/*
CreateUser TODO
*/
func (r MongoRepository) CreateUser(user domain.User) error {
	collection := r.client.Database(databaseName).Collection(usersCollection)
	_, err := collection.InsertOne(context.TODO(), user)
	fmt.Println(user)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

/*
UpdateUser TODO
*/
func (r MongoRepository) UpdateUser(user domain.User) error {
	collection := r.client.Database(databaseName).Collection(usersCollection)
	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
	_, err := collection.ReplaceOne(context.TODO(), filter, user)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

/*
CreateRecipe domain.Recipe
TODO
*/
func (r MongoRepository) CreateRecipe(recipe domain.Recipe) error {
	collection := r.client.Database(databaseName).Collection(recipesCollection)
	_, err := collection.InsertOne(context.TODO(), recipe)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

/*
FindRecipe TODO
*/
func (r MongoRepository) FindRecipe(id string) (*domain.Recipe, error) {
	collection := r.client.Database(databaseName).Collection(recipesCollection)

	var recipe domain.Recipe
	err := collection.FindOne(context.Background(), bson.M{"recipename": id}).Decode(&recipe)
	if err != nil {
		log.Fatal(err)
	}
	return &recipe, nil
}

/*
FindRecipes TODO
*/
func (r MongoRepository) FindRecipes(restrictions domain.Restrictions) ([]*domain.Recipe, error) {
	collection := r.client.Database(databaseName).Collection(recipesCollection)
	filter := bson.D{}
	//TODO: Build this filter
	findOptions := options.Find()
	var recipes []*domain.Recipe
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var elem domain.Recipe
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		recipes = append(recipes, &elem)
	}
	return recipes, nil
}
