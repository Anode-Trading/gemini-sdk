package api

type OrderBookResponse struct {
	Type                 string   `json:"type"`
	EventID              int      `json:"event_id"`
	SocketSequenceNumber int      `json:"socket_sequence"`
	Events               []Events `json:"events"`
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
