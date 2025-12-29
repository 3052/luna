package hls

import (
   "net/url"
   "os"
   "path/filepath"
   "strings"
   "testing"
)

const (
   mediaFilename  = "8500_complete-95fe4117-98fe-4ab7-8895-b2eec69b2b63.m3u8"
   masterFilename = "ctr-all-fb600154-a5e0-4125-ab89-01d627163485-b123e16f-c381-4335-bf76-dcca65425460.m3u8"
)

func TestDecodeMedia(t *testing.T) {
   path := filepath.Join("../testdata", mediaFilename)
   data, err := os.ReadFile(path)
   if err != nil {
      t.Fatalf("Failed to read file from %s: %v", path, err)
   }

   media, err := DecodeMedia(string(data))
   if err != nil {
      t.Fatalf("DecodeMedia failed: %v", err)
   }

   if media.TargetDuration != 9 {
      t.Errorf("Expected TargetDuration 9, got %d", media.TargetDuration)
   }

   // Test ResolveURIs
   baseURL, err := url.Parse("https://example.com/video/")
   if err != nil {
      t.Fatalf("Failed to parse base URL: %v", err)
   }

   media.ResolveURIs(baseURL)

   expectedURI := "https://example.com/video/H264_1_CMAF_CENC_CTR_8500K/95fe4117-98fe-4ab7-8895-b2eec69b2b63/pts_0.mp4"

   if media.Segments[0].URI == nil {
      t.Fatal("Expected URI, got nil")
   }
   if media.Segments[0].URI.String() != expectedURI {
      t.Errorf("Expected Absolute URI %s, got %s", expectedURI, media.Segments[0].URI.String())
   }
}

func TestDecodeMaster(t *testing.T) {
   path := filepath.Join("../testdata", masterFilename)
   data, err := os.ReadFile(path)
   if err != nil {
      t.Fatalf("Failed to read file from %s: %v", path, err)
   }

   master, err := DecodeMaster(string(data))
   if err != nil {
      t.Fatalf("DecodeMaster failed: %v", err)
   }
   // The sample manifest has 8 unique video stream URIs.
   if len(master.StreamInfs) != 8 {
      t.Errorf("Expected 8 unique streams, got %d", len(master.StreamInfs))
   }
   // Find a specific stream to verify grouping of audio tracks.
   var foundStream *StreamInf
   for _, stream := range master.StreamInfs {
      if strings.Contains(stream.URI.Path, "8500_complete") {
         foundStream = stream
         break
      }
   }
   if foundStream == nil {
      t.Fatal("Could not find expected stream '8500_complete' to test grouping")
   }
   // This specific stream has two #EXT-X-STREAM-INF tags with different audio.
   if len(foundStream.Audio) != 2 {
      t.Errorf("Expected stream to have 2 audio groups, got %d", len(foundStream.Audio))
   }

   // Sort the medias and streams
   master.Sort()

   // Print all Medias first
   t.Log("--- Medias (sorted by GroupID) ---")
   for _, media := range master.Medias {
      t.Logf("%s\n---", media)
   }

   // Print all streams and their grouped variants
   t.Log("\n--- StreamInfs (sorted by Average/Min Bandwidth) ---")
   for _, stream := range master.StreamInfs {
      t.Logf("%s\n---", stream)
   }
}
