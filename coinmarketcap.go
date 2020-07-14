package coinmarketcap_go

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

	query := url.Values{}
	if request.Start >= 1 {
		query.Add("start", strconv.Itoa(request.Start))
	}

	if request.ListingStatus != "" {
		query.Add("listing_status", request.ListingStatus)
	}

	if request.Limit >= 1 {
		query.Add("end", strconv.Itoa(request.Limit))
	}

	if request.Symbol != "" {
		query.Add("symbol", request.Symbol)
	}

	if request.Aux != "" {
		query.Add("aux", request.Aux)
	}

	httpRequest.Header.Set("Accepts", "application/json")
	httpRequest.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMC_PRO_API_KEY"))
	httpRequest.URL.RawQuery = query.Encode()

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

	query := url.Values{}
	if request.Id != "" {
		query.Add("id", request.Id)
	}

	if request.Slug != "" {
		query.Add("slug", request.Slug)
	}

	if request.Symbol != "" {
		query.Add("symbol", request.Symbol)
	}

	if request.Aux != "" {
		query.Add("aux", request.Aux)
	}

	httpRequest.Header.Set("Accepts", "application/json")
	httpRequest.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMC_PRO_API_KEY"))
	httpRequest.URL.RawQuery = query.Encode()

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

	query := url.Values{}
	if request.Start >= 1 {
		query.Add("start", strconv.Itoa(request.Start))
	}

	if request.Limit != 0 {
		query.Add("limit", strconv.Itoa(request.Limit))
	}

	if request.Aux != "" {
		query.Add("aux", request.Aux)
	}

	httpRequest.Header.Set("Accepts", "application/json")
	httpRequest.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMC_PRO_API_KEY"))
	httpRequest.URL.RawQuery = query.Encode()

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