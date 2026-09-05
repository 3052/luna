package dash

// mergeRepresentation merges all occurrences of one representation ID
// into a single Representation.
func mergeRepresentation(reps []Representation) (*Representation, error) {
   id := reps[0].ID
   out := Representation{ID: id}
   var err error

   if out.Codecs, err = firstValue(id, reps, "codecs", func(r Representation) string { return r.Codecs }); err != nil {
      return nil, err
   }
   if out.MimeType, err = firstValue(id, reps, "mimeType", func(r Representation) string { return r.MimeType }); err != nil {
      return nil, err
   }
   if out.SAR, err = firstValue(id, reps, "sar", func(r Representation) string { return r.SAR }); err != nil {
      return nil, err
   }
   if out.Width, err = firstValue(id, reps, "width", func(r Representation) uint32 { return r.Width }); err != nil {
      return nil, err
   }
   if out.Height, err = firstValue(id, reps, "height", func(r Representation) uint32 { return r.Height }); err != nil {
      return nil, err
   }
   if out.AudioSamplingRate, err = firstValue(id, reps, "audioSamplingRate", func(r Representation) uint32 { return r.AudioSamplingRate }); err != nil {
      return nil, err
   }
   out.Bandwidth = medianBandwidth(reps)
   if out.BaseURL, err = mergeBaseURL(id, reps); err != nil {
      return nil, err
   }
   if out.AudioChannelConfiguration, err = mergeAudioChannelConfiguration(id, reps); err != nil {
      return nil, err
   }
   if out.SegmentBase, err = mergeSegmentBase(id, reps); err != nil {
      return nil, err
   }
   if out.SegmentTemplate, err = mergeSegmentTemplate(id, reps); err != nil {
      return nil, err
   }
   return &out, nil
}

// mergeBaseURL merges the BaseURL field across occurrences.
func mergeBaseURL(id string, reps []Representation) (*BaseURL, error) {
   first := reps[0].BaseURL
   for _, r := range reps[1:] {
      if err := checkPresence(id, "BaseURL", first, r.BaseURL); err != nil {
         return nil, err
      }
      if first != nil && r.BaseURL.URL != first.URL {
         return nil, conflictErr(id, "BaseURL", first.URL, r.BaseURL.URL)
      }
   }
   return first, nil
}

// mergeAudioChannelConfiguration merges the AudioChannelConfiguration
// field across occurrences.
func mergeAudioChannelConfiguration(id string, reps []Representation) (*AudioChannelConfiguration, error) {
   first := reps[0].AudioChannelConfiguration
   for _, r := range reps[1:] {
      acc := r.AudioChannelConfiguration
      if err := checkPresence(id, "AudioChannelConfiguration", first, acc); err != nil {
         return nil, err
      }
      if first != nil && *acc != *first {
         return nil, conflictErr(id, "AudioChannelConfiguration",
            first.SchemeIDURI+"/"+first.Value, acc.SchemeIDURI+"/"+acc.Value)
      }
   }
   return first, nil
}

// mergeSegmentBase merges the SegmentBase field across occurrences.
// PresentationTimeOffset is period-relative: first value, not checked.
func mergeSegmentBase(id string, reps []Representation) (*SegmentBase, error) {
   first := reps[0].SegmentBase
   for _, r := range reps[1:] {
      if err := checkPresence(id, "SegmentBase", first, r.SegmentBase); err != nil {
         return nil, err
      }
   }
   if first == nil {
      return nil, nil
   }
   out := &SegmentBase{
      IndexRange:             first.IndexRange,
      Timescale:              first.Timescale,
      PresentationTimeOffset: first.PresentationTimeOffset,
   }
   if first.Initialization != nil {
      out.Initialization = &Initialization{Range: first.Initialization.Range}
   }
   for _, r := range reps[1:] {
      sb := r.SegmentBase
      switch {
      case sb.IndexRange != first.IndexRange:
         return nil, conflictErr(id, "SegmentBase indexRange", first.IndexRange, sb.IndexRange)
      case sb.Timescale != first.Timescale:
         return nil, conflictErr(id, "SegmentBase timescale", first.Timescale, sb.Timescale)
      case (first.Initialization == nil) != (sb.Initialization == nil):
         return nil, conflictErr(id, "SegmentBase Initialization presence", first.Initialization != nil, sb.Initialization != nil)
      case first.Initialization != nil && sb.Initialization.Range != first.Initialization.Range:
         return nil, conflictErr(id, "SegmentBase Initialization range", first.Initialization.Range, sb.Initialization.Range)
      }
   }
   return out, nil
}

// mergeSegmentTemplate merges the SegmentTemplate field across
// occurrences. StartNumber, PresentationTimeOffset and SegmentTimeline
// are period-relative: first value, not checked.
func mergeSegmentTemplate(id string, reps []Representation) (*SegmentTemplate, error) {
   first := reps[0].SegmentTemplate
   for _, r := range reps[1:] {
      if err := checkPresence(id, "SegmentTemplate", first, r.SegmentTemplate); err != nil {
         return nil, err
      }
   }
   if first == nil {
      return nil, nil
   }
   out := &SegmentTemplate{
      Media:                  first.Media,
      Duration:               first.Duration,
      Timescale:              first.Timescale,
      StartNumber:            first.StartNumber,
      PresentationTimeOffset: first.PresentationTimeOffset,
   }
   if first.SegmentTimeline != nil {
      out.SegmentTimeline = &SegmentTimeline{
         Entries: append([]SegmentTimelineEntry(nil), first.SegmentTimeline.Entries...),
      }
   }
   for _, r := range reps[1:] {
      st := r.SegmentTemplate
      switch {
      case st.Media != first.Media:
         return nil, conflictErr(id, "SegmentTemplate media", first.Media, st.Media)
      case st.Duration != first.Duration:
         return nil, conflictErr(id, "SegmentTemplate duration", first.Duration, st.Duration)
      case st.Timescale != first.Timescale:
         return nil, conflictErr(id, "SegmentTemplate timescale", first.Timescale, st.Timescale)
      }
   }
   return out, nil
}

// Representations returns one merged Representation per representation ID
// found anywhere in the manifest. The same ID typically occurs in every
// Period with period-specific values, so all occurrences are merged:
//
//   - Constant fields (codecs, mimeType, dimensions, addressing structure,
//     ...) keep the first observed value; conflicts error.
//   - Bandwidth keeps the median of all observed values.
//   - Period-relative fields (PresentationTimeOffset, StartNumber,
//     SegmentTimeline) are expected to differ between periods: they keep
//     the first observed value and are not conflict-checked.
//
// The order of the result is unspecified.
func (m *MPD) Representations() ([]*Representation, error) {
   groups := make(map[string][]Representation)
   for _, p := range m.Periods {
      for _, as := range p.AdaptationSets {
         for _, r := range as.Representations {
            groups[r.ID] = append(groups[r.ID], r)
         }
      }
   }
   out := make([]*Representation, 0, len(groups))
   for _, reps := range groups {
      merged, err := mergeRepresentation(reps)
      if err != nil {
         return nil, err
      }
      out = append(out, merged)
   }
   return out, nil
}

// dash/representations.go
