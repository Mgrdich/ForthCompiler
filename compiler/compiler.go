package compiler

import (
	"CompilerPlayground/generator"
	"CompilerPlayground/lexer"
	"CompilerPlayground/parser"
	"fmt"
	"os"
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

	dir, _ := os.Getwd()

	mainObjName := path.Join(dir, "tmp-main.o")
	printObjName := path.Join(dir, "print.o")

	cmdAsMain := exec.Command(fmt.Sprintln("as", "-o", mainObjName, path.Join(dir, gen.Source, gen.GetName())))
	cmdAsPrint := exec.Command(fmt.Sprintln("as", "-o", printObjName, path.Join(dir, gen.Source, "print.s")))
	cmdLink := exec.Command(fmt.Sprintln("ld", "-o", path.Join(dir, "testingExec"), mainObjName, printObjName, "-lc -dynamic-linker /lib64/ld-linux-x86-64.so.2"))

	_, err = cmdAsPrint.Output()
	if err != nil {
		panic(err)
	}

	_, err = cmdAsMain.Output()
	if err != nil {
		panic(err)
	}
	_, err = cmdLink.Output()
	if err != nil {
		panic(err)
	}
}

func GetCompiler() *Compiler {
	return &Compiler{}
}
