package dash

// mergeRepresentation merges all occurrences of one representation ID
// into a single Representation: the first occurrence's values, with
// Bandwidth set to the median across occurrences and SegmentTemplate
// set to the combined template across occurrences.
func mergeRepresentation(reps []Representation) *Representation {
   out := reps[0]
   out.Bandwidth = medianBandwidth(reps)
   out.SegmentTemplate = mergeSegmentTemplate(reps)
   return &out
}

// mergeSegmentTemplate combines the SegmentTemplates of all occurrences
// of one representation ID into a single template:
//
//   - If all occurrences share the same template (or only one has one),
//     it is returned as-is.
//   - Otherwise the occurrences carry per-period SegmentTimelines: the
//     result keeps the first occurrence's media, timescale and
//     startNumber, drops presentationTimeOffset, and concatenates all
//     timeline entries (their t values are already absolute). Numbering
//     then runs continuously from the first startNumber across the
//     combined timeline.
//   - For duration-based templates, segment numbers derive directly
//     from presentation time, so the first occurrence's template
//     (startNumber + duration) already covers the whole presentation.
func mergeSegmentTemplate(reps []Representation) *SegmentTemplate {
   var sts []*SegmentTemplate
   for _, r := range reps {
      if r.SegmentTemplate != nil {
         sts = append(sts, r.SegmentTemplate)
      }
   }
   switch len(sts) {
   case 0:
      return nil
   case 1:
      return sts[0]
   }
   for _, st := range sts[1:] {
      if st != sts[0] {
         goto combine
      }
   }
   return sts[0]
combine:
   first := sts[0]
   out := &SegmentTemplate{
      Media:       first.Media,
      Timescale:   first.Timescale,
      StartNumber: first.StartNumber,
   }
   if first.SegmentTimeline != nil {
      out.SegmentTimeline = &SegmentTimeline{}
      for _, st := range sts {
         out.SegmentTimeline.Entries = append(out.SegmentTimeline.Entries, st.SegmentTimeline.Entries...)
      }
   } else {
      out.Duration = first.Duration
   }
   return out
}

// Representations returns one merged Representation per representation ID
// found anywhere in the manifest. The same ID typically occurs in every
// Period with period-specific values. Occurrences are assumed to carry
// identical constant fields (codecs, mimeType, dimensions, addressing
// structure, ...) — the first occurrence wins — except Bandwidth, which
// is the median of all observed values, and SegmentTemplate, whose
// period-specific SegmentTimelines are concatenated into one covering
// the whole presentation.
//
// A SegmentTemplate declared on the AdaptationSet is inherited by its
// Representations, unless the Representation declares one itself.
//
// The order of the result is unspecified.
func (m *MPD) Representations() []*Representation {
   groups := make(map[string][]Representation)
   for _, p := range m.Periods {
      for _, as := range p.AdaptationSets {
         for _, r := range as.Representations {
            if r.SegmentTemplate == nil {
               r.SegmentTemplate = as.SegmentTemplate
            }
            groups[r.ID] = append(groups[r.ID], r)
         }
      }
   }
   out := make([]*Representation, 0, len(groups))
   for _, reps := range groups {
      out = append(out, mergeRepresentation(reps))
   }
   return out
}

// dash/representations.go
