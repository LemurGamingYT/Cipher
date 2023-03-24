package core

import (
	"strconv"
	"strings"
)

type StringObject struct {
	value string
}

func (s *StringObject) Repr(context any) string {
	return s.value
}

func (s *StringObject) Add(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject(s.value + other.(StringObject).value)
	} else {
		ReportError("Type", "Invalid types for operation '+'")

		return nil
	}
}

func (s *StringObject) Sub(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject(s.value[int(other.(IntObject).value):])
	} else {
		ReportError("Type", "Invalid types for operation '-'")

		return nil
	}
}

func (s *StringObject) Mul(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewStringObject(strings.Repeat(s.value, int(other.(IntObject).value)))
	} else {
		ReportError("Type", "Invalid types for operation '*'")

		return nil
	}
}

func (s *StringObject) Div(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject(s.value + other.(StringObject).value)
	} else {
		ReportError("Type", "Invalid types for operation '/'")

		return nil
	}
}

func (s *StringObject) Mod(other any) any {
	ReportError("Type", "Invalid types for operation '%'")

	return nil
}

func (s *StringObject) Pow(other any) any {
	ReportError("Type", "Invalid types for operation '/'")
	return nil
}

func (s *StringObject) Eq(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) == len(other.(StringObject).value)))
	} else {
		ReportError("Type", "Invalid types for operation '=='")

		return nil
	}
}

func (s *StringObject) Neq(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) != len(other.(StringObject).value)))
	} else {
		ReportError("Type", "Invalid types for operation '!='")

		return nil
	}
}

func (s *StringObject) Gt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) > len(other.(StringObject).value)))
	} else {
		ReportError("Type", "Invalid types for operation '>'")

		return nil
	}
}

func (s *StringObject) Lt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) < len(other.(StringObject).value)))
	} else {
		ReportError("Type", "Invalid types for operation '<'")

		return nil
	}
}

func (s *StringObject) Gte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) >= len(other.(StringObject).value)))
	} else {
		ReportError("Type", "Invalid types for operation '>='")

		return nil
	}
}

func (s *StringObject) Lte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) <= len(other.(StringObject).value)))
	} else {
		ReportError("Type", "Invalid types for operation '<='")

		return nil
	}
}

func (s *StringObject) And(other any) any {
	ReportError("Type", "Invalid types for operation '&&' ('and')")
	return nil
}

func (s *StringObject) Or(other any) any {
	ReportError("Type", "Invalid types for operation '||' ('or')")
	return nil
}

func NewStringObject(value string) *StringObject {
	return &StringObject{value[1 : len(value)-1]}
}
