package generator

import (
	"CompilerPlayground/lexer"
	"CompilerPlayground/token"
	"bufio"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Generator struct {
	Source string
	name   string
	Tokens []lexer.LexToken
	writer *bufio.Writer
}

func (generator *Generator) Generate() {
	if len(generator.Tokens) == 0 {
		panic("no tokens is defined in the Generator")
	}

	f, err := os.Create(filepath.Join(generator.Source, generator.name))
	if err != nil {
		panic(err)
	}

	generator.writer = bufio.NewWriter(f)
	generator.start()
}

func (generator *Generator) start() {
	var err error
	var stringBuilder strings.Builder

	stringBuilder.WriteString(".global _start\n")
	stringBuilder.WriteString("_start:\n")

	_, err = generator.writer.WriteString(stringBuilder.String())

	if err != nil {
		panic(err)
	}

	for _, tok := range generator.Tokens {
		switch {
		case tok.Tok.IsNumber():
			err = generator.generateNumber(tok.Lit)
		case tok.Tok.IsSimpleOperator():
			err = generator.generateOperation(tok)
		case tok.Tok.IsKeywordOperator():

		}

		if err != nil {
			panic(err)
		}
	}

	stringBuilder.Reset()
	stringBuilder.WriteString("mov $60, %rax\n")
	stringBuilder.WriteString("syscall\n")

	_, err = generator.writer.WriteString(stringBuilder.String())
	if err != nil {
		panic(err)
	}

	err = generator.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (generator *Generator) generateNumber(numberStr string) error {
	_, err := generator.writer.WriteString("pushq $" + numberStr + "\n")
	return err
}

func (generator *Generator) generateOperation(lexToken lexer.LexToken) error {
	operation := ""

	switch lexToken.Tok {
	case token.ADD:
		operation = "addq"
	case token.SUB:
		operation = "subq"
	case token.MUL:
		operation = "imulq"
	case token.QUO:
		panic("Currently we are not supporting division")
	default:
		panic("Synchronize with the token file something went wrong")
	}

	var stringBuilder strings.Builder

	stringBuilder.WriteString("popq %rax\n")
	stringBuilder.WriteString("popq %rbx\n")
	stringBuilder.WriteString(operation)
	stringBuilder.WriteString(" %rbx, %rax\n")
	stringBuilder.WriteString("pushq %rax\n\n")

	_, err := generator.writer.WriteString(stringBuilder.String())

	return err
}

func (generator *Generator) setName(name string) {
	if path.Ext(name) != ".s" {
		panic("can not set names without s extension")
	}
	generator.name = "tmp_" + name
}

func GetGenerator() *Generator {
	return &Generator{Source: "as", name: "tmp_source.s"}
}
