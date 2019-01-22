package user

import (
	"io"
	"net/http"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/routes/templates"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//user := vars["user"]
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	values := make(map[interface{}]interface{})
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	values["profile"] = session.Values["profile"]

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	templates.RenderTemplate(w, r, "profile", values)
}
