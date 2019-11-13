package main

import (
	"os"

	"github.com/akrylysov/algnhsa"
	"github.com/sorsby/gin-rating-api/api/routes"
)

const pkg = "github.com/sorsby/gin-rating-api/serverless/api"

func main() {
	h, err := routes.Create()
	if err != nil {
		os.Exit(-1)
	}
	algnhsa.ListenAndServe(h, nil)
}
