package core

import (
	"Cipher/Core/parser"
	"fmt"
	"reflect"
)

func IsInstance(val any, t reflect.Type) bool {
	if reflect.TypeOf(val) == t {
		return true
	}

	return false
}

func ReportOperatorError(op string) {
	ReportError("Type", fmt.Sprintf("Invalid operands for operation '%s'\n", op))
}

func LengthArgsCheck(args []any, needed int) {
	if len(args) < needed {
		if needed == 1 {
			ReportError("Index", fmt.Sprintf("Expected 1 argument, got %d\n", len(args)))
		}

		ReportError("Index", fmt.Sprintf("Expected %d arguments, got %d\n", needed, len(args)))
	}
}

func GetArgument(args []any, i int) any {
	return args[i]
}

func PassArgs(argsContext parser.IArgsContext, visitor *Visitor) []any {
	var args []any
	if argsContext != nil {
		args = visitor.VisitArgs(argsContext.(*parser.ArgsContext))
	} else {
		args = visitor.VisitArgs(nil)
	}

	return args
}

func PassParams(paramsContext parser.IParamsContext, visitor *Visitor) []string {
	var params []string
	if paramsContext != nil {
		params = visitor.VisitParams(paramsContext.(*parser.ParamsContext))
	} else {
		params = visitor.VisitParams(nil)
	}

	return params
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
	case *ArrayObject:
		return val.(*ArrayObject).Repr(val)
	default:
		return nil
	}
}
