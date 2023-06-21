package main

import (
	"log"

	webapp "github.com/Torgyn/web-app"
	"github.com/Torgyn/web-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(webapp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error dont : %s", err.Error())

	}
}
