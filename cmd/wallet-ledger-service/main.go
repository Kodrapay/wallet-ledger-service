package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/wallet-ledger-service/internal/config"
	"github.com/kodra-pay/wallet-ledger-service/internal/middleware"
	"github.com/kodra-pay/wallet-ledger-service/internal/repositories"
	"github.com/kodra-pay/wallet-ledger-service/internal/routes"
)

func main() {
	cfg := config.Load("wallet-ledger-service", "7007")

	// Initialize database
	db, err := repositories.InitDB(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	app := fiber.New()
	app.Use(middleware.RequestID())

	// Pass the database instance to the routes registration
	routes.Register(app, cfg.ServiceName, db)

	log.Printf("%s listening on :%s", cfg.ServiceName, cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
