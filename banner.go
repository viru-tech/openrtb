package openrtb

import "encoding/json"

// Banner object must be included directly in the impression object if the impression offered
// for auction is display or rich media, or it may be optionally embedded in the video object to
// describe the companion banners available for the linear or non-linear video ad.  The banner
// object may include a unique identifier; this can be useful if these IDs can be leveraged in the
// VAST response to dictate placement of the companion creatives when multiple companion ad
// opportunities of the same size are available on a page.
type Banner struct {
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// Array of format objects representing the banner sizes permitted.
	Formats []Format `json:"format,omitempty"`
	// Blocked banner types.
	BlockedTypes []BannerType `json:"btype,omitempty"`
	// Blocked creative attributes.
	BlockedAttrs []CreativeAttribute `json:"battr,omitempty"`
	// Whitelist of content MIME types supported.
	MIMEs []string `json:"mimes,omitempty"`
	// Specify properties for an expandable ad.
	ExpDirs []ExpDir `json:"expdir,omitempty"`
	// List of supported API frameworks.
	APIs     []APIFramework  `json:"api,omitempty"`
	Ext      json.RawMessage `json:"ext,omitempty"`
	Position AdPosition      `json:"pos,omitempty"`
	// Default: 0 ("1": Delivered in top frame, "0": Elsewhere).
	TopFrame int `json:"topframe,omitempty"`
	Width    int `json:"w,omitempty"`
	Height   int `json:"h,omitempty"`
	// Deprecated. Width maximum.
	WidthMax int `json:"wmax,omitempty"`
	// Deprecated. Height maximum.
	HeightMax int `json:"hmax,omitempty"`
	// Deprecated. Width minimum.
	WidthMin int `json:"wmin,omitempty"`
	// Deprecated. Height minimum.
	HeightMin int `json:"hmin,omitempty"`
	// Represents the relationship with video. 0 = concurrent, 1 = end-card.
	VCM int `json:"vcm,omitempty"`
}
