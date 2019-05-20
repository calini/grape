package flags

import "strconv"

// Int represents an integer flag
type Int struct {
	set   bool
	Value int
}

// Set sets the flag
func (sf *Int) Set(x string) error {
	sf.Value, _ = strconv.Atoi(x)
	sf.set = true
	return nil
}

// IsSet checks if the flag is set
func (sf *Int) IsSet() bool {
	return sf.set
}

// String returns the value as a string
func (sf *Int) String() string {
	return string(sf.Value)
}
