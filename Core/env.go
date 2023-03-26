package core

import (
	"math"
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

	functions := make(map[string]*FuncObject)
	functions["println"] = NewFuncObject("println", []string{"value"}, nil,
		functionprintln)
	functions["print"] = NewFuncObject("print", []string{"value"}, nil,
		functionprint)
	functions["min"] = NewFuncObject("min", []string{"value1", "value2"}, nil, functionmin)
	functions["max"] = NewFuncObject("max", []string{"value1", "value2"}, nil, functionmax)

	return &Env{variables, functions}
}
