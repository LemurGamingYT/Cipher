package core

type NullObject struct{}

func (n *NullObject) Repr(context any) string {
	return "null"
}

func NewNullObject() *NullObject {
	return &NullObject{}
}
