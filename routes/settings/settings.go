package settings

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/profile"
	"github.com/dsummers91/go-bmdc/user"
	"github.com/fatih/structs"
	"github.com/mongodb/mongo-go-driver/bson"
)

type SettingsData struct {
	profile.ProfileData
	Fields []Field
}

type Field struct {
	Title string
	Name  string
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	var data SettingsData
	data.Fields = []Field{
		Field{Title: "Username", Name: "username"},
		Field{Title: "Name", Name: "name"},
		Field{Title: "Partner", Name: "partner"},
		Field{Title: "PartnerFacebook", Name: "partnerFacebook"},
		Field{Title: "PartnerTwitter", Name: "partnerTwitter"},
		Field{Title: "PartnerInstagram", Name: "partnerInstagram"},
		Field{Title: "Location", Name: "location"},
	}

	renderTemplate(w, r, data)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, data SettingsData) {
	var user user.UserProfile

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

	cwd, _ := os.Getwd()
	t, err := template.New("settings.html").Funcs(template.FuncMap{
		"dict":     dict,
		"getField": getField,
	}).ParseFiles(
		filepath.Join(cwd, "./routes/settings/settings.html"),
		filepath.Join(cwd, "./routes/templates/header.html"),
		filepath.Join(cwd, "./routes/templates/navbar.html"),
		filepath.Join(cwd, "./routes/templates/footer.html"),
		filepath.Join(cwd, "./routes/templates/store.html"),
		filepath.Join(cwd, "./routes/templates/signin.html"),
		filepath.Join(cwd, "./routes/templates/input_field.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err = app.Store.Get(r, "auth-session")
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

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func getField(userStruct user.UserProfile, field string) interface{} {
	user := structs.Map(userStruct)
	return user[field]
}
