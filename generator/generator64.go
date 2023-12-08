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
	ASCIZ
	TEXT
	GLOBAL
	_START

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
	unaryArithmeticEnd

	binaryArithmeticBeg
	ADDQ
	SUBQ
	IMULQ
	IDIVQ
	XORQ
	LEAQ
	binaryArithmeticEnd
	arithmeticEnd

	// Data Movements
	MOVQ

	// Control transfer
	CMPQ
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

var AsEvalSyntax = map[AsSyntax]string{
	SECTION: ".section",
	RODATA:  ".rodata",
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

	CMPQ:  "cmpq",
	JMP:   "jmp",
	JE:    "JE",
	JNE:   "jne",
	JS:    "js",
	JNS:   "jns",
	JG:    "jg",
	JL:    "jl",
	PUSHQ: "pushq",
	POPQ:  "popq",
	CALL:  "call",
	RET:   "ret",
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
	pushq := AsEvalSyntax[PUSHQ]

	return fmt.Sprintln(pushq, "$"+numString)
}

func getPushQReg(syntax AsSyntax) string {
	checkRegister(syntax)

	pushq := AsEvalSyntax[PUSHQ]
	reg := AsEvalSyntax[syntax]

	return fmt.Sprintln(pushq, reg)
}

func getPopQReg(syntax AsSyntax) string {
	checkRegister(syntax)

	popq := AsEvalSyntax[POPQ]
	reg := AsEvalSyntax[syntax]

	return fmt.Sprintln(popq, reg)
}

func getUnaryArithmeticReg(operation AsSyntax, reg AsSyntax) string {
	if !isUnaryArithmetic(operation) {
		panic("operation should be of unary type")
	}

	checkRegister(reg)

	op := AsEvalSyntax[operation]
	r := AsEvalSyntax[reg]

	return fmt.Sprintln(op, r)
}

func getBinaryArithmeticRegs(operation AsSyntax, reg AsSyntax, anotherReg AsSyntax) string {
	if !isBinaryArithmetic(operation) {
		panic("operation should be of unary type")
	}

	checkRegister(reg)
	checkRegister(anotherReg)

	op := AsEvalSyntax[operation]
	r1 := AsEvalSyntax[reg]
	r2 := AsEvalSyntax[anotherReg]

	return fmt.Sprintln(op, r1, ",", r2)
}

func getBinaryArithmeticNumToReg(operation AsSyntax, number int, reg AsSyntax) string {
	if !isBinaryArithmetic(operation) {
		panic("operation should be of unary type")
	}

	checkRegister(reg)

	op := AsEvalSyntax[operation]
	r1 := AsEvalSyntax[reg]

	return fmt.Sprintln(op, "$"+strconv.Itoa(number), ",", r1)
}

func getLabel(label string) string {
	return label + ":\n"
}

func getAsciz(label string) string {
	return AsEvalSyntax[ASCIZ] + "\"" + label + "\"\n"
}

func getSectionROdata(syntax AsSyntax) string {
	return fmt.Sprintln(AsEvalSyntax[SECTION], AsEvalSyntax[RODATA])
}

func getSectionText(syntax AsSyntax) string {
	return fmt.Sprintln(AsEvalSyntax[SECTION], AsEvalSyntax[TEXT])
}

func getGlobalStart() string {
	return fmt.Sprintln(AsEvalSyntax[GLOBAL], AsEvalSyntax[_START])
}

func getStart() string {
	return AsEvalSyntax[_START] + ":\n"
}

func getNumber(number int) string {
	return "$" + strconv.Itoa(number)
}

func getMoveQRegToReg(reg AsSyntax, anotherReg AsSyntax) string {
	checkRegister(reg)
	checkRegister(anotherReg)

	movq := AsEvalSyntax[MOVQ]
	register := AsEvalSyntax[reg]
	anotherRegister := AsEvalSyntax[anotherReg]

	return fmt.Sprintln(movq, register, ",", anotherRegister)
}

func getMoveQDRefRegToReg(reg AsSyntax, anotherReg AsSyntax) string {
	checkRegister(reg)
	checkRegister(anotherReg)

	movq := AsEvalSyntax[MOVQ]
	register := AsEvalSyntax[reg]
	anotherRegister := AsEvalSyntax[anotherReg]

	return fmt.Sprintln(movq, "("+register+"),", anotherRegister)
}

func getMoveNumberToReg(number int, reg AsSyntax) string {
	checkRegister(reg)

	movq := AsEvalSyntax[MOVQ]
	register := AsEvalSyntax[reg]

	return fmt.Sprintln(movq, getNumber(number), register)
}

func getMoveVarToReg(strVar string, reg AsSyntax) string {
	movq := AsEvalSyntax[MOVQ]
	register := AsEvalSyntax[reg]

	return fmt.Sprintln(movq, "$"+strVar, register)
}

func getPrintSpace() string {
	return fmt.Sprintln(AsEvalSyntax[CALL], FN_PrintSpace)
}

func getPrintEOL() string {
	return fmt.Sprintln(AsEvalSyntax[CALL], FN_Printeol)
}

func getPrintReg(reg AsSyntax) string {
	checkRegister(reg)

	return fmt.Sprintln(getMoveQRegToReg(reg, RSI), "\n", AsEvalSyntax[CALL], FN_Print)
}

func getPrintNum(number int) string {
	return fmt.Sprintln(getMoveNumberToReg(number, RSI), "\n", AsEvalSyntax[CALL], FN_Print)
}

func getPrintlnReg(reg AsSyntax) string {
	checkRegister(reg)

	return fmt.Sprintln(getMoveQRegToReg(reg, RSI), "\n", AsEvalSyntax[CALL], FN_Println)
}

func getPrintlnNum(number int) string {
	return fmt.Sprintln(getMoveNumberToReg(number, RSI), "\n", AsEvalSyntax[CALL], FN_Println)
}

func getPrintW(strVar string) string {
	return fmt.Sprintln(getMoveVarToReg(strVar, RSI), "\n", AsEvalSyntax[CALL], FN_Printw)
}

func getPrintWln(strVar string) string {
	return fmt.Sprintln(getMoveVarToReg(strVar, RSI), "\n", AsEvalSyntax[CALL], FN_Printwln)
}

func getJmp(jmpType AsSyntax, label string) string {
	if !isJmpControlTransfer(jmpType) {
		panic("jmpType should be of type Jmp")
	}

	return fmt.Sprintln(jmpType, label)
}
