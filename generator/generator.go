package generator

import (
	"CompilerPlayground/lexer"
	"bufio"
	"os"
	"path"
	"path/filepath"
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
	for _, token := range generator.Tokens {
		if generator.writer.Available() == 0 {
			err = generator.writer.Flush()
			if err != nil {
				panic(err)
			}
		}

		switch {
		case token.Tok.IsNumber():
			err = generator.generateNumber(token.Lit)

		}

		if err != nil {
			panic(err)
		}
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

func (generator *Generator) setName(name string) {
	if path.Ext(name) != ".s" {
		panic("can not set names without s extension")
	}
	generator.name = "tmp_" + name
}

func GetGenerator() *Generator {
	return &Generator{Source: "as", name: "tmp_source.s"}
}
