package database

import 
(
	"fmt"
	"log"
	"time"
	"os"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client=DBinstance()

func DBinstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err !=nil{
		log.Fatal("Error loading in .env file")
	}

	MongoDb := os.Getenv("MONGODB URL")
	client , err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err !=nil{
		log.Fatal(err)
	}
	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
	
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("Authenication").Collection("signup")
	return collection
}

