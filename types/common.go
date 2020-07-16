package types

import (
	"time"
)



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
