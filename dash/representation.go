package dash

import (
   "fmt"
   "net/url"
   "strings"
)

// Representation describes a version of the media content.
type Representation struct {
   Bandwidth         int                  `xml:"bandwidth,attr"`
   BaseUrl           string               `xml:"BaseURL"`
   Codecs            string               `xml:"codecs,attr"`
   Height            int                  `xml:"height,attr"`
   Id                string               `xml:"id,attr"`
   MimeType          string               `xml:"mimeType,attr"`
   SegmentList       *SegmentList         `xml:"SegmentList"`
   SegmentTemplate   *SegmentTemplate     `xml:"SegmentTemplate"`
   Width             int                  `xml:"width,attr"`
   SegmentBase       *SegmentBase         `xml:"SegmentBase"`
   ContentProtection []*ContentProtection `xml:"ContentProtection"`
   Parent            *AdaptationSet       `xml:"-"`
}

func (r *Representation) GetCodecs() string {
   if r.Codecs != "" {
      return r.Codecs
   }
   if r.Parent != nil {
      return r.Parent.Codecs
   }
   return ""
}

func (r *Representation) GetContentProtection() []*ContentProtection {
   if len(r.ContentProtection) > 0 {
      return r.ContentProtection
   }
   if r.Parent != nil {
      return r.Parent.ContentProtection
   }
   return nil
}

func (r *Representation) GetHeight() int {
   if r.Height != 0 {
      return r.Height
   }
   if r.Parent != nil {
      return r.Parent.Height
   }
   return 0
}

func (r *Representation) GetLabel() string {
   if r.Parent != nil {
      return r.Parent.Label
   }
   return ""
}

func (r *Representation) GetLang() string {
   if r.Parent != nil {
      return r.Parent.Lang
   }
   return ""
}

func (r *Representation) GetMimeType() string {
   if r.MimeType != "" {
      return r.MimeType
   }
   if r.Parent != nil {
      return r.Parent.MimeType
   }
   return ""
}

func (r *Representation) GetPeriodDuration() string {
   if r.Parent != nil && r.Parent.Parent != nil {
      return r.Parent.Parent.Duration
   }
   return ""
}

func (r *Representation) GetPeriodId() string {
   if r.Parent != nil && r.Parent.Parent != nil {
      return r.Parent.Parent.Id
   }
   return ""
}

func (r *Representation) GetRole() string {
   if r.Parent != nil && r.Parent.Role != nil {
      return r.Parent.Role.Value
   }
   return ""
}

func (r *Representation) GetSegmentTemplate() *SegmentTemplate {
   if r.SegmentTemplate != nil {
      return r.SegmentTemplate
   }
   if r.Parent != nil {
      return r.Parent.SegmentTemplate
   }
   return nil
}

func (r *Representation) GetWidth() int {
   if r.Width != 0 {
      return r.Width
   }
   if r.Parent != nil {
      return r.Parent.Width
   }
   return 0
}

// ResolveBaseUrl resolves the Representation's BaseURL against the parent hierarchy.
func (r *Representation) ResolveBaseUrl() (*url.URL, error) {
   parentBase, err := r.Parent.getAbsoluteBaseUrl()
   if err != nil {
      return nil, err
   }
   return resolveRef(parentBase, r.BaseUrl)
}

// String returns a multi-line summary of the Representation.
func (r *Representation) String() string {
   data := &strings.Builder{}
   fmt.Fprintln(data, "bandwidth:", r.Bandwidth)
   if width := r.GetWidth(); width != 0 {
      fmt.Fprintln(data, "width:", width)
   }
   if height := r.GetHeight(); height != 0 {
      fmt.Fprintln(data, "height:", height)
   }
   if codecs := r.GetCodecs(); codecs != "" {
      fmt.Fprintln(data, "codecs:", codecs)
   }
   fmt.Fprintln(data, "mimeType:", r.GetMimeType())
   if label := r.GetLabel(); label != "" {
      fmt.Fprintln(data, "label:", label)
   } else if lang := r.GetLang(); lang != "" {
      fmt.Fprintln(data, "lang:", lang)
   }
   if role := r.GetRole(); role != "" {
      fmt.Fprintln(data, "role:", role)
   }
   if periodDuration := r.GetPeriodDuration(); periodDuration != "" {
      fmt.Fprintln(data, "period duration:", periodDuration)
   }
   data.WriteString("id: ")
   data.WriteString(r.Id)
   return data.String()
}

func (r *Representation) link() {
   if r.SegmentTemplate != nil {
      r.SegmentTemplate.ParentRepresentation = r
   }
   if r.SegmentList != nil {
      r.SegmentList.Parent = r
      r.SegmentList.link()
   }
}

// requiresOriginalId checks if the Representation ID should be preserved.
func (r *Representation) requiresOriginalId() bool {
   currentTemplate := r.GetSegmentTemplate()
   if currentTemplate == nil {
      return true
   }
   return strings.Contains(currentTemplate.Media, "$RepresentationID$")
}

// representation.go
