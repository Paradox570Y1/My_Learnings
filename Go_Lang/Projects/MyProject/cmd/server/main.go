package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"MyProject/internal/config"
	"MyProject/internal/database"
	"MyProject/internal/middleware"
	"MyProject/internal/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal(err)
	}
	
	r := gin.Default()
	r.Use(middleware.Logger())
	routes.RegisterRoutes(r, db)
	r.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}