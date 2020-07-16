package types

import "time"

type GlobalMetricsQuotesLatestRequest struct {
	//A comma-separated list of cryptocurrency or fiat currency symbols.
	Convert string `url:"convert,omitempty"`

	//A comma-separated list of CoinMarketCap IDs.
	ConvertId string `url:"convert_id,omitempty"`
}

type GlobalMetricsQuotesLatestResponse struct {
	Data   GlobalMetricsQuotesLatest `json:"data"`
	Status ResponseStatus            `json:"status"`
}

type GlobalMetricsQuotesLatest struct {
	BTCDominance           float64                        `json:"btc_dominance"`
	ETHDominance           float64                        `json:"eth_dominance"`
	ActiveCryptocurrencies float64                        `json:"active_cryptocurrencies"`
	TotalCryptocurrencies  float64                        `json:"total_cryptocurrencies"`
	ActiveMarketPairs      float64                        `json:"active_market_pairs"`
	ActiveExchanges        float64                        `json:"active_exchanges"`
	TotalExchanges         float64                        `json:"total_exchanges"`
	LastUpdated            *time.Time                     `json:"last_updated"`
	Quote                  map[string]*GlobalMetricsQuote `json:"quote"`
}

type GlobalMetricsQuote struct {
	TotalMarketCap           float64    `json:"total_market_cap"`
	TotalVolume24H           float64    `json:"total_volume_24h"`
	TotalVolume24HReported   float64    `json:"total_volume_24h_reported"`
	AltcoinVolume24H         float64    `json:"altcoin_volume_24h"`
	AltcoinVolume24HReported float64    `json:"altcoin_volume_24h_reported"`
	AltcoinMarketCap         float64    `json:"altcoin_market_cap"`
	LastUpdated              *time.Time `json:"last_updated"`
}

type GlobalMetricsQuotesHistoricalRequest struct {
	TimeStart string `url:"time_start,omitempty"`
	TimeEnd   string `url:"time_end,omitempty"`
	Count     int    `url:"count,omitempty"`
	Interval  string `url:"interval,omitempty"`
	Convert   string `url:"convert,omitempty"`
	ConvertId string `url:"convert_id,omitempty"`
	Aux       string `url:"aux,omitempty"`
}

type GlobalMetricsQuotesHistoricalResponse struct {
	Data struct {
		Quotes []AggregatedMarketQuote `json:"quotes"`
	} `json:"data"`
	Status ResponseStatus `json:"status"`
}

type AggregatedMarketQuote struct {
	Timestamp              *time.Time                     `json:"timestamp"`
	SearchInterval         *time.Time                     `json:"search_interval"`
	BTCDominance           float64                        `json:"btc_dominance"`
	ActiveCryptocurrencies int                            `json:"active_cryptocurrencies"`
	ActiveExchanges        int                            `json:"active_exchanges"`
	ActiveMarketPairs      int                            `json:"active_market_pairs"`
	Quote                  map[string]*GlobalMetricsQuote `json:"quote"`
}
