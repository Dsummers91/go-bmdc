package privacy

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dsummers91/go-bmdc/user"
)

func GetPrivacyHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r)
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	var data user.UserProfile
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(
		filepath.Join(cwd, "./routes/privacy/privacy.html"),
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
