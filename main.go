package main

import (
	"CompilerPlayground/compiler"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("path should be defined")
	}
	file := os.Args[1]

	cr := compiler.GetCompiler()
	cr.SetFile(file)
	cr.Compile()
}
