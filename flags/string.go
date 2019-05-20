package flags

// String represents a string flag
type String struct {
	set   bool
	Value string
}

// Set sets the flag
func (sf *String) Set(x string) error {
	sf.Value = x
	sf.set = true
	return nil
}

// IsSet checks if the flag is set
func (sf *String) IsSet() bool {
	return sf.set
}

// String returns the flag value
func (sf *String) String() string {
	return sf.Value
}
