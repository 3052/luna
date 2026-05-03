package hls

import (
   "cmp"
   "fmt"
   "net/url"
   "strconv"
   "strings"
)

// String returns a multi-line summary of the Media.
func (m *Media) String() string {
   data := &strings.Builder{}
   fmt.Fprintln(data, "type:", m.Type)
   if m.Name != "" {
      fmt.Fprintln(data, "name:", m.Name)
   }
   if m.Language != "" {
      fmt.Fprintln(data, "lang:", m.Language)
   }
   if m.GroupId != "" {
      fmt.Fprintln(data, "group:", m.GroupId)
   }
   fmt.Fprint(data, "id: ", m.Id)
   return data.String()
}

// String returns a multi-line summary of the StreamInf.
func (s *StreamInf) String() string {
   data := &strings.Builder{}
   if s.AverageBandwidth > 0 {
      fmt.Fprintln(data, "average bandwidth:", s.AverageBandwidth)
   }
   fmt.Fprintln(data, "bandwidth:", s.Bandwidth)
   if s.Resolution != "" {
      fmt.Fprintln(data, "resolution:", s.Resolution)
   }
   if s.Codecs != "" {
      videoCodec, _, _ := strings.Cut(s.Codecs, ",")
      fmt.Fprintln(data, "codecs:", videoCodec)
   }
   fmt.Fprint(data, "id: ", s.Id)
   return data.String()
}

// DecodeMaster parses a Master Playlist.
func DecodeMaster(content string) (*MasterPlaylist, error) {
   lines := splitLines(content)
   return parseMaster(lines)
}

func parseMediaTag(line string) (*Media, error) {
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
      parsedUrl, err := url.Parse(value)
      if err != nil {
         return nil, fmt.Errorf("invalid URI in EXT-X-MEDIA: %w", err)
      }
      newMedia.Uri = parsedUrl
   }
   return newMedia, nil
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

func parseMaster(lines []string) (*MasterPlaylist, error) {
   masterPlaylist := &MasterPlaylist{}
   streamCounter := 0
   streamMap := make(map[string]*StreamInf) // Map URL to StreamInf to handle grouping

   for i := 0; i < len(lines); i++ {
      line := lines[i]
      if strings.HasPrefix(line, "#EXT-X-MEDIA:") {
         media, err := parseMediaTag(line)
         if err != nil {
            return nil, err
         }
         media.Id = streamCounter
         streamCounter++
         masterPlaylist.Medias = append(masterPlaylist.Medias, media)
      } else if strings.HasPrefix(line, "#EXT-X-STREAM-INF:") {
         attrs := parseAttributes(line, "#EXT-X-STREAM-INF:")

         if i+1 >= len(lines) { // Malformed, missing URI
            return nil, fmt.Errorf("malformed EXT-X-STREAM-INF: missing URI")
         }
         i++
         uriLine := lines[i]

         stream, exists := streamMap[uriLine]
         if !exists {
            // First time seeing this URI, create a new StreamInf
            stream = &StreamInf{Id: streamCounter}
            streamCounter++
            parsedUrl, err := url.Parse(uriLine)
            if err != nil {
               return nil, fmt.Errorf("invalid URI in EXT-X-STREAM-INF: %w", err)
            }
            stream.Uri = parsedUrl
            streamMap[uriLine] = stream
            masterPlaylist.StreamInfs = append(masterPlaylist.StreamInfs, stream)

            // This is the first so it's automatically the lowest bandwidth; populate all fields
            if err := populateStreamInfAttributes(stream, attrs); err != nil {
               return nil, err
            }
         }

         // Always add the AUDIO group from the current tag to the list.
         if audioGroup := attrs["AUDIO"]; audioGroup != "" {
            stream.Audio = append(stream.Audio, audioGroup)
         }

         // Check if this variant has a lower bandwidth than the one stored.
         // If so, update the stream's primary attributes.
         if bwStr := attrs["BANDWIDTH"]; bwStr != "" {
            bw, err := strconv.Atoi(bwStr)
            if err != nil {
               return nil, fmt.Errorf("invalid BANDWIDTH in EXT-X-STREAM-INF: %w", err)
            }
            if exists && bw < stream.Bandwidth {
               if err := populateStreamInfAttributes(stream, attrs); err != nil {
                  return nil, err
               }
            }
         }
      }
   }
   return masterPlaylist, nil
}

// populateStreamInfAttributes updates a StreamInf's fields from a map of attributes.
func populateStreamInfAttributes(stream *StreamInf, attrs map[string]string) error {
   stream.Codecs = attrs["CODECS"]
   stream.Resolution = attrs["RESOLUTION"]
   stream.FrameRate = attrs["FRAME-RATE"]
   stream.Subtitles = attrs["SUBTITLES"]

   if val := attrs["BANDWIDTH"]; val != "" {
      bw, err := strconv.Atoi(val)
      if err != nil {
         return fmt.Errorf("invalid BANDWIDTH: %w", err)
      }
      stream.Bandwidth = bw
   }

   if val := attrs["AVERAGE-BANDWIDTH"]; val != "" {
      abw, err := strconv.Atoi(val)
      if err != nil {
         return fmt.Errorf("invalid AVERAGE-BANDWIDTH: %w", err)
      }
      stream.AverageBandwidth = abw
   }
   return nil
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
