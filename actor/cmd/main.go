package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/shakhboznorbekov/mytasks/golang_crud/actor/api"
	"github.com/shakhboznorbekov/mytasks/golang_crud/actor/config"
	"github.com/shakhboznorbekov/mytasks/golang_crud/actor/pkg/db"
)

func main() {

	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := gin.New()

	api.SetUpApi(r, db)

	log.Printf("Listening port %v...\n", cfg.HTTPPort)
	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
