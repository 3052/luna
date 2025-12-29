package hls

import (
   "fmt"
   "net/url"
   "sort"
   "strconv"
   "strings"
)

// StreamInf represents a single media playlist (URI) from a #EXT-X-STREAM-INF tag.
// It aggregates information from all tags that point to the same URI. The primary
// attributes are taken from the variant with the lowest bandwidth.
type StreamInf struct {
   URI              *url.URL
   ID               int
   Bandwidth        int
   AverageBandwidth int
   Codecs           string
   Resolution       string
   FrameRate        string
   Subtitles        string   // Refers to a Media GROUP-ID for subtitles
   Audio            []string // A list of associated audio Media GROUP-IDs
}

// String returns a multi-line summary of the StreamInf.
func (s *StreamInf) String() string {
   var builder strings.Builder

   if s.AverageBandwidth > 0 {
      builder.WriteString("average_bandwidth = ")
      builder.WriteString(strconv.Itoa(s.AverageBandwidth))
      builder.WriteString("\n")
   }

   builder.WriteString("bandwidth = ")
   builder.WriteString(strconv.Itoa(s.Bandwidth))

   if s.Resolution != "" {
      builder.WriteString("\nresolution = ")
      builder.WriteString(s.Resolution)
   }

   if s.Codecs != "" {
      videoCodec, _, _ := strings.Cut(s.Codecs, ",")
      builder.WriteString("\ncodecs = ")
      builder.WriteString(videoCodec)
   }

   builder.WriteString(fmt.Sprintf("\nid = %d", s.ID))
   return builder.String()
}

// SortBandwidth determines the value to use for sorting, prioritizing average bandwidth.
func (s *StreamInf) SortBandwidth() int {
   if s.AverageBandwidth > 0 {
      return s.AverageBandwidth
   }
   return s.Bandwidth
}

type MasterPlaylist struct {
   StreamInfs []*StreamInf
   Medias     []*Media
}

// ResolveURIs converts relative URLs to absolute URLs using the base URL.
func (mp *MasterPlaylist) ResolveURIs(base *url.URL) {
   for _, streamItem := range mp.StreamInfs {
      if streamItem.URI != nil {
         streamItem.URI = base.ResolveReference(streamItem.URI)
      }
   }
   for _, mediaItem := range mp.Medias {
      if mediaItem.URI != nil {
         mediaItem.URI = base.ResolveReference(mediaItem.URI)
      }
   }
}

// Sort sorts the StreamInfs and Medias slices in place.
// StreamInfs are sorted by their minimum average bandwidth (if available),
// otherwise falling back to minimum bandwidth.
// Medias are sorted by GroupID.
func (mp *MasterPlaylist) Sort() {
   sort.Slice(mp.StreamInfs, func(i, j int) bool {
      return mp.StreamInfs[i].SortBandwidth() < mp.StreamInfs[j].SortBandwidth()
   })
   sort.Slice(mp.Medias, func(i, j int) bool {
      return mp.Medias[i].GroupID < mp.Medias[j].GroupID
   })
}

// Media represents an #EXT-X-MEDIA tag.
type Media struct {
   Type            string
   GroupID         string
   Name            string
   Language        string
   URI             *url.URL
   AutoSelect      bool
   Default         bool
   Forced          bool
   Channels        string
   Characteristics string
   ID              int
}

// String returns a multi-line summary of the Media.
func (r *Media) String() string {
   var builder strings.Builder
   builder.WriteString("type = ")
   builder.WriteString(r.Type)
   if r.Name != "" {
      builder.WriteString("\nname = ")
      builder.WriteString(r.Name)
   }
   if r.Language != "" {
      builder.WriteString("\nlang = ")
      builder.WriteString(r.Language)
   }
   if r.GroupID != "" {
      builder.WriteString("\ngroup = ")
      builder.WriteString(r.GroupID)
   }
   builder.WriteString("\nid = ")
   builder.WriteString(strconv.Itoa(r.ID))
   return builder.String()
}

func parseMaster(lines []string) (*MasterPlaylist, error) {
   masterPlaylist := &MasterPlaylist{}
   streamCounter := 0
   streamMap := make(map[string]*StreamInf) // Map URL to StreamInf to handle grouping

   for i := 0; i < len(lines); i++ {
      line := lines[i]
      if strings.HasPrefix(line, "#EXT-X-MEDIA:") {
         media := parseMediaTag(line)
         media.ID = streamCounter
         streamCounter++
         masterPlaylist.Medias = append(masterPlaylist.Medias, media)
      } else if strings.HasPrefix(line, "#EXT-X-STREAM-INF:") {
         attrs := parseAttributes(line, "#EXT-X-STREAM-INF:")

         if i+1 >= len(lines) { // Malformed, missing URI
            continue
         }
         i++
         uriLine := lines[i]

         stream, exists := streamMap[uriLine]
         if !exists {
            // First time seeing this URI, create a new StreamInf
            stream = &StreamInf{ID: streamCounter}
            streamCounter++
            if parsedURL, err := url.Parse(uriLine); err == nil {
               stream.URI = parsedURL
            }
            streamMap[uriLine] = stream
            masterPlaylist.StreamInfs = append(masterPlaylist.StreamInfs, stream)

            // This is the first so it's automatically the lowest bandwidth; populate all fields
            populateStreamInfAttributes(stream, attrs)
         }

         // Always add the AUDIO group from the current tag to the list.
         if audioGroup := attrs["AUDIO"]; audioGroup != "" {
            stream.Audio = append(stream.Audio, audioGroup)
         }

         // Check if this variant has a lower bandwidth than the one stored.
         // If so, update the stream's primary attributes.
         if bw, _ := strconv.Atoi(attrs["BANDWIDTH"]); exists && bw < stream.Bandwidth {
            populateStreamInfAttributes(stream, attrs)
         }
      }
   }
   return masterPlaylist, nil
}

// populateStreamInfAttributes updates a StreamInf's fields from a map of attributes.
func populateStreamInfAttributes(stream *StreamInf, attrs map[string]string) {
   stream.Codecs = attrs["CODECS"]
   stream.Resolution = attrs["RESOLUTION"]
   stream.FrameRate = attrs["FRAME-RATE"]
   stream.Subtitles = attrs["SUBTITLES"]
   stream.Bandwidth, _ = strconv.Atoi(attrs["BANDWIDTH"])
   stream.AverageBandwidth, _ = strconv.Atoi(attrs["AVERAGE-BANDWIDTH"])
}

func parseMediaTag(line string) *Media {
   attrs := parseAttributes(line, "#EXT-X-MEDIA:")
   newMedia := &Media{
      Type:            attrs["TYPE"],
      GroupID:         attrs["GROUP-ID"],
      Name:            attrs["NAME"],
      Language:        attrs["LANGUAGE"],
      Channels:        attrs["CHANNELS"],
      Characteristics: attrs["CHARACTERISTICS"],
      AutoSelect:      attrs["AUTOSELECT"] == "YES",
      Default:         attrs["DEFAULT"] == "YES",
      Forced:          attrs["FORCED"] == "YES",
   }
   if value, ok := attrs["URI"]; ok && value != "" {
      if parsedURL, err := url.Parse(value); err == nil {
         newMedia.URI = parsedURL
      }
   }
   return newMedia
}
