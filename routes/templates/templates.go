package templates

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

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
	//session, err := app.Store.Get(r, "auth-session")
	//if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	//fmt.Println(data)
	//if _, ok := session.Values["profile"]; !ok {
	//data["profile"] = session.Values["profile"]
	//data["isLoggedIn"] = true
	//}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
