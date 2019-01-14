package settings

import (
	"log"
	"net/http"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/templates"
	"github.com/mongodb/mongo-go-driver/bson"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
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
	templates.RenderTemplate(w, "settings", values)
}