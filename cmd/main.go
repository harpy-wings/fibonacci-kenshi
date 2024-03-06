package main

import (
	"log"

	"github.com/harpy-wings/fibonacci-kenshi/app"
)

// main runs the server to listen and handle requests.
func main() {
	appServer, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = appServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
