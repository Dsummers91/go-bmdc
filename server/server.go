package server

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/callback"
	"github.com/dsummers91/go-bmdc/routes/home"
	"github.com/dsummers91/go-bmdc/routes/image"
	"github.com/dsummers91/go-bmdc/routes/join"
	"github.com/dsummers91/go-bmdc/routes/login"
	"github.com/dsummers91/go-bmdc/routes/logout"
	"github.com/dsummers91/go-bmdc/routes/middlewares"
	"github.com/dsummers91/go-bmdc/routes/profile"
	"github.com/dsummers91/go-bmdc/routes/settings"
	"github.com/gorilla/mux"
)

func StartServer() {
	database.Establish_connection()
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/join", join.GetJoinHandler).Methods("GET")
	r.HandleFunc("/join", join.PostJoinHandler).Methods("POST")
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/profile/settings", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(settings.SettingsHandler)),
	))
	r.HandleFunc("/profile", profile.PostProfileHandler).Methods("POST")
	r.HandleFunc("/profile/image", image.PostImageHandler).Methods("POST")
	r.Handle("/profile", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(profile.GetUserProfileHandler)),
	)).Methods("GET")
	r.HandleFunc("/profile/{id}", profile.GetProfileHandler).Methods("GET")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)

	log.Print("Server listening on http://localhost:" + os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}
