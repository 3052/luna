package dash

// mergeRepresentation merges all occurrences of one representation ID
// into a single Representation: the first occurrence's values, with
// Bandwidth set to the median across occurrences.
func mergeRepresentation(reps []Representation) *Representation {
   out := reps[0]
   out.Bandwidth = medianBandwidth(reps)
   return &out
}

// Representations returns one merged Representation per representation ID
// found anywhere in the manifest. The same ID typically occurs in every
// Period with period-specific values. Occurrences are assumed to carry
// identical constant fields (codecs, mimeType, dimensions, addressing
// structure, ...) — the first occurrence wins — except Bandwidth, which
// is the median of all observed values.
//
// The order of the result is unspecified.
func (m *MPD) Representations() []*Representation {
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
      out = append(out, mergeRepresentation(reps))
   }
   return out
}

// dash/representations.go
