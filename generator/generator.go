package generator

import (
	"CompilerPlayground/lexer"
	"bufio"
	"os"
)

type Generator struct {
	Source string
	Name   string
	Tokens []lexer.LexToken
	writer *bufio.Writer
}

func (generator *Generator) Generate() {
	if len(generator.Tokens) == 0 {
		panic("no tokens is defined in the Generator")
	}

	f, err := os.Create(generator.Source)
	if err != nil {
		panic("File Creation error from the Generator")
	}

	generator.writer = bufio.NewWriter(f)
}

func (generator *Generator) generateNumber() {

}

func GetGenerator() *Generator {
	return &Generator{Source: "./as/"}
}
