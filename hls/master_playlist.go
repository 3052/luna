package hls

import (
   "fmt"
   "net/url"
   "sort"
   "strconv"
   "strings"
)

// ExtStream represents a single media playlist (URI). It aggregates information
// from all #EXT-X-STREAM-INF tags that point to the same URI. The primary
// attributes are taken from the variant with the lowest bandwidth.
type ExtStream struct {
   URI              *url.URL
   ID               int
   Bandwidth        int
   AverageBandwidth int
   Codecs           string
   Resolution       string
   FrameRate        string
   Subtitles        string   // Refers to a ExtMedia GROUP-ID for subtitles
   Audio            []string // A list of associated audio ExtMedia GROUP-IDs
}

// String returns a multi-line summary of the ExtStream.
func (s *ExtStream) String() string {
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
func (s *ExtStream) SortBandwidth() int {
   if s.AverageBandwidth > 0 {
      return s.AverageBandwidth
   }
   return s.Bandwidth
}

type MasterPlaylist struct {
   Streams     []*ExtStream
   Medias      []*ExtMedia
   SessionKeys []*Key
}

// ResolveURIs converts relative URLs to absolute URLs using the base URL.
func (mp *MasterPlaylist) ResolveURIs(base *url.URL) {
   for _, streamItem := range mp.Streams {
      if streamItem.URI != nil {
         streamItem.URI = base.ResolveReference(streamItem.URI)
      }
   }
   for _, renditionItem := range mp.Medias {
      if renditionItem.URI != nil {
         renditionItem.URI = base.ResolveReference(renditionItem.URI)
      }
   }
   for _, keyItem := range mp.SessionKeys {
      keyItem.resolve(base)
   }
}

// Sort sorts the Streams and Medias slices in place.
// Streams are sorted by their minimum average bandwidth (if available),
// otherwise falling back to minimum bandwidth.
// ExtMedias (Medias) are sorted by GroupID.
func (mp *MasterPlaylist) Sort() {
   sort.Slice(mp.Streams, func(i, j int) bool {
      return mp.Streams[i].SortBandwidth() < mp.Streams[j].SortBandwidth()
   })
   sort.Slice(mp.Medias, func(i, j int) bool {
      return mp.Medias[i].GroupID < mp.Medias[j].GroupID
   })
}

// ExtMedia represents an #EXT-X-MEDIA tag.
type ExtMedia struct {
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

// String returns a multi-line summary of the ExtMedia.
func (r *ExtMedia) String() string {
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
   streamMap := make(map[string]*ExtStream) // Map URL to ExtStream to handle grouping

   for i := 0; i < len(lines); i++ {
      line := lines[i]
      if strings.HasPrefix(line, "#EXT-X-MEDIA:") {
         rendition := parseRendition(line)
         rendition.ID = streamCounter
         streamCounter++
         masterPlaylist.Medias = append(masterPlaylist.Medias, rendition)
      } else if strings.HasPrefix(line, "#EXT-X-SESSION-KEY:") {
         masterPlaylist.SessionKeys = append(masterPlaylist.SessionKeys, parseKey(line))
      } else if strings.HasPrefix(line, "#EXT-X-STREAM-INF:") {
         attrs := parseAttributes(line, "#EXT-X-STREAM-INF:")

         if i+1 >= len(lines) { // Malformed, missing URI
            continue
         }
         i++
         uriLine := lines[i]

         stream, exists := streamMap[uriLine]
         if !exists {
            // First time seeing this URI, create a new ExtStream
            stream = &ExtStream{ID: streamCounter}
            streamCounter++
            if parsedURL, err := url.Parse(uriLine); err == nil {
               stream.URI = parsedURL
            }
            streamMap[uriLine] = stream
            masterPlaylist.Streams = append(masterPlaylist.Streams, stream)

            // This is the first so it's automatically the lowest bandwidth; populate all fields
            populateStreamAttributes(stream, attrs)
         }

         // Always add the AUDIO group from the current tag to the list.
         if audioGroup := attrs["AUDIO"]; audioGroup != "" {
            stream.Audio = append(stream.Audio, audioGroup)
         }

         // Check if this variant has a lower bandwidth than the one stored.
         // If so, update the stream's primary attributes.
         if bw, _ := strconv.Atoi(attrs["BANDWIDTH"]); exists && bw < stream.Bandwidth {
            populateStreamAttributes(stream, attrs)
         }
      }
   }
   return masterPlaylist, nil
}

// populateStreamAttributes updates a ExtStream's fields from a map of attributes.
func populateStreamAttributes(stream *ExtStream, attrs map[string]string) {
   stream.Codecs = attrs["CODECS"]
   stream.Resolution = attrs["RESOLUTION"]
   stream.FrameRate = attrs["FRAME-RATE"]
   stream.Subtitles = attrs["SUBTITLES"]
   stream.Bandwidth, _ = strconv.Atoi(attrs["BANDWIDTH"])
   stream.AverageBandwidth, _ = strconv.Atoi(attrs["AVERAGE-BANDWIDTH"])
}

func parseRendition(line string) *ExtMedia {
   attrs := parseAttributes(line, "#EXT-X-MEDIA:")
   newRendition := &ExtMedia{
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
         newRendition.URI = parsedURL
      }
   }
   return newRendition
}
