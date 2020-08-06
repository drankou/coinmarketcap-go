package coinmarketcap_go

import (
	"github.com/drankou/coinmarketcap-go/types"
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

func initCmcClient() (*CoinmarketcapClient, error) {
	c := &CoinmarketcapClient{}
	err := c.Init(types.Basic)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func TestCoinmarketcapClient_Init(t *testing.T) {
	c := &CoinmarketcapClient{}
	err := c.Init(types.Basic)
	if err != nil {
		t.Error(err)
	}
}

func TestCoinmarketcapClient_CryptocurrencyIdMap(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyMapRequest{
		ListingStatus: "active,inactive,untracked",
	}

	assets, err := cmc.CryptocurrencyIdMap(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, assets, "empty assets response")
}

func TestCoinmarketcapClient_CryptocurrencyIdMap2(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyMapRequest{
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
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyInfoRequest{
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
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyListingsLatestRequest{}

	cryptocurrencyInfoMap, err := cmc.CryptocurrencyListingsLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, cryptocurrencyInfoMap, "empty response")
}

func TestCoinmarketcapClient_CryptocurrencyListingsHistorical(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyListingsHistoricalRequest{
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
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyQuotesLatestRequest{
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
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.FiatMapRequest{}
	assets, err := cmc.FiatMap(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, assets, "empty assets response")
}

func TestCoinmarketcapClient_GlobalMetricsQuotesLatest(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.GlobalMetricsQuotesLatestRequest{}
	globalMetrics, err := cmc.GlobalMetricsQuotesLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, globalMetrics, "global metrics response is nil")
}

func TestCoinmarketcapClient_GlobalMetricsQuotesHistorical(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.GlobalMetricsQuotesHistoricalRequest{}
	globalMetrics, err := cmc.GlobalMetricsQuotesHistorical(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, globalMetrics, "global metrics response is nil")
}

func TestCoinmarketcapClient_CryptocurrencyOHLCVHistorical(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyOHLCVHistoricalRequest{
		Symbol:    "BTC,ETH",
		TimeStart: "2019-08-19",
		TimeEnd:   "2019-08-20",
	}

	response, err := cmc.CryptocurrencyOHLCVHistorical(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, response, "empty response")

	assert.NotNil(t, response["BTC"], "cryptocurrency OHLCV data for BTC is empty")
	assert.NotNil(t, response["ETH"], "cryptocurrency OHLCV data for BTC is empty")

	for asset, ohlcv := range response {
		t.Logf("OHLCV response for %s: %+v", asset, ohlcv)
	}
}

func TestCoinmarketcapClient_CryptocurrencyOHLCVLatest(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyOHLCVLatestRequest{
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

		assert.NotNil(t, ohlcv.OHLCVQuote.Quote["USD"], "cryptocurrency OHLCV map is empty")
		t.Logf("OHLCV prices for %s :%+v", asset, ohlcv.OHLCVQuote.Quote["USD"])
	}

	assert.NotNil(t, cryptocurrencyOHLCVMap["BTC"], "missing data for BTC")
	assert.NotNil(t, cryptocurrencyOHLCVMap["ETH"], "missing data for ETH")
}

func TestCoinmarketcapClient_CryptocurrencyPricePerformanceStats(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.CryptocurrencyPricePerformanceStatsRequest{
		Symbol:  "BTC,ETH",
		Convert: "USD",
	}
	pricePerfomanceStatsMap, err := cmc.CryptocurrencyPricePerformanceStats(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, pricePerfomanceStatsMap, "empty response")

	assert.NotNil(t, pricePerfomanceStatsMap["BTC"], "missing data for BTC")
	assert.NotNil(t, pricePerfomanceStatsMap["ETH"], "missing data for ETH")

	for asset, stats := range pricePerfomanceStatsMap {
		t.Logf("Price stats for %s: %+v", asset, stats)
	}
}

func TestCoinmarketcapClient_ExchangeInfo(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.ExchangeInfoRequest{
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

func TestCoinmarketcapClient_ExchangeIdMap(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.ExchangeIdMapRequest{
		ListingStatus: "active",
		Slug:          "binance",
	}

	exchanges, err := cmc.ExchangeIdMap(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, exchanges, "empty exchanges response")
}

func TestCoinmarketcapClient_PartnersFCASListingsLatest(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.FCASListingsLatestRequest{}

	fcasAssets, err := cmc.PartnersFCASListingsLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, fcasAssets, "empty fcas ratings response")
}

func TestCoinmarketcapClient_PartnersFCASQuotesLatest(t *testing.T) {
	cmc, err := initCmcClient()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.FCASQuotesLatestRequest{
		Symbol: "BTC,ETH",
	}

	fcasAssetsMap, err := cmc.PartnersFCASQuotesLatest(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, fcasAssetsMap, "empty fcas ratings response")
	assert.NotNil(t, fcasAssetsMap["BTC"], "missing data for BTC")
	assert.NotNil(t, fcasAssetsMap["ETH"], "missing data for ETH")
}
