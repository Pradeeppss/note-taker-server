package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance()*mongo.Client{
	MongoDb:= "mongodb+srv://PradeepPS:<password>@myproject1.ntdejkm.mongodb.net/?retryWrites=true&w=majority&appName=myProject1"
	fmt.Print(MongoDb)
	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)

	client,err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil{
		log.Fatal(err)
	}

	defer cancel()

	fmt.Println("connected to mongoDB")
	return client
}

var Client*mongo.Client = DbInstance() 

func OpenCollection(client *mongo.Client ,collectionName string)*mongo.Collection{
	var collection*mongo.Collection = client.Database("MyNotesData").Collection(collectionName) 

	return collection
}