package settings

import (
	"net/http"

	"github.com/dsummers91/go-bmdc/routes/templates"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	values := make(map[interface{}]interface{})

	templates.RenderTemplate(w, r, "settings", values)
}
