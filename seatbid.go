package openrtb

import (
	"encoding/json"
	"errors"
)

// SeatBid contains seat information. At least one of Bid is required.
// A bid response can contain multiple "seatbid” objects, each on behalf of a different bidder seat.
// SeatBid object can contain multiple bids each pertaining to a different impression on behalf of a seat.
// Each "bid” object must include the impression ID to which it pertains as well as the bid price.
// Group attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if it is
// only interested in winning any if it can win them all (i.e., all or nothing).
type SeatBid struct {
	Seat  string          `json:"seat,omitempty"` // ID of the bidder seat optional string ID of the bidder seat on whose behalf this bid is made.
	Bids  []Bid           `json:"bid"`            // Array of bid objects; each relates to an imp, if exchange supported can have many bid objects.
	Ext   json.RawMessage `json:"ext,omitempty"`
	Group int             `json:"group,omitempty"` // '1' means impression must be won-lost as a group; default is '0'.
}

// Validation errors.
var (
	ErrInvalidSeatBidBid = errors.New("openrtb: seatbid is missing bids")
)

// Validate required attributes.
func (sb *SeatBid) Validate() error {
	if len(sb.Bids) == 0 {
		return ErrInvalidSeatBidBid
	}

	for i := range sb.Bids {
		bid := sb.Bids[i]
		if err := (&bid).Validate(); err != nil {
			return err
		}
	}

	return nil
}
