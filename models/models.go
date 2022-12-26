package models

// AvailableModels modela o retorno do endpoint -> https://brapi.dev/api/available
type AvailableModels struct {
	Stocks []string `json:"stocks"`
}

// Currency modela o retorno do endpoint -> https://brapi.dev/api/v2/currency?currency=USD-BRL%2CEUR-USD
type Currency struct {
	Currency []currencyData `json:"currency"`
}
type currencyData struct {
	FromCurrency       string `json:"fromCurrency"`
	ToCurrency         string `json:"toCurrency"`
	Name               string `json:"name"`
	High               string `json:"high"`
	Low                string `json:"low"`
	BidVariation       string `json:"bidVariation"`
	PercentageChange   string `json:"percentageChange"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	UpdatedAtTimestamp string `json:"updatedAtTimestamp"`
	UpdatedAtDate      string `json:"updatedAtDate"`
}

// QuoteList modela o retorno do endpoint -> https://brapi.dev/api/quote/list?sortOrder=desc&limit=1557
type QuoteList struct {
	Stocks []quoteListData `json:"stocks"`
}
type quoteListData struct {
	Stock      string  `json:"stock"`
	Name       string  `json:"name"`
	Close      float64 `json:"close"`
	Change     float64 `json:"change"`
	Volume     float64 `json:"volume"`
	Market_cap float64 `json:"market_cap"`
	Logo       string  `json:"logo"`
	Sector     string  `json:"sector"`
}
