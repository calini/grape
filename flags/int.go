package flags

import "strconv"

type Int struct {
	set   bool
	Value int
}

func (sf *Int) Set(x string) error {
	sf.Value, _ = strconv.Atoi(x)
	sf.set = true
	return nil
}

func (sf *Int) IsSet() bool {
	return sf.set
}

func (sf *Int) String() string {
	return string(sf.Value)
}