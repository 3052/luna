// Package dash models the data structures of an MPEG-DASH
// MPD (Media Presentation Description) manifest.
//
// This package currently contains data models only.
package dash

import (
   "encoding/xml"
   "time"
)

// MPD presentation types.
const (
   MPDTypeStatic  = "static"
   MPDTypeDynamic = "dynamic"
)

// Content types used by AdaptationSet.ContentType.
const (
   ContentTypeAudio = "audio"
   ContentTypeVideo = "video"
   ContentTypeText  = "text"
   ContentTypeImage = "image"
)

// AdaptationSet groups alternative representations of the same media
// content (e.g. different bitrates of the same video).
type AdaptationSet struct {
   // ID is the adaptation set identifier.
   ID uint32 `xml:"id,attr"`
   // ContentType is one of "audio", "video", "text" or "image".
   ContentType string `xml:"contentType,attr"`
   // Lang is the language of the content (e.g. "en-US").
   Lang string `xml:"lang,attr"`

   // SegmentAlignment indicates that segments are aligned across
   // representations.
   SegmentAlignment bool `xml:"segmentAlignment,attr"`
   // SubsegmentAlignment indicates subsegments are aligned across
   // representations.
   SubsegmentAlignment bool `xml:"subsegmentAlignment,attr"`
   // SubsegmentStartsWithSAP: 0 = no SAP required, 1 = each subsegment
   // starts with a SAP, 2 = SAP type 1 or 2.
   SubsegmentStartsWithSAP uint8 `xml:"subsegmentStartsWithSAP,attr"`

   // MaxWidth / MaxHeight are the maximum resolution across
   // representations.
   MaxWidth  uint32 `xml:"maxWidth,attr"`
   MaxHeight uint32 `xml:"maxHeight,attr"`
   // FrameRate is the frame rate, possibly as a ratio (e.g. "24000/1001").
   FrameRate string `xml:"frameRate,attr"`
   // PAR is the pixel aspect ratio (e.g. "24:13").
   PAR string `xml:"par,attr"`

   // ContentProtection entries hold DRM system information.
   ContentProtections []ContentProtection `xml:"ContentProtection"`
   // Roles describe the purpose of the content (e.g. "main", "caption").
   Roles []Role `xml:"Role"`
   // SupplementalProperties carry additional properties
   // (e.g. adaptation-set-switching).
   SupplementalProperties []SupplementalProperty `xml:"SupplementalProperty"`
   // Labels are human-readable labels (e.g. "en-US CC").
   Labels []Label `xml:"Label"`
   // SegmentTemplate, when present at the adaptation set level, applies
   // to all representations (e.g. image thumbnails).
   SegmentTemplate *SegmentTemplate `xml:"SegmentTemplate"`

   // Representations are the alternative encodings of this content.
   Representations []Representation `xml:"Representation"`
}

// ContentProtection describes a DRM system for the content.
type ContentProtection struct {
   // SchemeIDURI identifies the protection scheme
   // (e.g. "urn:mpeg:dash:mp4protection:2011" or a DRM system UUID).
   SchemeIDURI string `xml:"schemeIdUri,attr"`
   // Value is the scheme-specific value (e.g. "cenc").
   Value string `xml:"value,attr"`
   // DefaultKID is the CENC default key ID (cenc:default_KID).
   DefaultKID string `xml:"cenc:default_KID,attr"`
   // PSSH holds a base64-encoded CENC PSSH box (cenc:pssh).
   PSSH string `xml:"cenc:pssh"`
}

// EssentialProperty carries information the client must understand
// to process the content (e.g. thumbnail tile layout "1x1").
type EssentialProperty struct {
   SchemeIDURI string `xml:"schemeIdUri,attr"`
   Value       string `xml:"value,attr"`
}

// Label is a human-readable label for an adaptation set.
type Label struct {
   // Text is the label content (e.g. "en-US CC").
   Text string `xml:",chardata"`
}

// MPD is the root element of a DASH manifest.
type MPD struct {
   XMLName xml.Name `xml:"MPD"`

   // XMLNS is the default MPD namespace ("urn:mpeg:dash:schema:mpd:2011").
   XMLNS string `xml:"xmlns,attr"`
   // CencXMLNS is the CENC namespace ("urn:mpeg:cenc:2013").
   CencXMLNS string `xml:"xmlns:cenc,attr"`

   // Type is either "static" or "dynamic".
   Type string `xml:"type,attr"`
   // Profiles is a space-separated list of profile URIs.
   Profiles string `xml:"profiles,attr"`
   // MinBufferTime is the required buffer time (e.g. PT2S).
   MinBufferTime time.Duration `xml:"minBufferTime,attr"`
   // MediaPresentationDuration is the total duration of the presentation.
   MediaPresentationDuration time.Duration `xml:"mediaPresentationDuration,attr"`

   // Periods are the presentation intervals, in order.
   Periods []Period `xml:"Period"`
}

// Period is an interval of the media presentation.
type Period struct {
   // ID is the period identifier.
   ID string `xml:"id,attr"`
   // Start is the period start offset from the beginning of the presentation.
   Start time.Duration `xml:"start,attr"`
   // Duration is the duration of the period.
   Duration time.Duration `xml:"duration,attr"`

   // AdaptationSets are the alternative encodings of the media content
   // within this period.
   AdaptationSets []AdaptationSet `xml:"AdaptationSet"`
}

// Role describes the role of an adaptation set
// (e.g. "main" audio or "caption" text).
type Role struct {
   SchemeIDURI string `xml:"schemeIdUri,attr"`
   Value       string `xml:"value,attr"`
}

// SupplementalProperty carries supplementary information
// (e.g. adaptation-set-switching targets).
type SupplementalProperty struct {
   SchemeIDURI string `xml:"schemeIdUri,attr"`
   Value       string `xml:"value,attr"`
}

// dash/mpd.go
