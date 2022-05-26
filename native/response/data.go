package response

import "encoding/json"

// Data object contains response data.
type Data struct {
	// The optional formatted string name of the data type to be displayed.
	Label string `json:"label,omitempty"`
	// The formatted string of data to be displayed.
	// Can contain a formatted value such as “5 stars” or “$10” or “3.4 stars out of 5”.
	Value string          `json:"value"`
	Ext   json.RawMessage `json:"ext,omitempty"`
}
