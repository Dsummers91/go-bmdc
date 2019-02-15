package privacy

import (
	"net/http"

	"github.com/dsummers91/go-bmdc/routes/templates"
)

func GetPrivacyHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, "privacy", struct{}{})
}
