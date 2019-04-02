package flags

type String struct {
	set   bool
	Value string
}

func (sf *String) Set(x string) error {
	sf.Value = x
	sf.set = true
	return nil
}

func (sf *String) IsSet() bool {
	return sf.set
}

func (sf *String) String() string {
	return sf.Value
}
