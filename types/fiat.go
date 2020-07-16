package types

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
