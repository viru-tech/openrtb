package openrtb

import "encoding/json"

// Content object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
// content. This object may be useful when syndicated content contains impressions and does
// not necessarily match the publisher's general content. The exchange might or might not have
// knowledge of the page where the content is running, as a result of the syndication method. For
// example might be a video impression embedded in an iframe on an unknown web property or device.
type Content struct {
	// ID uniquely identifying the content.
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Series string `json:"series,omitempty"`
	Season string `json:"season,omitempty"`
	Artist string `json:"artist,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Album  string `json:"album,omitempty"`
	// International Standard Recording Code conforming to ISO - 3901.
	ISRC string `json:"isrc,omitempty"`
	// URL of the content, for buy-side contextualization or review.
	URL string `json:"url,omitempty"`
	// Content rating (e.g., MPAA).
	ContentRating string `json:"contentrating,omitempty"`
	// User rating of the content (e.g., number of stars, likes, etc.).
	UserRating string `json:"userrating,omitempty"`
	// Comma separated list of keywords describing the content.
	Keywords string `json:"keywords,omitempty"`
	// Content language using ISO-639-1-alpha-2.
	Language string    `json:"language,omitempty"`
	Producer *Producer `json:"producer,omitempty"`
	// Array of IAB content categories that describe the content.
	Categories []ContentCategory `json:"cat,omitempty"`
	// Additional content data.
	Data []Data          `json:"data,omitempty"`
	Ext  json.RawMessage `json:"ext,omitempty"`
	// Production quality per IAB's classification.
	ProductionQuality ProductionQuality `json:"prodq,omitempty"`
	// Deprecated. Video quality per IAB's classification.
	VideoQuality ProductionQuality `json:"videoquality,omitempty"`
	// Type of content (game, video, text, etc.).
	Context ContentContext `json:"context,omitempty"`
	// Media rating per QAG guidelines.
	MediaRating IQGRating `json:"qagmediarating,omitempty"`
	// Episode number (typically applies to video content).
	Episode int `json:"episode,omitempty"`
	// 0 = not live, 1 = content is live (e.g., stream, live blog).
	LiveStream int `json:"livestream,omitempty"`
	// 0 = indirect, 1 = direct.
	SourceRelationship int `json:"sourcerelationship,omitempty"`
	// Length of content in seconds; appropriate for video or audio.
	Length int `json:"len,omitempty"`
	// Indicator of whether the content is embeddable (e.g., an embeddable video player), where 0 = no, 1 = yes.
	Embeddable int `json:"embeddable,omitempty"`
}
