package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
)

type ErrorListener struct {
	antlr.DefaultErrorListener
}

func (d *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int,
	msg string, e antlr.RecognitionException) {
	ReportError("Syntax", fmt.Sprintf("'%s' at ln %d, column %d\n",
		offendingSymbol.(*antlr.CommonToken).GetText(), line, column))
}

func ReportError(errorType string, message string) {
	fmt.Printf("%sError: %s\n", errorType, message)
	os.Exit(1)
}
