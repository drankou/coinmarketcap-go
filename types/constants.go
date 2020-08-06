package types

import (
	"golang.org/x/time/rate"
	"time"
)

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

type ApiPlan int

const (
	Free ApiPlan = iota
	Basic
	Hobbyist
	Startup
	Standard
	Professional
	Enterprise
)

var APIRateLimits = map[ApiPlan]rate.Limit{
	Free:         rate.Limit(30 / time.Minute.Seconds()),
	Basic:        rate.Limit(30 / time.Minute.Seconds()),
	Hobbyist:     rate.Limit(30 / time.Minute.Seconds()),
	Startup:      rate.Limit(30 / time.Minute.Seconds()),
	Standard:     rate.Limit(60 / time.Minute.Seconds()),
	Professional: rate.Limit(90 / time.Minute.Seconds()),
}
