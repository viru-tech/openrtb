package request

import (
	"encoding/json"

	"github.com/bsm/openrtb/v3"
)

// Video is the native video object.
type Video struct {
	// Whitelist of content MIME types supported.
	MIMEs []string `json:"mimes,omitempty"`
	// Video bid response protocols.
	Protocols []openrtb.Protocol `json:"protocols,omitempty"`
	Ext       json.RawMessage    `json:"ext,omitempty"`
	// Minimum video ad duration in seconds.
	MinDuration int `json:"minduration,omitempty"`
	// Maximum video ad duration in seconds.
	MaxDuration int `json:"maxduration,omitempty"`
}
