package types

type CryptocurrencyStatus int
const (
	CryptocurrencyActive CryptocurrencyStatus = iota
	CryptocurrencyInActive
)

type CryptocurrencyCategory string
const (
	Coin  CryptocurrencyCategory = "coin"
	Token CryptocurrencyCategory = "token"
)

type ExchangeStatus string
const (
	ExchangeActive    ExchangeStatus = "active"
	ExchangeInActive  ExchangeStatus = "inactive"
	ExchangeUntracked ExchangeStatus = "untracked"
)
