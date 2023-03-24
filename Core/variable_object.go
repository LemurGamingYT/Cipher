package core

type VarObject struct {
	name     string
	value    any
	constant bool
}

func NewVar(name string, value any, constant bool) *VarObject {
	return &VarObject{name, value, constant}
}
