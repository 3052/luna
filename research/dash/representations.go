package dash

import (
   "fmt"
   "sort"
)

// conflictErr builds a conflict error for one representation ID.
func conflictErr(id, field string, first, other any) error {
   return fmt.Errorf("dash: representation %q: conflicting %s: %v vs %v", id, field, first, other)
}

// equalEssentialProperties reports whether two EssentialProperty lists
// are identical, including order.
func equalEssentialProperties(a, b []EssentialProperty) bool {
   if len(a) != len(b) {
      return false
   }
   for i := range a {
      if a[i] != b[i] {
         return false
      }
   }
   return true
}

// firstValue returns the first observed value of a field, or an error
// if any occurrence carries a different value.
func firstValue[T comparable](id string, reps []Representation, field string, get func(Representation) T) (T, error) {
   first := get(reps[0])
   for _, r := range reps[1:] {
      if v := get(r); v != first {
         return first, conflictErr(id, field, first, v)
      }
   }
   return first, nil
}

// medianBandwidth returns the median of the observed bandwidth values.
// Even count: average of the two middle values, rounded down.
func medianBandwidth(reps []Representation) uint32 {
   vals := make([]uint32, len(reps))
   for i, r := range reps {
      vals[i] = r.Bandwidth
   }
   sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })
   n := len(vals)
   if n%2 == 1 {
      return vals[n/2]
   }
   return uint32((uint64(vals[n/2-1]) + uint64(vals[n/2])) / 2)
}

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
   if out.EssentialProperties, err = mergeEssentialProperties(id, reps); err != nil {
      return nil, err
   }
   return &out, nil
}

// mergeBaseURL merges the BaseURL field across occurrences.
func mergeBaseURL(id string, reps []Representation) (*BaseURL, error) {
   first := reps[0].BaseURL
   for _, r := range reps[1:] {
      switch {
      case (first == nil) != (r.BaseURL == nil):
         return nil, conflictErr(id, "BaseURL presence", first != nil, r.BaseURL != nil)
      case first != nil && r.BaseURL.URL != first.URL:
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
      switch {
      case (first == nil) != (acc == nil):
         return nil, conflictErr(id, "AudioChannelConfiguration presence", first != nil, acc != nil)
      case first != nil && *acc != *first:
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
   if first == nil {
      for _, r := range reps[1:] {
         if r.SegmentBase != nil {
            return nil, conflictErr(id, "SegmentBase presence", false, true)
         }
      }
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
      case sb == nil:
         return nil, conflictErr(id, "SegmentBase presence", true, false)
      case sb.IndexRange != first.IndexRange:
         return nil, conflictErr(id, "SegmentBase indexRange", first.IndexRange, sb.IndexRange)
      case sb.Timescale != first.Timescale:
         return nil, conflictErr(id, "SegmentBase timescale", first.Timescale, sb.Timescale)
      case (first.Initialization == nil) != (sb.Initialization == nil):
         return nil, conflictErr(id, "SegmentBase Initialization presence",
            first.Initialization != nil, sb.Initialization != nil)
      case first.Initialization != nil && sb.Initialization.Range != first.Initialization.Range:
         return nil, conflictErr(id, "SegmentBase Initialization range",
            first.Initialization.Range, sb.Initialization.Range)
      }
   }
   return out, nil
}

// mergeSegmentTemplate merges the SegmentTemplate field across
// occurrences. StartNumber, PresentationTimeOffset and SegmentTimeline
// are period-relative: first value, not checked.
func mergeSegmentTemplate(id string, reps []Representation) (*SegmentTemplate, error) {
   first := reps[0].SegmentTemplate
   if first == nil {
      for _, r := range reps[1:] {
         if r.SegmentTemplate != nil {
            return nil, conflictErr(id, "SegmentTemplate presence", false, true)
         }
      }
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
      case st == nil:
         return nil, conflictErr(id, "SegmentTemplate presence", true, false)
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

// mergeEssentialProperties merges the EssentialProperty list across
// occurrences, requiring identical values in every occurrence.
func mergeEssentialProperties(id string, reps []Representation) ([]EssentialProperty, error) {
   first := reps[0].EssentialProperties
   for _, r := range reps[1:] {
      if !equalEssentialProperties(first, r.EssentialProperties) {
         return nil, conflictErr(id, "EssentialProperty values", first, r.EssentialProperties)
      }
   }
   return first, nil
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
// The result is ordered by first appearance in the manifest.
func (m *MPD) Representations() ([]Representation, error) {
   var order []string
   groups := make(map[string][]Representation)
   for i := range m.Periods {
      for j := range m.Periods[i].AdaptationSets {
         for _, r := range m.Periods[i].AdaptationSets[j].Representations {
            if _, seen := groups[r.ID]; !seen {
               order = append(order, r.ID)
            }
            groups[r.ID] = append(groups[r.ID], r)
         }
      }
   }
   out := make([]Representation, 0, len(order))
   for _, id := range order {
      merged, err := mergeRepresentation(groups[id])
      if err != nil {
         return nil, err
      }
      out = append(out, *merged)
   }
   return out, nil
}

// dash/representations.go
