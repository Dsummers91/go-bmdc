package terms

import (
	"net/http"

	"github.com/dsummers91/go-bmdc/routes/templates"
)

func GetTermsHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, "terms", struct{}{})
}
