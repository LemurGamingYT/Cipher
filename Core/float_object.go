package core

import (
	"math"
	"strconv"
)

type FloatObject struct {
	value float64
}

func (f *FloatObject) Repr(_ any) string {
	return strconv.FormatFloat(f.value, 'f', 3, 64)
}

func (f *FloatObject) AsInt(_ ...[]any) *IntObject {
	return NewIntObject(strconv.FormatInt(int64(f.value), 10))
}

func (f *FloatObject) Add(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value+other.(*FloatObject).value, 'f', 4, 64))
	} else if IsInstance(other, IntObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value+float64(other.(*IntObject).value), 'f', 4,
			64))
	} else {
		ReportOperatorError("+", f, other)
		return nil
	}
}

func (f *FloatObject) Sub(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value-other.(*FloatObject).value, 'f', 4, 64))
	} else if IsInstance(other, IntObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value-float64(other.(*IntObject).value), 'f', 4,
			64))
	} else {
		ReportOperatorError("-", f, other)
		return nil
	}
}

func (f *FloatObject) Mul(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value*other.(*FloatObject).value, 'f', 4, 64))
	} else if IsInstance(other, IntObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value*float64(other.(*IntObject).value), 'f', 4,
			64))
	} else {
		ReportOperatorError("*", f, other)
		return nil
	}
}

func (f *FloatObject) Div(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value/other.(*FloatObject).value, 'f', 4, 64))
	} else if IsInstance(other, IntObjectType) {
		return NewFloatObject(strconv.FormatFloat(f.value/float64(other.(*IntObject).value), 'f', 4,
			64))
	} else {
		ReportOperatorError("/", f, other)
		return nil
	}
}

func (f *FloatObject) Pow(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewFloatObject(strconv.FormatFloat(math.Pow(f.value, other.(*FloatObject).value), 'f', 4,
			64))
	} else if IsInstance(other, IntObjectType) {
		return NewFloatObject(strconv.FormatFloat(math.Pow(f.value, float64(other.(*IntObject).value)),
			'f', 4, 64))
	} else {
		ReportOperatorError("**", f, other)
		return nil
	}
}

func (f *FloatObject) Eq(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(f.value == other.(*FloatObject).value))
	} else {
		ReportOperatorError("==", f, other)
		return nil
	}
}

func (f *FloatObject) Neq(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(f.value != other.(*FloatObject).value))
	} else {
		ReportOperatorError("!=", f, other)
		return nil
	}
}

func (f *FloatObject) Gt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(f.value > other.(*FloatObject).value))
	} else {
		ReportOperatorError(">", f, other)
		return nil
	}
}

func (f *FloatObject) Lt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(f.value < other.(*FloatObject).value))
	} else {
		ReportOperatorError("<", f, other)
		return nil
	}
}

func (f *FloatObject) Gte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(f.value >= other.(*FloatObject).value))
	} else {
		ReportOperatorError(">=", f, other)
		return nil
	}
}

func (f *FloatObject) Lte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(f.value <= other.(*FloatObject).value))
	} else {
		ReportOperatorError("<=", f, other)
		return nil
	}
}

func NewFloatObject(value string) *FloatObject {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return &FloatObject{}
	}
	return &FloatObject{num}
}
