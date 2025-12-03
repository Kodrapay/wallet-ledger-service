package dto

type LedgerEntryRequest struct {
	DebitAccount  string  `json:"debit_account"`
	CreditAccount string  `json:"credit_account"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Reference     string  `json:"reference"`
}

type BalanceResponse struct {
	MerchantID string  `json:"merchant_id"`
	Available  float64 `json:"available"`
	Pending    float64 `json:"pending"`
	Currency   string  `json:"currency"`
}
