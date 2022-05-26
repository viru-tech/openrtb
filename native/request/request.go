package request

import (
	"bytes"
	"encoding/json"
)

// LayoutID enum.
type LayoutID int

// LayoutID enum values.
const (
	LayoutContentWall          LayoutID = 1
	LayoutAppWall              LayoutID = 2
	LayoutNewsFeed             LayoutID = 3
	LayoutChatList             LayoutID = 4
	LayoutCarousel             LayoutID = 5
	LayoutContentStream        LayoutID = 6
	LayoutGridAdjoiningContent LayoutID = 7
)

// AdUnitID enum.
type AdUnitID int

// AdUnitID enum values.
const (
	AdUnitPaidSearch            AdUnitID = 1
	AdUnitRecommendationWidget  AdUnitID = 2
	AdUnitPromotedListings      AdUnitID = 3
	AdUnitInAdWithNativeElement AdUnitID = 4
	AdUnitCustom                AdUnitID = 5
)

// ContextTypeID enum.
type ContextTypeID int

// ContextTypeID enum values.
const (
	ContextTypeContent ContextTypeID = 1 // newsfeed, article, image gallery, video gallery
	ContextTypeSocial  ContextTypeID = 2 // social network feed, email, chat
	ContextTypeProduct ContextTypeID = 3 // product listings, details, recommendations, reviews

	ContextSubTypeGeneral        ContextTypeID = 10
	ContextSubTypeArticle        ContextTypeID = 11
	ContextSubTypeVideo          ContextTypeID = 12
	ContextSubTypeAudio          ContextTypeID = 13
	ContextSubTypeImage          ContextTypeID = 14
	ContextSubTypeUserGenerated  ContextTypeID = 15
	ContextSubTypeSocial         ContextTypeID = 20
	ContextSubTypeEmail          ContextTypeID = 21
	ContextSubTypeChat           ContextTypeID = 22
	ContextSubTypeSelling        ContextTypeID = 30
	ContextSubTypeAppStore       ContextTypeID = 31
	ContextSubTypeProductReviews ContextTypeID = 32
)

// PlacementTypeID enum.
type PlacementTypeID int

// PlacementTypeID enum values.
const (
	// PlacementTypeInFeed is used when placement is in the feed of content,
	// for example as an item inside the organic feed/grid/listing/carousel.
	PlacementTypeInFeed PlacementTypeID = 1
	// PlacementTypeAtomic is used when placement is in the atomic unit of the
	// content, i.e. in the article page or single image page.
	PlacementTypeAtomic PlacementTypeID = 2
	// PlacementTypeOutside is used when placement is outside the core content,
	// for example in the ads section on the right rail, as a banner-style
	// placement near the content, etc.
	PlacementTypeOutside PlacementTypeID = 3
	// PlacementTypeRecommendation is used when placement is in a recommendation
	// widget, most commonly presented below the article content.
	PlacementTypeRecommendation PlacementTypeID = 4
)

// Request object represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as
// a native type impression. At the publisher’s discretion, that same impression may also be offered as
// banner and/or video by also including as Imp subordinates the Banner and/or Video objects,
// respectively. However, any given bid for the impression must conform to one of the offered types.
type Request struct {
	// Version of the Native Markup.
	Version string `json:"ver,omitempty"`
	// The number of identical placements in this Layout.
	PlacementCount json.Number `json:"plcmtcnt,omitempty"`
	// An array of Asset Objects.
	Assets []Asset         `json:"assets"`
	Ext    json.RawMessage `json:"ext,omitempty"`
	// Deprecated. The Layout ID of the native ad.
	LayoutID LayoutID `json:"layout,omitempty"`
	// Deprecated. The Ad unit ID of the native ad.
	AdUnitID AdUnitID `json:"adunit,omitempty"`
	// The context in which the ad appears.
	ContextTypeID ContextTypeID `json:"context,omitempty"`
	// A more detailed context in which the ad appears.
	ContextSubTypeID ContextTypeID `json:"contextsubtype,omitempty"`
	// The design/format/layout of the ad unit being offered.
	PlacementTypeID PlacementTypeID `json:"plcmttype,omitempty"`
	// 0 for the first ad, 1 for the second ad, and so on.
	Sequence int `json:"seq,omitempty"`

	// Prior to VERSION 1.1, the specification could be interpreted
	// as requiring the native request to have a root node with a single
	// field "native" that would contain the Request as its value.
	IsWrapped bool `json:"-"`
}

type (
	jsonRequest Request

	legacyNative struct {
		Native Request `json:"native"`
	}
)

//nolint:gochecknoglobals
var (
	nativeBytes = []byte("native")
	nullBytes   = []byte("null")
)

// UnmarshalJSON implements json.Unmarshaler.
func (r *Request) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		return nil
	}

	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		// this must be a quoted string
		var jsonS string
		if err := json.Unmarshal(data, &jsonS); err != nil {
			return err
		}

		data = []byte(jsonS)
	}

	if bytes.Contains(data, nativeBytes) {
		var jn legacyNative
		if err := json.Unmarshal(data, &jn); err != nil {
			return err
		}

		*r = jn.Native
		r.IsWrapped = true
		return nil
	}

	var jn jsonRequest
	if err := json.Unmarshal(data, &jn); err != nil {
		return err
	}

	*r = (Request)(jn)
	return nil
}
