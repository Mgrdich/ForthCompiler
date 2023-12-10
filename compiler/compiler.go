package compiler

import (
	"CompilerPlayground/generator"
	"CompilerPlayground/lexer"
	"CompilerPlayground/parser"
	"fmt"
	"os/exec"
	"path"
)

type Compiler struct {
	lexer     lexer.Lexer
	parser    parser.Parser
	directory string
}

func (compiler *Compiler) SetDirectory(dir string) {
	if path.Ext(dir) != ".mf" {
		panic("Cannot set a file with incorrect format")
	}

	compiler.directory = dir
}

func (compiler *Compiler) Compile() {
	if len(compiler.directory) == 0 {
		panic("define the directory")
	}

	lex := lexer.GetLexer(compiler.directory)
	lex.Tokenize()

	pars := parser.GetParser()
	pars.Tokens = lex.Tokens
	err := pars.Parse()

	if err != nil {
		panic(err)
	}

	gen := generator.GetGenerator()
	gen.Tokens = lex.Tokens
	gen.Generate()

	_, err = exec.Command("/bin/sh", "./as/compileLinkAssembly.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func GetCompiler() *Compiler {
	return &Compiler{}
}
