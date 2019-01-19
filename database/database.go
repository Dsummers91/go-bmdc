package database

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client *mongo.Client

func Establish_connection() {
	Client, _ = mongo.NewClient("mongodb://localhost:27017")
}

func Collection(table string) (*mongo.Collection, context.Context) {
	context := newContext()
	Client.Connect(context)
	collection := Client.Database("blackmendontcheat").Collection(table)
	return collection, context
}

func newContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func GetCollection(collection *mongo.Collection, ctx context.Context) []interface{} {
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var results []interface{}
	for cur.Next(ctx) {
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
