package compiler

import (
	"ForthCompiler/pkg/generator"
	"ForthCompiler/pkg/lexer"
	"ForthCompiler/pkg/parser"
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

type Compiler struct {
	lexer    lexer.Lexer
	parser   parser.Parser
	filePath string
}

func (compiler *Compiler) SetFile(dir string) {
	if path.Ext(dir) != ".mf" {
		panic("Cannot set a file with incorrect format")
	}

	compiler.filePath = dir
}

func (compiler *Compiler) assemblyCompilationAndLinking(gen *generator.Generator) {

	mainObjName := "tmp-main.o"
	printObjName := "print.o"
	fileName := path.Base(compiler.filePath)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]

	cmdAsMain := exec.Command("as", "-o", mainObjName, path.Join(gen.Source, gen.GetName()))
	cmdAsPrint := exec.Command("as", "-o", printObjName, path.Join(gen.Source, "print.s"))
	cmdLink := exec.Command("ld", "-o", fileNameWithoutExt, mainObjName, printObjName, "-lc", "-dynamic-linker", "/lib64/ld-linux-x86-64.so.2")

	err := cmdAsPrint.Run()
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

func (compiler *Compiler) Compile() {
	if len(compiler.filePath) == 0 {
		panic("define the File path")
	}

	lex := lexer.GetLexer(compiler.filePath)
	lex.Tokenize()
	fmt.Println("Tokenized Successfully")

	pars := parser.GetParser()
	pars.Tokens = lex.Tokens
	err := pars.Parse()

	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed Successfully")

	gen := generator.GetGenerator()
	gen.Tokens = lex.Tokens
	gen.Generate()

	if runtime.GOOS != "linux" {
		panic("Compiler does not work on this operating system")
	}

	if runtime.GOARCH != "amd64" {
		panic("Computer architecture is not currently supported")
	}
	compiler.assemblyCompilationAndLinking(gen)
}

func GetCompiler() *Compiler {
	return &Compiler{}
}
