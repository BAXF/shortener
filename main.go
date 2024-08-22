package main

import (
	"fmt"
	"github.com/BAXF/shortener/config"
	"github.com/BAXF/shortener/models"
	"github.com/BAXF/shortener/routes"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, cfg.PostgresPort)
	db, err := models.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	redisClient := models.ConnectRedis(cfg.RedisAddr)

	r := routes.SetupRouter(db, redisClient)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
