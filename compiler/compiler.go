package compiler

import (
	"CompilerPlayground/generator"
	"CompilerPlayground/lexer"
	"CompilerPlayground/parser"
	"strings"
)

type Compiler struct {
	lexer     lexer.Lexer
	parser    parser.Parser
	directory string
}

func (compiler *Compiler) SetDirectory(str string) {
	if !strings.HasSuffix(str, ".mf") {
		panic("Cannot set a file with incorrect format")
	}

	compiler.directory = str
}

func (compiler *Compiler) Compile() {
	if len(compiler.directory) == 0 {
		panic("please define the directory")
	}

	lex := lexer.GetLexer(compiler.directory)
	lex.Tokenize()

	pars := parser.GetParser(lex.Tokens)
	err := pars.Parse() // TODO return error if needed

	if err != nil {
		panic(err)
	}

	gen := generator.GetGenerator()
	gen.Generate()

	// TODO linker job and system call after the assembly file creation
}

func GetCompiler() *Compiler {
	return &Compiler{}
}
