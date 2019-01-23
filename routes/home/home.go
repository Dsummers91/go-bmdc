package home

import (
	"html/template"
	"net/http"
	"os"

	"github.com/dsummers91/go-bmdc/routes/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Auth0ClientId     string
		Auth0ClientSecret string
		Auth0Domain       string
		Auth0CallbackURL  template.URL
		IsLoggedIn        bool
	}{
		os.Getenv("AUTH0_CLIENT_ID"),
		os.Getenv("AUTH0_CLIENT_SECRET"),
		os.Getenv("AUTH0_DOMAIN"),
		template.URL(os.Getenv("AUTH0_CALLBACK_URL")),
		false,
	}

	templates.RenderTemplate(w, r, "home", data)
}
