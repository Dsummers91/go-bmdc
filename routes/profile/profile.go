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

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		//SHoudl error
	}

	oauthProfile := session.Values["profile"]
	oauthObject := oauthProfile.(map[string]interface{})
	oauth := oauthObject["sub"]

	user.Oauth = fmt.Sprintf("%v", oauth)
	data, _ := bson.Marshal(user)

	_, err = collection.ReplaceOne(context,
		bson.M{"oauth": oauth},
		data,
		options.Replace().SetUpsert(true),
	)

	if err != nil {
		// error
	}

	collection.FindOne(context, bson.M{"oauth": oauth}).Decode(&user)

	json.NewEncoder(w).Encode(user)
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user user.UserProfile
	var data ProfileData

	vars := mux.Vars(r)
	id := vars["id"]

	collection, context, cancel := database.Collection("members")
	defer cancel()

	collection.FindOne(context, bson.M{"username": id}).Decode(&user)

	data.User = user

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	renderTemplate(w, r, data)
}

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user user.UserProfile
	var data ProfileData

	collection, context, cancel := database.Collection("members")
	defer cancel()

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		//SHoudl error
	}

	oauthProfile := session.Values["profile"]
	oauthObject := oauthProfile.(map[string]interface{})
	oauth := oauthObject["sub"]

	collection.FindOne(context, bson.M{"oauth": oauth}).Decode(&user)

	data.User = user
	data.IsUser = true

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
