package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/kodra-pay/wallet-ledger-service/internal/dto"
	"github.com/kodra-pay/wallet-ledger-service/internal/models"
	"github.com/kodra-pay/wallet-ledger-service/internal/repositories"
)

// WalletService defines the business logic for wallet and ledger operations
type WalletService struct {
	repo repositories.WalletRepository
}

// NewWalletService creates a new wallet service
func NewWalletService(repo repositories.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(ctx context.Context, req dto.CreateWalletRequest) (*dto.WalletResponse, error) {
	// Check if wallet already exists for this user and currency
	existingWallet, err := s.repo.GetWalletByUserIDAndCurrency(ctx, req.UserID, req.Currency) // int
	if err != nil {
		return nil, fmt.Errorf("failed to check for existing wallet: %w", err)
	}
	if existingWallet != nil {
		return nil, errors.New("wallet already exists for this user and currency")
	}

	wallet := models.NewWallet(req.UserID, req.Currency) // int
	if err := s.repo.CreateWallet(ctx, wallet); err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	return &dto.WalletResponse{
		ID:        wallet.ID,        // int
		UserID:    wallet.UserID,    // int
		Currency:  wallet.Currency,
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}, nil
}

func (s *WalletService) GetWalletByID(ctx context.Context, walletID int) (*dto.WalletResponse, error) { // int
	wallet, err := s.repo.GetWalletByID(ctx, walletID) // int
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	return &dto.WalletResponse{
		ID:        wallet.ID,        // int
		UserID:    wallet.UserID,    // int
		Currency:  wallet.Currency,
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}, nil
}

func (s *WalletService) GetWalletByUserIDAndCurrency(ctx context.Context, userID int, currency string) (*dto.WalletResponse, error) { // int
	wallet, err := s.repo.GetWalletByUserIDAndCurrency(ctx, userID, currency) // int
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet by user ID and currency: %w", err)
	}
	if wallet == nil {
		return nil, errors.New("wallet not found for user and currency")
	}

	return &dto.WalletResponse{
		ID:        wallet.ID,        // int
		UserID:    wallet.UserID,    // int
		Currency:  wallet.Currency,
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}, nil
}

func (s *WalletService) UpdateWalletBalance(ctx context.Context, walletID int, req dto.UpdateBalanceRequest) (*dto.WalletResponse, error) { // int
	wallet, err := s.repo.GetWalletByID(ctx, walletID) // int
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet for update: %w", err)
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	var amountChange int64
	if req.Type == "credit" {
		amountChange = req.Amount
	} else if req.Type == "debit" {
		amountChange = -req.Amount
	} else {
		return nil, errors.New("invalid transaction type, must be 'credit' or 'debit'")
	}

	// Update wallet balance
	if err := s.repo.UpdateWalletBalance(ctx, walletID, amountChange); err != nil { // int
		return nil, fmt.Errorf("failed to update wallet balance: %w", err)
	}

	// Get updated wallet to record current balance in ledger
	updatedWallet, err := s.repo.GetWalletByID(ctx, walletID) // int
	if err != nil {
		return nil, fmt.Errorf("failed to get updated wallet: %w", err)
	}
	if updatedWallet == nil {
		return nil, errors.New("updated wallet not found after balance change")
	}

	// Create ledger entry
	ledgerEntry := models.NewLedgerEntry(
		walletID,        // int
		req.Reference,   // int
		req.Type,
		req.Amount,
		updatedWallet.Balance, // Current balance after update
		req.Description,
	)
	if err := s.repo.CreateLedgerEntry(ctx, ledgerEntry); err != nil {
		return nil, fmt.Errorf("failed to create ledger entry: %w", err)
	}

	return &dto.WalletResponse{
		ID:        updatedWallet.ID,        // int
		UserID:    updatedWallet.UserID,    // int
		Currency:  updatedWallet.Currency,
		Balance:   updatedWallet.Balance,
		CreatedAt: updatedWallet.CreatedAt,
		UpdatedAt: updatedWallet.UpdatedAt,
	}, nil
}

func (s *WalletService) GetWalletLedger(ctx context.Context, walletID int) ([]dto.LedgerEntryResponse, error) { // int
	entries, err := s.repo.GetLedgerEntriesByWalletID(ctx, walletID) // int
	if err != nil {
		return nil, fmt.Errorf("failed to get ledger entries: %w", err)
	}

	var resp []dto.LedgerEntryResponse
	for _, entry := range entries {
		resp = append(resp, dto.LedgerEntryResponse{
			ID:          entry.ID,          // int
			WalletID:    entry.WalletID,    // int
			Reference:   entry.Reference,   // int
			Type:        entry.Type,
			Amount:      entry.Amount,
			Balance:     entry.Balance,
			Description: entry.Description,
			CreatedAt:   entry.CreatedAt,
		})
	}
	return resp, nil
}
