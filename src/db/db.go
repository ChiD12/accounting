package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

type CreateUser struct {
	Email string
	User  string
	Pass  string
}

func MakeDB() DB {

	ctx := context.TODO()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	return DB{client}
}
func (db DB) GetUser() {
	coll := db.client.Database("Accounting").Collection("User")

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "User", Value: "hi"}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", "test")
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func (db DB) CreateUser(createUser CreateUser) string {
	coll := db.client.Database("Accounting").Collection("User")
	doc := bson.D{{"User", "test"}, {"Pass", "123"}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("result.InsertedID: %v\n", result.InsertedID)
	return fmt.Sprintf("%v", result.InsertedID)
}
