package server

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/callback"
	"github.com/dsummers91/go-bmdc/routes/contact"
	"github.com/dsummers91/go-bmdc/routes/image"
	"github.com/dsummers91/go-bmdc/routes/join"
	"github.com/dsummers91/go-bmdc/routes/login"
	"github.com/dsummers91/go-bmdc/routes/logout"
	"github.com/dsummers91/go-bmdc/routes/middlewares"
	"github.com/dsummers91/go-bmdc/routes/profile"
	"github.com/dsummers91/go-bmdc/routes/templates"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func InitServer() Server {
	database.Establish_connection()
	r := mux.NewRouter()

	r.HandleFunc("/", templates.RenderTemplate).Name("home")
	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/join", templates.RenderTemplate).Methods("GET").Name("join")
	r.HandleFunc("/join", join.PostJoinHandler).Methods("POST")
	r.HandleFunc("/login", login.LoginHandler).Name("login")
	r.HandleFunc("/logout", logout.LogoutHandler).Name("logout")
	r.HandleFunc("/terms", templates.RenderTemplate).Methods("GET").Name("terms")
	r.HandleFunc("/contact", templates.RenderTemplate).Methods("GET").Name("contact")
	r.HandleFunc("/contact", contact.PostContactHandler).Methods("POST")
	r.HandleFunc("/privacy", templates.RenderTemplate).Methods("GET").Name("privacy")
	r.HandleFunc("/shop", templates.RenderTemplate).Methods("GET").Name("shop")
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/profile/settings", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(templates.RenderTemplate)),
	)).Name("settings")
	r.HandleFunc("/profile", profile.PostProfileHandler).Methods("POST")
	r.HandleFunc("/profile/image", image.PostImageHandler).Methods("POST")
	r.Handle("/profile", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(profile.GetUserProfileHandler)),
	)).Methods("GET").Name("profile")
	r.HandleFunc("/profile/{id}", profile.GetProfileHandler).Methods("GET")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	return Server{r}
}

func (a *Server) RunServer() {
	r := a.Router
	http.Handle("/", r)

	log.Print("Server listening on http://localhost:" + os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
