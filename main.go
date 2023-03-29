package main

import (
	core "Cipher/Core"
	cipher "Cipher/Core/parser"
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var Operators = map[string]string{
	"+":   "Add",
	"-":   "Sub",
	"*":   "Mul",
	"/":   "Div",
	"^":   "Pow",
	"%":   "Mod",
	"==":  "Eq",
	"!=":  "Neq",
	">":   "Gt",
	"<":   "Lt",
	"<=":  "Lte",
	">=":  "Gte",
	"&&":  "And",
	"and": "And",
	"or":  "Or",
	"||":  "Or",
	"!":   "Not",
	"not": "Not",
	"~=":  "Neq",
}

func main() {
	f, _ := os.Create("main.profile")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	runtime.GC()
	err := pprof.StartCPUProfile(f)
	if err != nil {
		return
	}

	st := time.Now()

	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := cipher.NewCipherLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := cipher.NewCipherParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&core.ErrorListener{})
	p.BuildParseTrees = true

	tree := p.Parse()

	v := core.Visitor{Env: core.NewEnv(), Operators: Operators}
	v.VisitParse(tree.(*cipher.ParseContext))

	elapsed := time.Since(st)
	fmt.Printf("Time to execute %v\n", elapsed)

	reader := bufio.NewReader(os.Stdin)
	_, err = reader.ReadString('\n')
	if err != nil {
		return
	}

	pprof.StopCPUProfile()
}
