package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/user"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

func PostProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user user.UserProfile
	json.NewDecoder(r.Body).Decode(&user)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	collection, context, cancel := database.Collection("members")
	defer cancel()

	session, _ := app.Store.Get(r, "auth-session")
	fmt.Println(session.Values["profile"])

	res, _ := collection.ReplaceOne(context,
		bson.M{"city": user.City},
		bson.M{"city": user.City, "name": user.Name, "email": user.Email},
		options.Replace().SetUpsert(true),
	)

	json.NewEncoder(w).Encode(res)
}
