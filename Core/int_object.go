package core

import (
	"math"
	"strconv"
)

type IntObject struct {
	value int64
}

func (i *IntObject) Repr(context any) string {
	return strconv.FormatInt(i.value, 10)
}

func (i *IntObject) Add(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewIntObject(strconv.FormatInt(i.value+other.(*IntObject).value, 10))
	} else if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(float64(i.value)+other.(*FloatObject).value, 'f', 4, 64))
	} else {
		ReportOperatorError("+")
		return nil
	}
}

func (i *IntObject) Sub(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewIntObject(strconv.FormatInt(i.value-other.(*IntObject).value, 10))
	} else if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(float64(i.value)-other.(*FloatObject).value, 'f', 4,
			64))
	} else {
		ReportOperatorError("-")
		return nil
	}
}

func (i *IntObject) Mul(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewIntObject(strconv.FormatInt(i.value*other.(*IntObject).value, 10))
	} else if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(float64(i.value)*other.(*FloatObject).value, 'f', 4,
			64))
	} else {
		ReportOperatorError("*")
		return nil
	}
}

func (i *IntObject) Div(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewIntObject(strconv.FormatInt(i.value/other.(*IntObject).value, 10))
	} else if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(float64(i.value)/other.(*FloatObject).value, 'f', 4,
			64))
	} else {
		ReportOperatorError("/")
		return nil
	}
}

func (i *IntObject) Mod(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewIntObject(strconv.FormatInt(i.value%other.(*IntObject).value, 10))
	} else {
		ReportOperatorError("%")
		return nil
	}
}

func (i *IntObject) Pow(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewIntObject(strconv.FormatInt(int64(math.Pow(float64(i.value), float64(other.(*IntObject).value))),
			10))
	} else if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(math.Pow(float64(i.value), other.(*FloatObject).value), 'f',
			4, 64))
	} else {
		ReportOperatorError("**")
		return nil
	}
}

func (i *IntObject) Eq(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(i.value == other.(*IntObject).value))
	} else {
		ReportOperatorError("==")
		return nil
	}
}

func (i *IntObject) Neq(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(i.value != other.(*IntObject).value))
	} else {
		ReportOperatorError("!=")
		return nil
	}
}

func (i *IntObject) Gt(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(i.value > other.(*IntObject).value))
	} else {
		ReportOperatorError(">")
		return nil
	}
}

func (i *IntObject) Lt(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(i.value < other.(*IntObject).value))
	} else {
		ReportOperatorError("<")
		return nil
	}
}

func (i *IntObject) Gte(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(i.value >= other.(*IntObject).value))
	} else {
		ReportOperatorError(">=")
		return nil
	}
}

func (i *IntObject) Lte(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(i.value <= other.(*IntObject).value))
	} else {
		ReportOperatorError("<=")
		return nil
	}
}

func (i *IntObject) And(other any) any {
	ReportOperatorError("&&")
	return nil
}

func (i *IntObject) Or(other any) any {
	ReportOperatorError("||")
	return nil
}

func NewIntObject(value string) *IntObject {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return &IntObject{}
	}
	return &IntObject{num}
}
