package user

import (
	"io"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//user := vars["user"]
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
