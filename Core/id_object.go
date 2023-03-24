package core

type IdObject struct {
	value string
}

func (i *IdObject) Repr(context any) string {
	return i.value
}

func NewIdObject(value string) *IdObject {
	return &IdObject{value: value}
}
