package request

import "encoding/json"

// ImageTypeID enum.
type ImageTypeID int

// ImageTypeID enum values.
const (
	ImageTypeIcon ImageTypeID = 1 // Icon image
	ImageTypeLogo ImageTypeID = 2 // Logo image for the brand/app
	ImageTypeMain ImageTypeID = 3 // Large image preview for the ad
)

// Image is the native image object.
type Image struct {
	// Whitelist of content MIME types supported.
	MIMEs []string        `json:"mimes,omitempty"`
	Ext   json.RawMessage `json:"ext,omitempty"`
	// Type ID of the image element supported by the publisher.
	TypeID ImageTypeID `json:"type,omitempty"`

	// Either h/w or hmin/wmin should be transmitted. If only h/w is included, it
	// should be considered an exact requirement
	Width     int `json:"w,omitempty"`    // Width of the image in pixels.
	WidthMin  int `json:"wmin,omitempty"` // The minimum requested width of the image in pixels.
	Height    int `json:"h,omitempty"`    // Height of the image in pixels.
	HeightMin int `json:"hmin,omitempty"` // The minimum requested height of the image in pixels.
}
