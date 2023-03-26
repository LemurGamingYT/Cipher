package core

import (
	"strconv"
	"strings"
)

type StringObject struct {
	value string
}

func (s *StringObject) Repr(context any) string {
	if IsInstance(context, ArrayObjectType) {
		return "'" + s.value + "'"
	}

	return s.value
}

func (s *StringObject) Lower(args []any) *StringObject {
	return NewStringObject("\"" + strings.ToLower(s.value) + "\"")
}

func (s *StringObject) Upper(args []any) *StringObject {
	return NewStringObject("\"" + strings.ToUpper(s.value) + "\"")
}

func (s *StringObject) Title(args []any) *StringObject {
	return NewStringObject("\"" + strings.ToTitle(s.value) + "\"")
}

func (s *StringObject) Add(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject("\"" + s.value + other.(*StringObject).value + "\"")
	} else {
		ReportOperatorError("+")
		return nil
	}
}

func (s *StringObject) Sub(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject("\"" + s.value[int(other.(*IntObject).value):] + "\"")
	} else {
		ReportOperatorError("-")
		return nil
	}
}

func (s *StringObject) Mul(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewStringObject("\"" + strings.Repeat(s.value, int(other.(*IntObject).value)) + "\"")
	} else {
		ReportOperatorError("*")
		return nil
	}
}

func (s *StringObject) Div(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject("\"" + s.value + other.(*StringObject).value + "\"")
	} else {
		ReportOperatorError("/")
		return nil
	}
}

func (s *StringObject) Mod(other any) any {
	ReportOperatorError("%")
	return nil
}

func (s *StringObject) Pow(other any) any {
	ReportOperatorError("/")
	return nil
}

func (s *StringObject) Eq(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) == len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("==")
		return nil
	}
}

func (s *StringObject) Neq(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) != len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("!=")
		return nil
	}
}

func (s *StringObject) Gt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) > len(other.(*StringObject).value)))
	} else {
		ReportOperatorError(">")
		return nil
	}
}

func (s *StringObject) Lt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) < len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("<")
		return nil
	}
}

func (s *StringObject) Gte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) >= len(other.(*StringObject).value)))
	} else {
		ReportOperatorError(">=")
		return nil
	}
}

func (s *StringObject) Lte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) <= len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("<=")
		return nil
	}
}

func (s *StringObject) And(other any) any {
	ReportOperatorError("&&")
	return nil
}

func (s *StringObject) Or(other any) any {
	ReportOperatorError("||")
	return nil
}

func (s *StringObject) Not() any {
	r := []rune(s.value)
	n := len(r)

	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}

	return NewStringObject("\"" + string(r) + "\"")
}

func NewStringObject(value string) *StringObject {
	return &StringObject{value[1 : len(value)-1]}
}
