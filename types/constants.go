package types

type CryptocurrencyStatus int

const (
	CryptocurrencyInActive CryptocurrencyStatus = iota
	CryptocurrencyActive
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
