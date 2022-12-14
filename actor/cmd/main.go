package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/shakhboznorbekov/mytasks/30-dars/golang_crud-master/api"
	"github.com/shakhboznorbekov/mytasks/30-dars/golang_crud-master/config"
	"github.com/shakhboznorbekov/mytasks/30-dars/golang_crud-master/pkg/db"
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
