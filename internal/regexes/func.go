package regexes

// Ignore case prefix pattern for regex patterns
const ignoreCase string = "(?i)"

// Add ignoreCase flag any regex pattern
func makeCaseInsensitive(pattern string) string {
  return ignoreCase + pattern
}
