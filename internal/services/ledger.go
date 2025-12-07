package services

import (
	"context"

	"github.com/kodra-pay/wallet-ledger-service/internal/dto"
	"github.com/kodra-pay/wallet-ledger-service/internal/repositories"
)

type LedgerService struct {
	repo *repositories.LedgerRepository
}

func NewLedgerService(repo *repositories.LedgerRepository) *LedgerService {
	return &LedgerService{repo: repo}
}

func (s *LedgerService) GetBalance(_ context.Context, merchantID int) dto.BalanceResponse {
	return dto.BalanceResponse{
		MerchantID: merchantID, // int
		Available:  0,
		Pending:    0,
		Currency:   "NGN",
	}
}

func (s *LedgerService) CreateEntry(_ context.Context, req dto.LedgerEntryRequest) map[string]interface{} {
	return map[string]interface{}{
		"debit_account":  req.DebitAccount,  // int
		"credit_account": req.CreditAccount, // int
		"amount":         req.Amount,
		"currency":       req.Currency,
		"reference":      req.Reference,     // int
		"status":         "recorded",
	}
}
