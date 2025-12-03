package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/wallet-ledger-service/internal/dto"
	"github.com/kodra-pay/wallet-ledger-service/internal/services"
)

type LedgerHandler struct {
	svc *services.LedgerService
}

func NewLedgerHandler(svc *services.LedgerService) *LedgerHandler { return &LedgerHandler{svc: svc} }

func (h *LedgerHandler) GetBalance(c *fiber.Ctx) error {
	merchantID := c.Params("merchantId")
	return c.JSON(h.svc.GetBalance(c.Context(), merchantID))
}

func (h *LedgerHandler) CreateEntry(c *fiber.Ctx) error {
	var req dto.LedgerEntryRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	return c.JSON(h.svc.CreateEntry(c.Context(), req))
}
