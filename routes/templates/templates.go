package templates

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/user"
	"github.com/mongodb/mongo-go-driver/bson"
)

type TemplateData struct {
	Auth0ClientId     string
	Auth0ClientSecret string
	Auth0Domain       string
	Auth0CallbackURL  template.URL
	Profile           interface{}
	IsLoggedIn        bool
	User              user.UserProfile
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
	var user user.UserProfile

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
			collection, context, cancel := database.Collection("members")
			defer cancel()

			userProfile := profile.(map[string]interface{})
			oauth := userProfile["sub"]

			collection.FindOne(context, bson.M{"oauth": oauth}).Decode(&user)

			data = TemplateData{
				Profile:           profile,
				IsLoggedIn:        true,
				Auth0Domain:       os.Getenv("AUTH0_DOMAIN"),
				Auth0CallbackURL:  template.URL(os.Getenv("AUTH0_CALLBACK_URL")),
				Auth0ClientId:     os.Getenv("AUTH0_CLIENT_ID"),
				Auth0ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
				User:              user,
			}
		}
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
