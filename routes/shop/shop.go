package shop

import (
	"net/http"

	"github.com/dsummers91/go-bmdc/routes/templates"
)

func GetStoreHandler(w http.ResponseWriter, r *http.Request) {
	var data templates.TemplateData
	templates.RenderTemplate(w, r, "shop", data)
}
