package main

import (
	"CompilerPlayground/compiler"
)

func main() {
	//if len(os.Args) == 1 {
	//	panic("path should be defined")
	//}
	//file := os.Args[1]

	cr := compiler.GetCompiler()
	cr.SetDirectory("./test1.mf")
	cr.Compile()
}
