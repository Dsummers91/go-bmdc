package database

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

var Client *mongo.Client

func Establish_connection() {
	Client, _ = mongo.NewClient("mongodb://localhost:27017")
	collection := Client.Database("blackmendontcheat").Collection("members")
	ctx, cancel := newContext()
	Client.Connect(ctx)
	defer cancel()

	index := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "username", Value: bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	collection.Indexes().CreateOne(ctx, index, opts)
}

func Collection(table string) (*mongo.Collection, context.Context, context.CancelFunc) {
	context, cancel := newContext()
	Client.Connect(context)
	collection := Client.Database("blackmendontcheat").Collection(table)
	return collection, context, cancel
}

func newContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
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
