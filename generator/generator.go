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

	stringBuilder.WriteString(getSectionROdata())
	stringBuilder.WriteString(getLabel(VAR_okWord))
	stringBuilder.WriteString(getAsciz("ok"))
	stringBuilder.WriteString(getLabel(VAR_SymbolGreater))
	stringBuilder.WriteString(getAsciz("<"))
	stringBuilder.WriteString(getLabel(VAR_Symbollees))
	stringBuilder.WriteString(getAsciz(">"))

	stringBuilder.WriteString(getSectionText())
	stringBuilder.WriteString(getGlobalStart())
	stringBuilder.WriteString(getStart())

	// stack keepers
	stringBuilder.WriteString(getPushQ(RBP))
	stringBuilder.WriteString(getMovQ(RSP, RBP))

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
			//err = generator.generatePrintStack()
			err = generator.generateTopPrint()
		default:
			fmt.Println("we came here something is going wrong")
		}

		if err != nil {
			panic(err)
		}
	}

	stringBuilder.Reset()

	// Stack keepers
	stringBuilder.WriteString(getMovQ(RBP, RSP))
	stringBuilder.WriteString(getPopQ(RBP))

	stringBuilder.WriteString(getLabel("exit"))
	stringBuilder.WriteString(getMovQNumber(60, RAX))
	stringBuilder.WriteString(getXorQ(RDI, RDI))
	stringBuilder.WriteString(getSysCall())

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
	return generator.writeString(getPushQNum(lexToken.Lit))
}

func (generator *Generator) generateOperation(lexToken lexer.LexToken) error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getPopQ(RBX))

	switch lexToken.Tok {
	case token.ADD:
		stringBuilder.WriteString(getAddQ(RAX, RBX))
	case token.SUB:
		stringBuilder.WriteString(getSubQ(RAX, RBX))
	case token.MUL:
		stringBuilder.WriteString(getIMulQ(RAX, RBX))
	case token.QUO:
		panic("Currently we are not supporting division")
	default:
		panic("Synchronize with the token file something went wrong")
	}

	stringBuilder.WriteString(getPushQ(RBX))

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
	lSmaller := getUniqueLabel()
	lDoneMin := getUniqueLabel()

	var stringsBuilder strings.Builder
	stringsBuilder.WriteString(getPopQ(RAX))
	stringsBuilder.WriteString(getPopQ(RBX))

	stringsBuilder.WriteString(getCmpQ(RBX, RAX)) // %rax < %rbx
	stringsBuilder.WriteString(getJmp(JL, lSmaller))
	stringsBuilder.WriteString(getMovQ(RBX, RCX))
	stringsBuilder.WriteString(getJmp(JMP, lDoneMin))

	stringsBuilder.WriteString(getLabel(lSmaller))
	stringsBuilder.WriteString(getMovQ(RAX, RCX))

	stringsBuilder.WriteString(getLabel(lDoneMin))
	stringsBuilder.WriteString(getPushQ(RBX))
	stringsBuilder.WriteString(getPushQ(RAX))
	stringsBuilder.WriteString(getPushQ(RCX))

	return generator.writeString(stringsBuilder.String())
}

func (generator *Generator) generateKeywordOperationMax() error {
	lGreater := getUniqueLabel()
	lDoneMax := getUniqueLabel()

	var stringsBuilder strings.Builder
	stringsBuilder.WriteString(getPopQ(RAX))
	stringsBuilder.WriteString(getPopQ(RBX))

	stringsBuilder.WriteString(getCmpQ(RBX, RAX)) // %rax < %rbx
	stringsBuilder.WriteString(getJmp(JG, lGreater))
	stringsBuilder.WriteString(getMovQ(RBX, RCX))
	stringsBuilder.WriteString(getJmp(JMP, lDoneMax))

	stringsBuilder.WriteString(getLabel(lGreater))
	stringsBuilder.WriteString(getMovQ(RAX, RCX))

	stringsBuilder.WriteString(getLabel(lDoneMax))
	stringsBuilder.WriteString(getPushQ(RBX))
	stringsBuilder.WriteString(getPushQ(RAX))
	stringsBuilder.WriteString(getPushQ(RCX))

	return generator.writeString(stringsBuilder.String())
}

func (generator *Generator) generateKeywordOperationNegate() error {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getNegQ(RAX))
	stringBuilder.WriteString(getPushQ(RAX))

	return generator.writeString(stringBuilder.String())
}

func (generator *Generator) generateKeywordOperationAbs() error {
	lNegative := getUniqueLabel()
	lDoneAbs := getUniqueLabel()
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getMovQ(RAX, RCX))
	stringBuilder.WriteString(getTestQ(RAX, RAX))
	stringBuilder.WriteString(getJmp(JL, lNegative))
	stringBuilder.WriteString(getJmp(JMP, lDoneAbs))
	stringBuilder.WriteString(getLabel(lNegative))
	stringBuilder.WriteString(getNegQ(RCX))
	stringBuilder.WriteString(getLabel(lDoneAbs))
	stringBuilder.WriteString(getPushQ(RAX))
	stringBuilder.WriteString(getPushQ(RCX))

	return generator.writeString(stringBuilder.String())
}

func (generator *Generator) generateKeywordOperationSwap() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getPopQ(RBX))
	stringBuilder.WriteString(getPushQ(RAX))
	stringBuilder.WriteString(getPopQ(RBX))

	return generator.writeString(stringBuilder.String())
}

func (generator *Generator) generateKeywordOperationRot() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getPopQ(RBX))
	stringBuilder.WriteString(getPopQ(RCX))
	stringBuilder.WriteString(getPushQ(RAX))
	stringBuilder.WriteString(getPushQ(RBX))
	stringBuilder.WriteString(getPushQ(RCX))

	return generator.writeString(stringBuilder.String())
}

func (generator *Generator) generateKeywordOperationDrop() error {
	return generator.writeString(getPopQ(RAX))
}

func (generator *Generator) generateKeywordOperationNip() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getPopQ(RBX))
	stringBuilder.WriteString(getPushQ(RAX))

	return generator.writeString(stringBuilder.String())
}

func (generator *Generator) generateKeywordOperationTuck() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getPopQ(RBX))
	stringBuilder.WriteString(getPushQ(RAX))
	stringBuilder.WriteString(getPushQ(RBX))
	stringBuilder.WriteString(getPushQ(RAX))

	return generator.writeString(stringBuilder.String())
}

func (generator *Generator) generateKeywordOperationOver() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RAX))
	stringBuilder.WriteString(getPopQ(RBX))
	stringBuilder.WriteString(getPushQ(RBX))
	stringBuilder.WriteString(getPushQ(RAX))
	stringBuilder.WriteString(getPushQ(RBX))

	return generator.writeString(stringBuilder.String())
}

// TODO in the Future
func (generator *Generator) generateKeywordOperationRoll() error {
	return nil
}

// TODO in the Future
func (generator *Generator) generateKeywordOperationPick() error {
	return nil
}

func (generator *Generator) generatePopPrint() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RSI))
	stringBuilder.WriteString(getPrint())

	stringBuilder.WriteString(getPrintSpace())

	err := generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	return generator.printOk()
}

func (generator *Generator) generateTopPrint() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getPopQ(RSI))
	stringBuilder.WriteString(getMovQ(RSI, R15))
	stringBuilder.WriteString(getPrint())

	stringBuilder.WriteString(getPrintSpace())
	stringBuilder.WriteString(getPushQ(R15))

	err := generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	return generator.printOk()
}

func (generator *Generator) generatePrintStack() error {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(getMovQNumber(8, RCX))
	stringBuilder.WriteString(getXorQ(RDX, RDX))
	stringBuilder.WriteString(getMovQ(RBP, RAX))
	stringBuilder.WriteString(getSubQ(RSP, RAX))
	stringBuilder.WriteString(getIDivQ(RCX))
	stringBuilder.WriteString(getMovQ(RAX, R12))
	stringBuilder.WriteString(getMovQ(RBP, R14))
	stringBuilder.WriteString(getSubQNum(8, R14))

	err := generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	stringBuilder.Reset()

	// Print the stack number
	stringBuilder.WriteString(getMovQVar(VAR_SymbolGreater, RSI))
	stringBuilder.WriteString(getPrintW())

	stringBuilder.WriteString(getMovQ(R12, RSI))
	stringBuilder.WriteString(getPrint())

	stringBuilder.WriteString(getMovQVar(VAR_Symbollees, RSI))
	stringBuilder.WriteString(getPrintW())
	stringBuilder.WriteString(getPrintSpace())

	err = generator.writeString(stringBuilder.String())
	if err != nil {
		return err
	}

	stringBuilder.Reset()

	lLoop := getUniqueLabel()
	stringBuilder.WriteString(getLabel(lLoop))
	stringBuilder.WriteString(getMovQDRefReg(R14, RSI))
	stringBuilder.WriteString(getPrint())
	stringBuilder.WriteString(getPrintSpace())
	stringBuilder.WriteString(getSubQNum(8, R14))
	stringBuilder.WriteString(getSubQNum(1, R12))
	stringBuilder.WriteString(getCmpQNum(0, R12))
	stringBuilder.WriteString(getJmp(JNE, lLoop))

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
