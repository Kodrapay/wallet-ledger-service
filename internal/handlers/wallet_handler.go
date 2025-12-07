package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/wallet-ledger-service/internal/dto"
	"github.com/kodra-pay/wallet-ledger-service/internal/services"
)

type WalletHandler struct {
	svc *services.WalletService
}

func NewWalletHandler(svc *services.WalletService) *WalletHandler {
	return &WalletHandler{svc: svc}
}

// CreateWallet handles requests to create a new wallet
func (h *WalletHandler) CreateWallet(c *fiber.Ctx) error {
	var req dto.CreateWalletRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if req.UserID == "" || req.Currency == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user_id and currency are required")
	}

	resp, err := h.svc.CreateWallet(c.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetWalletByID handles requests to get a wallet by its ID
func (h *WalletHandler) GetWalletByID(c *fiber.Ctx) error {
	walletID := c.Params("id")
	if walletID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "wallet ID is required")
	}

	resp, err := h.svc.GetWalletByID(c.Context(), walletID)
	if err != nil {
		if errors.Is(err, errors.New("wallet not found")) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

// GetWalletByUserIDAndCurrency handles requests to get a wallet by user ID and currency
func (h *WalletHandler) GetWalletByUserIDAndCurrency(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	currency := c.Query("currency")
	if userID == "" || currency == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user_id and currency are required query parameters")
	}

	resp, err := h.svc.GetWalletByUserIDAndCurrency(c.Context(), userID, currency)
	if err != nil {
		if errors.Is(err, errors.New("wallet not found for user and currency")) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

// UpdateWalletBalance handles requests to credit or debit a wallet
func (h *WalletHandler) UpdateWalletBalance(c *fiber.Ctx) error {
	walletID := c.Params("id")
	if walletID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "wallet ID is required")
	}

	var req dto.UpdateBalanceRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if req.Amount <= 0 || req.Reference == "" || (req.Type != "credit" && req.Type != "debit") {
		return fiber.NewError(fiber.StatusBadRequest, "positive amount, reference, and valid type ('credit'/'debit') are required")
	}

	resp, err := h.svc.UpdateWalletBalance(c.Context(), walletID, req)
	if err != nil {
		if errors.Is(err, errors.New("wallet not found")) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

// GetWalletLedger handles requests to get ledger entries for a wallet
func (h *WalletHandler) GetWalletLedger(c *fiber.Ctx) error {
	walletID := c.Params("id")
	if walletID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "wallet ID is required")
	}

	resp, err := h.svc.GetWalletLedger(c.Context(), walletID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
