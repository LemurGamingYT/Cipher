package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
)

type ErrorListener struct {
	antlr.DefaultErrorListener
}

func (d *ErrorListener) SyntaxError(_ antlr.Recognizer, offendingSymbol interface{}, line, column int,
	_ string, _ antlr.RecognitionException) {
	ReportError("Syntax", fmt.Sprintf("Unexpected '%s' at ln %d, column %d\n",
		offendingSymbol.(*antlr.CommonToken).GetText(), line, column))
}

func ReportError(errorType string, message string) {
	fmt.Printf("%sError: %s\n", errorType, message)
	os.Exit(1)
}
