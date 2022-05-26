package response

import (
	"encoding/json"

	"github.com/viru-tech/openrtb/v3"
)

// Response is the native object is the top level JSON object which identifies a native response.
type Response struct {
	// Version of the Native Markup.
	Version openrtb.StringOrNumber `json:"ver,omitempty"`
	// An array of Asset Objects.
	Assets []Asset `json:"assets"`
	// Destination Link. This is default link object for the ad.
	Link Link `json:"link"`
	// Array of impression tracking URLs, expected to return a 1x1 image or 204 response.
	ImpTrackers []string `json:"imptrackers,omitempty"`
	// Optional JavaScript impression tracker. This is a valid HTML,
	// Javascript is already wrapped in <script> tags. It should be
	// executed at impression time, where it can be supported.
	JSTracker string          `json:"jstracker,omitempty"`
	Ext       json.RawMessage `json:"ext,omitempty"`
}
