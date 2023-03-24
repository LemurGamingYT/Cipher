package core

import (
	"Cipher/Core/parser"
)

type FuncObject struct {
	name   string
	params []string

	block    *parser.BlockContext
	function func(args []any, v *Visitor) any
}

func (f *FuncObject) Call(args []any, v *Visitor) any {
	if f.block != nil {
		return v.VisitBlock(f.block)
	} else if f.function != nil {
		return f.function(args, v)
	} else {
		return NewNullObject()
	}
}

func NewFuncObject(name string, params []string, block *parser.BlockContext,
	function func(args []any, v *Visitor) any) *FuncObject {
	return &FuncObject{name, params, block, function}
}
