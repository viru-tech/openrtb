package request

import "encoding/json"

// Title is the native title object.
type Title struct {
	Ext    json.RawMessage `json:"ext,omitempty"`
	Length int             `json:"len"` // Maximum length of the text in the title element
}
