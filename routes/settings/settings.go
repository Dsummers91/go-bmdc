package settings

import (
	"net/http"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/templates"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	values := make(map[interface{}]interface{})
	collection := database.Collection("users")
	users := database.GetCollection(collection)
	values["users"] = users

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	values["profile"] = session.Values["profile"]

	templates.RenderTemplate(w, "settings", values)
}
