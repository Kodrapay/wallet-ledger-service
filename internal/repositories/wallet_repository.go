package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/kodra-pay/wallet-ledger-service/internal/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// WalletRepository defines the interface for wallet and ledger data operations
type WalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	GetWalletByUserIDAndCurrency(ctx context.Context, userID int, currency string) (*models.Wallet, error)
	GetWalletByID(ctx context.Context, id int) (*models.Wallet, error)
	UpdateWalletBalance(ctx context.Context, walletID int, amount int64) error
	CreateLedgerEntry(ctx context.Context, entry *models.LedgerEntry) error
	GetLedgerEntriesByWalletID(ctx context.Context, walletID int) ([]models.LedgerEntry, error)
}

// postgresWalletRepository implements WalletRepository for PostgreSQL
type postgresWalletRepository struct {
	db *sql.DB
}

// NewPostgresWalletRepository creates a new PostgreSQL repository
func NewPostgresWalletRepository(db *sql.DB) WalletRepository {
	return &postgresWalletRepository{db: db}
}

func (r *postgresWalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	query := `INSERT INTO wallets (user_id, currency, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query, wallet.UserID, wallet.Currency, wallet.Balance, wallet.CreatedAt, wallet.UpdatedAt).Scan(&id)
	if err == nil {
		wallet.ID = id
	}
	return err
}

func (r *postgresWalletRepository) GetWalletByUserIDAndCurrency(ctx context.Context, userID int, currency string) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	query := `SELECT id, user_id, currency, balance, created_at, updated_at FROM wallets WHERE user_id = $1 AND currency = $2`
	err := r.db.QueryRowContext(ctx, query, userID, currency).Scan(&wallet.ID, &wallet.UserID, &wallet.Currency, &wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil // Wallet not found
	}
	return wallet, err
}

func (r *postgresWalletRepository) GetWalletByID(ctx context.Context, id int) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	query := `SELECT id, user_id, currency, balance, created_at, updated_at FROM wallets WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&wallet.ID, &wallet.UserID, &wallet.Currency, &wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil // Wallet not found
	}
	return wallet, err
}

func (r *postgresWalletRepository) UpdateWalletBalance(ctx context.Context, walletID int, amount int64) error {
	query := `UPDATE wallets SET balance = balance + $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, amount, time.Now(), walletID)
	return err
}

func (r *postgresWalletRepository) CreateLedgerEntry(ctx context.Context, entry *models.LedgerEntry) error {
	query := `INSERT INTO ledger_entries (wallet_id, reference, type, amount, balance, description, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query, entry.WalletID, entry.Reference, entry.Type, entry.Amount, entry.Balance, entry.Description, entry.CreatedAt).Scan(&id)
	if err == nil {
		entry.ID = id
	}
	return err
}

func (r *postgresWalletRepository) GetLedgerEntriesByWalletID(ctx context.Context, walletID int) ([]models.LedgerEntry, error) {
	var entries []models.LedgerEntry
	query := `SELECT id, wallet_id, reference, type, amount, balance, description, created_at FROM ledger_entries WHERE wallet_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		entry := models.LedgerEntry{}
		if err := rows.Scan(&entry.ID, &entry.WalletID, &entry.Reference, &entry.Type, &entry.Amount, &entry.Balance, &entry.Description, &entry.CreatedAt); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
