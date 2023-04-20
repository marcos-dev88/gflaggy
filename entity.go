package gflaggy

const DefaultNullValue = -1

const (
	Types FlagReturnType = iota
	Boolean
	Integer
	Float32
	Float64
	String
	JSON
)

type FlagReturnType int

type flag struct {
	Name     string
	Type     FlagReturnType
	Required bool
}

func NewFlag(name string, required ...bool) Flag {
	req := false
	if len(required) > 0 {
		req = required[0]
	}
	return &flag{Name: name, Required: req}
}
