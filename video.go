package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors.
var (
	ErrInvalidVideoNoMIMEs     = errors.New("openrtb: video has no mimes")
	ErrInvalidVideoNoLinearity = errors.New("openrtb: video linearity missing")
	ErrInvalidVideoNoProtocols = errors.New("openrtb: video protocols missing")
)

// Video object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	// If exchange publisher has rules preventing letter boxing.
	BoxingAllowed *int `json:"boxingallowed,omitempty"`
	// List of allowed playback methods.
	PlaybackMethods []VideoPlayback `json:"playbackmethod,omitempty"`
	// List of supported delivery methods.
	Delivery []ContentDelivery `json:"delivery,omitempty"`
	// Content MIME types supported..
	MIMEs []string `json:"mimes,omitempty"`
	// Video bid response protocols.
	Protocols []Protocol `json:"protocols,omitempty"`
	// Blocked creative attributes.
	BlockedAttrs []CreativeAttribute `json:"battr,omitempty"`
	CompanionAds []Banner            `json:"companionad,omitempty"`
	// List of supported API frameworks.
	APIs           []APIFramework  `json:"api,omitempty"`
	CompanionTypes []CompanionType `json:"companiontype,omitempty"`
	Ext            json.RawMessage `json:"ext,omitempty"`
	// Deprecated. Video bid response protocols.
	Protocol Protocol `json:"protocol,omitempty"`
	// Minimum video ad duration in seconds.
	MinDuration int `json:"minduration,omitempty"`
	// Maximum video ad duration in seconds.
	MaxDuration int `json:"maxduration,omitempty"`
	// Width of the player in pixels.
	Width int `json:"w,omitempty"`
	// Height of the player in pixels.
	Height int `json:"h,omitempty"`
	// Indicates the start delay in seconds.
	StartDelay StartDelay `json:"startdelay,omitempty"`
	// Indicates whether the ad impression is linear or non-linear.
	Linearity VideoLinearity `json:"linearity,omitempty"`
	// Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	Skip int `json:"skip,omitempty"`
	// Videos of total duration greater than this number of seconds can be skippable.
	SkipMin int `json:"skipmin,omitempty"`
	// Number of seconds a video must play before skipping is enabled.
	SkipAfter int `json:"skipafter,omitempty"`
	// Default: 1.
	Sequence int `json:"sequence,omitempty"`
	// Maximum extended video ad duration.
	MaxExtended int `json:"maxextended,omitempty"`
	// Minimum bit rate in Kbps.
	MinBitrate int `json:"minbitrate,omitempty"`
	// Maximum bit rate in Kbps.
	MaxBitrate int `json:"maxbitrate,omitempty"`
	// Ad Position.
	Position AdPosition `json:"pos,omitempty"`
	// Video placement type.
	Placement VideoPlacement `json:"placement,omitempty"`
}

type jsonVideo Video

// Validate the object.
func (v *Video) Validate() error {
	if len(v.MIMEs) == 0 {
		return ErrInvalidVideoNoMIMEs
	}
	if v.Linearity == 0 {
		return ErrInvalidVideoNoLinearity
	}
	if v.Protocol == 0 && len(v.Protocols) == 0 {
		return ErrInvalidVideoNoProtocols
	}
	return nil
}

// GetBoxingAllowed returns the boxing-allowed indicator.
func (v *Video) GetBoxingAllowed() int {
	if v.BoxingAllowed != nil {
		return *v.BoxingAllowed
	}
	return 1
}

// MarshalJSON custom marshaling with normalization.
func (v *Video) MarshalJSON() ([]byte, error) {
	v.normalize()
	return json.Marshal((*jsonVideo)(v))
}

// UnmarshalJSON custom unmarshaling with normalization.
func (v *Video) UnmarshalJSON(data []byte) error {
	var h jsonVideo
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*v = (Video)(h)
	v.normalize()
	return nil
}

func (v *Video) normalize() {
	if v.Sequence == 0 {
		v.Sequence = 1
	}
	if v.Linearity == 0 {
		v.Linearity = VideoLinearityLinear
	}
}
