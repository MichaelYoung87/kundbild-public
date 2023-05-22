package valueobjects

// Name represents a Name value object
type Name string

// NewName creates a new Name value object
func NewName(name string) Name {
	return Name(name)
}

// String returns a string value that represents a Name value object
func (n Name) String() string {
	return string(n)
}
