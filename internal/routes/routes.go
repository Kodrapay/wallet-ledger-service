package routes

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/wallet-ledger-service/internal/handlers"
	"github.com/kodra-pay/wallet-ledger-service/internal/repositories"
	"github.com/kodra-pay/wallet-ledger-service/internal/services"
)

func Register(app *fiber.App, serviceName string, db *sql.DB) {
	health := handlers.NewHealthHandler(serviceName)
	health.Register(app)

	walletRepo := repositories.NewPostgresWalletRepository(db)
	walletService := services.NewWalletService(walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	// API Group for wallets
	walletGroup := app.Group("/api/v1/wallets")
	walletGroup.Post("/", walletHandler.CreateWallet)
	walletGroup.Get("/:id", walletHandler.GetWalletByID)
	walletGroup.Get("/", walletHandler.GetWalletByUserIDAndCurrency) // Query params: user_id, currency
	walletGroup.Post("/:id/update-balance", walletHandler.UpdateWalletBalance)
	walletGroup.Get("/:id/ledger", walletHandler.GetWalletLedger)
}