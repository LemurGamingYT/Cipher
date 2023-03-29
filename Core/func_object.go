package core

import (
	"Cipher/Core/parser"
	"fmt"
)

type FuncObject struct {
	name   string
	params []string

	block    *parser.BlockContext
	function func(args []any, v *Visitor) any
}

func (f *FuncObject) Call(args []any, v *Visitor) any {
	if f.block != nil {
		return v.VisitBlock(f.block, false)
	} else if f.function != nil {
		return f.function(args, v)
	} else {
		return NewNullObject()
	}
}

func (f *FuncObject) Repr(_ any) string {
	return fmt.Sprintf("Function(name=%s, at=%p, params=%s)\n", f.name, f, f.params)
}

func NewFuncObject(name string, params []string, block *parser.BlockContext,
	function func(args []any, v *Visitor) any) *FuncObject {
	return &FuncObject{name, params, block, function}
}
