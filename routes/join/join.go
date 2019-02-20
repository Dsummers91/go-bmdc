package join

import (
	"encoding/json"
	"net/http"

	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/user"
	"github.com/mongodb/mongo-go-driver/bson"
)

func PostJoinHandler(w http.ResponseWriter, r *http.Request) {
	var user user.User
	json.NewDecoder(r.Body).Decode(&user)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	collection, context, cancel := database.Collection("users")
	defer cancel()
	res, _ := collection.InsertOne(context, bson.M{"name": user.Name, "email": user.Email})

	json.NewEncoder(w).Encode(res)
}
