package openrtb

import "encoding/json"

// Source object describes the nature and behavior of the entity
// that is the source of the bid request upstream from the exchange.
type Source struct {
	// Transaction ID that must be common across all participants
	// in this bid request (e.g., potentially multiple exchanges).
	TransactionID string `json:"tid,omitempty"`
	// Payment ID chain string containing embedded syntax
	// described in the TAG Payment ID Protocol v1.0.
	PaymentChain string `json:"pchain,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
	// Entity responsible for the final impression sale decision, where 0 = exchange, 1 = upstream source.
	FinalSaleDecision int `json:"fd,omitempty"`
}
