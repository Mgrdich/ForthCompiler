package generator

import (
	"CompilerPlayground/lexer"
	"CompilerPlayground/token"
	"bufio"
	"fmt"
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

	for _, lexTok := range generator.Tokens {
		switch {
		case lexTok.Tok.IsNumber():
			err = generator.generateNumber(lexTok)
		case lexTok.Tok.IsSimpleOperator():
			err = generator.generateOperation(lexTok)
		case lexTok.Tok.IsKeywordOperator():
			err = generator.generateKeywordOperation(lexTok)
		case lexTok.Tok == token.DOT:
			err = generator.generatePopPrint()
		case lexTok.Tok == token.PRINT:
			err = generator.generatePrintStack()
		default:
			fmt.Println("we came here something is going wrong")
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

func (generator *Generator) generateNumber(lexToken lexer.LexToken) error {
	_, err := generator.writer.WriteString("pushq $" + lexToken.Lit + "\n")
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

func (generator *Generator) generateKeywordOperation(lexToken lexer.LexToken) error {
	var err error

	switch lexToken.Tok {
	case token.MIN:
		err = generator.generateKeywordOperationMin()
	case token.MAX:
		err = generator.generateKeywordOperationMax()
	case token.NEGATE:
		err = generator.generateKeywordOperationNegate()
	case token.ABS:
		err = generator.generateKeywordOperationAbs()
	case token.SWAP:
		err = generator.generateKeywordOperationSwap()
	case token.ROT:
		err = generator.generateKeywordOperationRot()
	case token.DROP:
		err = generator.generateKeywordOperationDrop()
	case token.NIP:
		err = generator.generateKeywordOperationNip()
	case token.TUCK:
		err = generator.generateKeywordOperationTuck()
	case token.OVER:
		err = generator.generateKeywordOperationOver()
	case token.ROLL:
		err = generator.generateKeywordOperationRoll()
	case token.PICK:
		err = generator.generateKeywordOperationPick()
	default:
		// this is not error case whether some parsing or code
		fmt.Println("Generation for this code is not written yet", lexToken)
	}

	return err
}

func (generator *Generator) generateKeywordOperationMin() error {
	return nil
}

func (generator *Generator) generateKeywordOperationMax() error {
	return nil
}

func (generator *Generator) generateKeywordOperationNegate() error {
	return nil
}

func (generator *Generator) generateKeywordOperationAbs() error {
	return nil
}

func (generator *Generator) generateKeywordOperationSwap() error {
	return nil
}

func (generator *Generator) generateKeywordOperationRot() error {
	return nil
}

func (generator *Generator) generateKeywordOperationDrop() error {
	return nil
}

func (generator *Generator) generateKeywordOperationNip() error {
	return nil
}

func (generator *Generator) generateKeywordOperationTuck() error {
	return nil
}

func (generator *Generator) generateKeywordOperationOver() error {
	return nil
}

func (generator *Generator) generateKeywordOperationRoll() error {
	return nil
}

func (generator *Generator) generateKeywordOperationPick() error {
	return nil
}

func (generator *Generator) generatePopPrint() error {
	return nil
}

func (generator *Generator) generatePrintStack() error {
	return nil
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
