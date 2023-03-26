package core

import "fmt"

type VarObject struct {
	name     string
	value    any
	constant bool
}

func (v *VarObject) Repr(context any) string {
	return fmt.Sprintf("Variable(name=%s, value=%s, constant=%v, at=%p)\n", v.name, v.value, v.constant, v)
}

func NewVar(name string, value any, constant bool) *VarObject {
	return &VarObject{name, value, constant}
}
