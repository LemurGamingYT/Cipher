// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Cipher

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by CipherParser.
type CipherVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CipherParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by CipherParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by CipherParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by CipherParser#iterationStmts.
	VisitIterationStmts(ctx *IterationStmtsContext) interface{}

	// Visit a parse tree produced by CipherParser#functionStmts.
	VisitFunctionStmts(ctx *FunctionStmtsContext) interface{}

	// Visit a parse tree produced by CipherParser#allStmts.
	VisitAllStmts(ctx *AllStmtsContext) interface{}

	// Visit a parse tree produced by CipherParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by CipherParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by CipherParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by CipherParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by CipherParser#undefineStmt.
	VisitUndefineStmt(ctx *UndefineStmtContext) interface{}

	// Visit a parse tree produced by CipherParser#inheritList.
	VisitInheritList(ctx *InheritListContext) interface{}

	// Visit a parse tree produced by CipherParser#classdef.
	VisitClassdef(ctx *ClassdefContext) interface{}

	// Visit a parse tree produced by CipherParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by CipherParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by CipherParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by CipherParser#assignments.
	VisitAssignments(ctx *AssignmentsContext) interface{}

	// Visit a parse tree produced by CipherParser#varAssign.
	VisitVarAssign(ctx *VarAssignContext) interface{}

	// Visit a parse tree produced by CipherParser#funcAssign.
	VisitFuncAssign(ctx *FuncAssignContext) interface{}

	// Visit a parse tree produced by CipherParser#getAttributes.
	VisitGetAttributes(ctx *GetAttributesContext) interface{}

	// Visit a parse tree produced by CipherParser#funcExpr.
	VisitFuncExpr(ctx *FuncExprContext) interface{}

	// Visit a parse tree produced by CipherParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by CipherParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by CipherParser#atom.
	VisitAtom(ctx *AtomContext) interface{}
}
