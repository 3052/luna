package dash

import (
   "cmp"
   "net/url"
   "strconv"
   "strings"
)

func Bandwidth(r1, r2 *Representation) int {
   return cmp.Or(
      r1.MedianBandwidth-r2.MedianBandwidth,
      r1.Bandwidth-r2.Bandwidth,
   )
}

// String returns a multi-line summary of the Representation.
func (r *Representation) String() string {
   var data strings.Builder
   if r.MedianBandwidth >= 1 {
      data.WriteString("median bandwidth = ")
      data.WriteString(strconv.Itoa(r.MedianBandwidth))
   }
   if data.Len() >= 1 {
      data.WriteByte('\n')
   }
   data.WriteString("bandwidth = ")
   data.WriteString(strconv.Itoa(r.Bandwidth))
   if width := r.GetWidth(); width != 0 {
      data.WriteString("\nwidth = ")
      data.WriteString(strconv.Itoa(width))
   }
   if height := r.GetHeight(); height != 0 {
      data.WriteString("\nheight = ")
      data.WriteString(strconv.Itoa(height))
   }
   if codecs := r.GetCodecs(); codecs != "" {
      data.WriteString("\ncodecs = ")
      data.WriteString(codecs)
   }
   data.WriteString("\nmimeType = ")
   data.WriteString(r.GetMimeType())
   if label := r.GetLabel(); label != "" {
      data.WriteString("\nlabel = ")
      data.WriteString(label)
   } else if lang := r.GetLang(); lang != "" {
      data.WriteString("\nlang = ")
      data.WriteString(lang)
   }
   if role := r.GetRole(); role != "" {
      data.WriteString("\nrole = ")
      data.WriteString(role)
   }
   if periodDuration := r.GetPeriodDuration(); periodDuration != "" {
      data.WriteString("\nperiod duration = ")
      data.WriteString(periodDuration)
   }
   data.WriteString("\nid = ")
   data.WriteString(r.Id)
   return data.String()
}

// ResolveBaseUrl resolves the Representation's BaseURL against the parent hierarchy.
func (r *Representation) ResolveBaseUrl() (*url.URL, error) {
   parentBase, err := r.Parent.getAbsoluteBaseUrl()
   if err != nil {
      return nil, err
   }
   return resolveRef(parentBase, r.BaseUrl)
}

// requiresOriginalId checks if the Representation ID should be preserved.
func (r *Representation) requiresOriginalId() bool {
   currentTemplate := r.GetSegmentTemplate()
   if currentTemplate == nil {
      return true
   }
   return strings.Contains(currentTemplate.Media, "$RepresentationID$")
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

func (r *Representation) GetHeight() int {
   if r.Height != 0 {
      return r.Height
   }
   if r.Parent != nil {
      return r.Parent.Height
   }
   return 0
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

func (r *Representation) GetMimeType() string {
   if r.MimeType != "" {
      return r.MimeType
   }
   if r.Parent != nil {
      return r.Parent.MimeType
   }
   return ""
}

func (r *Representation) GetLang() string {
   if r.Parent != nil {
      return r.Parent.Lang
   }
   return ""
}

func (r *Representation) GetLabel() string {
   if r.Parent != nil {
      return r.Parent.Label
   }
   return ""
}

func (r *Representation) GetRole() string {
   if r.Parent != nil && r.Parent.Role != nil {
      return r.Parent.Role.Value
   }
   return ""
}

func (r *Representation) GetPeriodId() string {
   if r.Parent != nil && r.Parent.Parent != nil {
      return r.Parent.Parent.Id
   }
   return ""
}

func (r *Representation) GetPeriodDuration() string {
   if r.Parent != nil && r.Parent.Parent != nil {
      return r.Parent.Parent.Duration
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

func (r *Representation) GetSegmentTemplate() *SegmentTemplate {
   if r.SegmentTemplate != nil {
      return r.SegmentTemplate
   }
   if r.Parent != nil {
      return r.Parent.SegmentTemplate
   }
   return nil
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

// Representation describes a version of the media content.
type Representation struct {
   Bandwidth         int                  `xml:"bandwidth,attr"`
   BaseUrl           string               `xml:"BaseURL"`
   Codecs            string               `xml:"codecs,attr"`
   ContentProtection []*ContentProtection `xml:"ContentProtection"`
   Height            int                  `xml:"height,attr"`
   Id                string               `xml:"id,attr"`
   MedianBandwidth   int                  `xml:"-"`
   MimeType          string               `xml:"mimeType,attr"`
   Parent            *AdaptationSet       `xml:"-"`
   SegmentBase       *SegmentBase         `xml:"SegmentBase"`
   SegmentList       *SegmentList         `xml:"SegmentList"`
   SegmentTemplate   *SegmentTemplate     `xml:"SegmentTemplate"`
   Width             int                  `xml:"width,attr"`
}
