package hls

import (
   "cmp"
   "fmt"
   "net/url"
   "strconv"
   "strings"
)

func parseMediaTag(line string) *Media {
   attrs := parseAttributes(line, "#EXT-X-MEDIA:")
   newMedia := &Media{
      Type:       attrs["TYPE"],
      GroupId:    attrs["GROUP-ID"],
      Name:       attrs["NAME"],
      Language:   attrs["LANGUAGE"],
      Channels:   attrs["CHANNELS"],
      AutoSelect: attrs["AUTOSELECT"] == "YES",
   }
   if value, ok := attrs["URI"]; ok && value != "" {
      if parsedUrl, err := url.Parse(value); err == nil {
         newMedia.Uri = parsedUrl
      }
   }
   return newMedia
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

   builder.WriteString(fmt.Sprintf("\nid = %d", s.Id))
   return builder.String()
}

func (mp *MasterPlaylist) ResolveUris(base *url.URL) {
   for _, streamItem := range mp.StreamInfs {
      if streamItem.Uri != nil {
         streamItem.Uri = base.ResolveReference(streamItem.Uri)
      }
   }
   for _, mediaItem := range mp.Medias {
      if mediaItem.Uri != nil {
         mediaItem.Uri = base.ResolveReference(mediaItem.Uri)
      }
   }
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
   if r.GroupId != "" {
      builder.WriteString("\ngroup = ")
      builder.WriteString(r.GroupId)
   }
   builder.WriteString("\nid = ")
   builder.WriteString(strconv.Itoa(r.Id))
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
         media.Id = streamCounter
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
            stream = &StreamInf{Id: streamCounter}
            streamCounter++
            if parsedUrl, err := url.Parse(uriLine); err == nil {
               stream.Uri = parsedUrl
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

type MasterPlaylist struct {
   Medias     []*Media
   StreamInfs []*StreamInf
}

// Media represents an #EXT-X-MEDIA tag.
type Media struct {
   AutoSelect bool
   Channels   string
   GroupId    string
   Id         int
   Language   string
   Name       string
   Type       string
   Uri        *url.URL
}

// StreamInf represents a single media playlist (URI) from a #EXT-X-STREAM-INF tag.
// It aggregates information from all tags that point to the same URI. The primary
// attributes are taken from the variant with the lowest bandwidth.
type StreamInf struct {
   Audio            []string // A list of associated audio Media GROUP-IDs
   AverageBandwidth int
   Bandwidth        int
   Codecs           string
   FrameRate        string
   Id               int
   Resolution       string
   Subtitles        string // Refers to a Media GROUP-ID for subtitles
   Uri              *url.URL
}

func Bandwidth(s1, s2 *StreamInf) int {
   return cmp.Or(
      s1.AverageBandwidth-s2.AverageBandwidth,
      s1.Bandwidth-s2.Bandwidth,
   )
}

func GroupId(m1, m2 *Media) int {
   return cmp.Compare(m1.GroupId, m2.GroupId)
}
