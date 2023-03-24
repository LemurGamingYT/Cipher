package core

import (
	"Cipher/Core/parser"
)

type Visitor struct {
	parser.BaseCipherVisitor

	Env *Env
}

func (v *Visitor) VisitParse(ctx *parser.ParseContext) interface{} {
	for i := 0; i < ctx.GetChildCount()-1; i++ {
		v.VisitStmt(ctx.Stmt(i).(*parser.StmtContext))
	}

	return nil
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	if ctx.GetChildCount() > 1 {
		for i := 0; i < ctx.GetChildCount()-2; i++ {
			v.VisitStmt(ctx.Stmt(i).(*parser.StmtContext))
		}
	} else {
		v.VisitStmt(ctx.Stmt(0).(*parser.StmtContext))
	}

	return nil
}

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.Expr() != nil {
		return v.VisitExpr(ctx.Expr().(*parser.ExprContext))
	} else if ctx.Assignments() != nil {
		return v.VisitAssignments(ctx.Assignments().(*parser.AssignmentsContext))
	} else if ctx.AllStmts() != nil {
		return v.VisitAllStmts(ctx.AllStmts().(*parser.AllStmtsContext))
	} else {
		return nil
	}
}

func (v *Visitor) VisitAllStmts(ctx *parser.AllStmtsContext) interface{} {
	if ctx.IfStmt() != nil {
		return v.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
	} else if ctx.WhileStmt() != nil {
		return v.VisitWhileStmt(ctx.WhileStmt().(*parser.WhileStmtContext))
	} else if ctx.UndefineStmt() != nil {
		return v.VisitUndefineStmt(ctx.UndefineStmt().(*parser.UndefineStmtContext))
	}

	return nil
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	condition := v.VisitCondition(ctx.Condition(0).(*parser.ConditionContext))

	if condition {
		v.VisitBlock(ctx.Block(0).(*parser.BlockContext))
	} else {
	}

	return nil
}

func (v *Visitor) VisitCondition(ctx *parser.ConditionContext) bool {
	expr := v.VisitExpr(ctx.Expr().(*parser.ExprContext))

	if b, ok := expr.(*BoolObject); ok {
		return b.value
	} else {
		ReportError("Type", "Conditions must be type 'bool'")
		return false
	}
}

func (v *Visitor) VisitUndefineStmt(ctx *parser.UndefineStmtContext) interface{} {
	variable := ctx.ID().GetText()

	if _, ok := v.Env.variables[variable]; ok {
		delete(v.Env.variables, variable)
	} else {
		ReportError("Name", "Cannot undefine a variable that is not defined")
	}

	return nil
}

func (v *Visitor) VisitAssignments(ctx *parser.AssignmentsContext) interface{} {
	if ctx.VarAssign() != nil {
		return v.VisitVarAssign(ctx.VarAssign().(*parser.VarAssignContext))
	} else if ctx.FuncAssign() != nil {
		return v.VisitFuncAssign(ctx.FuncAssign().(*parser.FuncAssignContext))
	}

	return nil
}

func (v *Visitor) VisitArgs(ctx *parser.ArgsContext) []any {
	var args []any
	if ctx != nil {
		for i := 0; i < ctx.GetChildCount(); i++ {
			args = append(args, v.VisitExpr(ctx.Expr(i).(*parser.ExprContext)))
		}
	}

	return args
}

func (v *Visitor) VisitParams(ctx *parser.ParamsContext) []string {
	var params []string
	if ctx != nil {
		for i := 0; i < ctx.GetChildCount(); i++ {
			params = append(params, ctx.ID(i).GetText())
		}
	}

	return params
}

func (v *Visitor) VisitCall(ctx *parser.CallContext) interface{} {
	name := ctx.ID().GetText()
	var args []any
	if ctx.Args() != nil {
		args = v.VisitArgs(ctx.Args().(*parser.ArgsContext))
	} else {
		args = v.VisitArgs(nil)
	}

	for _, s := range v.Env.functions {
		if s.name == name {
			return s.Call(args, v)
		}
	}

	return nil
}

func (v *Visitor) VisitVarAssign(ctx *parser.VarAssignContext) interface{} {
	name := ctx.ID().GetText()
	value := v.VisitExpr(ctx.Expr().(*parser.ExprContext))
	constant := ctx.CONST() != nil

	v.Env.variables[name] = NewVar(name, value, constant)
	return nil
}

func (v *Visitor) VisitFuncAssign(ctx *parser.FuncAssignContext) interface{} {
	name := ctx.ID().GetText()
	block := ctx.Block().(*parser.BlockContext)
	var params []string
	if ctx.Params() != nil {
		params = v.VisitParams(ctx.Params().(*parser.ParamsContext))
	} else {
		params = v.VisitParams(nil)
	}
	/*
		override := ctx.OVERRIDE() != nil
		public := ctx.PRIVATE() == nil*/

	v.Env.functions[name] = NewFuncObject(name, params, block, nil)
	return nil
}

func (v *Visitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	if ctx.Atom() != nil {
		return v.VisitAtom(ctx.Atom().(*parser.AtomContext))
	} else if ctx.Call() != nil {
		return v.VisitCall(ctx.Call().(*parser.CallContext))
	}

	return nil
}

func (v *Visitor) VisitAtom(ctx *parser.AtomContext) interface{} {
	txt := ctx.GetText()
	if ctx.ID() != nil {
		if val, ok := v.Env.variables[txt]; ok {
			return val.value
		} else {
			if val, ok := v.Env.functions[txt]; ok {
				return val
			} else {
				ReportError("Name", "Unknown object '"+txt+"'")
			}
		}
	} else if ctx.STRING() != nil {
		return NewStringObject(txt)
	} else if ctx.INT() != nil {
		return NewIntObject(txt)
	} else if ctx.FLOAT() != nil {
		return NewFloatObject(txt)
	} else if ctx.BOOL() != nil {
		return NewBoolObject(txt)
	}

	return NewNullObject()
}
