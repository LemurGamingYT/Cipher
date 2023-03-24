package main

import (
	core "Cipher/Core"
	cipher "Cipher/Core/parser"
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"time"
)

func main() {
	/*
		f, err := os.Create("profile.prof")
		if err != nil {
			log.Fatal(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()*/

	st := time.Now()

	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := cipher.NewCipherLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := cipher.NewCipherParser(stream)
	p.BuildParseTrees = true
	/*
		interpreter := p.Interpreter
		if interpreter != nil {
			interpreter.SetPredictionMode(antlr.PredictionModeLL)
			interpreter.SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
		}*/

	tree := p.Parse()
	// println(tree.ToStringTree([]string{}, p))

	v := core.Visitor{Env: core.NewEnv()}
	v.VisitParse(tree.(*cipher.ParseContext))

	elapsed := time.Since(st)
	fmt.Printf("Time to execute %v\n", elapsed)

	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		return
	}
}
