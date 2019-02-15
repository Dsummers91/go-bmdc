package store

import (
	"net/http"

	"github.com/dsummers91/go-bmdc/routes/templates"
)

func GetStoreHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, "store", struct{}{})
}
