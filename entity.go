package gflaggy

const DefaultNullValue = -1

type flag struct {
	Name     string
	Required bool
}

func NewFlag(name string, required ...bool) Flag {
	req := false
	if len(required) > 0 {
		req = required[0]
	}
	return &flag{Name: name, Required: req}
}
