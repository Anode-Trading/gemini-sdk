package api

type OrderBookResponse struct {
	Type                 string   `json:"type"`
	EventID              int      `json:"event_id"`
	SocketSequenceNumber int      `json:"socket_sequence"`
	Events               []Events `json:"events"`
	Timestamp            int64    `json:"timestamp"`
}

type Events struct {
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Price     string `json:"price"`
	Delta     string `json:"delta"`
	Remaining string `json:"remaining"`
	Side      string `json:"side"`
}

type WsDepthEvent struct {
	Asks        []PriceLevel
	Bids        []PriceLevel
	ChannelName string
	Pair        string
	ChannelID   int
	Checksum    *string
	IsSnapshot  bool
}

type PriceLevel struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type SymbolDetails struct {
	Symbol                string  `json:"symbol"`
	BaseCurrency          string  `json:"base_currency"`
	QuoteCurrency         string  `json:"quote_currency"`
	TickSize              float64 `json:"tick_size"`
	QuoteIncrement        float64 `json:"quote_increment"`
	MinOrderSize          string  `json:"min_order_size"`
	Status                string  `json:"status"`
	WrapEnabled           bool    `json:"wrap_enabled"`
	ProductType           string  `json:"product_type"`
	ContractType          string  `json:"contract_type"`
	ContractPriceCurrency string  `json:"contract_price_currency"`
}
