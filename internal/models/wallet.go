package models

import (
	"time"
)

// Wallet represents a customer's wallet
type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Currency  string    `json:"currency"`
	Balance   int64     `json:"balance"` // Stored in cents/smallest unit
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewWallet creates a new Wallet instance
func NewWallet(userID int, currency string) *Wallet {
	return &Wallet{
		ID:        0, // Will be set by DB
		UserID:    userID,
		Currency:  currency,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// LedgerEntry represents an entry in the transaction ledger for a wallet
type LedgerEntry struct {
	ID          int       `json:"id"`
	WalletID    int       `json:"wallet_id"`
	Reference   int       `json:"reference"` // Reference to the external transaction
	Type        string    `json:"type"`      // "credit" or "debit"
	Amount      int64     `json:"amount"`    // Stored in cents/smallest unit
	Balance     int64     `json:"balance"`   // Balance of the wallet after this entry
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewLedgerEntry creates a new LedgerEntry instance
func NewLedgerEntry(walletID, reference int, entryType string, amount, balance int64, description string) *LedgerEntry {
	return &LedgerEntry{
		ID:          0, // Will be set by DB
		WalletID:    walletID,
		Reference:   reference,
		Type:        entryType,
		Amount:      amount,
		Balance:     balance,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
