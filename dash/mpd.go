package dash

import (
   "encoding/xml"
   "fmt"
   "hash/fnv"
   "net/url"
)

func resolveRef(base *url.URL, relStr string) (*url.URL, error) {
   if relStr == "" {
      return base, nil
   }
   rel, err := url.Parse(relStr)
   if err != nil {
      return nil, err
   }
   if base == nil {
      return rel, nil
   }
   return base.ResolveReference(rel), nil
}

// Mpd represents the root element of the DASH MPD file.
type Mpd struct {
   MediaPresentationDuration string    `xml:"mediaPresentationDuration,attr"`
   BaseUrl                   string    `xml:"BaseURL"`
   Periods                   []*Period `xml:"Period"`
   mpdUrl                    *url.URL
}

// Parse takes a byte slice of an MPD file, unmarshals it,
// links navigation parents, and normalizes Representation IDs.
func Parse(data []byte, mpdUrl *url.URL) (*Mpd, error) {
   var manifest Mpd
   err := xml.Unmarshal(data, &manifest)
   if err != nil {
      return nil, err
   }
   manifest.mpdUrl = mpdUrl
   manifest.link()
   manifest.normalizeIds()
   return &manifest, nil
}

// GetRepresentations returns a map of all Representations keyed by their Id.
func (m *Mpd) GetRepresentations() map[string][]*Representation {
   grouped := make(map[string][]*Representation)
   for _, manifestPeriod := range m.Periods {
      for _, currentSet := range manifestPeriod.AdaptationSets {
         for _, mediaRep := range currentSet.Representations {
            grouped[mediaRep.Id] = append(grouped[mediaRep.Id], mediaRep)
         }
      }
   }
   return grouped
}

// ResolveBaseUrl resolves the MPD's BaseURL against the mpdUrl.
func (m *Mpd) ResolveBaseUrl() (*url.URL, error) {
   return resolveRef(m.mpdUrl, m.BaseUrl)
}

func (m *Mpd) link() {
   for _, manifestPeriod := range m.Periods {
      manifestPeriod.Parent = m
      manifestPeriod.link()
   }
}

// normalizeIds iterates through the MPD and rewrites Representation IDs using a 32-bit hash.
func (m *Mpd) normalizeIds() {
   for _, manifestPeriod := range m.Periods {
      for _, currentSet := range manifestPeriod.AdaptationSets {
         for _, mediaRep := range currentSet.Representations {
            if mediaRep.requiresOriginalId() {
               continue
            }
            currentTemplate := mediaRep.GetSegmentTemplate()
            if currentTemplate == nil {
               continue
            }

            // We use FNV-1a (New32a) instead of FNV-1 (New32) because the FNV-1a XOR-then-multiply
            // approach provides significantly better avalanche properties. This drastically reduces
            // collisions for DASH Media templates, which often share long identical URL prefixes.
            hasher := fnv.New32a()
            fmt.Fprint(hasher, currentTemplate.Media)
            mediaRep.Id = fmt.Sprintf("%x", hasher.Sum32())
         }
      }
   }
}
