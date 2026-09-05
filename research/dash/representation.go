package dash

import "slices"

// medianBandwidth returns the median bandwidth; even count averages the
// two middle values, rounded down.
func medianBandwidth(reps []Representation) uint32 {
   vals := make([]uint32, len(reps))
   for i, r := range reps {
      vals[i] = r.Bandwidth
   }
   slices.Sort(vals)
   n := len(vals)
   if n%2 == 1 {
      return vals[n/2]
   }
   return uint32((uint64(vals[n/2-1]) + uint64(vals[n/2])) / 2)
}

// BaseURL is a media URL for single-file addressing.
type BaseURL struct {
   // URL is the relative or absolute media URL (e.g. "v/2fda2a/v0.mp4").
   URL string `xml:",chardata"`
}

// Initialization is the byte range of the initialization data.
type Initialization struct {
   // Range is the byte range within the media file (e.g. "0-657").
   Range string `xml:"range,attr"`
}

// Representation is one encoding variant within an adaptation set.
type Representation struct {
   // ID is the representation identifier.
   ID string `xml:"id,attr"`
   // Codecs is the codec identifier (e.g. "avc1.64001f", "hvc1.2.4.L120.90",
   // "mp4a.40.5", "ec-3").
   Codecs string `xml:"codecs,attr"`
   // MimeType is the media type (e.g. "video/mp4", "audio/mp4",
   // "text/vtt", "image/jpeg").
   MimeType string `xml:"mimeType,attr"`
   // Bandwidth is the peak bandwidth in bits per second.
   Bandwidth uint32 `xml:"bandwidth,attr"`

   // Width / Height are the video dimensions.
   Width  uint32 `xml:"width,attr"`
   Height uint32 `xml:"height,attr"`
   // SAR is the sample aspect ratio (e.g. "1:1").
   SAR string `xml:"sar,attr"`
   // AudioSamplingRate is the audio sample rate in Hz.
   AudioSamplingRate uint32 `xml:"audioSamplingRate,attr"`

   // BaseURL points to the media resource (single-file addressing).
   BaseURL *BaseURL `xml:"BaseURL"`
   // SegmentBase describes byte-range addressing for a single file.
   SegmentBase *SegmentBase `xml:"SegmentBase"`
   // SegmentTemplate describes template-based multi-segment addressing.
   SegmentTemplate *SegmentTemplate `xml:"SegmentTemplate"`
}

// SegmentBase describes byte-range addressing of a single media file.
type SegmentBase struct {
   // IndexRange is the byte range of the segment index
   // (e.g. "658-17705").
   IndexRange string `xml:"indexRange,attr"`
   // Timescale is the number of ticks per second.
   Timescale uint64 `xml:"timescale,attr"`
   // PresentationTimeOffset is the offset of the first presentation
   // time in ticks.
   PresentationTimeOffset uint64 `xml:"presentationTimeOffset,attr"`
   // Initialization is the byte range of the initialization segment
   // (e.g. "0-657").
   Initialization *Initialization `xml:"Initialization"`
}

// SegmentTemplate describes template-based, multi-segment addressing.
type SegmentTemplate struct {
   // Media is the segment URL template
   // (e.g. "t/aa517e/t0/$Number$.vtt" or "i/1971b3/images-$Number$.jpg").
   Media string `xml:"media,attr"`
   // StartNumber is the number of the first segment.
   StartNumber uint32 `xml:"startNumber,attr"`
   // Duration is the constant segment duration.
   Duration uint64 `xml:"duration,attr"`
   // Timescale is the number of ticks per second.
   Timescale uint64 `xml:"timescale,attr"`
   // PresentationTimeOffset is the offset of the first presentation
   // time in ticks.
   PresentationTimeOffset uint64 `xml:"presentationTimeOffset,attr"`

   // SegmentTimeline lists explicit segment timing entries, used when
   // segments have non-uniform durations.
   SegmentTimeline *SegmentTimeline `xml:"SegmentTimeline"`
}

// SegmentTimeline is a list of explicit segment timing entries.
type SegmentTimeline struct {
   // Entries are the segment timing entries, in order.
   Entries []SegmentTimelineEntry `xml:"S"`
}

// SegmentTimelineEntry is a single <S> element of a SegmentTimeline.
type SegmentTimelineEntry struct {
   // T is the presentation time of the first segment in ticks.
   T uint64 `xml:"t,attr"`
   // D is the segment duration in ticks.
   D uint64 `xml:"d,attr"`
   // R is the repeat count: the entry describes R+1 consecutive segments.
   R int `xml:"r,attr"`
}

// dash/representation.go
