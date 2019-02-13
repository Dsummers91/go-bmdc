package contact

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsummers91/go-bmdc/routes/templates"
	"github.com/dsummers91/go-bmdc/user"
)

func GetContactHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, "contact", struct{}{})
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	var data user.UserProfile
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(
		filepath.Join(cwd, "./routes/contact/contact.html"),
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

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
