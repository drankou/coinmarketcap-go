package coinmarketcap_go

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (c *CoinmarketcapClient) CryptocurrencyIdMap(request *CryptocurrencyMapRequest) ([]Cryptocurrency, error) {
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

		var cmcIdMapResponse CryptocurrencyMapResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyIdMap: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyInfo(request *CryptocurrencyInfoRequest) (map[string]*CryptocurrencyInfo, error) {
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

		var cmcIdMapResponse CryptocurrencyInfoResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyInfo: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyListingsLatest(request *CryptocurrencyListingsLatestRequest) ([]CryptocurrencyListing, error) {
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

		var cmcIdMapResponse CryptocurrencyListingsLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyListingsLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyListingsHistorical(request *CryptocurrencyListingsHistoricalRequest) ([]CryptocurrencyListing, error) {
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

		var cmcIdMapResponse CryptocurrencyListingsHistoricalResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyListingsHistorical: %d: %s:", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyOHLCVLatest(request *CryptocurrencyOHLCVLatestRequest) (map[string]*CryptocurrencyOHLCV, error) {
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

		var cmcIdMapResponse CryptocurrencyOHLCVLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyOHLCVLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) CryptocurrencyQuotesLatest(request *CryptocurrencyQuotesLatestRequest) (map[string]CryptocurrencyQuote, error) {
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

		var cmcIdMapResponse CryptocurrencyQuotesLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("CryptocurrencyQuotesLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) FiatMap(request *FiatMapRequest) ([]Fiat, error) {
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

		var cmcIdMapResponse FiatMapResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("FiatMap: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) GlobalMetricsQuotesLatest(request *GlobalMetricsQuotesLatestRequest) (*GlobalMetricsQuotesLatest, error) {
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

		var cmcIdMapResponse GlobalMetricsQuotesLatestResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return &cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("GlobalMetricsQuotesLatest: %d: %s", resp.StatusCode, resp.Status))
	}
}

func (c *CoinmarketcapClient) ExchangeInfo(request *ExchangeInfoRequest) (map[string]*ExchangeInfo, error) {
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

		var cmcIdMapResponse ExchangeInfoResponse
		err = json.Unmarshal(respBody, &cmcIdMapResponse)
		if err != nil {
			return nil, err
		}

		return cmcIdMapResponse.Data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("ExchangeInfo: %d: %s", resp.StatusCode, resp.Status))
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
