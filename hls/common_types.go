// common_types.go
package hls

import (
   "encoding/base64"
   "errors"
   "fmt"
   "net/url"
   "strings"
)

// Key represents encryption info from a #EXT-X-KEY tag.
type Key struct {
   Method            string
   Uri               *url.URL
   KeyFormat         string
   KeyFormatVersions string
   Characteristics   string
}

func parseKey(line string) (*Key, error) {
   prefix := "#EXT-X-KEY:"
   attrs := parseAttributes(line, prefix)
   newKey := &Key{
      Method:            attrs["METHOD"],
      KeyFormat:         attrs["KEYFORMAT"],
      KeyFormatVersions: attrs["KEYFORMATVERSIONS"],
      Characteristics:   attrs["CHARACTERISTICS"],
   }
   if value, ok := attrs["URI"]; ok && value != "" {
      parsedUrl, err := url.Parse(value)
      if err != nil {
         return nil, fmt.Errorf("invalid URI in EXT-X-KEY: %w", err)
      }
      newKey.Uri = parsedUrl
   }
   return newKey, nil
}

// DecodeData extracts and decodes the Base64 data directly from the URL Opaque field.
func (k *Key) DecodeData() ([]byte, error) {
   if k.Uri == nil {
      return nil, errors.New("URI is nil")
   }
   if k.Uri.Scheme != "data" {
      return nil, errors.New("URI is not a data URI")
   }
   // For data URIs, net/url stores the content (mime+encoding+data) in Opaque.
   // Format: [<mediatype>][;base64],<data>
   meta, dataString, found := strings.Cut(k.Uri.Opaque, ",")
   if !found {
      return nil, errors.New("invalid data URI: missing comma separator")
   }
   // Verify base64 encoding is specified in the metadata (before the comma)
   if !strings.Contains(meta, ";base64") {
      return nil, errors.New("data URI does not contain base64 indicator")
   }
   return base64.StdEncoding.DecodeString(dataString)
}

func (k *Key) resolve(base *url.URL) {
   if k.Uri != nil {
      k.Uri = base.ResolveReference(k.Uri)
   }
}
