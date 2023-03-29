package core

import (
	"Cipher/Core/parser"
	"fmt"
	"reflect"
	"strings"
)

func IsInstance(val any, t reflect.Type) bool {
	if reflect.TypeOf(val) == t {
		return true
	}

	return false
}

func ReportOperatorError(op string, typeOne any, typeTwo any) {
	ReportError("Type", fmt.Sprintf("Invalid operands (type '%s' and '%s') for operation '%s'\n",
		TypeOfObject(typeOne), TypeOfObject(typeTwo), op))
}

func LengthArgsCheck(args []any, needed int) {
	if len(args) < needed {
		if needed == 1 {
			ReportError("Index", fmt.Sprintf("Expected 1 argument, got %d\n", len(args)))
		}

		ReportError("Index", fmt.Sprintf("Expected %d arguments, got %d\n", needed, len(args)))
	}
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

func ReprOfObject(val any, context any) string {
	switch val.(type) {
	case *StringObject:
		return val.(*StringObject).Repr(context)
	case *IntObject:
		return val.(*IntObject).Repr(context)
	case *FloatObject:
		return val.(*FloatObject).Repr(context)
	case *BoolObject:
		return val.(*BoolObject).Repr(context)
	case *ArrayObject:
		return val.(*ArrayObject).Repr(context)
	default:
		return "null"
	}
}

func TypeOfObject(val any) string {
	switch (val).(type) {
	case *ArrayObject:
		return "array"
	case *IntObject:
		return "int"
	case *FloatObject:
		return "float"
	case *BoolObject:
		return "bool"
	case *StringObject:
		return "string"
	default:
		return "null"
	}
}

func TitleString(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word[0:1]) + word[1:]
	}

	return strings.Join(words, " ")
}
