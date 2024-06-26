package response

// EventTracker is the tracking object to run with the ad.
type EventTracker struct {
	URL    string `json:"url"`
	Event  uint8  `json:"event"`
	Method uint8  `json:"method"`
}
