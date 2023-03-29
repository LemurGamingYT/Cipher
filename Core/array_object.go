package core

import (
	"strings"
)

type ArrayObject struct {
	value []any
}

func (a *ArrayObject) Repr(_ any) string {
	var builder strings.Builder

	builder.WriteRune('[')

	for i, obj := range a.value {
		repr := ReprOfObject(obj, a)
		if i != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(repr)
	}

	builder.WriteRune(']')

	return builder.String()
}

func (a *ArrayObject) Add(other any) *ArrayObject {
	if IsInstance(other, ArrayObjectType) {
		values := other.(*ArrayObject).value
		a.value = append(a.value[:len(a.value)], values...)
		return a
		/*values := other.(*ArrayObject).value
		newLen := len(a.value) + len(values)
		newValue := make([]interface{}, newLen)

		copy(newValue, a.value)
		copy(newValue[len(a.value):], values)

		return NewArrayObject(newValue)*/
	} else {
		ReportOperatorError("+", a, other)
		return nil
	}
}

func (a *ArrayObject) NewArrayObject(args []any) *ArrayObject {
	return &ArrayObject{value: args}
}
