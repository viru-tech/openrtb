package openrtb

import "encoding/json"

// PMP is the Private Marketplace Object.
type PMP struct {
	Deals   []Deal          `json:"deals,omitempty"`
	Ext     json.RawMessage `json:"ext,omitempty"`
	Private int             `json:"private_auction,omitempty"`
}

// Deal contains PMP deal information.
type Deal struct {
	// Unique deal ID.
	ID string `json:"id,omitempty"`
	// Currency of bid floor.
	BidFloorCurrency string `json:"bidfloorcur,omitempty"`
	// Array of buyer seats allowed to bid on this Direct Deal.
	Seats []string `json:"wseat,omitempty"`
	// Array of advertiser domains allowed to bid on this Direct Deal.
	AdvDomains []string        `json:"wadomain,omitempty"`
	Ext        json.RawMessage `json:"ext,omitempty"`
	// Optional override of the overall auction type of the bid request, where
	// 1 = First Price,
	// 2 = Second Price Plus,
	// 3 = the value passed in bidfloor is the agreed upon deal price.
	// Additional auction types can be defined by the exchange.
	AuctionType int     `json:"at,omitempty"`
	BidFloor    float64 `json:"bidfloor,omitempty"`
}

type jsonDeal Deal

// MarshalJSON custom marshaling with normalization.
func (d *Deal) MarshalJSON() ([]byte, error) {
	d.normalize()
	return json.Marshal((*jsonDeal)(d))
}

// UnmarshalJSON custom unmarshalling with normalization.
func (d *Deal) UnmarshalJSON(data []byte) error {
	var h jsonDeal
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*d = (Deal)(h)
	d.normalize()
	return nil
}

func (d *Deal) normalize() {
	if d.AuctionType == 0 {
		d.AuctionType = 2
	}
}
