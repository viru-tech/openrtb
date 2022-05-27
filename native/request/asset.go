package request

import "encoding/json"

// Asset is the main container object for each asset requested or supported by Exchange
// on behalf of the rendering client.  Only one of the {title,img,video,data}
// objects should be present in each object.  The id is to be unique within the
// AssetObject array so that the response can be aligned.
type Asset struct {
	// Title object for title assets.
	Title *Title `json:"title,omitempty"`
	// Image object for image assets.
	Image *Image `json:"img,omitempty"`
	// Video object for video assets.
	Video *Video `json:"video,omitempty"`
	// Data object for brand name, description, ratings, prices etc.
	Data *Data           `json:"data,omitempty"`
	Ext  json.RawMessage `json:"ext,omitempty"`
	// Unique asset ID, assigned by exchange.
	ID int `json:"id"`
	// Set to 1 if asset is required.
	Required int `json:"required,omitempty"`
}
