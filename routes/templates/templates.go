package templates

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsummers91/go-bmdc/app"
)

type TemplateData struct {
	Auth0ClientId     string
	Auth0ClientSecret string
	Auth0Domain       string
	Auth0CallbackURL  template.URL
	Profile           interface{}
	IsLoggedIn        bool
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(
		filepath.Join(cwd, "./routes/"+tmpl+"/"+tmpl+".html"),
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
		if profile, ok := session.Values["profile"]; ok {
			data = TemplateData{
				Profile:           profile,
				IsLoggedIn:        true,
				Auth0Domain:       os.Getenv("AUTH0_DOMAIN"),
				Auth0CallbackURL:  template.URL(os.Getenv("AUTH0_CALLBACK_URL")),
				Auth0ClientId:     os.Getenv("AUTH0_CLIENT_ID"),
				Auth0ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
			}
		}
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
