package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors.
var (
	ErrInvalidBidNoID    = errors.New("openrtb: bid is missing ID")
	ErrInvalidBidNoImpID = errors.New("openrtb: bid is missing impression ID")
)

// Bid object contains bid information.
// ID, ImpID and Price are required; all other optional.
// If the bidder wins the impression, the exchange calls notice URL (nurl)
// a) to inform the bidder of the win;
// b) to convey certain information using substitution macros.
// Adomain can be used to check advertiser block list compliance.
// Cid can be used to block ads that were previously identified as inappropriate.
// Substitution macros may allow a bidder to use a static notice URL for all of its bids.
type Bid struct {
	ID string `json:"id"`
	// Required string ID of the impression object to which this bid applies.
	ImpID string `json:"impid"`
	// References the ad to be served if the bid wins.
	AdID string `json:"adid,omitempty"`
	// Win notice URL.
	NoticeURL string `json:"nurl,omitempty"`
	// Billing notice URL.
	BillingURL string `json:"burl,omitempty"`
	// Loss notice URL.
	LossURL string `json:"lurl,omitempty"`
	// Actual ad markup. XHTML if a response to a banner
	// object, or VAST XML if a response to a video object.
	AdMarkup string `json:"adm,omitempty"`
	// A platform-specific application identifier intended
	// to be unique to the app and independent of the exchange.
	Bundle string `json:"bundle,omitempty"`
	// Sample image URL.
	ImageURL string `json:"iurl,omitempty"`
	// Campaign ID that appears with the Ad markup.
	CampaignID StringOrNumber `json:"cid,omitempty"`
	// Creative ID for reporting content issues or defects. This could also
	// be used as a reference to a creative ID that is posted with an exchange.
	CreativeID string `json:"crid,omitempty"`
	// Tactic ID to enable buyers to label bids for reporting to
	// the exchange the tactic through which their bid was submitted.
	Tactic string `json:"tactic,omitempty"`
	// Language of the creative using ISO-639-1-alpha-2.
	Language string `json:"language,omitempty"`
	// DealID extension of private marketplace deals.
	DealID string `json:"dealid,omitempty"`
	// Advertiser’s primary or top-level domain for
	// advertiser checking; or multiple if imp rotating.
	AdvDomains []string `json:"adomain,omitempty"`
	// IAB content categories of the creative. Refer to List 5.1.
	Categories []ContentCategory `json:"cat,omitempty"`
	// Array of creative attributes.
	Attrs []CreativeAttribute `json:"attr,omitempty"`
	// Native ads markup, specially for AdNow.
	Native json.RawMessage `json:"native,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
	// API required by the markup if applicable.
	API APIFramework `json:"api,omitempty"`
	// Video response protocol of the markup if applicable.
	Protocol Protocol `json:"protocol,omitempty"`
	// Creative media rating per IQG guidelines.
	MediaRating IQGRating `json:"qagmediarating,omitempty"`
	// Width of the ad in pixels.
	Width int `json:"w,omitempty"`
	// Height of the ad in pixels.
	Height int `json:"h,omitempty"`
	// Relative width of the creative when expressing size as a ratio.
	WidthRatio int `json:"wratio,omitempty"`
	// Relative height of the creative when expressing size as a ratio.
	HeightRatio int `json:"hratio,omitempty"`
	// Advisory as to the number of seconds the bidder is willing
	// to wait between the auction and the actual impression.
	Exp int `json:"exp,omitempty"`
	// Bid price in CPM. Suggests using integer math for accounting to avoid rounding errors.
	Price float64 `json:"price"`
}

// Validate required attributes.
func (bid *Bid) Validate() error {
	if bid.ID == "" {
		return ErrInvalidBidNoID
	}
	if bid.ImpID == "" {
		return ErrInvalidBidNoImpID
	}

	return nil
}
