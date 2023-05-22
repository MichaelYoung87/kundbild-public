package valueobjects

// URL represents a URL value object
type URL string

// NewURL creates a new URL value object
func NewURL(url string) URL {
	return URL(url)
}

// String returns a string value that represents a URL value object
func (u URL) String() string {
	return string(u)
}
