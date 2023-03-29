package core

import (
	"math"
	"os"
	"strconv"
)

type Env struct {
	variables map[string]*VarObject
	functions map[string]*FuncObject
}

func NewEnv() *Env {
	variables := make(map[string]*VarObject)
	variables["Pi"] = NewVar("Pi", NewFloatObject(strconv.FormatFloat(math.Pi, 'f', 4, 64)),
		true)
	dir, err := os.Getwd()
	if err != nil {
		variables["Cd"] = NewVar("Cd", NewStringObject("\"C:\\\""), true)
	} else {
		variables["Cd"] = NewVar("Cd", NewStringObject("\""+dir+"\""),
			true)
	}

	functions := make(map[string]*FuncObject)
	functions["println"] = NewFuncObject("println", []string{"value"}, nil,
		functionprintln)
	functions["print"] = NewFuncObject("print", []string{"value"}, nil,
		functionprint)
	functions["min"] = NewFuncObject("min", []string{"value1", "value2"}, nil, functionmin)
	functions["max"] = NewFuncObject("max", []string{"value1", "value2"}, nil, functionmax)
	functions["bin"] = NewFuncObject("bin", []string{"string"}, nil, functionbin)
	functions["randomInt"] = NewFuncObject("randomInt", []string{"min", "max"}, nil, functionrandomInt)
	functions["prompt"] = NewFuncObject("prompt", []string{"prompt"}, nil, functionprompt)
	functions["any"] = NewFuncObject("any", []string{"array"}, nil, functionany)
	functions["all"] = NewFuncObject("all", []string{"array"}, nil, functionall)

	return &Env{variables, functions}
}
