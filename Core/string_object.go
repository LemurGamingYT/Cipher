package core

import (
	"fmt"
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

func (s *StringObject) Lower(_ []any) *StringObject {
	return NewStringObject("\"" + strings.ToLower(s.value) + "\"")
}

func (s *StringObject) Upper(_ []any) *StringObject {
	return NewStringObject("\"" + strings.ToUpper(s.value) + "\"")
}

func (s *StringObject) Title(_ []any) *StringObject {
	return NewStringObject("\"" + TitleString(s.value) + "\"")
}

func (s *StringObject) Clone(_ []any) *StringObject {
	return NewStringObject("\"" + strings.Clone(s.value) + "\"")
}

func (s *StringObject) Count(args []any) *IntObject {
	if substr, ok := args[0].(*StringObject); ok {
		return NewIntObject(strconv.FormatInt(int64(strings.Count(s.value, substr.value)), 10))
	} else {
		ReportError("Type", "'Count' method requires type 'string' argument")
		return nil
	}
}

func (s *StringObject) AsInt(_ ...[]any) *IntObject {
	i, err := strconv.ParseInt(s.value, 10, 64)
	if err != nil {
		ReportError("Type", fmt.Sprintf("Invalid cast '%s' to type 'int'\n", s.value))
	}

	return NewIntObject(strconv.FormatInt(i, 10))
}

func (s *StringObject) AsFloat(_ ...[]any) *FloatObject {
	f, err := strconv.ParseFloat(s.value, 64)
	if err != nil {
		ReportError("Type", fmt.Sprintf("Invalid cast '%s' to type 'float'\n", s.value))
	}

	return NewFloatObject(strconv.FormatFloat(f, 'f', 4, 64))
}

func (s *StringObject) AsBool(_ ...[]any) *BoolObject {
	b, err := strconv.ParseBool(s.value)
	if err != nil {
		ReportError("Type", fmt.Sprintf("Invalid cast '%s' to type 'bool'\n", s.value))
	}

	return NewBoolObject(strconv.FormatBool(b))
}

func (s *StringObject) Contains(args []any) *BoolObject {
	if pattern, ok := args[0].(*StringObject); ok {
		return NewBoolObject(strconv.FormatBool(strings.Contains(s.value, pattern.value)))
	} else {
		ReportError("Type", "'Contains' method requires type 'string' argument")
		return nil
	}
}

func (s *StringObject) Endswith(args []any) *BoolObject {
	if suffix, ok := args[0].(*StringObject); ok {
		return NewBoolObject(strconv.FormatBool(strings.HasSuffix(s.value, suffix.value)))
	} else {
		ReportError("Type", "'Startswith' method requires type 'string' argument")
		return nil
	}
}

func (s *StringObject) Startswith(args []any) *BoolObject {
	if prefix, ok := args[0].(*StringObject); ok {
		return NewBoolObject(strconv.FormatBool(strings.HasPrefix(s.value, prefix.value)))
	} else {
		ReportError("Type", "'Startswith' method requires type 'string' argument")
		return nil
	}
}

/*
func (s *StringObject) Start(_ []any) *StringObject {
	return NewStringObject(string(s.value[0]))
}

func (s *StringObject) End(_ []any) *StringObject {
	return NewStringObject(string([]rune(s.value)[len(s.value)-1]))
}*/

func (s *StringObject) Replace(args []any) *StringObject {
	if old, ok := args[0].(*StringObject); ok {
		if newStr, ok := args[1].(*StringObject); ok {
			return NewStringObject("\"" + strings.ReplaceAll(s.value, old.value, newStr.value) + "\"")
		} else {
			ReportError("Type", "'Replace' method requires type 'string' and 'string' arguments")
			return nil
		}
	} else {
		ReportError("Type", "'Replace' method requires type 'string' and 'string' arguments")
		return nil
	}
}

func (s *StringObject) Add(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject("\"" + s.value + other.(*StringObject).value + "\"")
	} else {
		ReportOperatorError("+", s, other)
		return nil
	}
}

func (s *StringObject) Sub(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject("\"" + s.value[int(other.(*IntObject).value):] + "\"")
	} else {
		ReportOperatorError("-", s, other)
		return nil
	}
}

func (s *StringObject) Mul(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewStringObject("\"" + strings.Repeat(s.value, int(other.(*IntObject).value)) + "\"")
	} else {
		ReportOperatorError("*", s, other)
		return nil
	}
}

func (s *StringObject) Div(other any) any {
	if IsInstance(other, StringObjectType) {
		return NewStringObject("\"" + s.value + other.(*StringObject).value + "\"")
	} else {
		ReportOperatorError("/", s, other)
		return nil
	}
}

func (s *StringObject) Eq(other any) any {
	if IsInstance(other, IntObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) == len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("==", s, other)
		return nil
	}
}

func (s *StringObject) Neq(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) != len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("!=", s, other)
		return nil
	}
}

func (s *StringObject) Gt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) > len(other.(*StringObject).value)))
	} else {
		ReportOperatorError(">", s, other)
		return nil
	}
}

func (s *StringObject) Lt(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) < len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("<", s, other)
		return nil
	}
}

func (s *StringObject) Gte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) >= len(other.(*StringObject).value)))
	} else {
		ReportOperatorError(">=", s, other)
		return nil
	}
}

func (s *StringObject) Lte(other any) any {
	if IsInstance(other, FloatObjectType) {
		return NewBoolObject(strconv.FormatBool(len(s.value) <= len(other.(*StringObject).value)))
	} else {
		ReportOperatorError("<=", s, other)
		return nil
	}
}

func (s *StringObject) Not() any {
	r := []rune(s.value)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return NewStringObject(`"` + string(r) + `"`)
}

func NewStringObject(value string) *StringObject {
	return &StringObject{value[1 : len(value)-1]}
}
