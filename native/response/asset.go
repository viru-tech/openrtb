package response

import "encoding/json"

// Asset corresponds to the Asset Object in the request. The main container object for
// each asset requested or supported by Exchange on behalf of the rendering
// client. Any object that is required is to be flagged as such. Only one of the
// {title,img,video,data} objects should be present in each object. All others
// should be null/absent. The id is to be unique within the AssetObject array so
// that the response can be aligned.
type Asset struct {
	// Title object for title assets.
	Title *Title `json:"title,omitempty"`
	// Image object for image assets.
	Image *Image `json:"img,omitempty"`
	// Video object for video assets.
	Video *Video `json:"video,omitempty"`
	// Data object for brand name, description, ratings, prices etc.
	Data *Data `json:"data,omitempty"`
	// Link object for call to actions.
	// The link object applies if the asset item is activated (clicked).
	Link *Link           `json:"link,omitempty"`
	Ext  json.RawMessage `json:"ext,omitempty"`
	// Unique asset ID, assigned by exchange, must match one of the asset IDs in request.
	ID int `json:"id"`
	// Set to 1 if asset is required.
	Required int `json:"required,omitempty"`
}
