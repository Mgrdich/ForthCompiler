package main

import "CompilerPlayground/compiler"

func main() {
	cr := compiler.GetCompiler()
	cr.SetDirectory("./test1.mf")
	cr.Compile()
}
