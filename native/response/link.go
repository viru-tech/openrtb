package response

import "encoding/json"

// Link object contains response link.
type Link struct {
	// Landing URL of the clickable link.
	URL string `json:"url"`
	// List of third-party tracker URLs to be fired on click of the URL.
	ClickTrackers []string `json:"clicktrackers,omitempty"`
	// Fallback URL for deeplink. To be used if the URL given in url is not supported by the device.
	FallbackURL string          `json:"fallback,omitempty"`
	Ext         json.RawMessage `json:"ext,omitempty"`
}
