package types

import "time"

type FCASListingsLatestRequest struct {
	Start int    `url:"start,omitempty"`
	Limit int    `url:"limit,omitempty"`
	Aux   string `url:"aux,omitempty"`
}

type FCASListingsLatestResponse struct {
	Data   []FCASRating   `json:"data"`
	Status ResponseStatus `json:"status"`
}

type FCASRating struct {
	Id               int        `json:"id"`
	Name             string     `json:"name"`
	Symbol           string     `json:"symbol"`
	Slug             string     `json:"slug"`
	Score            int        `json:"score"`
	Grade            string     `json:"grade"`
	PercentChange24H float64    `json:"percent_change_24h"`
	PointChange24H   float64    `json:"point_change_24h"`
	LastUpdated      *time.Time `json:"last_updated"`
}

type FCASQuotesLatestRequest struct {
	Id     string `url:"id,omitempty"`
	Slug   string `url:"slug,omitempty"`
	Symbol string `url:"symbol,omitempty"`
	Aux    string `url:"aux,omitempty"`
}

type FCASQuotesLatestResponse struct {
	Data   map[string]*FCASRating `json:"data"`
	Status ResponseStatus         `json:"status"`
}
