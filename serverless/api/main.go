package main

import (
	"log"

	"github.com/akrylysov/algnhsa"
	"github.com/andreaperizzato/go-config"
	"github.com/sorsby/gin-rating-api/api/routes"
	"github.com/sorsby/gin-rating-api/settings"
)

const pkg = "github.com/sorsby/gin-rating-api/serverless/api"

func main() {
	l := config.NewLoader(
		config.NewEnvSource(),
	)
	var s settings.APISettings
	err := l.Load(&s)
	if err != nil {
		log.Fatal(err)
	}

	h, err := routes.Create(s)
	if err != nil {
		log.Fatal(err)
	}
	algnhsa.ListenAndServe(h, nil)
}
