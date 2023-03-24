package core

import "strconv"

type BoolObject struct {
	value bool
}

func (b *BoolObject) Repr(context any) string {
	return strconv.FormatBool(b.value)
}

func NewBoolObject(value string) *BoolObject {
	var v bool
	v, _ = strconv.ParseBool(value)
	return &BoolObject{v}
}
