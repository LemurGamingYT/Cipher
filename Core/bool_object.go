package core

import "strconv"

type BoolObject struct {
	value bool
}

func (b *BoolObject) Repr(context any) string {
	return strconv.FormatBool(b.value)
}

func (b *BoolObject) Add(other any) any {
	ReportOperatorError("+")
	return nil
}

func (b *BoolObject) Sub(other any) any {
	ReportOperatorError("-")
	return nil
}

func (b *BoolObject) Mul(other any) any {
	ReportOperatorError("*")
	return nil
}

func (b *BoolObject) Div(other any) any {
	ReportOperatorError("/")
	return nil
}

func (b *BoolObject) Mod(other any) any {
	ReportOperatorError("%")
	return nil
}

func (b *BoolObject) Pow(other any) any {
	ReportOperatorError("**")
	return nil
}

func (b *BoolObject) Eq(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value == other.(*BoolObject).value))
	} else {
		ReportOperatorError("==")
		return nil
	}
}

func (b *BoolObject) Neq(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value != other.(*BoolObject).value))
	} else {
		ReportOperatorError("!=")
		return nil
	}
}

func (b *BoolObject) Gt(other any) any {
	ReportOperatorError(">")
	return nil
}

func (b *BoolObject) Lt(other any) any {
	ReportOperatorError("<")
	return nil
}

func (b *BoolObject) Gte(other any) any {
	ReportOperatorError(">=")
	return nil
}

func (b *BoolObject) Lte(other any) any {
	ReportOperatorError("<=")
	return nil
}

func (b *BoolObject) And(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value && other.(*BoolObject).value))
	} else {
		ReportOperatorError("&&")
		return nil
	}
}

func (b *BoolObject) Or(other any) any {
	if IsInstance(other, BoolObjectType) {
		return NewBoolObject(strconv.FormatBool(b.value || other.(*BoolObject).value))
	} else {
		ReportOperatorError("||")
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
