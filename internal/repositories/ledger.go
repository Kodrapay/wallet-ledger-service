package repositories

import "log"

type LedgerRepository struct {
    dsn string
}

func NewLedgerRepository(dsn string) *LedgerRepository {
    log.Printf("LedgerRepository using DSN: %s", dsn)
    return &LedgerRepository{dsn: dsn}
}

// TODO: implement persistence for ledger accounts, entries, balances.
