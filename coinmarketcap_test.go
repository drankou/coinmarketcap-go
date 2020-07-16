package coinmarketcap_go

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func TestCoinmarketcapClient_Init(t *testing.T) {
	c := &CoinmarketcapClient{}
	err := c.Init()
	if err != nil {
		t.Error(err)
	}
}

func TestCoinmarketcapClient_CryptocurrencyIdMap(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
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
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
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
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
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

	assert.Len(t, cryptocurrencyInfoMap, 2, "number of assets in response doesnt match")
	assert.NotNil(t, cryptocurrencyInfoMap["BTC"], "missing cryptocurrency info for BTC")
	assert.NotNil(t, cryptocurrencyInfoMap["ETH"], "missing cryptocurrency info for ETH")
}

func TestCoinmarketcapClient_CryptocurrencyListingsLatest(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
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

func TestCoinmarketcapClient_CryptocurrencyListingsHistorical(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &CryptocurrencyListingsHistoricalRequest{
		Convert: "USD,BTC",
		Date:    "2019-08-19",
	}

	cryptocurrencyInfoMap, err := cmc.CryptocurrencyListingsHistorical(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, cryptocurrencyInfoMap, "empty response")
}

func TestCoinmarketcapClient_CryptocurrencyQuotesLatest(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &CryptocurrencyQuotesLatestRequest{
		Symbol: "BTC,ETH",
	}

	cryptocurrencyQuotesLatestMap, err := cmc.CryptocurrencyQuotesLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, cryptocurrencyQuotesLatestMap, "empty response")
	assert.Len(t, cryptocurrencyQuotesLatestMap, 2, "number of assets in response doesnt match")
	assert.NotNil(t, cryptocurrencyQuotesLatestMap["BTC"], "missing cryptocurrency quotes for BTC")
	assert.NotNil(t, cryptocurrencyQuotesLatestMap["ETH"], "missing cryptocurrency quotes for ETH")
}

func TestCoinmarketcapClient_FiatMap(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &FiatMapRequest{}
	assets, err := cmc.FiatMap(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, assets, "empty assets response")
}

func TestCoinmarketcapClient_GlobalMetricsQuotesLatest(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &GlobalMetricsQuotesLatestRequest{}
	globalMetrics, err := cmc.GlobalMetricsQuotesLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, globalMetrics, "global metrics response is nil")
}

func TestCoinmarketcapClient_CryptocurrencyOHLCVLatest(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &CryptocurrencyOHLCVLatestRequest{
		Symbol:  "BTC,ETH",
		Convert: "USD",
	}
	cryptocurrencyOHLCVMap, err := cmc.CryptocurrencyOHLCVLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, cryptocurrencyOHLCVMap, "empty response")

	for asset, ohlcv := range cryptocurrencyOHLCVMap {
		t.Logf("OHLCV response for %s: %+v", asset, ohlcv)

		assert.NotNil(t, ohlcv.QuoteOHLCV["USD"], "cryptocurrency OHLCV map is empty")
		t.Logf("OHLCV prices for %s :%+v", asset, ohlcv.QuoteOHLCV["USD"])
	}

	assert.NotNil(t, cryptocurrencyOHLCVMap["BTC"], "missing data for BTC")
	assert.NotNil(t, cryptocurrencyOHLCVMap["ETH"], "missing data for ETH")
}

func TestCoinmarketcapClient_ExchangeInfo(t *testing.T) {
	cmc := &CoinmarketcapClient{}
	err := cmc.Init()
	if err != nil {
		t.Fatal(err)
	}

	req := &ExchangeInfoRequest{
		Slug: "binance",
	}

	exchangeInfoMap, err := cmc.ExchangeInfo(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, exchangeInfoMap, "empty response")

	for exchange, metadata := range exchangeInfoMap {
		t.Logf("Metadata for %s: %+v", exchange, metadata)
	}

	assert.NotNil(t, exchangeInfoMap["binance"], "missing data for binance")
}
