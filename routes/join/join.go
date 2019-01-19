package join

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/templates"
	"github.com/mongodb/mongo-go-driver/bson"
)

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GetJoinHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Auth0ClientId     string
		Auth0ClientSecret string
		Auth0Domain       string
		Auth0CallbackURL  template.URL
	}{
		os.Getenv("AUTH0_CLIENT_ID"),
		os.Getenv("AUTH0_CLIENT_SECRET"),
		os.Getenv("AUTH0_DOMAIN"),
		template.URL(os.Getenv("AUTH0_CALLBACK_URL")),
	}

	templates.RenderTemplate(w, "join", data)
}

func PostJoinHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	collection, context := database.Collection("users")
	res, _ := collection.InsertOne(context, bson.M{"name": user.Name, "email": user.Email})

	json.NewEncoder(w).Encode(res)
}
