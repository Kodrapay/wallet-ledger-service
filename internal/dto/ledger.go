package dto

type LedgerEntryRequest struct {
	DebitAccount  int     `json:"debit_account"`
	CreditAccount int     `json:"credit_account"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Reference     int     `json:"reference"`
}

type BalanceResponse struct {
	MerchantID int     `json:"merchant_id"`
	Available  float64 `json:"available"`
	Pending    float64 `json:"pending"`
	Currency   string  `json:"currency"`
}
