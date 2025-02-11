package constants

type (
	Currency          string
	TransactionMode   string
	TransactionStatus string
)

const (
	CurrencyNGN Currency = "NGN"
	CurrencyUSD Currency = "USD"

	TransactionModeCredit TransactionMode = "credit"
	TransactionModeDebit  TransactionMode = "debit"

	TransactionStatusSuccess TransactionStatus = "success"
)
