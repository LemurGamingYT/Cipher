package core

import "fmt"

func functionprintln(args []any, visitor *Visitor) any {
	value := args[0]

	fmt.Printf("%s\n", ReprOfObject(value))

	/*
		switch value.(type) {
		case *StringObject:
			println(value.(*StringObject).value)
		case *IntObject:
			println(value.(*IntObject).value)
		case *FloatObject:
			println(value.(*FloatObject).value)
		case *BoolObject:
			println(value.(*BoolObject).value)
		case *IdObject:
			println(visitor.Env.variables[value.(*IdObject).value].value)
		case NullObject:
			println("null")
		default:
			r := reflect.ValueOf(val).MethodByName("Repr")
			return r.Call([]reflect.Value{reflect.ValueOf(val)})
		}*/

	return NewNullObject()
}

func functionprint(args []any, visitor *Visitor) any {
	value := args[0]

	fmt.Printf("%s\n", ReprOfObject(value))

	return NewNullObject()
}
