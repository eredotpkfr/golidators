package regexes

// Ignore case prefix pattern for regex patterns
const ignore_case string = "(?i)"

// Add IGNORE_CASE flag any regex pattern
func make_case_insensitive(pattern string) string {
  return ignore_case + pattern
}
