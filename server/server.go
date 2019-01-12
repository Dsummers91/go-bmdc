package server

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/dsummers91/go-bmdc/routes/callback"
	"github.com/dsummers91/go-bmdc/routes/home"
	"github.com/dsummers91/go-bmdc/routes/login"
	"github.com/dsummers91/go-bmdc/routes/logout"
	"github.com/dsummers91/go-bmdc/routes/middlewares"
	"github.com/dsummers91/go-bmdc/routes/user"
	"github.com/gorilla/mux"
)

func StartServer() {
	database.Establish_connection()
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	log.Print("Server listening on http://localhost:3000/")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
