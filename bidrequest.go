package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors.
var (
	ErrInvalidReqNoID     = errors.New("openrtb: request ID missing")
	ErrInvalidReqNoImps   = errors.New("openrtb: request has no impressions")
	ErrInvalidReqMultiInv = errors.New("openrtb: request has multiple inventory sources") // has site and app
)

// BidRequest is the top-level bid request object contains a globally unique bid request or auction ID.  This "id"
// attribute is required as is at least one "imp" (i.e., impression) object.  Other attributes are
// optional since an exchange may establish default values.
type BidRequest struct {
	// Unique ID of the bid request.
	ID     string  `json:"id"`
	Site   *Site   `json:"site,omitempty"`
	App    *App    `json:"app,omitempty"`
	Device *Device `json:"device,omitempty"`
	User   *User   `json:"user,omitempty"`
	// A Source object that provides data about the inventory
	// source and which entity makes the final decision.
	Source      *Source      `json:"source,omitempty"`
	Regulations *Regulations `json:"regs,omitempty"`
	Impressions []Impression `json:"imp,omitempty"`
	// Array of buyer seats allowed to bid on this auction.
	Seats []string `json:"wseat,omitempty"`
	// Array of buyer seats blocked to bid on this auction.
	BlockedSeats []string `json:"bseat,omitempty"`
	// Array of languages for creatives using ISO-639-1-alpha-2.
	Languages []string `json:"wlang,omitempty"`
	// Array of allowed currencies.
	Currencies []string `json:"cur,omitempty"`
	// Blocked Advertiser Categories.
	BlockedCategories []ContentCategory `json:"bcat,omitempty"`
	// Array of strings of blocked toplevel domains of advertisers.
	BlockedAdvDomains []string `json:"badv,omitempty"`
	// Block list of applications by their platform-specific exchange-independent application identifiers.
	// On Android, these should be bundle or package names (e.g., com.foo.mygame).
	// On iOS, these are numeric IDs.
	BlockedApps []string        `json:"bapp,omitempty"`
	Ext         json.RawMessage `json:"ext,omitempty"`
	// Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	Test int `json:"test,omitempty"`
	// Auction type, where 1 = First Price, 2 = Second Price Plus.
	// Exchange-specific auction types can be defined using values greater than 500.
	AuctionType int `json:"at"`
	// Maximum amount of time in milliseconds to submit a bid.
	TimeMax int `json:"tmax,omitempty"`
	// Flag to indicate whether exchange can verify that all impressions
	// offered represent all the impressions available in context, default: 0.
	AllImpressions int `json:"allimps,omitempty"`
}

// Validate validates the bid request.
func (req *BidRequest) Validate() error {
	if req.ID == "" {
		return ErrInvalidReqNoID
	}
	if len(req.Impressions) == 0 {
		return ErrInvalidReqNoImps
	}
	if req.Site != nil && req.App != nil {
		return ErrInvalidReqMultiInv
	}

	for i := range req.Impressions {
		imp := req.Impressions[i]
		if err := (&imp).Validate(); err != nil {
			return err
		}
	}

	return nil
}
