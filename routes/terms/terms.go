package terms

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/user"
)

type ProfileData struct {
	User       user.UserProfile
	Terms      interface{}
	IsLoggedIn bool
	IsUser     bool
}

func GetTermsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r)
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	var data ProfileData
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(
		filepath.Join(cwd, "./routes/terms/terms.html"),
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
		data.Terms = session.Values["terms"]
		data.IsLoggedIn = true
	}

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
