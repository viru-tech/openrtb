package openrtb

import "encoding/json"

// Device object provides information pertaining to the device including its hardware,
// platform, location, and carrier. This device can refer to a mobile handset, a desktop computer,
// set top box or other digital device.
type Device struct {
	UA        string `json:"ua,omitempty"`
	IP        string `json:"ip,omitempty"`
	IPv6      string `json:"ipv6,omitempty"`
	Make      string `json:"make,omitempty"`
	Model     string `json:"model,omitempty"`
	OS        string `json:"os,omitempty"`
	OSVersion string `json:"osv,omitempty"`
	// Hardware version of the device (e.g., "5S" for iPhone 5S).
	HWVersion    string `json:"hwv,omitempty"`
	FlashVersion string `json:"flashver,omitempty"`
	Language     string `json:"language,omitempty"`
	// Carrier or ISP derived from the IP address.
	Carrier string `json:"carrier,omitempty"`
	// Mobile carrier as the concatenated MCC-MNC code,
	// e.g. "310-005" identifies Verizon Wireless CDMA in the USA.
	MCCMNC string `json:"mccmnc,omitempty"`
	// Native identifier for advertisers.
	IFA string `json:"ifa,omitempty"`
	// SHA1 hashed device ID.
	IDSHA1 string `json:"didsha1,omitempty"`
	// MD5 hashed device ID.
	IDMD5 string `json:"didmd5,omitempty"`
	// SHA1 hashed platform device ID.
	PIDSHA1 string `json:"dpidsha1,omitempty"`
	// MD5 hashed platform device ID.
	PIDMD5 string `json:"dpidmd5,omitempty"`
	// SHA1 hashed device ID; IMEI when available, else MEID or ESN.
	MacSHA1 string `json:"macsha1,omitempty"`
	// MD5 hashed device ID; IMEI when available, else MEID or ESN.
	MacMD5 string `json:"macmd5,omitempty"`
	// Location of the device assumed to be the user’s current location.
	Geo *Geo            `json:"geo,omitempty"`
	Ext json.RawMessage `json:"ext,omitempty"`
	// "1": Do not track
	DNT int `json:"dnt,omitempty"`
	// "1": Limit Ad Tracking
	LMT        int        `json:"lmt,omitempty"`
	DeviceType DeviceType `json:"devicetype,omitempty"`
	// Physical height of the screen in pixels.
	Height int `json:"h,omitempty"`
	// Physical width of the screen in pixels.
	Width int `json:"w,omitempty"`
	// Screen size as pixels per linear inch.
	PPI int `json:"ppi,omitempty"`
	// The ratio of physical pixels to device independent pixels.
	PixelRatio float64 `json:"pxratio,omitempty"`
	// Javascript status ("0": Disabled, "1": Enabled)
	JS int `json:"js,omitempty"`
	// Indicates if the geolocation API will be available to JavaScript code running in the banner,
	GeoFetch int      `json:"geofetch,omitempty"`
	ConnType ConnType `json:"connectiontype,omitempty"`
}
