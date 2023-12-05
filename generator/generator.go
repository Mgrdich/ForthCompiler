package generator

type Generator struct {
}

func (generator *Generator) Generate() {

}

func GetGenerator() *Generator {
	return &Generator{}
}
