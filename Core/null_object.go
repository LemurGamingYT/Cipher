package core

type NullObject struct{}

func (n *NullObject) Repr(_ any) string {
	return "null"
}

func NewNullObject() *NullObject {
	return &NullObject{}
}
