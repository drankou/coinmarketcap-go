package coinmarketcap_go

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinmarketcapClient_Init(t *testing.T) {
	c := &CoinmarketcapClient{}
	err := c.Init()
	if err != nil {
		t.Error(err)
	}
}

func TestCoinmarketcapClient_CryptocurrencyIdMap(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	cmc := &CoinmarketcapClient{}
	err = cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &CryptocurrencyMapRequest{
		ListingStatus: "active,inactive,untracked",
	}

	assets, err := cmc.CryptocurrencyIdMap(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, assets, "empty assets response")
}

func TestCoinmarketcapClient_CryptocurrencyIdMap2(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	cmc := &CoinmarketcapClient{}
	err = cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &CryptocurrencyMapRequest{
		Symbol: "BTC",
	}

	assets, err := cmc.CryptocurrencyIdMap(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, assets, 1, "unexpected assets response")
	assert.Equal(t, assets[0].Symbol, "BTC", "specified assets doesn't match")
}

func TestCoinmarketcapClient_CryptocurrencyInfo(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	cmc := &CoinmarketcapClient{}
	err = cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &CryptocurrencyInfoRequest{
		Symbol: "BTC,ETH",
	}

	cryptocurrencyInfoMap, err := cmc.CryptocurrencyInfo(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, cryptocurrencyInfoMap, 2, "nubmer of assets in response doesnt match")
	assert.NotNil(t, cryptocurrencyInfoMap["BTC"], "missing cryptocurrency info for BTC")
	assert.NotNil(t, cryptocurrencyInfoMap["ETH"], "missing cryptocurrency info for ETH")
}

func TestCoinmarketcapClient_CryptocurrencyListingsLatest(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	cmc := &CoinmarketcapClient{}
	err = cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &CryptocurrencyListingsLatestRequest{}

	cryptocurrencyInfoMap, err := cmc.CryptocurrencyListingsLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, cryptocurrencyInfoMap, "empty response")
}
