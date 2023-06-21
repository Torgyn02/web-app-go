package main

import (
	"log"

	webapp "github.com/Torgyn/web-app"
	"github.com/Torgyn/web-app/pkg/handler"
	"github.com/Torgyn/web-app/pkg/repository"
	"github.com/Torgyn/web-app/pkg/service"
)

func main() {
	rep := repository.NewRepository()
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	srv := new(webapp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error dont : %s", err.Error())

	}
}
