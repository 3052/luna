package dash

import "encoding/base64"

// ContentProtection specifies DRM schemes.
type ContentProtection struct {
   SchemeIdUri string `xml:"schemeIdUri,attr"`
   DefaultKid  string `xml:"default_KID,attr"`
   Pssh        string `xml:"pssh"`
}

// GetPssh returns the PSSH data as a byte slice.
func (cp *ContentProtection) GetPssh() ([]byte, error) {
   if cp.Pssh == "" {
      return nil, nil
   }
   return base64.StdEncoding.DecodeString(cp.Pssh)
}

// content_protection.go
