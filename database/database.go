package database

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client *mongo.Client
var Context context.Context

func Establish_connection() {
	Client, _ = mongo.NewClient("mongodb://localhost:27017")
	Context, _ = context.WithTimeout(context.Background(), 10*time.Second)

	Client.Connect(Context)
}

func Collection(table string) *mongo.Collection {
	collection := Client.Database("blackmendontcheat").Collection(table)
	return collection
}

func GetCollection(collection *mongo.Collection) []interface{} {
	cur, err := collection.Find(Context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(Context)
	var results []interface{}
	for cur.Next(Context) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}
