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
		"", "'{'", "'}'", "'('", "')'", "'&'", "'['", "']'", "'='", "'.'", "','",
		"'?'", "", "", "", "", "'func'", "'if'", "'else'", "'while'", "'use'",
		"'override'", "'public'", "'private'", "'return'", "'break'", "'continue'",
		"'const'", "", "", "", "", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "ASSIGN", "DOT", "COMMA", "QUESTION",
		"NOT", "PREDTWO", "PREDONE", "COMPARATIVE", "FUNC", "IF", "ELSE", "WHILE",
		"USE", "OVERRIDE", "PUBLIC", "PRIVATE", "RETURN", "BREAK", "CONTINUE",
		"CONST", "WS", "COMMENT", "MULTILINECOMMENT", "BOOL", "NULL", "ID",
		"INT", "FLOAT", "STRING",
	}
	staticData.ruleNames = []string{
		"parse", "block", "stmt", "keywordStmts", "allStmts", "useStmt", "ifStmt",
		"whileStmt", "condition", "args", "params", "call", "assignments", "varAssign",
		"funcAssign", "getAttributes", "memoryAddress", "cast", "expr", "array",
		"atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 219, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 5, 0, 44, 8, 0, 10, 0, 12, 0, 47, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1,
		53, 8, 1, 10, 1, 12, 1, 56, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 64, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 70, 8, 3, 1, 4, 1, 4, 1, 4,
		3, 4, 75, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 5, 6, 88, 8, 6, 10, 6, 12, 6, 91, 9, 6, 1, 6, 1, 6, 3, 6, 95,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 106,
		8, 9, 10, 9, 12, 9, 109, 9, 9, 1, 10, 1, 10, 1, 10, 5, 10, 114, 8, 10,
		10, 10, 12, 10, 117, 9, 10, 1, 11, 1, 11, 1, 11, 3, 11, 122, 8, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 3, 12, 128, 8, 12, 1, 13, 3, 13, 131, 8, 13, 1,
		13, 3, 13, 134, 8, 13, 1, 13, 1, 13, 3, 13, 138, 8, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 3, 14, 144, 8, 14, 1, 14, 1, 14, 3, 14, 148, 8, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 153, 8, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 3, 15, 163, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 188, 8, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 199,
		8, 18, 10, 18, 12, 18, 202, 9, 18, 1, 19, 1, 19, 3, 19, 206, 8, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 217,
		8, 20, 1, 20, 0, 1, 36, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 3, 1, 0, 22, 23, 1, 0, 13, 14, 2,
		0, 33, 33, 36, 36, 235, 0, 45, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 63, 1,
		0, 0, 0, 6, 69, 1, 0, 0, 0, 8, 74, 1, 0, 0, 0, 10, 76, 1, 0, 0, 0, 12,
		79, 1, 0, 0, 0, 14, 96, 1, 0, 0, 0, 16, 100, 1, 0, 0, 0, 18, 102, 1, 0,
		0, 0, 20, 110, 1, 0, 0, 0, 22, 118, 1, 0, 0, 0, 24, 127, 1, 0, 0, 0, 26,
		130, 1, 0, 0, 0, 28, 143, 1, 0, 0, 0, 30, 157, 1, 0, 0, 0, 32, 166, 1,
		0, 0, 0, 34, 169, 1, 0, 0, 0, 36, 187, 1, 0, 0, 0, 38, 203, 1, 0, 0, 0,
		40, 216, 1, 0, 0, 0, 42, 44, 3, 4, 2, 0, 43, 42, 1, 0, 0, 0, 44, 47, 1,
		0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0, 47,
		45, 1, 0, 0, 0, 48, 49, 5, 0, 0, 1, 49, 1, 1, 0, 0, 0, 50, 54, 5, 1, 0,
		0, 51, 53, 3, 4, 2, 0, 52, 51, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0,
		57, 58, 5, 2, 0, 0, 58, 3, 1, 0, 0, 0, 59, 64, 3, 36, 18, 0, 60, 64, 3,
		24, 12, 0, 61, 64, 3, 8, 4, 0, 62, 64, 3, 6, 3, 0, 63, 59, 1, 0, 0, 0,
		63, 60, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 62, 1, 0, 0, 0, 64, 5, 1, 0,
		0, 0, 65, 70, 5, 25, 0, 0, 66, 70, 5, 26, 0, 0, 67, 68, 5, 24, 0, 0, 68,
		70, 3, 36, 18, 0, 69, 65, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0, 69, 67, 1, 0,
		0, 0, 70, 7, 1, 0, 0, 0, 71, 75, 3, 12, 6, 0, 72, 75, 3, 14, 7, 0, 73,
		75, 3, 10, 5, 0, 74, 71, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 73, 1, 0,
		0, 0, 75, 9, 1, 0, 0, 0, 76, 77, 5, 20, 0, 0, 77, 78, 5, 36, 0, 0, 78,
		11, 1, 0, 0, 0, 79, 80, 5, 17, 0, 0, 80, 81, 3, 16, 8, 0, 81, 89, 3, 2,
		1, 0, 82, 83, 5, 18, 0, 0, 83, 84, 5, 17, 0, 0, 84, 85, 3, 16, 8, 0, 85,
		86, 3, 2, 1, 0, 86, 88, 1, 0, 0, 0, 87, 82, 1, 0, 0, 0, 88, 91, 1, 0, 0,
		0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 94, 1, 0, 0, 0, 91, 89,
		1, 0, 0, 0, 92, 93, 5, 18, 0, 0, 93, 95, 3, 2, 1, 0, 94, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 13, 1, 0, 0, 0, 96, 97, 5, 19, 0, 0, 97, 98, 3,
		16, 8, 0, 98, 99, 3, 2, 1, 0, 99, 15, 1, 0, 0, 0, 100, 101, 3, 36, 18,
		0, 101, 17, 1, 0, 0, 0, 102, 107, 3, 36, 18, 0, 103, 104, 5, 10, 0, 0,
		104, 106, 3, 36, 18, 0, 105, 103, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107,
		105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 19, 1, 0, 0, 0, 109, 107, 1,
		0, 0, 0, 110, 115, 5, 33, 0, 0, 111, 112, 5, 10, 0, 0, 112, 114, 5, 33,
		0, 0, 113, 111, 1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0,
		115, 116, 1, 0, 0, 0, 116, 21, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 119,
		5, 33, 0, 0, 119, 121, 5, 3, 0, 0, 120, 122, 3, 18, 9, 0, 121, 120, 1,
		0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 5, 4, 0,
		0, 124, 23, 1, 0, 0, 0, 125, 128, 3, 26, 13, 0, 126, 128, 3, 28, 14, 0,
		127, 125, 1, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 25, 1, 0, 0, 0, 129, 131,
		7, 0, 0, 0, 130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 133, 1, 0,
		0, 0, 132, 134, 5, 27, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 137, 5, 33, 0, 0, 136, 138, 7, 1, 0, 0, 137,
		136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140,
		5, 8, 0, 0, 140, 141, 3, 36, 18, 0, 141, 27, 1, 0, 0, 0, 142, 144, 7, 0,
		0, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0,
		145, 147, 5, 16, 0, 0, 146, 148, 5, 21, 0, 0, 147, 146, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 5, 33, 0, 0, 150, 152,
		5, 3, 0, 0, 151, 153, 3, 20, 10, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1,
		0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 5, 4, 0, 0, 155, 156, 3, 2, 1,
		0, 156, 29, 1, 0, 0, 0, 157, 158, 7, 2, 0, 0, 158, 159, 5, 9, 0, 0, 159,
		160, 5, 33, 0, 0, 160, 162, 5, 3, 0, 0, 161, 163, 3, 18, 9, 0, 162, 161,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 5, 4,
		0, 0, 165, 31, 1, 0, 0, 0, 166, 167, 5, 5, 0, 0, 167, 168, 5, 33, 0, 0,
		168, 33, 1, 0, 0, 0, 169, 170, 3, 40, 20, 0, 170, 171, 5, 9, 0, 0, 171,
		172, 5, 3, 0, 0, 172, 173, 5, 33, 0, 0, 173, 174, 5, 4, 0, 0, 174, 35,
		1, 0, 0, 0, 175, 176, 6, 18, -1, 0, 176, 188, 3, 22, 11, 0, 177, 188, 3,
		40, 20, 0, 178, 179, 5, 3, 0, 0, 179, 180, 3, 36, 18, 0, 180, 181, 5, 4,
		0, 0, 181, 188, 1, 0, 0, 0, 182, 183, 5, 12, 0, 0, 183, 188, 3, 36, 18,
		7, 184, 188, 3, 34, 17, 0, 185, 188, 3, 30, 15, 0, 186, 188, 3, 32, 16,
		0, 187, 175, 1, 0, 0, 0, 187, 177, 1, 0, 0, 0, 187, 178, 1, 0, 0, 0, 187,
		182, 1, 0, 0, 0, 187, 184, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 186,
		1, 0, 0, 0, 188, 200, 1, 0, 0, 0, 189, 190, 10, 6, 0, 0, 190, 191, 5, 14,
		0, 0, 191, 199, 3, 36, 18, 7, 192, 193, 10, 5, 0, 0, 193, 194, 5, 13, 0,
		0, 194, 199, 3, 36, 18, 6, 195, 196, 10, 4, 0, 0, 196, 197, 5, 15, 0, 0,
		197, 199, 3, 36, 18, 5, 198, 189, 1, 0, 0, 0, 198, 192, 1, 0, 0, 0, 198,
		195, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201,
		1, 0, 0, 0, 201, 37, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 205, 5, 6,
		0, 0, 204, 206, 3, 18, 9, 0, 205, 204, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0,
		206, 207, 1, 0, 0, 0, 207, 208, 5, 7, 0, 0, 208, 39, 1, 0, 0, 0, 209, 217,
		3, 38, 19, 0, 210, 217, 5, 33, 0, 0, 211, 217, 5, 34, 0, 0, 212, 217, 5,
		35, 0, 0, 213, 217, 5, 36, 0, 0, 214, 217, 5, 32, 0, 0, 215, 217, 5, 31,
		0, 0, 216, 209, 1, 0, 0, 0, 216, 210, 1, 0, 0, 0, 216, 211, 1, 0, 0, 0,
		216, 212, 1, 0, 0, 0, 216, 213, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216,
		215, 1, 0, 0, 0, 217, 41, 1, 0, 0, 0, 23, 45, 54, 63, 69, 74, 89, 94, 107,
		115, 121, 127, 130, 133, 137, 143, 147, 152, 162, 187, 198, 200, 205, 216,
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
	CipherParserT__0             = 1
	CipherParserT__1             = 2
	CipherParserT__2             = 3
	CipherParserT__3             = 4
	CipherParserT__4             = 5
	CipherParserT__5             = 6
	CipherParserT__6             = 7
	CipherParserASSIGN           = 8
	CipherParserDOT              = 9
	CipherParserCOMMA            = 10
	CipherParserQUESTION         = 11
	CipherParserNOT              = 12
	CipherParserPREDTWO          = 13
	CipherParserPREDONE          = 14
	CipherParserCOMPARATIVE      = 15
	CipherParserFUNC             = 16
	CipherParserIF               = 17
	CipherParserELSE             = 18
	CipherParserWHILE            = 19
	CipherParserUSE              = 20
	CipherParserOVERRIDE         = 21
	CipherParserPUBLIC           = 22
	CipherParserPRIVATE          = 23
	CipherParserRETURN           = 24
	CipherParserBREAK            = 25
	CipherParserCONTINUE         = 26
	CipherParserCONST            = 27
	CipherParserWS               = 28
	CipherParserCOMMENT          = 29
	CipherParserMULTILINECOMMENT = 30
	CipherParserBOOL             = 31
	CipherParserNULL             = 32
	CipherParserID               = 33
	CipherParserINT              = 34
	CipherParserFLOAT            = 35
	CipherParserSTRING           = 36
)

// CipherParser rules.
const (
	CipherParserRULE_parse         = 0
	CipherParserRULE_block         = 1
	CipherParserRULE_stmt          = 2
	CipherParserRULE_keywordStmts  = 3
	CipherParserRULE_allStmts      = 4
	CipherParserRULE_useStmt       = 5
	CipherParserRULE_ifStmt        = 6
	CipherParserRULE_whileStmt     = 7
	CipherParserRULE_condition     = 8
	CipherParserRULE_args          = 9
	CipherParserRULE_params        = 10
	CipherParserRULE_call          = 11
	CipherParserRULE_assignments   = 12
	CipherParserRULE_varAssign     = 13
	CipherParserRULE_funcAssign    = 14
	CipherParserRULE_getAttributes = 15
	CipherParserRULE_memoryAddress = 16
	CipherParserRULE_cast          = 17
	CipherParserRULE_expr          = 18
	CipherParserRULE_array         = 19
	CipherParserRULE_atom          = 20
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135557484648) != 0 {
		{
			p.SetState(42)
			p.Stmt()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(CipherParserT__0)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135557484648) != 0 {
		{
			p.SetState(51)
			p.Stmt()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(CipherParserT__1)
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

func (s *StmtContext) KeywordStmts() IKeywordStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordStmtsContext)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Assignments()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.AllStmts()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.KeywordStmts()
		}

	}

	return localctx
}

// IKeywordStmtsContext is an interface to support dynamic dispatch.
type IKeywordStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordStmtsContext differentiates from other interfaces.
	IsKeywordStmtsContext()
}

type KeywordStmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordStmtsContext() *KeywordStmtsContext {
	var p = new(KeywordStmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_keywordStmts
	return p
}

func (*KeywordStmtsContext) IsKeywordStmtsContext() {}

func NewKeywordStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordStmtsContext {
	var p = new(KeywordStmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_keywordStmts

	return p
}

func (s *KeywordStmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordStmtsContext) BREAK() antlr.TerminalNode {
	return s.GetToken(CipherParserBREAK, 0)
}

func (s *KeywordStmtsContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(CipherParserCONTINUE, 0)
}

func (s *KeywordStmtsContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CipherParserRETURN, 0)
}

func (s *KeywordStmtsContext) Expr() IExprContext {
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

func (s *KeywordStmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordStmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordStmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitKeywordStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) KeywordStmts() (localctx IKeywordStmtsContext) {
	this := p
	_ = this

	localctx = NewKeywordStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CipherParserRULE_keywordStmts)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserBREAK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(CipherParserBREAK)
		}

	case CipherParserCONTINUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(CipherParserCONTINUE)
		}

	case CipherParserRETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(CipherParserRETURN)
		}
		{
			p.SetState(68)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *AllStmtsContext) UseStmt() IUseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseStmtContext)
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
	p.EnterRule(localctx, 8, CipherParserRULE_allStmts)

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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.IfStmt()
		}

	case CipherParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.WhileStmt()
		}

	case CipherParserUSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.UseStmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUseStmtContext is an interface to support dynamic dispatch.
type IUseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUseStmtContext differentiates from other interfaces.
	IsUseStmtContext()
}

type UseStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseStmtContext() *UseStmtContext {
	var p = new(UseStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_useStmt
	return p
}

func (*UseStmtContext) IsUseStmtContext() {}

func NewUseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseStmtContext {
	var p = new(UseStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_useStmt

	return p
}

func (s *UseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UseStmtContext) USE() antlr.TerminalNode {
	return s.GetToken(CipherParserUSE, 0)
}

func (s *UseStmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(CipherParserSTRING, 0)
}

func (s *UseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitUseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) UseStmt() (localctx IUseStmtContext) {
	this := p
	_ = this

	localctx = NewUseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CipherParserRULE_useStmt)

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
		p.SetState(76)
		p.Match(CipherParserUSE)
	}
	{
		p.SetState(77)
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
	p.EnterRule(localctx, 12, CipherParserRULE_ifStmt)
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
		p.SetState(79)
		p.Match(CipherParserIF)
	}
	{
		p.SetState(80)
		p.Condition()
	}
	{
		p.SetState(81)
		p.Block()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(82)
				p.Match(CipherParserELSE)
			}
			{
				p.SetState(83)
				p.Match(CipherParserIF)
			}
			{
				p.SetState(84)
				p.Condition()
			}
			{
				p.SetState(85)
				p.Block()
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserELSE {
		{
			p.SetState(92)
			p.Match(CipherParserELSE)
		}
		{
			p.SetState(93)
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
	p.EnterRule(localctx, 14, CipherParserRULE_whileStmt)

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
		p.SetState(96)
		p.Match(CipherParserWHILE)
	}
	{
		p.SetState(97)
		p.Condition()
	}
	{
		p.SetState(98)
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
	p.EnterRule(localctx, 16, CipherParserRULE_condition)

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
		p.SetState(100)
		p.expr(0)
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
	p.EnterRule(localctx, 18, CipherParserRULE_args)
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
		p.SetState(102)
		p.expr(0)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(103)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(104)
			p.expr(0)
		}

		p.SetState(109)
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
	p.EnterRule(localctx, 20, CipherParserRULE_params)
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
		p.SetState(110)
		p.Match(CipherParserID)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(111)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(112)
			p.Match(CipherParserID)
		}

		p.SetState(117)
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
	p.EnterRule(localctx, 22, CipherParserRULE_call)
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
		p.SetState(118)
		p.Match(CipherParserID)
	}
	{
		p.SetState(119)
		p.Match(CipherParserT__2)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291474024) != 0 {
		{
			p.SetState(120)
			p.Args()
		}

	}
	{
		p.SetState(123)
		p.Match(CipherParserT__3)
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
	p.EnterRule(localctx, 24, CipherParserRULE_assignments)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.VarAssign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
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

func (s *VarAssignContext) PREDONE() antlr.TerminalNode {
	return s.GetToken(CipherParserPREDONE, 0)
}

func (s *VarAssignContext) PREDTWO() antlr.TerminalNode {
	return s.GetToken(CipherParserPREDTWO, 0)
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
	p.EnterRule(localctx, 26, CipherParserRULE_varAssign)
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
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserPUBLIC || _la == CipherParserPRIVATE {
		{
			p.SetState(129)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CipherParserPUBLIC || _la == CipherParserPRIVATE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserCONST {
		{
			p.SetState(132)
			p.Match(CipherParserCONST)
		}

	}
	{
		p.SetState(135)
		p.Match(CipherParserID)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserPREDTWO || _la == CipherParserPREDONE {
		{
			p.SetState(136)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CipherParserPREDTWO || _la == CipherParserPREDONE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(139)
		p.Match(CipherParserASSIGN)
	}
	{
		p.SetState(140)
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
	p.EnterRule(localctx, 28, CipherParserRULE_funcAssign)
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
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserPUBLIC || _la == CipherParserPRIVATE {
		{
			p.SetState(142)
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
		p.SetState(145)
		p.Match(CipherParserFUNC)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserOVERRIDE {
		{
			p.SetState(146)
			p.Match(CipherParserOVERRIDE)
		}

	}
	{
		p.SetState(149)
		p.Match(CipherParserID)
	}
	{
		p.SetState(150)
		p.Match(CipherParserT__2)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserID {
		{
			p.SetState(151)
			p.Params()
		}

	}
	{
		p.SetState(154)
		p.Match(CipherParserT__3)
	}
	{
		p.SetState(155)
		p.Block()
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

func (s *GetAttributesContext) DOT() antlr.TerminalNode {
	return s.GetToken(CipherParserDOT, 0)
}

func (s *GetAttributesContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CipherParserID)
}

func (s *GetAttributesContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserID, i)
}

func (s *GetAttributesContext) STRING() antlr.TerminalNode {
	return s.GetToken(CipherParserSTRING, 0)
}

func (s *GetAttributesContext) Args() IArgsContext {
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
	p.EnterRule(localctx, 30, CipherParserRULE_getAttributes)
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
		p.SetState(157)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CipherParserID || _la == CipherParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(158)
		p.Match(CipherParserDOT)
	}
	{
		p.SetState(159)
		p.Match(CipherParserID)
	}
	{
		p.SetState(160)
		p.Match(CipherParserT__2)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291474024) != 0 {
		{
			p.SetState(161)
			p.Args()
		}

	}
	{
		p.SetState(164)
		p.Match(CipherParserT__3)
	}

	return localctx
}

// IMemoryAddressContext is an interface to support dynamic dispatch.
type IMemoryAddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemoryAddressContext differentiates from other interfaces.
	IsMemoryAddressContext()
}

type MemoryAddressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemoryAddressContext() *MemoryAddressContext {
	var p = new(MemoryAddressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_memoryAddress
	return p
}

func (*MemoryAddressContext) IsMemoryAddressContext() {}

func NewMemoryAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemoryAddressContext {
	var p = new(MemoryAddressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_memoryAddress

	return p
}

func (s *MemoryAddressContext) GetParser() antlr.Parser { return s.parser }

func (s *MemoryAddressContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *MemoryAddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemoryAddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemoryAddressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitMemoryAddress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) MemoryAddress() (localctx IMemoryAddressContext) {
	this := p
	_ = this

	localctx = NewMemoryAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CipherParserRULE_memoryAddress)

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
		p.SetState(166)
		p.Match(CipherParserT__4)
	}
	{
		p.SetState(167)
		p.Match(CipherParserID)
	}

	return localctx
}

// ICastContext is an interface to support dynamic dispatch.
type ICastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCastContext differentiates from other interfaces.
	IsCastContext()
}

type CastContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCastContext() *CastContext {
	var p = new(CastContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_cast
	return p
}

func (*CastContext) IsCastContext() {}

func NewCastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastContext {
	var p = new(CastContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_cast

	return p
}

func (s *CastContext) GetParser() antlr.Parser { return s.parser }

func (s *CastContext) Atom() IAtomContext {
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

func (s *CastContext) DOT() antlr.TerminalNode {
	return s.GetToken(CipherParserDOT, 0)
}

func (s *CastContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *CastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitCast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) Cast() (localctx ICastContext) {
	this := p
	_ = this

	localctx = NewCastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CipherParserRULE_cast)

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
		p.SetState(169)
		p.Atom()
	}
	{
		p.SetState(170)
		p.Match(CipherParserDOT)
	}
	{
		p.SetState(171)
		p.Match(CipherParserT__2)
	}
	{
		p.SetState(172)
		p.Match(CipherParserID)
	}
	{
		p.SetState(173)
		p.Match(CipherParserT__3)
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

func (s *ExprContext) Cast() ICastContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICastContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICastContext)
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

func (s *ExprContext) MemoryAddress() IMemoryAddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemoryAddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemoryAddressContext)
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
	_startState := 36
	p.EnterRecursionRule(localctx, 36, CipherParserRULE_expr, _p)

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
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(176)
			p.Call()
		}

	case 2:
		{
			p.SetState(177)
			p.Atom()
		}

	case 3:
		{
			p.SetState(178)
			p.Match(CipherParserT__2)
		}
		{
			p.SetState(179)
			p.expr(0)
		}
		{
			p.SetState(180)
			p.Match(CipherParserT__3)
		}

	case 4:
		{
			p.SetState(182)

			var _m = p.Match(CipherParserNOT)

			localctx.(*ExprContext).op = _m
		}
		{
			p.SetState(183)
			p.expr(7)
		}

	case 5:
		{
			p.SetState(184)
			p.Cast()
		}

	case 6:
		{
			p.SetState(185)
			p.GetAttributes()
		}

	case 7:
		{
			p.SetState(186)
			p.MemoryAddress()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(190)

					var _m = p.Match(CipherParserPREDONE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(191)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(193)

					var _m = p.Match(CipherParserPREDTWO)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(194)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(196)

					var _m = p.Match(CipherParserCOMPARATIVE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(197)
					p.expr(5)
				}

			}

		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 38, CipherParserRULE_array)
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
		p.SetState(203)
		p.Match(CipherParserT__5)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291474024) != 0 {
		{
			p.SetState(204)
			p.Args()
		}

	}
	{
		p.SetState(207)
		p.Match(CipherParserT__6)
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
	p.EnterRule(localctx, 40, CipherParserRULE_atom)

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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Array()
		}

	case CipherParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.Match(CipherParserID)
		}

	case CipherParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(211)
			p.Match(CipherParserINT)
		}

	case CipherParserFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(212)
			p.Match(CipherParserFLOAT)
		}

	case CipherParserSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(213)
			p.Match(CipherParserSTRING)
		}

	case CipherParserNULL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(214)
			p.Match(CipherParserNULL)
		}

	case CipherParserBOOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(215)
			p.Match(CipherParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *CipherParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
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
