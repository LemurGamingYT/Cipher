package core

import "fmt"

func functionprintln(args []any, visitor *Visitor) any {
	LengthArgsCheck(args, 1)
	value := GetArgument(args, 0)

	fmt.Printf("%s\n", ReprOfObject(value))

	return NewNullObject()
}

func functionprint(args []any, visitor *Visitor) any {
	LengthArgsCheck(args, 1)
	value := GetArgument(args, 0)

	fmt.Printf("%s", ReprOfObject(value))

	return NewNullObject()
}

func functionmin(args []any, visitor *Visitor) any {
	if val, ok := args[0].(*IntObject); ok {
		if val2, ok := args[1].(*IntObject); ok {
			if val.value > val2.value {
				return val2
			}
			return val
		}
	}

	return args[0]
}

func functionmax(args []any, visitor *Visitor) any {
	if val, ok := args[0].(*IntObject); ok {
		if val2, ok := args[1].(*IntObject); ok {
			if val.value > val2.value {
				return val
			} else {
				return val2
			}
		}
	}

	return args[0]
}
