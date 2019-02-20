package main

import (
	"log"

	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	app.Init()
	server := server.InitServer()
	server.RunServer()
}
