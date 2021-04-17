package extensions

const (
	// Json extension
	Json = ".json"
)

// All supported extensions
var All = []string{
	Json,
}

// IsValid true if ext is a valid extension, false otherwise.
func IsValid(ext string) bool {
	for _, l := range All {
		if l == ext {
			return true
		}
	}
	return false
}
