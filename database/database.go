package database

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client *mongo.Client
var Context context.Context

func Establish_connection() {
	Client, _ = mongo.NewClient("mongodb://localhost:27017")
	Context, _ = context.WithTimeout(context.Background(), 10*time.Second)

	Client.Connect(Context)
}
