package request

import "encoding/json"

// DataTypeID enum.
type DataTypeID int

// DataTypeID enum values.
const (
	// DataTypeSponsored is used for data that should contain the brand name of the sponsor.
	DataTypeSponsored DataTypeID = 1
	// DataTypeDesc is used for descriptive text associated with the product or service being advertised.
	DataTypeDesc DataTypeID = 2
	// DataTypeRating is used for rating of the product being offered to the user.
	DataTypeRating DataTypeID = 3
	// DataTypeLikes is used for number of social ratings or “likes” of the product being offered to the user.
	DataTypeLikes DataTypeID = 4
	// DataTypeDownloads is used for number downloads/installs of this product.
	DataTypeDownloads DataTypeID = 5
	// DataTypePrice is used for price for product / app / in-app purchase.
	// Value should include currency symbol in localized format.
	DataTypePrice DataTypeID = 6
	// DataTypeSalePrice is used for sale price, that can be used together
	// with price to indicate a discounted price compared to a regular price.
	// Value should include currency symbol in localised format.
	DataTypeSalePrice DataTypeID = 7
	// DataTypePhone is used for phone number.
	DataTypePhone DataTypeID = 8
	// DataTypeAddress is used for address.
	DataTypeAddress DataTypeID = 9
	// DataTypeDescAdditional is used for additional descriptive text
	// associated with the product or service being advertised.
	DataTypeDescAdditional DataTypeID = 10
	// DataTypeDisplayURL is used for display URL for the text ad.
	// To be used when sponsoring entity does not own the content,
	// i.e. sponsored by BRAND on SITE (where SITE is transmitted in this field).
	DataTypeDisplayURL DataTypeID = 11
	// DataTypeCTADesc is used for CTA description - descriptive text
	// describing a ‘call to action’ button for the destination URL.
	DataTypeCTADesc DataTypeID = 12
)

// Data is the native data object.
type Data struct {
	Ext    json.RawMessage `json:"ext,omitempty"`
	TypeID DataTypeID      `json:"type"` // Type ID of the element supported by the publisher. The publisher can display this information in an appropriate format
	Length int             `json:"len"`  // Maximum length of the text in the element’s response
}
