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
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "'.'", "','", "'?'",
		"", "", "", "", "'func'", "'if'", "'else'", "'while'", "'use'", "'override'",
		"'public'", "'private'", "'return'", "'break'", "'continue'", "'undefine'",
		"'const'", "", "", "", "", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LLIST", "RLIST", "ASSIGN",
		"DOT", "COMMA", "QUESTION", "NOT", "PREDTWO", "PREDONE", "COMPARATIVE",
		"FUNC", "IF", "ELSE", "WHILE", "USE", "OVERRIDE", "PUBLIC", "PRIVATE",
		"RETURN", "BREAK", "CONTINUE", "UNDEFINE", "CONST", "WS", "COMMENT",
		"MULTILINECOMMENT", "BOOL", "NULL", "ID", "INT", "FLOAT", "STRING",
	}
	staticData.ruleNames = []string{
		"parse", "block", "stmt", "keywordStmts", "allStmts", "useList", "useStmt",
		"ifStmt", "whileStmt", "condition", "args", "params", "call", "assignments",
		"varAssign", "funcAssign", "getAttributes", "expr", "array", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 211, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 5, 0, 42,
		8, 0, 10, 0, 12, 0, 45, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 51, 8, 1, 10,
		1, 12, 1, 54, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 62, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 73, 8,
		4, 1, 5, 1, 5, 1, 5, 5, 5, 78, 8, 5, 10, 5, 12, 5, 81, 9, 5, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 94, 8, 7, 10,
		7, 12, 7, 97, 9, 7, 1, 7, 1, 7, 3, 7, 101, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 112, 8, 10, 10, 10, 12, 10, 115,
		9, 10, 1, 11, 1, 11, 1, 11, 5, 11, 120, 8, 11, 10, 11, 12, 11, 123, 9,
		11, 1, 12, 1, 12, 1, 12, 3, 12, 128, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		3, 13, 134, 8, 13, 1, 14, 3, 14, 137, 8, 14, 1, 14, 3, 14, 140, 8, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 3, 15, 147, 8, 15, 1, 15, 1, 15, 3,
		15, 151, 8, 15, 1, 15, 1, 15, 1, 15, 3, 15, 156, 8, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 166, 8, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3,
		17, 180, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 5, 17, 191, 8, 17, 10, 17, 12, 17, 194, 9, 17, 1, 18, 1, 18, 3,
		18, 198, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 3, 19, 209, 8, 19, 1, 19, 0, 1, 34, 20, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 1, 1, 0, 21, 22, 226,
		0, 43, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4, 61, 1, 0, 0, 0, 6, 67, 1, 0, 0,
		0, 8, 72, 1, 0, 0, 0, 10, 74, 1, 0, 0, 0, 12, 82, 1, 0, 0, 0, 14, 85, 1,
		0, 0, 0, 16, 102, 1, 0, 0, 0, 18, 106, 1, 0, 0, 0, 20, 108, 1, 0, 0, 0,
		22, 116, 1, 0, 0, 0, 24, 124, 1, 0, 0, 0, 26, 133, 1, 0, 0, 0, 28, 136,
		1, 0, 0, 0, 30, 146, 1, 0, 0, 0, 32, 160, 1, 0, 0, 0, 34, 179, 1, 0, 0,
		0, 36, 195, 1, 0, 0, 0, 38, 208, 1, 0, 0, 0, 40, 42, 3, 4, 2, 0, 41, 40,
		1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0,
		44, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 5, 0, 0, 1, 47, 1, 1, 0,
		0, 0, 48, 52, 5, 3, 0, 0, 49, 51, 3, 4, 2, 0, 50, 49, 1, 0, 0, 0, 51, 54,
		1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 55, 1, 0, 0, 0,
		54, 52, 1, 0, 0, 0, 55, 56, 5, 4, 0, 0, 56, 3, 1, 0, 0, 0, 57, 62, 3, 34,
		17, 0, 58, 62, 3, 26, 13, 0, 59, 62, 3, 8, 4, 0, 60, 62, 3, 6, 3, 0, 61,
		57, 1, 0, 0, 0, 61, 58, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 60, 1, 0, 0,
		0, 62, 5, 1, 0, 0, 0, 63, 68, 5, 24, 0, 0, 64, 68, 5, 25, 0, 0, 65, 66,
		5, 23, 0, 0, 66, 68, 3, 34, 17, 0, 67, 63, 1, 0, 0, 0, 67, 64, 1, 0, 0,
		0, 67, 65, 1, 0, 0, 0, 68, 7, 1, 0, 0, 0, 69, 73, 3, 14, 7, 0, 70, 73,
		3, 16, 8, 0, 71, 73, 3, 12, 6, 0, 72, 69, 1, 0, 0, 0, 72, 70, 1, 0, 0,
		0, 72, 71, 1, 0, 0, 0, 73, 9, 1, 0, 0, 0, 74, 79, 5, 36, 0, 0, 75, 76,
		5, 9, 0, 0, 76, 78, 5, 36, 0, 0, 77, 75, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0,
		79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 11, 1, 0, 0, 0, 81, 79, 1,
		0, 0, 0, 82, 83, 5, 19, 0, 0, 83, 84, 3, 10, 5, 0, 84, 13, 1, 0, 0, 0,
		85, 86, 5, 16, 0, 0, 86, 87, 3, 18, 9, 0, 87, 95, 3, 2, 1, 0, 88, 89, 5,
		17, 0, 0, 89, 90, 5, 16, 0, 0, 90, 91, 3, 18, 9, 0, 91, 92, 3, 2, 1, 0,
		92, 94, 1, 0, 0, 0, 93, 88, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1,
		0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 100, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98,
		99, 5, 17, 0, 0, 99, 101, 3, 2, 1, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1,
		0, 0, 0, 101, 15, 1, 0, 0, 0, 102, 103, 5, 18, 0, 0, 103, 104, 3, 18, 9,
		0, 104, 105, 3, 2, 1, 0, 105, 17, 1, 0, 0, 0, 106, 107, 3, 34, 17, 0, 107,
		19, 1, 0, 0, 0, 108, 113, 3, 34, 17, 0, 109, 110, 5, 9, 0, 0, 110, 112,
		3, 34, 17, 0, 111, 109, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1,
		0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 21, 1, 0, 0, 0, 115, 113, 1, 0, 0,
		0, 116, 121, 5, 33, 0, 0, 117, 118, 5, 9, 0, 0, 118, 120, 5, 33, 0, 0,
		119, 117, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121,
		122, 1, 0, 0, 0, 122, 23, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 125, 5,
		33, 0, 0, 125, 127, 5, 1, 0, 0, 126, 128, 3, 20, 10, 0, 127, 126, 1, 0,
		0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 5, 2, 0, 0,
		130, 25, 1, 0, 0, 0, 131, 134, 3, 28, 14, 0, 132, 134, 3, 30, 15, 0, 133,
		131, 1, 0, 0, 0, 133, 132, 1, 0, 0, 0, 134, 27, 1, 0, 0, 0, 135, 137, 7,
		0, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0,
		0, 138, 140, 5, 27, 0, 0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 142, 5, 33, 0, 0, 142, 143, 5, 7, 0, 0, 143, 144,
		3, 34, 17, 0, 144, 29, 1, 0, 0, 0, 145, 147, 7, 0, 0, 0, 146, 145, 1, 0,
		0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 5, 15, 0, 0,
		149, 151, 5, 20, 0, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		152, 1, 0, 0, 0, 152, 153, 5, 33, 0, 0, 153, 155, 5, 1, 0, 0, 154, 156,
		3, 22, 11, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1,
		0, 0, 0, 157, 158, 5, 2, 0, 0, 158, 159, 3, 2, 1, 0, 159, 31, 1, 0, 0,
		0, 160, 161, 3, 38, 19, 0, 161, 162, 5, 8, 0, 0, 162, 163, 5, 33, 0, 0,
		163, 165, 5, 1, 0, 0, 164, 166, 3, 20, 10, 0, 165, 164, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 5, 2, 0, 0, 168, 33, 1,
		0, 0, 0, 169, 170, 6, 17, -1, 0, 170, 180, 3, 24, 12, 0, 171, 180, 3, 38,
		19, 0, 172, 173, 5, 1, 0, 0, 173, 174, 3, 34, 17, 0, 174, 175, 5, 2, 0,
		0, 175, 180, 1, 0, 0, 0, 176, 177, 5, 11, 0, 0, 177, 180, 3, 34, 17, 5,
		178, 180, 3, 32, 16, 0, 179, 169, 1, 0, 0, 0, 179, 171, 1, 0, 0, 0, 179,
		172, 1, 0, 0, 0, 179, 176, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 192,
		1, 0, 0, 0, 181, 182, 10, 4, 0, 0, 182, 183, 5, 13, 0, 0, 183, 191, 3,
		34, 17, 5, 184, 185, 10, 3, 0, 0, 185, 186, 5, 12, 0, 0, 186, 191, 3, 34,
		17, 4, 187, 188, 10, 2, 0, 0, 188, 189, 5, 14, 0, 0, 189, 191, 3, 34, 17,
		3, 190, 181, 1, 0, 0, 0, 190, 184, 1, 0, 0, 0, 190, 187, 1, 0, 0, 0, 191,
		194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 35, 1,
		0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 197, 5, 5, 0, 0, 196, 198, 3, 20, 10,
		0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199,
		200, 5, 6, 0, 0, 200, 37, 1, 0, 0, 0, 201, 209, 3, 36, 18, 0, 202, 209,
		5, 33, 0, 0, 203, 209, 5, 34, 0, 0, 204, 209, 5, 35, 0, 0, 205, 209, 5,
		36, 0, 0, 206, 209, 5, 32, 0, 0, 207, 209, 5, 31, 0, 0, 208, 201, 1, 0,
		0, 0, 208, 202, 1, 0, 0, 0, 208, 203, 1, 0, 0, 0, 208, 204, 1, 0, 0, 0,
		208, 205, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 207, 1, 0, 0, 0, 209,
		39, 1, 0, 0, 0, 23, 43, 52, 61, 67, 72, 79, 95, 100, 113, 121, 127, 133,
		136, 139, 146, 150, 155, 165, 179, 190, 192, 197, 208,
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
	CipherParserQUESTION         = 10
	CipherParserNOT              = 11
	CipherParserPREDTWO          = 12
	CipherParserPREDONE          = 13
	CipherParserCOMPARATIVE      = 14
	CipherParserFUNC             = 15
	CipherParserIF               = 16
	CipherParserELSE             = 17
	CipherParserWHILE            = 18
	CipherParserUSE              = 19
	CipherParserOVERRIDE         = 20
	CipherParserPUBLIC           = 21
	CipherParserPRIVATE          = 22
	CipherParserRETURN           = 23
	CipherParserBREAK            = 24
	CipherParserCONTINUE         = 25
	CipherParserUNDEFINE         = 26
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
	CipherParserRULE_useList       = 5
	CipherParserRULE_useStmt       = 6
	CipherParserRULE_ifStmt        = 7
	CipherParserRULE_whileStmt     = 8
	CipherParserRULE_condition     = 9
	CipherParserRULE_args          = 10
	CipherParserRULE_params        = 11
	CipherParserRULE_call          = 12
	CipherParserRULE_assignments   = 13
	CipherParserRULE_varAssign     = 14
	CipherParserRULE_funcAssign    = 15
	CipherParserRULE_getAttributes = 16
	CipherParserRULE_expr          = 17
	CipherParserRULE_array         = 18
	CipherParserRULE_atom          = 19
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135491586082) != 0 {
		{
			p.SetState(40)
			p.Stmt()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(46)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(CipherParserLBRACE)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135491586082) != 0 {
		{
			p.SetState(49)
			p.Stmt()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(CipherParserRBRACE)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Assignments()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.AllStmts()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(60)
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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserBREAK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(CipherParserBREAK)
		}

	case CipherParserCONTINUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(CipherParserCONTINUE)
		}

	case CipherParserRETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.Match(CipherParserRETURN)
		}
		{
			p.SetState(66)
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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.IfStmt()
		}

	case CipherParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.WhileStmt()
		}

	case CipherParserUSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.UseStmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUseListContext is an interface to support dynamic dispatch.
type IUseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUseListContext differentiates from other interfaces.
	IsUseListContext()
}

type UseListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseListContext() *UseListContext {
	var p = new(UseListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CipherParserRULE_useList
	return p
}

func (*UseListContext) IsUseListContext() {}

func NewUseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseListContext {
	var p = new(UseListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CipherParserRULE_useList

	return p
}

func (s *UseListContext) GetParser() antlr.Parser { return s.parser }

func (s *UseListContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(CipherParserSTRING)
}

func (s *UseListContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserSTRING, i)
}

func (s *UseListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CipherParserCOMMA)
}

func (s *UseListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CipherParserCOMMA, i)
}

func (s *UseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CipherVisitor:
		return t.VisitUseList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CipherParser) UseList() (localctx IUseListContext) {
	this := p
	_ = this

	localctx = NewUseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CipherParserRULE_useList)
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
		p.SetState(74)
		p.Match(CipherParserSTRING)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(75)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(76)
			p.Match(CipherParserSTRING)
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *UseStmtContext) UseList() IUseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseListContext)
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
	p.EnterRule(localctx, 12, CipherParserRULE_useStmt)

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
		p.SetState(82)
		p.Match(CipherParserUSE)
	}
	{
		p.SetState(83)
		p.UseList()
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
		p.SetState(85)
		p.Match(CipherParserIF)
	}
	{
		p.SetState(86)
		p.Condition()
	}
	{
		p.SetState(87)
		p.Block()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(88)
				p.Match(CipherParserELSE)
			}
			{
				p.SetState(89)
				p.Match(CipherParserIF)
			}
			{
				p.SetState(90)
				p.Condition()
			}
			{
				p.SetState(91)
				p.Block()
			}

		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserELSE {
		{
			p.SetState(98)
			p.Match(CipherParserELSE)
		}
		{
			p.SetState(99)
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
		p.SetState(102)
		p.Match(CipherParserWHILE)
	}
	{
		p.SetState(103)
		p.Condition()
	}
	{
		p.SetState(104)
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
		p.SetState(106)
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
	p.EnterRule(localctx, 20, CipherParserRULE_args)
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
		p.SetState(108)
		p.expr(0)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(109)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(110)
			p.expr(0)
		}

		p.SetState(115)
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
	p.EnterRule(localctx, 22, CipherParserRULE_params)
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
		p.SetState(116)
		p.Match(CipherParserID)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CipherParserCOMMA {
		{
			p.SetState(117)
			p.Match(CipherParserCOMMA)
		}
		{
			p.SetState(118)
			p.Match(CipherParserID)
		}

		p.SetState(123)
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
	p.EnterRule(localctx, 24, CipherParserRULE_call)
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
		p.SetState(124)
		p.Match(CipherParserID)
	}
	{
		p.SetState(125)
		p.Match(CipherParserLPAREN)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291471906) != 0 {
		{
			p.SetState(126)
			p.Args()
		}

	}
	{
		p.SetState(129)
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
	p.EnterRule(localctx, 26, CipherParserRULE_assignments)

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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.VarAssign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
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
	p.EnterRule(localctx, 28, CipherParserRULE_varAssign)
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
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserPUBLIC || _la == CipherParserPRIVATE {
		{
			p.SetState(135)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CipherParserPUBLIC || _la == CipherParserPRIVATE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserCONST {
		{
			p.SetState(138)
			p.Match(CipherParserCONST)
		}

	}
	{
		p.SetState(141)
		p.Match(CipherParserID)
	}
	{
		p.SetState(142)
		p.Match(CipherParserASSIGN)
	}
	{
		p.SetState(143)
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
	p.EnterRule(localctx, 30, CipherParserRULE_funcAssign)
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
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserPUBLIC || _la == CipherParserPRIVATE {
		{
			p.SetState(145)
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
		p.SetState(148)
		p.Match(CipherParserFUNC)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserOVERRIDE {
		{
			p.SetState(149)
			p.Match(CipherParserOVERRIDE)
		}

	}
	{
		p.SetState(152)
		p.Match(CipherParserID)
	}
	{
		p.SetState(153)
		p.Match(CipherParserLPAREN)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CipherParserID {
		{
			p.SetState(154)
			p.Params()
		}

	}
	{
		p.SetState(157)
		p.Match(CipherParserRPAREN)
	}
	{
		p.SetState(158)
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

func (s *GetAttributesContext) DOT() antlr.TerminalNode {
	return s.GetToken(CipherParserDOT, 0)
}

func (s *GetAttributesContext) ID() antlr.TerminalNode {
	return s.GetToken(CipherParserID, 0)
}

func (s *GetAttributesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, 0)
}

func (s *GetAttributesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, 0)
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
	p.EnterRule(localctx, 32, CipherParserRULE_getAttributes)
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
		p.SetState(160)
		p.Atom()
	}
	{
		p.SetState(161)
		p.Match(CipherParserDOT)
	}
	{
		p.SetState(162)
		p.Match(CipherParserID)
	}
	{
		p.SetState(163)
		p.Match(CipherParserLPAREN)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291471906) != 0 {
		{
			p.SetState(164)
			p.Args()
		}

	}
	{
		p.SetState(167)
		p.Match(CipherParserRPAREN)
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

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserLPAREN, 0)
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

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CipherParserRPAREN, 0)
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
	_startState := 34
	p.EnterRecursionRule(localctx, 34, CipherParserRULE_expr, _p)

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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(170)
			p.Call()
		}

	case 2:
		{
			p.SetState(171)
			p.Atom()
		}

	case 3:
		{
			p.SetState(172)
			p.Match(CipherParserLPAREN)
		}
		{
			p.SetState(173)
			p.expr(0)
		}
		{
			p.SetState(174)
			p.Match(CipherParserRPAREN)
		}

	case 4:
		{
			p.SetState(176)

			var _m = p.Match(CipherParserNOT)

			localctx.(*ExprContext).op = _m
		}
		{
			p.SetState(177)
			p.expr(5)
		}

	case 5:
		{
			p.SetState(178)
			p.GetAttributes()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(182)

					var _m = p.Match(CipherParserPREDONE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(183)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(185)

					var _m = p.Match(CipherParserPREDTWO)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(186)
					p.expr(4)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CipherParserRULE_expr)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(188)

					var _m = p.Match(CipherParserCOMPARATIVE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(189)
					p.expr(3)
				}

			}

		}
		p.SetState(194)
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

func (s *ArrayContext) LLIST() antlr.TerminalNode {
	return s.GetToken(CipherParserLLIST, 0)
}

func (s *ArrayContext) RLIST() antlr.TerminalNode {
	return s.GetToken(CipherParserRLIST, 0)
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
	p.EnterRule(localctx, 36, CipherParserRULE_array)
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
		p.SetState(195)
		p.Match(CipherParserLLIST)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291471906) != 0 {
		{
			p.SetState(196)
			p.Args()
		}

	}
	{
		p.SetState(199)
		p.Match(CipherParserRLIST)
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
	p.EnterRule(localctx, 38, CipherParserRULE_atom)

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

	p.SetState(208)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CipherParserLLIST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Array()
		}

	case CipherParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(CipherParserID)
		}

	case CipherParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.Match(CipherParserINT)
		}

	case CipherParserFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.Match(CipherParserFLOAT)
		}

	case CipherParserSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(205)
			p.Match(CipherParserSTRING)
		}

	case CipherParserNULL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(206)
			p.Match(CipherParserNULL)
		}

	case CipherParserBOOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(207)
			p.Match(CipherParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *CipherParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
