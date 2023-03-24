package core

import (
	"reflect"
)

func IsInstance(val any, t reflect.Type) bool {
	if reflect.TypeOf(val) == t {
		return true
	}

	return false
}

func ReprOfObject(val any) any {
	switch val.(type) {
	case *StringObject:
		return val.(*StringObject).Repr(val)
	case *IntObject:
		return val.(*IntObject).Repr(val)
	case *FloatObject:
		return val.(*FloatObject).Repr(val)
	case *BoolObject:
		return val.(*BoolObject).Repr(val)
	case *IdObject:
		return val.(*IdObject).Repr(val)
	default:
		return nil
	}
}
