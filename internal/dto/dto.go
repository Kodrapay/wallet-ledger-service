package dto

import "time"

// CreateWalletRequest DTO for creating a new wallet
type CreateWalletRequest struct {
	UserID   int    `json:"user_id"`
	Currency string `json:"currency"`
}

// UpdateBalanceRequest DTO for updating a wallet's balance (credit/debit)
type UpdateBalanceRequest struct {
	Amount      int64  `json:"amount"`
	Reference   int    `json:"reference"` // Unique reference for the transaction
	Description string `json:"description"`
	Type        string `json:"type"`      // "credit" or "debit"
}

// WalletResponse DTO for returning wallet information
type WalletResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Currency  string    `json:"currency"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LedgerEntryResponse DTO for returning a ledger entry
type LedgerEntryResponse struct {
	ID          int       `json:"id"`
	WalletID    int       `json:"wallet_id"`
	Reference   int       `json:"reference"`
	Type        string    `json:"type"`
	Amount      int64     `json:"amount"`
	Balance     int64     `json:"balance"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}