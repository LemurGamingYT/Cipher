package core

import (
	"reflect"
)

type ArrayObject struct {
	value []any
}

func (a *ArrayObject) Repr(context any) string {
	repr := make([]byte, 0, 2+len(a.value)*10)
	repr = append(repr, '[')
	for i, obj := range a.value {
		m, _ := reflect.TypeOf(obj).MethodByName("Repr")
		args := []reflect.Value{reflect.ValueOf(a), reflect.ValueOf(obj)}
		repr = m.Func.Call(args)[0].Interface().([]byte)
		if i < len(a.value)-1 {
			repr = append(repr, ',')
		}
		repr = append(repr)
	}
	repr = append(repr, ']')
	return string(repr)
}

func (a *ArrayObject) NewArrayObject(args []any) *ArrayObject {
	return &ArrayObject{value: args}
}
