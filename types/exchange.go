package types

import "time"

type ExchangeInfoRequest struct {
	Id   string `url:"id,omitempty"`
	Slug string `url:"slug,omitempty"`
	Aux  string `url:"aux,omitempty"`
}

type ExchangeInfoResponse struct {
	Data   map[string]*ExchangeInfo `json:"data"`
	Status ResponseStatus           `json:"status"`
}

type ExchangeInfo struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Slug         string       `json:"slug"`
	Logo         string       `json:"logo"`
	Description  string       `json:"description"`
	DateLaunched *time.Time   `json:"date_launched"`
	Notice       string       `json:"notice"`
	Urls         ExchangeUrls `json:"urls"`
}

type ExchangeUrls struct {
	Website []string `json:"website"`
	Blog    []string `json:"blog"`
	Chat    []string `json:"chat"`
	Fee     []string `json:"fee"`
	Twitter []string `json:"twitter"`
}

type ExchangeIdMapRequest struct {
	ListingStatus ExchangeStatus `url:"listing_status,omitempty"`
	Slug          string         `url:"slug,omitempty"`
	Start         int            `url:"start,omitempty"`
	Limit         int            `url:"limit,omitempty"`
	Sort          string         `url:"sort,omitempty"`
	Aux           string         `url:"aux,omitempty"`
}

type ExchangeIdMapResponse struct {
	Data   []Exchange     `json:"data"`
	Status ResponseStatus `json:"status"`
}

type Exchange struct {
	Id             int                       `json:"id"`
	Name           string                    `json:"name"`
	Slug           string                    `json:"slug"`
	NumMarketPairs int                       `json:"num_market_pairs"`
	LastUpdated    *time.Time                `json:"last_updated"`
	Quote          map[string]*ExchangeQuote `json:"quote"`
}

type ExchangeQuote struct {
	Volume24H              float64    `json:"volume_24h"`
	Volume24HAdjusted      float64    `json:"volume_24h_adjusted"`
	Volume7d               float64    `json:"volume_7d"`
	Volume30D              float64    `json:"volume_30d"`
	MarketCap              float64    `json:"market_cap"`
	PercentChangeVolume24H float64    `json:"percent_change_volume_24h"`
	PercentChangeVolume7D  float64    `json:"percent_change_volume_7d"`
	PercentChangeVolume30D float64    `json:"percent_change_volume_30d"`
	LastUpdated            *time.Time `json:"last_updated"`
}
