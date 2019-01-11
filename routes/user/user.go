package user

import (
	"log"
	"net/http"

	templates ".."
	"../../app"
	"../../database"
	"github.com/mongodb/mongo-go-driver/bson"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	values := make(map[interface{}]interface{})
	database.Establish_connection()
	client := database.Client
	ctx := database.Context
	collection := client.Database("blackmendontcheat").Collection("members")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var users []interface{}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, result)
	}
	values["users"] = users
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	values["profile"] = session.Values["profile"]
	templates.RenderTemplate(w, "user", values)
}
