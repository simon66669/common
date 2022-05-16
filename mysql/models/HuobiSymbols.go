package models

type HuobiSymbols struct {
	Id              string `json:"id"`
	BaseCurrency    string `json:"base-currency"`
	QuoteCurrency   string `json:"quote-currency"`
	PricePrecision  int    `json:"price-precision"`
	AmountPrecision int    `json:"amount-precision"`
	SymbolPartition string `json:"symbol-partition"`
	Symbol          string `json:"symbol"`
}
