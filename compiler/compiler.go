package compiler

import (
	"CompilerPlayground/generator"
	"CompilerPlayground/lexer"
	"CompilerPlayground/parser"
	"os/exec"
	"path"
	"runtime"
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

	if runtime.GOOS == "windows" {
		panic("Compiler does not work on windows")
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

	mainObjName := "tmp-main.o"
	printObjName := "print.o"

	cmdAsMain := exec.Command("as", "-o", mainObjName, path.Join(gen.Source, gen.GetName()))
	cmdAsPrint := exec.Command("as", "-o", printObjName, path.Join(gen.Source, "print.s"))
	cmdLink := exec.Command("ld", "-o", "testingExec", mainObjName, printObjName, "-lc -dynamic-linker /lib64/ld-linux-x86-64.so.2")

	err = cmdAsPrint.Run()
	if err != nil {
		panic(err)
	}

	err = cmdAsMain.Run()
	if err != nil {
		panic(err)
	}
	err = cmdLink.Run()
	if err != nil {
		panic(err)
	}
}

func GetCompiler() *Compiler {
	return &Compiler{}
}
