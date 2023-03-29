package core

import (
	"Cipher/Core/parser"
	"fmt"
	"reflect"
)

type Visitor struct {
	*parser.BaseCipherVisitor

	Env       *Env
	Operators map[string]string
}

func (v *Visitor) VisitParse(ctx *parser.ParseContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		v.VisitStmt(stmt.(*parser.StmtContext))
	}

	return nil
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext, canUseKeywords bool) interface{} {
	for _, stmt := range ctx.AllStmt() {
		out := v.VisitStmt(stmt.(*parser.StmtContext))
		if msg, ok := out.(*Messenger); ok && !canUseKeywords {
			switch msg.storeMessage {
			case "BREAK":
				break
			case "CONTINUE":
				continue
			default:
				return msg.storeValue
			}
		}
	}

	return nil
}

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	switch {
	case ctx.Expr() != nil:
		return v.VisitExpr(ctx.Expr().(*parser.ExprContext))
	case ctx.Assignments() != nil:
		return v.VisitAssignments(ctx.Assignments().(*parser.AssignmentsContext))
	case ctx.AllStmts() != nil:
		return v.VisitAllStmts(ctx.AllStmts().(*parser.AllStmtsContext))
	case ctx.KeywordStmts() != nil:
		return v.VisitKeywordStmts(ctx.KeywordStmts().(*parser.KeywordStmtsContext))
	default:
		return nil
	}
}

func (v *Visitor) VisitKeywordStmts(ctx *parser.KeywordStmtsContext) *Messenger {
	switch {
	case ctx.RETURN() != nil:
		return NewMessenger("RETURN", v.VisitExpr(ctx.Expr().(*parser.ExprContext)))
	case ctx.BREAK() != nil:
		return NewMessenger("BREAK", nil)
	case ctx.CONTINUE() != nil:
		return NewMessenger("CONTINUE", nil)
	default:
		return NewMessenger("", nil)
	}
}

func (v *Visitor) VisitAllStmts(ctx *parser.AllStmtsContext) interface{} {
	switch {
	case ctx.IfStmt() != nil:
		return v.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
	case ctx.WhileStmt() != nil:
		return v.VisitWhileStmt(ctx.WhileStmt().(*parser.WhileStmtContext))
	case ctx.UseStmt() != nil:
		return v.VisitUseStmt(ctx.UseStmt().(*parser.UseStmtContext))
	default:
		return nil
	}
}

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	for v.VisitCondition(ctx.Condition().(*parser.ConditionContext)) {
		v.VisitBlock(ctx.Block().(*parser.BlockContext), true)
	}

	return nil
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	if v.VisitCondition(ctx.Condition(0).(*parser.ConditionContext)) {
		v.VisitBlock(ctx.Block(0).(*parser.BlockContext), false)
	} else {
		lengthOfIfs := len(ctx.AllIF())
		for i := 1; i < lengthOfIfs; i++ {
			if v.VisitCondition(ctx.Condition(i).(*parser.ConditionContext)) {
				v.VisitBlock(ctx.Block(i).(*parser.BlockContext), false)
				return nil
			}
		}
		if len(ctx.AllELSE()) > 0 {
			v.VisitBlock(ctx.Block(len(ctx.AllBlock())-1).(*parser.BlockContext), false)
		}
	}

	return nil
}

func (v *Visitor) VisitCondition(ctx *parser.ConditionContext) bool {
	expr := v.VisitExpr(ctx.Expr().(*parser.ExprContext))
	if b, ok := expr.(*BoolObject); ok {
		return b.value
	}
	ReportError("Type", "Conditions must be type 'bool'")
	return false
}

func (v *Visitor) VisitAssignments(ctx *parser.AssignmentsContext) interface{} {
	switch {
	case ctx.VarAssign() != nil:
		return v.VisitVarAssign(ctx.VarAssign().(*parser.VarAssignContext))
	case ctx.FuncAssign() != nil:
		return v.VisitFuncAssign(ctx.FuncAssign().(*parser.FuncAssignContext))
	default:
		return nil
	}
}

func (v *Visitor) VisitUseStmt(ctx *parser.UseStmtContext) interface{} {
	module := ctx.STRING().GetText()[1 : len(ctx.STRING().GetText())-1]

	switch module {
	case "time":
		lib := reflect.TypeOf(NewTimeLib())

		for i := 0; i < lib.NumField(); i++ {
			f := lib.Field(i)
			fmt.Printf("%s\n", f.Type)
		}

		for i := 0; i < lib.NumMethod(); i++ {
			m := lib.Method(i)
			fmt.Printf("%s\n", m.Name)
		}
	}

	return nil
}

func (v *Visitor) VisitArgs(ctx *parser.ArgsContext) []any {
	if ctx != nil {
		args := make([]any, len(ctx.AllExpr()))
		for i, expr := range ctx.AllExpr() {
			args[i] = v.VisitExpr(expr.(*parser.ExprContext))
		}
		return args
	}

	return []any{}
}

func (v *Visitor) VisitParams(ctx *parser.ParamsContext) []string {
	if ctx != nil {
		params := make([]string, len(ctx.AllID()))
		for i, id := range ctx.AllID() {
			params[i] = id.GetText()
		}
		return params
	}

	return []string{}
}

func (v *Visitor) VisitCall(ctx *parser.CallContext) interface{} {
	name := ctx.ID().GetText()
	args := PassArgs(ctx.Args(), v)

	if fn, ok := v.Env.functions[name]; ok {
		return fn.Call(args, v)
	}
	ReportError("Name", "Unknown function '"+name+"'")
	return nil
}

func (v *Visitor) VisitVarAssign(ctx *parser.VarAssignContext) interface{} {
	name := ctx.ID().GetText()
	value := v.VisitExpr(ctx.Expr().(*parser.ExprContext))
	constant := ctx.CONST() != nil

	if variable, ok := v.Env.variables[name]; ok {
		if variable.constant {
			ReportError("Type", "Cannot assign to a constant variable (immutable)")
		} else {
			if constant && variable.constant {
				ReportError("Assign", "Variable is already assigned as immutable (constant)")
			} else if constant && !variable.constant {
				ReportError("Assign", "Variable is already assigned as mutable")
			}

			v.Env.variables[name].value = value
		}
	}

	v.Env.variables[name] = NewVar(name, value, constant)
	return nil
}

func (v *Visitor) VisitFuncAssign(ctx *parser.FuncAssignContext) interface{} {
	name := ctx.ID().GetText()
	block := ctx.Block().(*parser.BlockContext)
	params := PassParams(ctx.Params(), v)

	/*
		override := ctx.OVERRIDE() != nil
		public := ctx.PRIVATE() == nil*/

	v.Env.functions[name] = NewFuncObject(name, params, block, nil)
	return nil
}

func (v *Visitor) VisitGetAttributes(ctx *parser.GetAttributesContext) interface{} {
	var obj any
	var attr string

	if len(ctx.AllID()) == 2 {
		obj = v.Env.variables[ctx.ID(0).GetText()].value
		attr = ctx.ID(1).GetText()
	} else if ctx.STRING() != nil {
		obj = NewStringObject(ctx.STRING().GetText())
		attr = ctx.ID(0).GetText()
	}

	args := PassArgs(ctx.Args(), v)

	m, _ := reflect.TypeOf(obj).MethodByName(attr)
	a := []reflect.Value{reflect.ValueOf(obj), reflect.ValueOf(args)}

	if !m.Func.IsValid() {
		ReportError("Attribute", fmt.Sprintf("No attribute '%s'", attr))
		return nil
	} else {
		return m.Func.Call(a)[0].Interface()
	}
}

func (v *Visitor) VisitMemoryAddress(ctx *parser.MemoryAddressContext) *StringObject {
	variableName := ctx.ID().GetText()
	if variable, ok := v.Env.variables[variableName]; ok {
		return NewStringObject(fmt.Sprintf("%p", &variable))
	} else {
		return NewStringObject("0x0")
	}
}

func (v *Visitor) VisitCast(ctx *parser.CastContext) interface{} {
	atom := v.VisitAtom(ctx.Atom().(*parser.AtomContext))
	castType := ctx.ID().GetText()

	m, _ := reflect.TypeOf(atom).MethodByName("As" + TitleString(castType))
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(atom)

	if !m.Func.IsValid() {
		ReportError("Type", "Invalid cast type '"+castType+"'")
		return nil
	} else {
		return m.Func.Call(args)[0].Interface()
	}
}

func (v *Visitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	switch {
	case ctx.Atom() != nil:
		return v.VisitAtom(ctx.Atom().(*parser.AtomContext))
	case ctx.Call() != nil:
		return v.VisitCall(ctx.Call().(*parser.CallContext))
	case ctx.GetAttributes() != nil:
		return v.VisitGetAttributes(ctx.GetAttributes().(*parser.GetAttributesContext))
	case ctx.MemoryAddress() != nil:
		return v.VisitMemoryAddress(ctx.MemoryAddress().(*parser.MemoryAddressContext))
	case ctx.Cast() != nil:
		return v.VisitCast(ctx.Cast().(*parser.CastContext))
	case len(ctx.AllExpr()) == 1 && ctx.GetOp() != nil:
		e := v.VisitExpr(ctx.Expr(0).(*parser.ExprContext))
		op := v.Operators[ctx.GetOp().GetText()]

		m, _ := reflect.TypeOf(e).MethodByName(op)
		args := make([]reflect.Value, 1)
		args[0] = reflect.ValueOf(e)

		if !m.Func.IsValid() {
			ReportError("Type", fmt.Sprintf("Invalid operands for type '%s'\n", ReprOfObject(e,
				nil)))
			return nil
		} else {
			return m.Func.Call(args)[0].Interface()
		}
	case len(ctx.AllExpr()) == 2 && ctx.GetOp() != nil:
		op := v.Operators[ctx.GetOp().GetText()]
		e1 := v.VisitExpr(ctx.Expr(0).(*parser.ExprContext))
		e2 := v.VisitExpr(ctx.Expr(1).(*parser.ExprContext))

		m, _ := reflect.TypeOf(e1).MethodByName(op)

		args := make([]reflect.Value, 2)
		args[0] = reflect.ValueOf(e1)
		args[1] = reflect.ValueOf(e2)

		if !m.Func.IsValid() {
			ReportOperatorError(ctx.GetOp().GetText(), e1, e2)
			return nil
		} else {
			return m.Func.Call(args)[0].Interface()
		}
	}

	return nil
}

func (v *Visitor) VisitArray(ctx *parser.ArrayContext) []any {
	return PassArgs(ctx.Args(), v)
}

func (v *Visitor) VisitAtom(ctx *parser.AtomContext) interface{} {
	txt := ctx.GetText()
	switch {
	case ctx.ID() != nil:
		if val, ok := v.Env.variables[txt]; ok {
			return val.value
		} else {
			if val, ok := v.Env.functions[txt]; ok {
				return val
			}

			ReportError("Name", "Unknown Object '"+txt+"'")
			return nil
		}
	case ctx.Array() != nil:
		return NewArrayObject(v.VisitArray(ctx.Array().(*parser.ArrayContext)))
	case ctx.STRING() != nil:
		return NewStringObject(txt)
	case ctx.INT() != nil:
		return NewIntObject(txt)
	case ctx.FLOAT() != nil:
		return NewFloatObject(txt)
	case ctx.BOOL() != nil:
		return NewBoolObject(txt)
	default:
		return NewNullObject()
	}
}

func NewArrayObject(args []any) *ArrayObject {
	return &ArrayObject{value: args}
}
