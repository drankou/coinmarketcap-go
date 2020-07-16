package coinmarketcap_go

import "time"

type CryptocurrencyStatus int

const (
	Active CryptocurrencyStatus = iota
	InActive
)

type CryptocurrencyCategory string

const (
	Coin CryptocurrencyCategory = "coin"
	Token
)

type CryptocurrencyInfoRequest struct {
	//One or more comma-separated CoinMarketCap cryptocurrency IDs
	Id string `url:"id,omitempty"`
	//A comma-separated list of cryptocurrency slugs
	Slug string `url:"slug,omitempty"`

	//One or more comma-separated cryptocurrency symbols
	Symbol string `url:"symbol,omitempty"`

	//Default: "urls,logo,description,tags,platform,date_added,notice"
	//A comma-separated list of supplemental data fields to return.
	//Pass urls,logo,description,tags,platform,date_added,notice,status to include all auxiliary fields.
	Aux string `url:"aux,omitempty"`
}

type CryptocurrencyInfoResponse struct {
	Data   map[string]*CryptocurrencyInfo `json:"data"`
	Status ResponseStatus                 `json:"status"`
}

type CryptocurrencyInfo struct {
	//The unique CoinMarketCap ID for this cryptocurrency.
	Id int `json:"id"`

	//The name of this cryptocurrency.
	Name string `json:"name"`

	//The ticker symbol for this cryptocurrency.
	Symbol string `json:"symbol"`

	//The category for this cryptocurrency.
	Category CryptocurrencyCategory `json:"category"`

	//The web URL friendly shorthand version of this cryptocurrency name.
	Slug string `json:"slug"`

	//Link to a CoinMarketCap hosted logo png for this cryptocurrency.
	//64px is default size returned.
	//Replace "64x64" in the image path with these alternative sizes: 16, 32, 64, 128, 200
	Logo string `json:"logo"`

	//Brief description of this cryptocurrency.
	Description string `json:"description"`

	//Timestamp (ISO 8601) of when this cryptocurrency was added to CoinMarketCap.
	DateAdded *time.Time `json:"date_added"`

	//A Markdown formatted notice that may highlight a significant event
	//or condition that is impacting the cryptocurrency or how it is displayed.
	Notice string `json:"notice"`

	//Tags associated with this cryptocurrency.
	Tags []string `json:"tags"`

	//Metadata about the parent cryptocurrency platform this cryptocurrency belongs to if it is a token.
	Platform Platform `json:"platform"`

	//Various resource URLs for this cryptocurrency.
	Urls Urls `json:"urls"`
}

type Urls struct {
	Website      []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Explorer     []string `json:"explorer"`
	SourceCode   []string `json:"source_code"`
	MessageBoard []string `json:"message_board"`
	Chat         []string `json:"chat"`
	Announcement []string `json:"announcement"`
	Reddit       []string `json:"reddit"`
	Twitter      []string `json:"twitter"`
}

type CryptocurrencyMapRequest struct {
	//Default: "active"
	//One or more comma-separated values: "active", "inactive", "untracked".
	ListingStatus string `url:"listing_status,omitempty"`

	//Default: "1"
	//Offset the start (1-based index) of the paginated list of items to return.
	Start int `url:"start,omitempty"`

	//The number of results to return. Range [1..5000]
	Limit int `url:"limit,omitempty"`

	//Default: "id"
	//Field to sort the list of cryptocurrencies by
	Sort string `url:"sort,omitempty"`

	//A comma-separated list of cryptocurrency symbols to return CoinMarketCap IDs for.
	//If this option is passed, other options will be ignored.
	Symbol string `url:"symbol,omitempty"`

	//Default: "platform,first_historical_data,last_historical_data,is_active"
	//A comma-separated list of supplemental data fields to return. Pass platform,first_historical_data,last_historical_data,is_active,status to include all auxiliary fields.
	Aux string `url:"aux,omitempty"`
}

type CryptocurrencyMapResponse struct {
	Data   []Cryptocurrency `json:"data"`
	Status ResponseStatus   `json:"status"`
}

type Cryptocurrency struct {
	//The unique cryptocurrency ID for this cryptocurrency.
	Id int `json:"id"`

	//The name of this cryptocurrency.
	Name string `json:"name"`

	//The ticker symbol for this cryptocurrency, always in all caps.
	Symbol string `json:"symbol"`

	//The web URL friendly shorthand version of this cryptocurrency name.
	Slug string `json:"slug"`

	//1 if this cryptocurrency has at least 1 active market currently being tracked by the platform, otherwise 0.
	IsActive CryptocurrencyStatus `json:"is_active"`

	//The listing status of the cryptocurrency.
	//This field is only returned if requested through the aux request parameter.
	Status string `json:"status"`

	//Timestamp (ISO 8601) of the date this cryptocurrency was first available on the platform.
	FirstHistoricalData *time.Time `json:"first_historical_data"`

	//Timestamp (ISO 8601) of the last time this cryptocurrency's market data was updated.
	LastHistoricalData *time.Time `json:"last_historical_data"`

	//Metadata about the parent cryptocurrency platform this cryptocurrency belongs to if it is a token, otherwise null.
	Platform Platform `json:"platform"`
}

type Platform struct {
	//The unique CoinMarketCap ID for the parent platform cryptocurrency.
	Id int `json:"id"`

	//The name of the parent platform cryptocurrency.
	Name string `json:"name"`

	//The ticker symbol for the parent platform cryptocurrency.
	Symbol string `json:"symbol"`

	//The web URL friendly shorthand version of the parent platform cryptocurrency name.
	Slug string `json:"slug"`

	//The token address on the parent platform cryptocurrency.
	TokenAddress string `json:"token_address"`
}

type ResponseStatus struct {
	//Current timestamp (ISO 8601) on the server.
	Timestamp *time.Time `json:"timestamp"`

	//An internal error code for the current error. If a unique platform error code is not available the HTTP status code is returned
	ErrorCode int `json:"error_code"`

	//An error message to go along with the error code.
	ErrorMessage string `json:"error_message"`

	//Number of milliseconds taken to generate this response.
	Elapsed int `json:"elapsed"`

	//Number of API call credits that were used for this call.
	CreditCount int `json:"credit_count"`
}

type CryptocurrencyListingsLatestRequest struct {
	Start                int     `url:"start,omitempty"`
	Limit                int     `url:"limit,omitempty"`
	PriceMin             float64 `url:"price_min,omitempty"`
	PriceMax             float64 `url:"price_max,omitempty"`
	MarketCapMin         float64 `url:"market_cap_min,omitempty"`
	MarketCapMax         float64 `url:"market_cap_max,omitempty"`
	Volume24HMin         float64 `url:"volume_24h_min,omitempty"`
	Volume24HMax         float64 `url:"volume_24h_max,omitempty"`
	CirculatingSupplyMin float64 `url:"circulating_supply_min,omitempty"`
	CirculatingSupplyMax float64 `url:"circulating_supply_max,omitempty"`
	PercentChange24HMin  float64 `url:"percent_change_24h_min,omitempty"`
	PercentChange24HMax  float64 `url:"percent_change_24h_max,omitempty"`
	Convert              string  `url:"convert,omitempty"`
	ConvertId            string  `url:"convert_id,omitempty"`
	Sort                 string  `url:"sort,omitempty"`
	SortDir              string  `url:"sort_dir,omitempty"`
	CryptocurrencyType   string  `url:"cryptocurrency_type,omitempty"`
	Tag                  string  `url:"tag,omitempty"`
	Aux                  string  `url:"aux,omitempty"`
}

type CryptocurrencyListingsLatestResponse struct {
	Data   []CryptocurrencyListing
	Status ResponseStatus
}

type CryptocurrencyListingsHistoricalRequest struct {
	Date               string `url:"date,omitempty"`
	Start              int    `url:"start,omitempty"`
	Limit              int    `url:"limit,omitempty"`
	Convert            string `url:"convert,omitempty"`
	ConvertId          string `url:"convert_id,omitempty"`
	Sort               string `url:"sort,omitempty"`
	SortDir            string `url:"sort_dir,omitempty"`
	CryptocurrencyType string `url:"cryptocurrency_type,omitempty"`
	Aux                string `url:"aux,omitempty"`
}

type CryptocurrencyListingsHistoricalResponse struct {
	Data   []CryptocurrencyListing
	Status ResponseStatus
}

type CryptocurrencyListing struct {
	Id                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	Symbol                 string                 `json:"symbol"`
	Slug                   string                 `json:"slug"`
	CmcRank                int                    `json:"cmc_rank"`
	NumMarketPairs         int                    `json:"num_market_pairs"`
	CirculatingSupply      float64                `json:"circulating_supply"`
	TotalSupply            float64                `json:"total_supply"`
	MarketCapByTotalSupply float64                `json:"market_cap_by_total_supply"`
	MaxSupply              float64                `json:"max_supply"`
	LastUpdated            *time.Time             `json:"last_updated"`
	DateAdded              *time.Time             `json:"date_added"`
	Tags                   []string               `json:"tags"`
	Platform               Platform               `json:"platform"`
	Quote                  map[string]MarketQuote `json:"quote"`
}

type MarketQuote struct {
	Price             float64    `json:"price"`
	Volume24H         float64    `json:"volume_24h"`
	Volume24HReported float64    `json:"volume_24h_reported"`
	Volume7d          float64    `json:"volume_7d"`
	Volume7dReported  float64    `json:"volume_7d_reported"`
	Volume30D         float64    `json:"volume_30d"`
	Volume30DReported float64    `json:"volume_30d_reported"`
	MarketCap         float64    `json:"market_cap"`
	PercentChange1H   float64    `json:"percent_change_1h"`
	PercentChange24H  float64    `json:"percent_change_24h"`
	PercentChange7D   float64    `json:"percent_change_7d"`
	LastUpdated       *time.Time `json:"last_updated"`
}

type CryptocurrencyQuotesLatestRequest struct {
	//One or more comma-separated cryptocurrency CoinMarketCap IDs.
	Id string `url:"id,omitempty"`

	//One or more comma-separated cryptocurrency slugs
	Slug string `url:"slug,omitempty"`

	//One or more comma-separated cryptocurrency symbols.
	Symbol string `url:"symbol,omitempty"`

	//A comma-separated list of cryptocurrency or fiat currency symbols.
	Convert string `url:"convert,omitempty"`

	//A comma-separated list of CoinMarketCap IDs.
	ConvertId string `url:"convert_id,omitempty"`

	//A comma-separated list of supplemental data fields to return
	Aux string `url:"aux,omitempty"`

	//If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned.
	SkipInvalid bool `url:"skip_invalid,omitempty"`
}

type CryptocurrencyQuotesLatestResponse struct {
	Data   map[string]CryptocurrencyQuote `json:"data"`
	Status ResponseStatus                 `json:"status"`
}

type CryptocurrencyQuote struct {
	//The unique CoinMarketCap ID for the parent platform cryptocurrency.
	Id int `json:"id"`

	//The name of the parent platform cryptocurrency.
	Name string `json:"name"`

	//The ticker symbol for this cryptocurrency.
	Symbol string `json:"symbol"`

	//The web URL friendly shorthand version of the parent platform cryptocurrency name.
	Slug string `json:"slug"`

	//1 if this cryptocurrency has at least 1 active market currently being tracked by the platform, otherwise 0.
	IsActive CryptocurrencyStatus `json:"is_active"`

	IsFiat int `json:"is_fiat"`

	//The cryptocurrency's CoinMarketCap rank by market cap
	CmcRank int `json:"cmc_rank"`

	//The number of active trading pairs available for this cryptocurrency across supported exchanges.
	NumMarketPairs int `json:"num_market_pairs"`

	CirculatingSupply float64 `json:"circulating_supply"`

	//The approximate total amount of coins in existence right now (minus any coins that have been verifiably burned).
	TotalSupply float64 `json:"total_supply"`

	//The market cap by total supply. This field is only returned if requested through the aux request parameter.
	MarketCapByTotalSupply float64 `json:"market_cap_by_total_supply"`

	//The expected maximum limit of coins ever to be available for this cryptocurrency.
	MaxSupply float64 `json:"max_supply"`

	//Timestamp (ISO 8601) of when this cryptocurrency was added to CoinMarketCap.
	DateAdded *time.Time `json:"date_added"`

	//Array of tags associated with this cryptocurrency.
	Tags []string `json:"tags"`

	//Metadata about the parent cryptocurrency platform this cryptocurrency belongs to if it is a token, otherwise null.
	Platform Platform `json:"platform"`

	//Timestamp (ISO 8601) of the last time this cryptocurrency's market data was updated.
	LastUpdated *time.Time `json:"last_updated"`

	Quote map[string]MarketQuote `json:"quote"`
}

type CryptocurrencyOHLCVLatestRequest struct {
	Id          string `url:"id,omitempty"`
	Symbol      string `url:"symbol,omitempty"`
	Convert     string `url:"convert,omitempty"`
	ConvertId   string `url:"convert_id,omitempty"`
	SkipInvalid string `url:"skip_invalid,omitempty"`
}

type CryptocurrencyOHLCVLatestResponse struct {
	Data   map[string]*CryptocurrencyOHLCV `json:"data"`
	Status ResponseStatus                  `json:"status"`
}

type CryptocurrencyOHLCV struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Symbol      string            `json:"symbol"`
	LastUpdated *time.Time        `json:"last_updated"`
	TimeOpen    *time.Time        `json:"time_open"`
	TimeHigh    *time.Time        `json:"time_high"`
	TimeLow     *time.Time        `json:"time_low"`
	TimeClose   *time.Time        `json:"time_close"`
	QuoteOHLCV  map[string]*OHLCV `json:"quote"`
}

type OHLCV struct {
	Open        float64    `json:"open"`
	High        float64    `json:"high"`
	Low         float64    `json:"low"`
	Close       float64    `json:"close"`
	Volume      float64    `json:"volume"`
	LastUpdated *time.Time `json:"last_updated"`
}

type FiatMapRequest struct {
	Start         int    `url:"start,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Sort          string `url:"sort,omitempty"`
	IncludeMetals bool   `url:"include_metals,omitempty"`
}

type FiatMapResponse struct {
	Data   []Fiat         `json:"data"`
	Status ResponseStatus `json:"status"`
}

type Fiat struct {
	//The unique CoinMarketCap ID for this asset.
	Id int `json:"id"`

	//The name of this asset.
	Name string `json:"name"`

	//The currency sign for this asset.
	Sign string `json:"sign"`

	//The ticker symbol for this asset, always in all caps.
	Symbol string `json:"symbol"`
}

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
