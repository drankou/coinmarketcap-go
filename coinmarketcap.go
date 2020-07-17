package coinmarketcap_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drankou/coinmarketcap-go/types"
	"github.com/google/go-querystring/query"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	API_URL     = "https://pro-api.coinmarketcap.com"
	SANDBOX_URL = "https://sandbox-api.coinmarketcap.com"
)

type CoinmarketcapClient struct {
	client *http.Client
}

func (c *CoinmarketcapClient) Init() error {
	c.client = &http.Client{}

	return nil
}

// ------ Cryptocurrency ------ //

// Returns a mapping of all cryptocurrencies to unique CoinMarketCap ids.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1CryptocurrencyMap
func (c *CoinmarketcapClient) CryptocurrencyIdMap(request *types.CryptocurrencyMapRequest) ([]types.Cryptocurrency, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/map", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyMapResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyIdMap: %d: %s", resp.StatusCode, resp.Status))
	}
}

// Returns all static metadata available for one or more cryptocurrencies.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1CryptocurrencyInfo
func (c *CoinmarketcapClient) CryptocurrencyInfo(request *types.CryptocurrencyInfoRequest) (map[string]*types.CryptocurrencyInfo, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/info", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyInfoResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyInfo: %d: %s", resp.StatusCode, resp.Status))
	}
}

// Returns a ranked and sorted list of all cryptocurrencies for a historical UTC date.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1CryptocurrencyListingsHistorical
func (c *CoinmarketcapClient) CryptocurrencyListingsHistorical(request *types.CryptocurrencyListingsHistoricalRequest) ([]types.CryptocurrencyListing, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/listings/historical", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyListingsHistoricalResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyListingsHistorical: %d: %s:", resp.StatusCode, resp.Status))
	}
}

// Returns a paginated list of all active cryptocurrencies with latest market data.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1CryptocurrencyListingsLatest
func (c *CoinmarketcapClient) CryptocurrencyListingsLatest(request *types.CryptocurrencyListingsLatestRequest) ([]types.CryptocurrencyListing, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyListingsLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyListingsLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyOHLCVHistorical(request *types.CryptocurrencyOHLCVHistoricalRequest) (map[string]*types.OHLCVHistoricalResult, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/ohlcv/historical", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyOHLCVHistoricalResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyOHLCVHistorical: %d: %s", resp.StatusCode, resp.Status))
	}
}

// Returns the latest OHLCV (Open, High, Low, Close, Volume) market values for one or more cryptocurrencies for the current UTC day.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1CryptocurrencyOhlcvLatest
func (c *CoinmarketcapClient) CryptocurrencyOHLCVLatest(request *types.CryptocurrencyOHLCVLatestRequest) (map[string]*types.CryptocurrencyOHLCV, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/ohlcv/latest", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyOHLCVLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyOHLCVLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyQuotesLatest(request *types.CryptocurrencyQuotesLatestRequest) (map[string]types.CryptocurrencyQuote, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyQuotesLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyQuotesLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyPricePerformanceStats(request *types.CryptocurrencyPricePerformanceStatsRequest) (map[string]*types.PricePerformanceStats, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/cryptocurrency/price-performance-stats/latest", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.CryptocurrencyPricePerformanceStatsResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyPricePerformanceStats: %d: %s", resp.StatusCode, resp.Status))
	}
}

// ------ Fiat ------ //

// Returns a mapping of all supported fiat currencies to unique CoinMarketCap ids.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1FiatMap
func (c *CoinmarketcapClient) FiatMap(request *types.FiatMapRequest) ([]types.Fiat, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/fiat/map", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.FiatMapResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("FiatMap: %d: %s", resp.StatusCode, resp.Status))
	}
}

// ------ Exchange ------ //

// Returns all static metadata for one or more exchanges.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1ExchangeInfo
func (c *CoinmarketcapClient) ExchangeInfo(request *types.ExchangeInfoRequest) (map[string]*types.ExchangeInfo, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/exchange/info", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)

	log.Print(httpRequest.URL.String())
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.ExchangeInfoResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("ExchangeInfo: %d: %s", resp.StatusCode, resp.Status))
	}
}

// Returns a paginated list of all cryptocurrency exchanges by CoinMarketCap ID.
// By default listing_status=active
// https://pro.coinmarketcap.com/api/v1/#operation/getV1ExchangeMap
func (c *CoinmarketcapClient) ExchangeIdMap(request *types.ExchangeIdMapRequest) ([]types.Exchange, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/exchange/map", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)

	log.Print(httpRequest.URL.String())
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.ExchangeIdMapResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("ExchangeInfo: %d: %s", resp.StatusCode, resp.Status))
	}
}

// ------ Global-Metrics ------ //

// Returns the latest global cryptocurrency market metrics.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1GlobalmetricsQuotesLatest
func (c *CoinmarketcapClient) GlobalMetricsQuotesLatest(request *types.GlobalMetricsQuotesLatestRequest) (*types.GlobalMetricsQuotesLatest, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/global-metrics/quotes/latest", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.GlobalMetricsQuotesLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return &cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("GlobalMetricsQuotesLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

// Returns an interval of historical global cryptocurrency market metrics based on time and interval parameters.
// https://pro.coinmarketcap.com/api/v1/#operation/getV1GlobalmetricsQuotesHistorical
func (c *CoinmarketcapClient) GlobalMetricsQuotesHistorical(request *types.GlobalMetricsQuotesHistoricalRequest) ([]types.AggregatedMarketQuote, error) {
	httpRequest, err := http.NewRequest("GET", API_URL+"/v1/global-metrics/quotes/historical", nil)
	if err != nil {
		log.Error(err)
	}

	err = prepareHttpRequest(httpRequest, request)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var cmcIdMapResponse types.GlobalMetricsQuotesHistoricalResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data.Quotes, nil
	} else {
		return nil, errors.New(fmt.Sprintf("GlobalMetricsQuotesHistorical: %d: %s", resp.StatusCode, resp.Status))
	}
}

func prepareHttpRequest(httpRequest *http.Request, request interface{}) error {
	values, err := query.Values(request)
	if err != nil {
		return err
	}

	httpRequest.Header.Set("Accepts", "application/json")
	httpRequest.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMC_PRO_API_KEY"))
	httpRequest.URL.RawQuery = values.Encode()

	return nil
}
