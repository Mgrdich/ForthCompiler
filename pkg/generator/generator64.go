package generator

import (
	"fmt"
	"strconv"
)

type AsSyntax int

// only the used ones are here and all of the registers
const (
	SECTION AsSyntax = iota
	RODATA
	DATA
	QUAD
	ASCIZ
	TEXT
	GLOBAL
	_START
	SYSCALL

	// registers
	regBegin
	RAX
	RBX
	RCX
	RDX
	RSI
	RDI
	RBP
	RSP
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	regEnd

	// Arithmetic Operations
	arithmeticBeg
	unaryArithmeticBeg
	INCQ
	DECQ
	NEGQ
	IDIVQ
	unaryArithmeticEnd

	binaryArithmeticBeg
	ADDQ
	SUBQ
	IMULQ
	XORQ
	LEAQ
	binaryArithmeticEnd
	arithmeticEnd

	// Data Movements
	MOVQ

	// Control transfer
	CMPQ
	TESTQ
	controlJmpBeg
	JMP
	JE
	JNE
	JS
	JNS
	JG
	JL
	controlJmpEnd
	PUSHQ
	POPQ
	CALL
	RET
)

var counter = 0

var AsEvalSyntax = map[AsSyntax]string{
	SECTION: ".section",
	RODATA:  ".rodata",
	DATA:    ".data",
	QUAD:    ".quad",
	ASCIZ:   ".asciz",
	TEXT:    ".text",
	GLOBAL:  ".global",
	_START:  "_start",

	RAX: "%rax",
	RBX: "%rbx",
	RCX: "%rcx",
	RDX: "%rdx",
	RSI: "%rsi",
	RDI: "%rdi",
	RBP: "%rbp",
	RSP: "%rsp",
	R8:  "%r8",
	R9:  "%r9",
	R10: "%r10",
	R11: "%r11",
	R12: "%r12",
	R13: "%r13",
	R14: "%r14",
	R15: "%r15",

	INCQ:  "incq",
	DECQ:  "decq",
	NEGQ:  "negq",
	LEAQ:  "leaq",
	ADDQ:  "addq",
	SUBQ:  "subq",
	IMULQ: "imulq",
	IDIVQ: "idivq",
	XORQ:  "xorq",

	MOVQ: "movq",

	CMPQ:    "cmpq",
	TESTQ:   "testq",
	JMP:     "jmp",
	JE:      "JE",
	JNE:     "jne",
	JS:      "js",
	JNS:     "jns",
	JG:      "jg",
	JL:      "jl",
	PUSHQ:   "pushq",
	POPQ:    "popq",
	CALL:    "call",
	RET:     "ret",
	SYSCALL: "syscall",
}

func getUniqueLabel() string {
	counter++
	return "l" + strconv.Itoa(counter)
}

func getValue(str string) string {
	return "$" + str
}

func getNumber(number int) string {
	return getValue(strconv.Itoa(number))
}

func isRegister(syntax AsSyntax) bool {
	return syntax > regBegin && syntax < regEnd
}

func checkRegister(syntax AsSyntax) {
	if !isRegister(syntax) {
		panic("param should be a register")
	}
}

func isArithmetic(syntax AsSyntax) bool {
	return syntax > arithmeticBeg && syntax < arithmeticEnd
}

func isUnaryArithmetic(syntax AsSyntax) bool {
	return isArithmetic(syntax) && syntax > unaryArithmeticBeg && syntax < unaryArithmeticEnd
}

func isBinaryArithmetic(syntax AsSyntax) bool {
	return isArithmetic(syntax) && syntax > binaryArithmeticBeg && syntax < binaryArithmeticEnd
}

func isJmpControlTransfer(syntax AsSyntax) bool {
	return syntax > controlJmpBeg && syntax < controlJmpEnd
}

func getPushQNum(numString string) string {
	return fmt.Sprintln(AsEvalSyntax[PUSHQ], getValue(numString))
}

func getPushQ(syntax AsSyntax) string {
	checkRegister(syntax)

	return fmt.Sprintln(AsEvalSyntax[PUSHQ], AsEvalSyntax[syntax])
}

func getPopQ(syntax AsSyntax) string {
	checkRegister(syntax)

	return fmt.Sprintln(AsEvalSyntax[POPQ], AsEvalSyntax[syntax])
}

func getUnaryArithmeticReg(operation AsSyntax, reg AsSyntax) string {
	if !isUnaryArithmetic(operation) {
		panic("operation should be of unary type")
	}
	checkRegister(reg)

	return fmt.Sprintln(AsEvalSyntax[operation], AsEvalSyntax[reg])
}

func getBinaryArithmeticRegs(operation AsSyntax, reg AsSyntax, anotherReg AsSyntax) string {
	if !isBinaryArithmetic(operation) {
		panic("operation should be of unary type")
	}
	checkRegister(reg)
	checkRegister(anotherReg)

	return fmt.Sprintln(AsEvalSyntax[operation], AsEvalSyntax[reg], ",", AsEvalSyntax[anotherReg])
}

func getBinaryArithmeticNumToReg(operation AsSyntax, number int, reg AsSyntax) string {
	if !isBinaryArithmetic(operation) {
		panic("operation should be of unary type")
	}
	checkRegister(reg)

	return fmt.Sprintln(AsEvalSyntax[operation], getNumber(number), ",", AsEvalSyntax[reg])
}

func getAddQ(reg AsSyntax, anotherReg AsSyntax) string {
	return getBinaryArithmeticRegs(ADDQ, reg, anotherReg)
}

func getSubQ(reg AsSyntax, anotherReg AsSyntax) string {
	return getBinaryArithmeticRegs(SUBQ, reg, anotherReg)
}

func getSubQNum(number int, reg AsSyntax) string {
	return getBinaryArithmeticNumToReg(SUBQ, number, reg)
}

func getIMulQ(reg AsSyntax, anotherReg AsSyntax) string {
	return getBinaryArithmeticRegs(IMULQ, reg, anotherReg)
}

func getIDivQ(reg AsSyntax) string {
	return getUnaryArithmeticReg(IDIVQ, reg)
}

func getXorQ(reg AsSyntax, anotherReg AsSyntax) string {
	return getBinaryArithmeticRegs(XORQ, reg, anotherReg)
}

func getNegQ(reg AsSyntax) string {
	return getUnaryArithmeticReg(NEGQ, reg)
}

func getLabel(label string) string {
	return label + ":\n"
}

func getAsciz(label string) string {
	return fmt.Sprintln(AsEvalSyntax[ASCIZ], "\""+label+"\"")
}

func getSectionROdata() string {
	return fmt.Sprintln(AsEvalSyntax[SECTION], AsEvalSyntax[RODATA])
}

func getSectionText() string {
	return fmt.Sprintln(AsEvalSyntax[SECTION], AsEvalSyntax[TEXT])
}

func getGlobalStart() string {
	return fmt.Sprintln(AsEvalSyntax[GLOBAL], AsEvalSyntax[_START])
}

func getStart() string {
	return AsEvalSyntax[_START] + ":\n"
}

func getMovQ(reg AsSyntax, anotherReg AsSyntax) string {
	checkRegister(reg)
	checkRegister(anotherReg)

	return fmt.Sprintln(AsEvalSyntax[MOVQ], AsEvalSyntax[reg], ",", AsEvalSyntax[anotherReg])
}

func getMovQDRefReg(reg AsSyntax, anotherReg AsSyntax) string {
	checkRegister(reg)
	checkRegister(anotherReg)

	return fmt.Sprintln(AsEvalSyntax[MOVQ], "("+AsEvalSyntax[reg]+")", ",", AsEvalSyntax[anotherReg])
}

func getMovQNumber(number int, reg AsSyntax) string {
	checkRegister(reg)

	return fmt.Sprintln(AsEvalSyntax[MOVQ], getNumber(number), ",", AsEvalSyntax[reg])
}

func getMovQVar(strVar string, reg AsSyntax) string {
	checkRegister(reg)

	return fmt.Sprintln(AsEvalSyntax[MOVQ], getValue(strVar), ",", AsEvalSyntax[reg])
}

func getPrintSpace() string {
	return fmt.Sprintln(AsEvalSyntax[CALL], FN_PrintSpace)
}

func getPrintEOL() string {
	return fmt.Sprintln(AsEvalSyntax[CALL], FN_Printeol)
}

func getPrint() string {

	return fmt.Sprintln(AsEvalSyntax[CALL], FN_Print)
}

func getPrintln() string {
	return fmt.Sprintln(AsEvalSyntax[CALL], FN_Println)
}

func getPrintW() string {
	return fmt.Sprintln(AsEvalSyntax[CALL], FN_Printw)
}

func getPrintWln() string {
	return fmt.Sprintln(AsEvalSyntax[CALL], FN_Printwln)
}

func getJmp(jmpType AsSyntax, label string) string {
	if !isJmpControlTransfer(jmpType) {
		panic("jmpType should be of type Jmp")
	}

	return fmt.Sprintln(AsEvalSyntax[jmpType], label)
}

func getCmpQ(reg AsSyntax, anotherReg AsSyntax) string {
	checkRegister(reg)
	checkRegister(anotherReg)

	return fmt.Sprintln(AsEvalSyntax[CMPQ], AsEvalSyntax[reg], ",", AsEvalSyntax[anotherReg])
}

func getCmpQNum(number int, reg AsSyntax) string {
	checkRegister(reg)

	return fmt.Sprintln(AsEvalSyntax[CMPQ], getNumber(number), ",", AsEvalSyntax[reg])
}

func getTestQ(reg AsSyntax, anotherReg AsSyntax) string {
	checkRegister(reg)
	checkRegister(anotherReg)

	return fmt.Sprintln(AsEvalSyntax[TESTQ], AsEvalSyntax[reg], ",", AsEvalSyntax[anotherReg])
}

func getSysCall() string {
	return fmt.Sprint(AsEvalSyntax[SYSCALL])
}
