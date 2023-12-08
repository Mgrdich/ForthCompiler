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

func (generator *Generator) writeString(str string) error {
	_, err := generator.writer.WriteString(str)
	return err
}

func (generator *Generator) start() {
	var err error
	var stringBuilder strings.Builder

	stringBuilder.WriteString(".section .rodata\n")
	stringBuilder.WriteString("okWord:\n")
	stringBuilder.WriteString(".asciz \"ok\" \n")
	stringBuilder.WriteString("p1:\n")
	stringBuilder.WriteString(".asciz \"<\"\n")
	stringBuilder.WriteString("p2:\n")
	stringBuilder.WriteString(".asciz \">\"\n")

	stringBuilder.WriteString(".section .text\n")
	stringBuilder.WriteString(".global _start\n")
	stringBuilder.WriteString("_start:\n")

	// stack keepers
	stringBuilder.WriteString("pushq %rbp\n")
	stringBuilder.WriteString("movq %rsp , %rbp\n")

	err = generator.writeString(stringBuilder.String())

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

	// Stack keepers
	stringBuilder.WriteString("movq %rbp , %rsp\n")
	stringBuilder.WriteString("popq %rbp\n")

	stringBuilder.WriteString("exit:\n")
	stringBuilder.WriteString("mov $60, %rax\n")
	stringBuilder.WriteString("xor %rdi, %rdi\n")
	stringBuilder.WriteString("syscall\n")

	err = generator.writeString(stringBuilder.String())
	if err != nil {
		panic(err)
	}

	err = generator.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (generator *Generator) printOk() error {
	return generator.writeString("movq $okWord, %rsi\n" + "call printwln\n")
}

func (generator *Generator) generateNumber(lexToken lexer.LexToken) error {
	return generator.writeString("pushq $" + lexToken.Lit + "\n")
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
	stringBuilder.WriteString(" %rax, %rbx\n")
	stringBuilder.WriteString("pushq %rbx\n")

	return generator.writeString(stringBuilder.String())
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
	var stringsBuilder strings.Builder
	stringsBuilder.WriteString("popq %rax\n")
	stringsBuilder.WriteString("popq %rbx\n")

	stringsBuilder.WriteString("cmp %rbx, %rax\n") // %rax < %rbx
	stringsBuilder.WriteString("JL smaller\n")
	stringsBuilder.WriteString("movq %rbx, %rcx\n")
	stringsBuilder.WriteString("jmp doneMin\n")

	stringsBuilder.WriteString("smaller:\n")
	stringsBuilder.WriteString("movq %rax, %rcx\n")

	stringsBuilder.WriteString("doneMin:\n")
	stringsBuilder.WriteString("pushq %rbx\n")
	stringsBuilder.WriteString("pushq %rax\n")
	stringsBuilder.WriteString("pushq %rcx\n")

	return generator.writeString(stringsBuilder.String())
}

func (generator *Generator) generateKeywordOperationMax() error {
	var stringsBuilder strings.Builder
	stringsBuilder.WriteString("popq %rax\n")
	stringsBuilder.WriteString("popq %rbx\n")

	stringsBuilder.WriteString("cmp %rbx, %rax\n") // %rax < %rbx
	stringsBuilder.WriteString("JG greater\n")
	stringsBuilder.WriteString("movq %rbx, %rcx\n")
	stringsBuilder.WriteString("jmp doneMax\n")

	stringsBuilder.WriteString("greater:\n")
	stringsBuilder.WriteString("movq %rax, %rcx\n")

	stringsBuilder.WriteString("doneMax:\n")
	stringsBuilder.WriteString("pushq %rbx\n")
	stringsBuilder.WriteString("pushq %rax\n")
	stringsBuilder.WriteString("pushq %rcx\n")

	return generator.writeString(stringsBuilder.String())
}

func (generator *Generator) generateKeywordOperationNegate() error {
	var stringBuilder strings.Builder
	stringBuilder.WriteString("popq %rax\n")
	stringBuilder.WriteString("negq %rax\n")
	stringBuilder.WriteString("pushq %rax\n")

	return generator.writeString(stringBuilder.String())
}

func (generator *Generator) generateKeywordOperationAbs() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString("popq %rax\n")
	stringBuilder.WriteString("movq %rax, %rcx\n")
	stringBuilder.WriteString("testq %rax\n")
	stringBuilder.WriteString("JL negative\n")
	stringBuilder.WriteString("jmp doneAbs\n")
	stringBuilder.WriteString("negative:\n")
	stringBuilder.WriteString("negq %rcx\n")
	stringBuilder.WriteString("doneAbs:\n")
	stringBuilder.WriteString("pushq %rax\n")
	stringBuilder.WriteString("pushq %rcx\n")

	return generator.writeString(stringBuilder.String())
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
	var stringBuilder strings.Builder

	stringBuilder.WriteString("popq %rsi\n")
	stringBuilder.WriteString("call print\n")

	stringBuilder.WriteString("call printSpace\n")

	err := generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	return generator.printOk()
}

func (generator *Generator) generatePrintStack() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString("movq $8 , %rcx\n")
	stringBuilder.WriteString("xor %rdx, %rdx\n")
	stringBuilder.WriteString("movq %rbp , %rax\n")
	stringBuilder.WriteString("subq %rsp , %rax\n")
	stringBuilder.WriteString("idiv %rcx\n")
	stringBuilder.WriteString("movq %rax , %r12\n")
	stringBuilder.WriteString("movq %rbp , %r14\n")
	stringBuilder.WriteString("subq $8 , %r14\n")

	err := generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	stringBuilder.Reset()

	// Print the stack number
	stringBuilder.WriteString("movq $p1 , %rsi\n")
	stringBuilder.WriteString("call printw\n")

	stringBuilder.WriteString("movq %r12 , %rsi\n")
	stringBuilder.WriteString("call print\n")

	stringBuilder.WriteString("movq $p2 , %rsi\n")
	stringBuilder.WriteString("call printw\n")
	stringBuilder.WriteString("call printSpace\n")

	err = generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	stringBuilder.Reset()

	stringBuilder.WriteString("loop:\n")
	stringBuilder.WriteString("movq (%r14), %rsi\n")
	stringBuilder.WriteString("call print\n")
	stringBuilder.WriteString("call printSpace\n")
	stringBuilder.WriteString("subq $8, %r14\n")
	stringBuilder.WriteString("subq $1 ,%r12\n")
	stringBuilder.WriteString("cmp $0 , %r12\n")
	stringBuilder.WriteString("jne loop\n")

	err = generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	return generator.printOk()
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
