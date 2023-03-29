package core

import "strconv"

type BoolObject struct {
	value bool
}

func (b *BoolObject) Repr(_ any) string {
	return strconv.FormatBool(b.value)
}

func (b *BoolObject) Eq(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value == other.(*BoolObject).value))
	} else {
		ReportOperatorError("==", b, other)
		return nil
	}
}

func (b *BoolObject) Neq(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value != other.(*BoolObject).value))
	} else {
		ReportOperatorError("!=", b, other)
		return nil
	}
}

func (b *BoolObject) And(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value && other.(*BoolObject).value))
	} else {
		ReportOperatorError("&&", b, other)
		return nil
	}
}

func (b *BoolObject) Or(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value || other.(*BoolObject).value))
	} else {
		ReportOperatorError("||", b, other)
		return nil
	}
}

func (b *BoolObject) Not() any {
	return NewBoolObject(strconv.FormatBool(!b.value))
}

func NewBoolObject(value string) *BoolObject {
	var v bool
	v, _ = strconv.ParseBool(value)
	return &BoolObject{v}
}
