package main

import (
	"log"

	"github.com/andrsj/go_anime_crud/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal("App can't be run")
	}

	log.Fatal(app.Run())
}
