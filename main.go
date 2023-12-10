package main

import (
	"CompilerPlayground/compiler"
	"os"
)

func main() {
	file := os.Args[1]
	if len(file) == 0 {
		panic("path should be defined")
	}

	cr := compiler.GetCompiler()
	cr.SetDirectory(file)
	cr.Compile()
}
