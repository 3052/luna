package hls

import (
   "fmt"
   "net/url"
   "strconv"
   "strings"
)

type MediaPlaylist struct {
   TargetDuration int
   MediaSequence  int
   Version        int
   PlaylistType   string
   Segments       []*Segment
   Keys           []*Key   // A slice of all keys found in the playlist
   Map            *url.URL // The playlist's initialization map
   EndList        bool
}

// ResolveURIs converts relative URLs to absolute URLs using the base URL.
func (mp *MediaPlaylist) ResolveURIs(base *url.URL) {
   for _, keyItem := range mp.Keys {
      keyItem.resolve(base)
   }
   for _, segmentItem := range mp.Segments {
      segmentItem.resolve(base)
   }
   if mp.Map != nil {
      mp.Map = base.ResolveReference(mp.Map)
   }
}

type Segment struct {
   URI      *url.URL
   Duration float64
   Title    string
}

// resolve updates the Segment's URI to be absolute.
func (s *Segment) resolve(base *url.URL) {
   if s.URI != nil {
      s.URI = base.ResolveReference(s.URI)
   }
}

func parseMedia(lines []string) (*MediaPlaylist, error) {
   mediaPlaylist := &MediaPlaylist{}

   for i := 0; i < len(lines); i++ {
      line := lines[i]
      switch {
      case strings.HasPrefix(line, "#EXT-X-VERSION:"):
         version, err := strconv.Atoi(strings.TrimPrefix(line, "#EXT-X-VERSION:"))
         if err != nil {
            return nil, fmt.Errorf("invalid EXT-X-VERSION: %w", err)
         }
         mediaPlaylist.Version = version
      case strings.HasPrefix(line, "#EXT-X-TARGETDURATION:"):
         duration, err := strconv.Atoi(strings.TrimPrefix(line, "#EXT-X-TARGETDURATION:"))
         if err != nil {
            return nil, fmt.Errorf("invalid EXT-X-TARGETDURATION: %w", err)
         }
         mediaPlaylist.TargetDuration = duration
      case strings.HasPrefix(line, "#EXT-X-MEDIA-SEQUENCE:"):
         sequence, err := strconv.Atoi(strings.TrimPrefix(line, "#EXT-X-MEDIA-SEQUENCE:"))
         if err != nil {
            return nil, fmt.Errorf("invalid EXT-X-MEDIA-SEQUENCE: %w", err)
         }
         mediaPlaylist.MediaSequence = sequence
      case strings.HasPrefix(line, "#EXT-X-PLAYLIST-TYPE:"):
         mediaPlaylist.PlaylistType = strings.TrimPrefix(line, "#EXT-X-PLAYLIST-TYPE:")
      case strings.HasPrefix(line, "#EXT-X-ENDLIST"):
         mediaPlaylist.EndList = true
      case strings.HasPrefix(line, "#EXT-X-KEY:"):
         newKey := parseKey(line)
         mediaPlaylist.Keys = append(mediaPlaylist.Keys, newKey)
      case strings.HasPrefix(line, "#EXT-X-MAP:"):
         attrs := parseAttributes(line, "#EXT-X-MAP:")
         if value, ok := attrs["URI"]; ok && value != "" {
            if parsedURL, err := url.Parse(value); err == nil {
               mediaPlaylist.Map = parsedURL
            }
         }
      case strings.HasPrefix(line, "#EXTINF:"):
         // Parse duration and title
         // Format: #EXTINF:duration,[title]
         raw := strings.TrimPrefix(line, "#EXTINF:")
         durationStr, title, _ := strings.Cut(raw, ",")
         duration, err := strconv.ParseFloat(durationStr, 64)
         if err != nil {
            return nil, fmt.Errorf("invalid EXTINF duration: %w", err)
         }
         newSegment := &Segment{
            Duration: duration,
            Title:    strings.TrimSpace(title),
         }
         // The URI is on the next line
         if i+1 < len(lines) {
            nextLine := lines[i+1]
            if !strings.HasPrefix(nextLine, "#") && nextLine != "" {
               if parsedURL, err := url.Parse(nextLine); err == nil {
                  newSegment.URI = parsedURL
               }
               i++
            }
         }
         mediaPlaylist.Segments = append(mediaPlaylist.Segments, newSegment)
      }
   }
   return mediaPlaylist, nil
}
