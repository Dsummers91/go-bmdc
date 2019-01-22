package settings

import (
	"net/http"

	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/templates"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	values := make(map[interface{}]interface{})
	collection, context := database.Collection("members")
	users := database.GetCollection(collection, context)
	values["users"] = users

	templates.RenderTemplate(w, r, "settings", values)
}
