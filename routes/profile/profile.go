package profile

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/user"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type ProfileData struct {
	User       user.UserProfile
	Profile    interface{}
	IsLoggedIn bool
	IsUser     bool
}

func PostProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user user.UserProfile
	json.NewDecoder(r.Body).Decode(&user)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	collection, context, cancel := database.Collection("members")
	defer cancel()

	fmt.Println(user.Username)

	res, _ := collection.ReplaceOne(context,
		bson.M{"city": user.City},
		bson.M{"city": user.City, "name": user.Name, "email": user.Email, "username": user.Username},
		options.Replace().SetUpsert(true),
	)

	json.NewEncoder(w).Encode(res)
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user user.UserProfile
	var data ProfileData

	vars := mux.Vars(r)
	id := vars["id"]

	collection, context, cancel := database.Collection("members")
	defer cancel()

	collection.FindOne(context, bson.M{"city": id}).Decode(&user)

	data.User = user
	data.IsUser = true

	fmt.Println(user)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	renderTemplate(w, r, data)
}

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user user.UserProfile
	var data ProfileData

	vars := mux.Vars(r)
	id := vars["id"]

	collection, context, cancel := database.Collection("members")
	defer cancel()

	collection.FindOne(context, bson.M{"city": id}).Decode(&user)

	data.User = user

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	renderTemplate(w, r, data)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, data ProfileData) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(
		filepath.Join(cwd, "./routes/profile/profile.html"),
		filepath.Join(cwd, "./routes/templates/header.html"),
		filepath.Join(cwd, "./routes/templates/navbar.html"),
		filepath.Join(cwd, "./routes/templates/footer.html"),
		filepath.Join(cwd, "./routes/templates/store.html"),
		filepath.Join(cwd, "./routes/templates/signin.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := app.Store.Get(r, "auth-session")
	if err == nil {
		data.Profile = session.Values["profile"]
		data.IsLoggedIn = true
	}

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
