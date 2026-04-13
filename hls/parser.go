package hls

import "strings"

// Helper to split lines, ignoring empty lines and carriage returns,
// but preserving leading/trailing spaces.
func splitLines(content string) []string {
   return strings.FieldsFunc(content, func(r rune) bool {
      return r == '\n' || r == '\r'
   })
}
