package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/wallet-ledger-service/internal/config"
	"github.com/kodra-pay/wallet-ledger-service/internal/handlers"
	"github.com/kodra-pay/wallet-ledger-service/internal/repositories"
	"github.com/kodra-pay/wallet-ledger-service/internal/services"
)

func Register(app *fiber.App, cfg config.Config, repo *repositories.LedgerRepository) {
	health := handlers.NewHealthHandler(cfg.ServiceName)
	health.Register(app)

	svc := services.NewLedgerService(repo)
	h := handlers.NewLedgerHandler(svc)
	api := app.Group("/")
	api.Get("balances/:merchantId", h.GetBalance)
	api.Post("ledger/entries", h.CreateEntry)
}
