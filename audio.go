package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors.
var (
	ErrInvalidAudioNoMIMEs = errors.New("openrtb: audio has no mimes")
)

// Audio object must be included directly in the impression object.
type Audio struct {
	// Content MIME types supported.
	MIMEs []string `json:"mimes"`
	// Video bid response protocols.
	Protocols []Protocol `json:"protocols,omitempty"`
	// Blocked creative attributes.
	BlockedAttrs []CreativeAttribute `json:"battr,omitempty"`
	// List of supported delivery methods.
	Delivery       []ContentDelivery `json:"delivery,omitempty"`
	CompanionAds   []Banner          `json:"companionad,omitempty"`
	APIs           []APIFramework    `json:"api,omitempty"`
	CompanionTypes []CompanionType   `json:"companiontype,omitempty"`
	Ext            json.RawMessage   `json:"ext,omitempty"`
	// Minimum video ad duration in seconds.
	MinDuration int `json:"minduration,omitempty"`
	// Maximum video ad duration in seconds.
	MaxDuration int `json:"maxduration,omitempty"`
	// Indicates the start delay in seconds.
	StartDelay StartDelay `json:"startdelay,omitempty"`
	// Default: 1.
	Sequence int `json:"sequence,omitempty"`
	// Maximum extended video ad duration.
	MaxExtended int `json:"maxextended,omitempty"`
	// Minimum bit rate in Kbps.
	MinBitrate int `json:"minbitrate,omitempty"`
	// Maximum bit rate in Kbps.
	MaxBitrate int `json:"maxbitrate,omitempty"`
	// The maximum number of ads that can be played in an ad pod.
	MaxSequence int `json:"maxseq,omitempty"`
	// Type of audio feed.
	Feed FeedType `json:"feed,omitempty"`
	// Indicates if the ad is stitched with audio content or delivered independently.
	Stitched int `json:"stitched,omitempty"`
	// Volume normalization mode.
	VolumeNorm VolumeNorm `json:"nvol,omitempty"`
}

type jsonAudio Audio

// Validate validates the audio object.
func (a *Audio) Validate() error {
	if len(a.MIMEs) == 0 {
		return ErrInvalidAudioNoMIMEs
	}
	return nil
}

// MarshalJSON custom marshaling with normalization.
func (a *Audio) MarshalJSON() ([]byte, error) {
	a.normalize()
	return json.Marshal((*jsonAudio)(a))
}

// UnmarshalJSON custom unmarshalling with normalization.
func (a *Audio) UnmarshalJSON(data []byte) error {
	var h jsonAudio
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*a = (Audio)(h)
	a.normalize()
	return nil
}

func (a *Audio) normalize() {
	if a.Sequence == 0 {
		a.Sequence = 1
	}
}
