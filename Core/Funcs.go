package core

import (
	"bytes"
	"fmt"
	rand2 "math/rand"
	"strconv"
	"time"
)

func functionprintln(args []any, _ *Visitor) any {
	LengthArgsCheck(args, 1)
	value := args[0]

	fmt.Printf("%s\n", ReprOfObject(value, nil))

	return NewNullObject()
}

func functionprint(args []any, _ *Visitor) any {
	LengthArgsCheck(args, 1)
	value := args[0]

	fmt.Printf("%s", ReprOfObject(value, nil))

	return NewNullObject()
}

func functionany(args []any, _ *Visitor) any {
	if obj, ok := args[0].(*ArrayObject); ok {
		for _, val := range obj.value {
			if o, ok := val.(*BoolObject); ok {
				if o.value {
					return NewBoolObject("true")
				}
			}
		}
	}

	return NewBoolObject("false")
}

func functionall(args []any, _ *Visitor) any {
	if obj, ok := args[0].(*ArrayObject); ok {
		for _, val := range obj.value {
			if o, ok := val.(*BoolObject); ok {
				if !o.value {
					return NewBoolObject("false")
				}
			}
		}
	}

	return NewBoolObject("true")
}

func functionbin(args []any, _ *Visitor) any {
	LengthArgsCheck(args, 1)
	value := args[0]

	switch v := value.(type) {
	case *StringObject:
		s := v.value
		var buf bytes.Buffer
		buf.Grow(len(s) * 8)

		for _, char := range s {
			buf.WriteString(fmt.Sprintf("%08s", strconv.FormatInt(int64(char), 2)))
		}

		return NewStringObject(buf.String())
	default:
		ReportError("Type", "'bin' function requires type 'string' argument")
		return nil
	}
}

func functionrandomInt(args []any, _ *Visitor) any {
	LengthArgsCheck(args, 2)

	if val, ok := args[1].(*IntObject); ok {
		if val2, ok := args[0].(*IntObject); ok {
			rand2.Seed(time.Now().UnixNano())
			return NewIntObject(strconv.FormatInt(rand2.Int63n(val.value-val2.value+1)+val2.value, 10))
		} else {
			ReportError("Type", "'randomInt' function requires type 'int' and 'int' arguments")
			return nil
		}
	} else {
		ReportError("Type", "'randomInt' function requires type 'int' and 'int' arguments")
		return nil
	}
}

func functionprompt(args []any, _ *Visitor) any {
	LengthArgsCheck(args, 1)

	if prompt, ok := args[0].(*StringObject); ok {
		var u string
		println(prompt.value)
		_, err := fmt.Scanln(&u)
		if err != nil {
			ReportError("Internal", "An internal prompt error occurred")
			return nil
		}

		return NewStringObject("\"" + fmt.Sprintf("%s\n", u) + "\"")
	} else {
		ReportError("Type", "'prompt' function requires type 'string' as argument one")
		return nil
	}
}

func functionmin(args []any, _ *Visitor) any {
	LengthArgsCheck(args, 2)

	if val, ok := args[0].(*IntObject); ok {
		if val2, ok := args[1].(*IntObject); ok {
			if val.value > val2.value {
				return val2
			}
			return val
		} else {
			ReportError("Type", "'min' function requires type 'int' and 'int' arguments")
			return nil
		}
	} else {
		ReportError("Type", "'min' function requires type 'int' and 'int' arguments")
		return nil
	}
}

func functionmax(args []any, _ *Visitor) any {
	LengthArgsCheck(args, 2)

	if val, ok := args[0].(*IntObject); ok {
		if val2, ok := args[1].(*IntObject); ok {
			if val.value > val2.value {
				return val
			} else {
				return val2
			}
		} else {
			ReportError("Type", "'max' function requires type 'int' and 'int' arguments")
			return nil
		}
	} else {
		ReportError("Type", "'max' function requires type 'int' and 'int' arguments")
		return nil
	}
}
