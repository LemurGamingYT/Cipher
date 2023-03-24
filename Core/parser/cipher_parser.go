// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Cipher

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CipherParser struct {
	*antlr.BaseParser
}

var cipherParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cipherParserInit() {
	staticData := &cipherParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "'.'", "','", "':'",
		"';'", "'->'", "'<-'", "'=>'", "'?'", "", "", "", "", "'func'", "'if'",
		"'else'", "'while'", "'import'", "'from'", "'override'", "'new'", "'class'",
		"'public'", "'private'", "'return'", "'break'", "'continue'", "'undefine'",
		"'const'", "", "", "", "", "'null'", "'''",
	}
	staticData.symbolicNames = []string{
		"", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LLIST", "RLIST", "ASSIGN",
		"DOT", "COMMA", "COLON", "SEMI", "LARROW", "RARROW", "ARROWASSIGN",
		"QUESTION", "NOT", "PREDTWO", "PREDONE", "COMPARATIVE", "FUNC", "IF",
		"ELSE", "WHILE", "IMPORT", "FROM", "OVERRIDE", "NEW", "CLASS", "PUBLIC",
		"PRIVATE", "RETURN", "BREAK", "CONTINUE", "UNDEFINE", "CONST", "WS",
		"COMMENT", "MULTILINECOMMENT", "BOOL", "NULL", "APOSTROPHE", "ID", "INT",
		"FLOAT", "STRING",
	}
	staticData.ruleNames = []string{
		"parse", "block", "stmt", "iterationStmts", "functionStmts", "allStmts",
		"importStmt", "ifStmt", "whileStmt", "condition", "undefineStmt", "inheritList",
		"classdef", "args", "params", "call", "assignments", "varAssign", "funcAssign",
		"getAttributes", "funcExpr", "expr", "array", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 269, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 5, 0, 50, 8, 0, 10, 0, 12,
		0, 53, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 59, 8, 1, 10, 1, 12, 1, 62,
		9, 1, 1, 1, 1, 1, 3, 1, 66, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 74, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		85, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 91, 8, 6, 10, 6, 12, 6, 94, 9,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 107, 8, 7, 10, 7, 12, 7, 110, 9, 7, 1, 7, 1, 7, 3, 7, 114, 8, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 5, 11, 130, 8, 11, 10, 11, 12, 11, 133, 9, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 139, 8, 12, 1, 13, 1, 13, 1, 13, 5, 13, 144,
		8, 13, 10, 13, 12, 13, 147, 9, 13, 1, 14, 1, 14, 1, 14, 5, 14, 152, 8,
		14, 10, 14, 12, 14, 155, 9, 14, 1, 15, 1, 15, 1, 15, 3, 15, 160, 8, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 166, 8, 16, 1, 17, 3, 17, 169, 8, 17,
		1, 17, 3, 17, 172, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 3, 18, 179,
		8, 18, 1, 18, 1, 18, 3, 18, 183, 8, 18, 1, 18, 1, 18, 1, 18, 3, 18, 188,
		8, 18, 1, 18, 1, 18, 1, 18, 3, 18, 193, 8, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 198, 8, 18, 1, 18, 1, 18, 1, 18, 3, 18, 203, 8, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 3, 19, 210, 8, 19, 1, 19, 5, 19, 213, 8, 19, 10, 19,
		12, 19, 216, 9, 19, 1, 20, 1, 20, 1, 20, 3, 20, 221, 8, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 238, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 249, 8, 21, 10, 21, 12, 21, 252, 9,
		21, 1, 22, 1, 22, 3, 22, 256, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 267, 8, 23, 1, 23, 0, 1, 42, 24, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 0, 2, 1, 0, 32, 33, 1, 0, 29, 30, 290, 0, 51, 1, 0, 0,
		0, 2, 65, 1, 0, 0, 0, 4, 73, 1, 0, 0, 0, 6, 75, 1, 0, 0, 0, 8, 77, 1, 0,
		0, 0, 10, 84, 1, 0, 0, 0, 12, 86, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 115,
		1, 0, 0, 0, 18, 119, 1, 0, 0, 0, 20, 121, 1, 0, 0, 0, 22, 126, 1, 0, 0,
		0, 24, 134, 1, 0, 0, 0, 26, 140, 1, 0, 0, 0, 28, 148, 1, 0, 0, 0, 30, 156,
		1, 0, 0, 0, 32, 165, 1, 0, 0, 0, 34, 168, 1, 0, 0, 0, 36, 202, 1, 0, 0,
		0, 38, 204, 1, 0, 0, 0, 40, 217, 1, 0, 0, 0, 42, 237, 1, 0, 0, 0, 44, 253,
		1, 0, 0, 0, 46, 266, 1, 0, 0, 0, 48, 50, 3, 4, 2, 0, 49, 48, 1, 0, 0, 0,
		50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1,
		0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 5, 0, 0, 1, 55, 1, 1, 0, 0, 0, 56,
		60, 5, 3, 0, 0, 57, 59, 3, 4, 2, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0,
		0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 63, 1, 0, 0, 0, 62, 60,
		1, 0, 0, 0, 63, 66, 5, 4, 0, 0, 64, 66, 3, 4, 2, 0, 65, 56, 1, 0, 0, 0,
		65, 64, 1, 0, 0, 0, 66, 3, 1, 0, 0, 0, 67, 74, 3, 42, 21, 0, 68, 74, 3,
		32, 16, 0, 69, 74, 3, 24, 12, 0, 70, 74, 3, 10, 5, 0, 71, 74, 3, 6, 3,
		0, 72, 74, 3, 8, 4, 0, 73, 67, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 73, 69,
		1, 0, 0, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0,
		74, 5, 1, 0, 0, 0, 75, 76, 7, 0, 0, 0, 76, 7, 1, 0, 0, 0, 77, 78, 5, 31,
		0, 0, 78, 79, 3, 42, 21, 0, 79, 9, 1, 0, 0, 0, 80, 85, 3, 14, 7, 0, 81,
		85, 3, 16, 8, 0, 82, 85, 3, 20, 10, 0, 83, 85, 3, 12, 6, 0, 84, 80, 1,
		0, 0, 0, 84, 81, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85,
		11, 1, 0, 0, 0, 86, 87, 5, 24, 0, 0, 87, 92, 5, 45, 0, 0, 88, 89, 5, 9,
		0, 0, 89, 91, 5, 45, 0, 0, 90, 88, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92,
		90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0,
		0, 95, 96, 5, 25, 0, 0, 96, 97, 5, 45, 0, 0, 97, 13, 1, 0, 0, 0, 98, 99,
		5, 21, 0, 0, 99, 100, 3, 18, 9, 0, 100, 108, 3, 2, 1, 0, 101, 102, 5, 22,
		0, 0, 102, 103, 5, 21, 0, 0, 103, 104, 3, 18, 9, 0, 104, 105, 3, 2, 1,
		0, 105, 107, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108,
		106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 113, 1, 0, 0, 0, 110, 108,
		1, 0, 0, 0, 111, 112, 5, 22, 0, 0, 112, 114, 3, 2, 1, 0, 113, 111, 1, 0,
		0, 0, 113, 114, 1, 0, 0, 0, 114, 15, 1, 0, 0, 0, 115, 116, 5, 23, 0, 0,
		116, 117, 3, 18, 9, 0, 117, 118, 3, 2, 1, 0, 118, 17, 1, 0, 0, 0, 119,
		120, 3, 42, 21, 0, 120, 19, 1, 0, 0, 0, 121, 122, 5, 34, 0, 0, 122, 123,
		5, 1, 0, 0, 123, 124, 5, 42, 0, 0, 124, 125, 5, 2, 0, 0, 125, 21, 1, 0,
		0, 0, 126, 131, 5, 42, 0, 0, 127, 128, 5, 9, 0, 0, 128, 130, 5, 42, 0,
		0, 129, 127, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131,
		132, 1, 0, 0, 0, 132, 23, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 135, 5,
		28, 0, 0, 135, 136, 5, 42, 0, 0, 136, 138, 5, 13, 0, 0, 137, 139, 3, 22,
		11, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 25, 1, 0, 0, 0,
		140, 145, 3, 42, 21, 0, 141, 142, 5, 9, 0, 0, 142, 144, 3, 42, 21, 0, 143,
		141, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146,
		1, 0, 0, 0, 146, 27, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 153, 5, 42,
		0, 0, 149, 150, 5, 9, 0, 0, 150, 152, 5, 42, 0, 0, 151, 149, 1, 0, 0, 0,
		152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		29, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 5, 42, 0, 0, 157, 159,
		5, 1, 0, 0, 158, 160, 3, 26, 13, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1,
		0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 2, 0, 0, 162, 31, 1, 0, 0,
		0, 163, 166, 3, 34, 17, 0, 164, 166, 3, 36, 18, 0, 165, 163, 1, 0, 0, 0,
		165, 164, 1, 0, 0, 0, 166, 33, 1, 0, 0, 0, 167, 169, 7, 1, 0, 0, 168, 167,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 171, 1, 0, 0, 0, 170, 172, 5, 35,
		0, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0,
		173, 174, 5, 42, 0, 0, 174, 175, 5, 7, 0, 0, 175, 176, 3, 42, 21, 0, 176,
		35, 1, 0, 0, 0, 177, 179, 7, 1, 0, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1,
		0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 5, 20, 0, 0, 181, 183, 5, 26,
		0, 0, 182, 181, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0,
		184, 185, 5, 42, 0, 0, 185, 187, 5, 1, 0, 0, 186, 188, 3, 28, 14, 0, 187,
		186, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190,
		5, 2, 0, 0, 190, 203, 3, 2, 1, 0, 191, 193, 7, 1, 0, 0, 192, 191, 1, 0,
		0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 5, 42, 0, 0,
		195, 197, 5, 1, 0, 0, 196, 198, 3, 28, 14, 0, 197, 196, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 5, 2, 0, 0, 200, 201,
		5, 14, 0, 0, 201, 203, 3, 2, 1, 0, 202, 178, 1, 0, 0, 0, 202, 192, 1, 0,
		0, 0, 203, 37, 1, 0, 0, 0, 204, 214, 3, 46, 23, 0, 205, 206, 5, 8, 0, 0,
		206, 207, 5, 42, 0, 0, 207, 209, 5, 1, 0, 0, 208, 210, 3, 26, 13, 0, 209,
		208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213,
		5, 2, 0, 0, 212, 205, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0,
		0, 0, 214, 215, 1, 0, 0, 0, 215, 39, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0,
		217, 218, 5, 42, 0, 0, 218, 220, 5, 1, 0, 0, 219, 221, 3, 28, 14, 0, 220,
		219, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223,
		5, 2, 0, 0, 223, 224, 5, 14, 0, 0, 224, 225, 3, 2, 1, 0, 225, 41, 1, 0,
		0, 0, 226, 227, 6, 21, -1, 0, 227, 238, 3, 30, 15, 0, 228, 238, 3, 46,
		23, 0, 229, 230, 5, 16, 0, 0, 230, 238, 3, 42, 21, 7, 231, 238, 3, 38,
		19, 0, 232, 238, 3, 40, 20, 0, 233, 234, 5, 1, 0, 0, 234, 235, 3, 42, 21,
		0, 235, 236, 5, 2, 0, 0, 236, 238, 1, 0, 0, 0, 237, 226, 1, 0, 0, 0, 237,
		228, 1, 0, 0, 0, 237, 229, 1, 0, 0, 0, 237, 231, 1, 0, 0, 0, 237, 232,
		1, 0, 0, 0, 237, 233, 1, 0, 0, 0, 238, 250, 1, 0, 0, 0, 239, 240, 10, 6,
		0, 0, 240, 241, 5, 18, 0, 0, 241, 249, 3, 42, 21, 7, 242, 243, 10, 5, 0,
		0, 243, 244, 5, 17, 0, 0, 244, 249, 3, 42, 21, 6, 245, 246, 10, 4, 0, 0,
		246, 247, 5, 19, 0, 0, 247, 249, 3, 42, 21, 5, 248, 239, 1, 0, 0, 0, 248,
		242, 1, 0, 0, 0, 248, 245, 1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250, 248,
		1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 43, 1, 0, 0, 0, 252, 250, 1, 0,
		0, 0, 253, 255, 5, 3, 0, 0, 254, 256, 3, 26, 13, 0, 255, 254, 1, 0, 0,
		0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 5, 4, 0, 0, 258,
		45, 1, 0, 0, 0, 259, 267, 3, 44, 22, 0, 260, 267, 5, 42, 0, 0, 261, 267,
		5, 43, 0, 0, 262, 267, 5, 44, 0, 0, 263, 267, 5, 45, 0, 0, 264, 267, 5,
		40, 0, 0, 265, 267, 5, 39, 0, 0, 266, 259, 1, 0, 0, 0, 266, 260, 1, 0,
		0, 0, 266, 261, 1, 0, 0, 0, 266, 262, 1, 0, 0, 0, 266, 263, 1, 0, 0, 0,
		266, 264, 1, 0, 0, 0, 266, 265, 1, 0, 0, 0, 267, 47, 1, 0, 0, 0, 30, 51,
		60, 65, 73, 84, 92, 108, 113, 131, 138, 145, 153, 159, 165, 168, 171, 178,
		182, 187, 192, 197, 202, 209, 214, 220, 237, 248, 250, 255, 266,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CipherParserInit initializes any static state used to implement CipherParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCipherParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CipherParserInit() {
	staticData := &cipherParserStaticData
	staticData.once.Do(cipherParserInit)
}

// NewCipherParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCipherParser(input antlr.TokenStream) *CipherParser {
	CipherParserInit()
	this := new(CipherParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cipherParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// CipherParser tokens.
const (
	CipherParserEOF              = antlr.TokenEOF
	CipherParserLPAREN           = 1
	CipherParserRPAREN           = 2
	CipherParserLBRACE           = 3
	CipherParserRBRACE           = 4
	CipherParserLLIST            = 5
	CipherParserRLIST            = 6
	CipherParserASSIGN           = 7
	CipherParserDOT              = 8
	CipherParserCOMMA            = 9
	CipherParserCOLON            = 10
	CipherParserSEMI             = 11
	CipherParserLARROW           = 12
	CipherParserRARROW           = 13
	CipherParserARROWASSIGN      = 14
	CipherParserQUESTION         = 15
	CipherParserNOT              = 16
	CipherParserPREDTWO          = 17
	CipherParserPREDONE          = 18
	CipherParserCOMPARATIVE      = 19
	CipherParserFUNC             = 20
	CipherParserIF               = 21
	CipherParserELSE             = 22
	CipherParserWHILE            = 23
	CipherParserIMPORT           = 24
	CipherParserFROM             = 25
	CipherParserOVERRIDE         = 26
	CipherParserNEW              = 27
	CipherParserCLASS            = 28
	CipherParserPUBLIC           = 29
	CipherParserPRIVATE          = 30
	CipherParserRETURN           = 31
	CipherParserBREAK            = 32
	CipherParserCONTINUE         = 33
	CipherParserUNDEFINE         = 34
	CipherParserCONST            = 35
	CipherParserWS               = 36
	CipherParserCOMMENT          = 37
	CipherParserMULTILINECOMMENT = 38
	CipherParserBOOL             = 39
	CipherParserNULL             = 40
	CipherParserAPOSTROPHE       = 41
	CipherParserID               = 42
	CipherParserINT              = 43
	CipherParserFLOAT            = 44
	CipherParserSTRING           = 45
)

// CipherParser rules.
const (
	CipherParserRULE_parse          = 0
	CipherParserRULE_block          = 1
	CipherParserRULE_stmt           = 2
	CipherParserRULE_iterationStmts = 3
	CipherParserRULE_functionStmts  = 4
	CipherParserRULE_allStmts       = 5
	CipherParserRULE_importStmt     = 6
	CipherParserRULE_ifStmt         = 7
	CipherParserRULE_whileStmt      = 8
	CipherParserRULE_condition      = 9
	CipherParserRULE_undefineStmt   = 10
	CipherParserRULE_inheritList    = 11
	CipherParserRULE_classdef       = 12
	CipherParserRULE_args           = 13
	CipherParserRULE_params         = 14
	CipherParserRULE_call           = 15
	CipherParserRULE_assignments    = 16
	CipherParserRULE_varAssign      = 17
	CipherParserRULE_funcAssign     = 18
	CipherParserRULE_getAttributes  = 19
	CipherParserRULE_funcExpr       = 20
	CipherParserRULE_expr           = 21
	CipherParserRULE_array          = 22
	CipherParserRULE_atom           = 23
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(CipherParserEOF, 0)
}

func (s *ParseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Parse() (localctx IParseContext) {
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CipherParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67688444526602) != 0 {
		{
			p.SetState(48)
			p.Stmt()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(CipherParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(CipherParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(CipherParserRBRACE, 0)
}

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CipherParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(CipherParserLBRACE)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67688444526602) != 0 {
			{
				p.SetState(57)
				p.Stmt()
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(63)
			p.Match(CipherParserRBRACE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Stmt()
		}

	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtContext) Assignments() IAssignmentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentsContext)
}

func (s *StmtContext) Classdef() IClassdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassdefContext)
}

func (s *StmtContext) AllStmts() IAllStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllStmtsContext)
}

func (s *StmtContext) IterationStmts() IIterationStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterationStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIterationStmtsContext)
}

func (s *StmtContext) FunctionStmts() IFunctionStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStmtsContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CipherParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Assignments()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Classdef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.AllStmts()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.IterationStmts()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)
			p.FunctionStmts()
		}

	}

	return localctx
}

// IIterationStmtsContext is an interface to support dynamic dispatch.
type IIterationStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStmtsContext differentiates from other interfaces.
	IsIterationStmtsContext()
}

type IterationStmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStmtsContext() *IterationStmtsContext {
	var p = new(IterationStmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_iterationStmts
	return p
}

func (*IterationStmtsContext) IsIterationStmtsContext() {}

func NewIterationStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStmtsContext {
	var p = new(IterationStmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_iterationStmts

	return p
}

func (s *IterationStmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStmtsContext) BREAK() antlr.TerminalNode {
	return s.GetToken(CipherParserBREAK, 0)
}

func (s *IterationStmtsContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(CipherParserCONTINUE, 0)
}

func (s *IterationStmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationStmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitIterationStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) IterationStmts() (localctx IIterationStmtsContext) {
	this := p
	_ = this

	localctx = NewIterationStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CipherParserRULE_iterationStmts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CipherParserBREAK || _la == CipherParserCONTINUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionStmtsContext is an interface to support dynamic dispatch.
type IFunctionStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionStmtsContext differentiates from other interfaces.
	IsFunctionStmtsContext()
}

type FunctionStmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStmtsContext() *FunctionStmtsContext {
	var p = new(FunctionStmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_functionStmts
	return p
}

func (*FunctionStmtsContext) IsFunctionStmtsContext() {}

func NewFunctionStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStmtsContext {
	var p = new(FunctionStmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_functionStmts

	return p
}

func (s *FunctionStmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStmtsContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CipherParserRETURN, 0)
}

func (s *FunctionStmtsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionStmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionStmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitFunctionStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) FunctionStmts() (localctx IFunctionStmtsContext) {
	this := p
	_ = this

	localctx = NewFunctionStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CipherParserRULE_functionStmts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(CipherParserRETURN)
	}
	{
		p.SetState(78)
		p.expr(0)
	}

	return localctx
}

// IAllStmtsContext is an interface to support dynamic dispatch.
type IAllStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllStmtsContext differentiates from other interfaces.
	IsAllStmtsContext()
}

type AllStmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllStmtsContext() *AllStmtsContext {
	var p = new(AllStmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_allStmts
	return p
}

func (*AllStmtsContext) IsAllStmtsContext() {}

func NewAllStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllStmtsContext {
	var p = new(AllStmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_allStmts

	return p
}

func (s *AllStmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *AllStmtsContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *AllStmtsContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *AllStmtsContext) UndefineStmt() IUndefineStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndefineStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndefineStmtContext)
}

func (s *AllStmtsContext) ImportStmt() IImportStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *AllStmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllStmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllStmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitAllStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) AllStmts() (localctx IAllStmtsContext) {
	this := p
	_ = this

	localctx = NewAllStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CipherParserRULE_allStmts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.IfStmt()
		}

	case CipherParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.WhileStmt()
		}

	case CipherParserUNDEFINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.UndefineStmt()
		}

	case CipherParserIMPORT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.ImportStmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_importStmt
	return p
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(CipherParserIMPORT, 0)
}

func (s *ImportStmtContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(CipherParserSTRING)
}

func (s *ImportStmtContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserSTRING, i)
}

func (s *ImportStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(CipherParserFROM, 0)
}

func (s *ImportStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CipherParserCOMMA)
}

func (s *ImportStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserCOMMA, i)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) ImportStmt() (localctx IImportStmtContext) {
	this := p
	_ = this

	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CipherParserRULE_importStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(CipherParserIMPORT)
	}
	{
		p.SetState(87)
		p.Match(CipherParserSTRING)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(88)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(89)
			p.Match(CipherParserSTRING)
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(CipherParserFROM)
	}
	{
		p.SetState(96)
		p.Match(CipherParserSTRING)
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(CipherParserIF)
}

func (s *IfStmtContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserIF, i)
}

func (s *IfStmtContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(CipherParserELSE)
}

func (s *IfStmtContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserELSE, i)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CipherParserRULE_ifStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(CipherParserIF)
	}
	{
		p.SetState(99)
		p.Condition()
	}
	{
		p.SetState(100)
		p.Block()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(101)
				p.Match(CipherParserELSE)
			}
			{
				p.SetState(102)
				p.Match(CipherParserIF)
			}
			{
				p.SetState(103)
				p.Condition()
			}
			{
				p.SetState(104)
				p.Block()
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(111)
			p.Match(CipherParserELSE)
		}
		{
			p.SetState(112)
			p.Block()
		}

	}

	return localctx
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_whileStmt
	return p
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(CipherParserWHILE, 0)
}

func (s *WhileStmtContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) WhileStmt() (localctx IWhileStmtContext) {
	this := p
	_ = this

	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CipherParserRULE_whileStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(CipherParserWHILE)
	}
	{
		p.SetState(116)
		p.Condition()
	}
	{
		p.SetState(117)
		p.Block()
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CipherParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.expr(0)
	}

	return localctx
}

// IUndefineStmtContext is an interface to support dynamic dispatch.
type IUndefineStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUndefineStmtContext differentiates from other interfaces.
	IsUndefineStmtContext()
}

type UndefineStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUndefineStmtContext() *UndefineStmtContext {
	var p = new(UndefineStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_undefineStmt
	return p
}

func (*UndefineStmtContext) IsUndefineStmtContext() {}

func NewUndefineStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UndefineStmtContext {
	var p = new(UndefineStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_undefineStmt

	return p
}

func (s *UndefineStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UndefineStmtContext) UNDEFINE() antlr.TerminalNode {
	return s.GetToken(CipherParserUNDEFINE, 0)
}

func (s *UndefineStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, 0)
}

func (s *UndefineStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *UndefineStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, 0)
}

func (s *UndefineStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndefineStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UndefineStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitUndefineStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) UndefineStmt() (localctx IUndefineStmtContext) {
	this := p
	_ = this

	localctx = NewUndefineStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CipherParserRULE_undefineStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(CipherParserUNDEFINE)
	}
	{
		p.SetState(122)
		p.Match(CipherParserLPAREN)
	}
	{
		p.SetState(123)
		p.Match(CipherParserID)
	}
	{
		p.SetState(124)
		p.Match(CipherParserRPAREN)
	}

	return localctx
}

// IInheritListContext is an interface to support dynamic dispatch.
type IInheritListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInheritListContext differentiates from other interfaces.
	IsInheritListContext()
}

type InheritListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInheritListContext() *InheritListContext {
	var p = new(InheritListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_inheritList
	return p
}

func (*InheritListContext) IsInheritListContext() {}

func NewInheritListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InheritListContext {
	var p = new(InheritListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_inheritList

	return p
}

func (s *InheritListContext) GetParser() antlr.Parser { return s.parser }

func (s *InheritListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CipherParserID)
}

func (s *InheritListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserID, i)
}

func (s *InheritListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CipherParserCOMMA)
}

func (s *InheritListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserCOMMA, i)
}

func (s *InheritListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InheritListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitInheritList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) InheritList() (localctx IInheritListContext) {
	this := p
	_ = this

	localctx = NewInheritListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CipherParserRULE_inheritList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(CipherParserID)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(127)
				p.Match(CipherParserCOMMA)
			}
			{
				p.SetState(128)
				p.Match(CipherParserID)
			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IClassdefContext is an interface to support dynamic dispatch.
type IClassdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassdefContext differentiates from other interfaces.
	IsClassdefContext()
}

type ClassdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassdefContext() *ClassdefContext {
	var p = new(ClassdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_classdef
	return p
}

func (*ClassdefContext) IsClassdefContext() {}

func NewClassdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassdefContext {
	var p = new(ClassdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_classdef

	return p
}

func (s *ClassdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassdefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(CipherParserCLASS, 0)
}

func (s *ClassdefContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *ClassdefContext) RARROW() antlr.TerminalNode {
	return s.GetToken(CipherParserRARROW, 0)
}

func (s *ClassdefContext) InheritList() IInheritListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInheritListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInheritListContext)
}

func (s *ClassdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitClassdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Classdef() (localctx IClassdefContext) {
	this := p
	_ = this

	localctx = NewClassdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CipherParserRULE_classdef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(CipherParserCLASS)
	}
	{
		p.SetState(135)
		p.Match(CipherParserID)
	}
	{
		p.SetState(136)
		p.Match(CipherParserRARROW)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(137)
			p.InheritList()
		}

	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CipherParserCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserCOMMA, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Args() (localctx IArgsContext) {
	this := p
	_ = this

	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CipherParserRULE_args)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.expr(0)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(141)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(142)
			p.expr(0)
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CipherParserID)
}

func (s *ParamsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserID, i)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CipherParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserCOMMA, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Params() (localctx IParamsContext) {
	this := p
	_ = this

	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CipherParserRULE_params)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(CipherParserID)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(149)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(150)
			p.Match(CipherParserID)
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *CallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, 0)
}

func (s *CallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, 0)
}

func (s *CallContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Call() (localctx ICallContext) {
	this := p
	_ = this

	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CipherParserRULE_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(CipherParserID)
	}
	{
		p.SetState(157)
		p.Match(CipherParserLPAREN)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67619965173770) != 0 {
		{
			p.SetState(158)
			p.Args()
		}

	}
	{
		p.SetState(161)
		p.Match(CipherParserRPAREN)
	}

	return localctx
}

// IAssignmentsContext is an interface to support dynamic dispatch.
type IAssignmentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentsContext differentiates from other interfaces.
	IsAssignmentsContext()
}

type AssignmentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentsContext() *AssignmentsContext {
	var p = new(AssignmentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_assignments
	return p
}

func (*AssignmentsContext) IsAssignmentsContext() {}

func NewAssignmentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentsContext {
	var p = new(AssignmentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_assignments

	return p
}

func (s *AssignmentsContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentsContext) VarAssign() IVarAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAssignContext)
}

func (s *AssignmentsContext) FuncAssign() IFuncAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncAssignContext)
}

func (s *AssignmentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitAssignments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Assignments() (localctx IAssignmentsContext) {
	this := p
	_ = this

	localctx = NewAssignmentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CipherParserRULE_assignments)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.VarAssign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.FuncAssign()
		}

	}

	return localctx
}

// IVarAssignContext is an interface to support dynamic dispatch.
type IVarAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarAssignContext differentiates from other interfaces.
	IsVarAssignContext()
}

type VarAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignContext() *VarAssignContext {
	var p = new(VarAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_varAssign
	return p
}

func (*VarAssignContext) IsVarAssignContext() {}

func NewVarAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignContext {
	var p = new(VarAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_varAssign

	return p
}

func (s *VarAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *VarAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CipherParserASSIGN, 0)
}

func (s *VarAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarAssignContext) CONST() antlr.TerminalNode {
	return s.GetToken(CipherParserCONST, 0)
}

func (s *VarAssignContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(CipherParserPUBLIC, 0)
}

func (s *VarAssignContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(CipherParserPRIVATE, 0)
}

func (s *VarAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitVarAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) VarAssign() (localctx IVarAssignContext) {
	this := p
	_ = this

	localctx = NewVarAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CipherParserRULE_varAssign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserPUBLIC || _la == CipherParserPRIVATE {
		{
			p.SetState(167)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CipherParserPUBLIC || _la == CipherParserPRIVATE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserCONST {
		{
			p.SetState(170)
			p.Match(CipherParserCONST)
		}

	}
	{
		p.SetState(173)
		p.Match(CipherParserID)
	}
	{
		p.SetState(174)
		p.Match(CipherParserASSIGN)
	}
	{
		p.SetState(175)
		p.expr(0)
	}

	return localctx
}

// IFuncAssignContext is an interface to support dynamic dispatch.
type IFuncAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncAssignContext differentiates from other interfaces.
	IsFuncAssignContext()
}

type FuncAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncAssignContext() *FuncAssignContext {
	var p = new(FuncAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_funcAssign
	return p
}

func (*FuncAssignContext) IsFuncAssignContext() {}

func NewFuncAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncAssignContext {
	var p = new(FuncAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_funcAssign

	return p
}

func (s *FuncAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncAssignContext) FUNC() antlr.TerminalNode {
	return s.GetToken(CipherParserFUNC, 0)
}

func (s *FuncAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *FuncAssignContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, 0)
}

func (s *FuncAssignContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, 0)
}

func (s *FuncAssignContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncAssignContext) OVERRIDE() antlr.TerminalNode {
	return s.GetToken(CipherParserOVERRIDE, 0)
}

func (s *FuncAssignContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FuncAssignContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(CipherParserPUBLIC, 0)
}

func (s *FuncAssignContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(CipherParserPRIVATE, 0)
}

func (s *FuncAssignContext) ARROWASSIGN() antlr.TerminalNode {
	return s.GetToken(CipherParserARROWASSIGN, 0)
}

func (s *FuncAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitFuncAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) FuncAssign() (localctx IFuncAssignContext) {
	this := p
	_ = this

	localctx = NewFuncAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CipherParserRULE_funcAssign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CipherParserPUBLIC || _la == CipherParserPRIVATE {
			{
				p.SetState(177)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CipherParserPUBLIC || _la == CipherParserPRIVATE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(180)
			p.Match(CipherParserFUNC)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CipherParserOVERRIDE {
			{
				p.SetState(181)
				p.Match(CipherParserOVERRIDE)
			}

		}
		{
			p.SetState(184)
			p.Match(CipherParserID)
		}
		{
			p.SetState(185)
			p.Match(CipherParserLPAREN)
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CipherParserID {
			{
				p.SetState(186)
				p.Params()
			}

		}
		{
			p.SetState(189)
			p.Match(CipherParserRPAREN)
		}
		{
			p.SetState(190)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CipherParserPUBLIC || _la == CipherParserPRIVATE {
			{
				p.SetState(191)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CipherParserPUBLIC || _la == CipherParserPRIVATE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(194)
			p.Match(CipherParserID)
		}
		{
			p.SetState(195)
			p.Match(CipherParserLPAREN)
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CipherParserID {
			{
				p.SetState(196)
				p.Params()
			}

		}
		{
			p.SetState(199)
			p.Match(CipherParserRPAREN)
		}
		{
			p.SetState(200)
			p.Match(CipherParserARROWASSIGN)
		}
		{
			p.SetState(201)
			p.Block()
		}

	}

	return localctx
}

// IGetAttributesContext is an interface to support dynamic dispatch.
type IGetAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetAttributesContext differentiates from other interfaces.
	IsGetAttributesContext()
}

type GetAttributesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetAttributesContext() *GetAttributesContext {
	var p = new(GetAttributesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_getAttributes
	return p
}

func (*GetAttributesContext) IsGetAttributesContext() {}

func NewGetAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetAttributesContext {
	var p = new(GetAttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_getAttributes

	return p
}

func (s *GetAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *GetAttributesContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *GetAttributesContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(CipherParserDOT)
}

func (s *GetAttributesContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserDOT, i)
}

func (s *GetAttributesContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CipherParserID)
}

func (s *GetAttributesContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserID, i)
}

func (s *GetAttributesContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(CipherParserLPAREN)
}

func (s *GetAttributesContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, i)
}

func (s *GetAttributesContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(CipherParserRPAREN)
}

func (s *GetAttributesContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, i)
}

func (s *GetAttributesContext) AllArgs() []IArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgsContext); ok {
			len++
		}
	}

	tst := make([]IArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgsContext); ok {
			tst[i] = t.(IArgsContext)
			i++
		}
	}

	return tst
}

func (s *GetAttributesContext) Args(i int) IArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *GetAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetAttributesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitGetAttributes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) GetAttributes() (localctx IGetAttributesContext) {
	this := p
	_ = this

	localctx = NewGetAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CipherParserRULE_getAttributes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Atom()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(205)
				p.Match(CipherParserDOT)
			}
			{
				p.SetState(206)
				p.Match(CipherParserID)
			}
			{
				p.SetState(207)
				p.Match(CipherParserLPAREN)
			}
			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67619965173770) != 0 {
				{
					p.SetState(208)
					p.Args()
				}

			}
			{
				p.SetState(211)
				p.Match(CipherParserRPAREN)
			}

		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncExprContext is an interface to support dynamic dispatch.
type IFuncExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExprContext differentiates from other interfaces.
	IsFuncExprContext()
}

type FuncExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExprContext() *FuncExprContext {
	var p = new(FuncExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_funcExpr
	return p
}

func (*FuncExprContext) IsFuncExprContext() {}

func NewFuncExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExprContext {
	var p = new(FuncExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_funcExpr

	return p
}

func (s *FuncExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *FuncExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, 0)
}

func (s *FuncExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, 0)
}

func (s *FuncExprContext) ARROWASSIGN() antlr.TerminalNode {
	return s.GetToken(CipherParserARROWASSIGN, 0)
}

func (s *FuncExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncExprContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) FuncExpr() (localctx IFuncExprContext) {
	this := p
	_ = this

	localctx = NewFuncExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CipherParserRULE_funcExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(CipherParserID)
	}
	{
		p.SetState(218)
		p.Match(CipherParserLPAREN)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserID {
		{
			p.SetState(219)
			p.Params()
		}

	}
	{
		p.SetState(222)
		p.Match(CipherParserRPAREN)
	}
	{
		p.SetState(223)
		p.Match(CipherParserARROWASSIGN)
	}
	{
		p.SetState(224)
		p.Block()
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *ExprContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(CipherParserNOT, 0)
}

func (s *ExprContext) GetAttributes() IGetAttributesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetAttributesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetAttributesContext)
}

func (s *ExprContext) FuncExpr() IFuncExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncExprContext)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, 0)
}

func (s *ExprContext) PREDONE() antlr.TerminalNode {
	return s.GetToken(CipherParserPREDONE, 0)
}

func (s *ExprContext) PREDTWO() antlr.TerminalNode {
	return s.GetToken(CipherParserPREDTWO, 0)
}

func (s *ExprContext) COMPARATIVE() antlr.TerminalNode {
	return s.GetToken(CipherParserCOMPARATIVE, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CipherParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, CipherParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(227)
			p.Call()
		}

	case 2:
		{
			p.SetState(228)
			p.Atom()
		}

	case 3:
		{
			p.SetState(229)

			var _m = p.Match(CipherParserNOT)

			localctx.(*ExprContext).op = _m
		}
		{
			p.SetState(230)
			p.expr(7)
		}

	case 4:
		{
			p.SetState(231)
			p.GetAttributes()
		}

	case 5:
		{
			p.SetState(232)
			p.FuncExpr()
		}

	case 6:
		{
			p.SetState(233)
			p.Match(CipherParserLPAREN)
		}
		{
			p.SetState(234)
			p.expr(0)
		}
		{
			p.SetState(235)
			p.Match(CipherParserRPAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(240)

					var _m = p.Match(CipherParserPREDONE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(241)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(243)

					var _m = p.Match(CipherParserPREDTWO)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(244)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(246)

					var _m = p.Match(CipherParserCOMPARATIVE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(247)
					p.expr(5)
				}

			}

		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(CipherParserLBRACE, 0)
}

func (s *ArrayContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(CipherParserRBRACE, 0)
}

func (s *ArrayContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CipherParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(CipherParserLBRACE)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67619965173770) != 0 {
		{
			p.SetState(254)
			p.Args()
		}

	}
	{
		p.SetState(257)
		p.Match(CipherParserRBRACE)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *AtomContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *AtomContext) INT() antlr.TerminalNode {
	return s.GetToken(CipherParserINT, 0)
}

func (s *AtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CipherParserFLOAT, 0)
}

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(CipherParserSTRING, 0)
}

func (s *AtomContext) NULL() antlr.TerminalNode {
	return s.GetToken(CipherParserNULL, 0)
}

func (s *AtomContext) BOOL() antlr.TerminalNode {
	return s.GetToken(CipherParserBOOL, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CipherParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Array()
		}

	case CipherParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Match(CipherParserID)
		}

	case CipherParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(261)
			p.Match(CipherParserINT)
		}

	case CipherParserFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.Match(CipherParserFLOAT)
		}

	case CipherParserSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(263)
			p.Match(CipherParserSTRING)
		}

	case CipherParserNULL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(264)
			p.Match(CipherParserNULL)
		}

	case CipherParserBOOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(265)
			p.Match(CipherParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *CipherParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CipherParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
