package godin

import (
	"github.com/lukasjarosch/protoc-gen-godin/generator"
)

func init() {
	generator.RegisterPlugin(new(godin))
}

type godin struct {
	gen *generator.Generator
}

func (g *godin) Name() string {
	return "godin"
}

// P forwards to g.gen.P.
func (g *godin) P(args ...interface{}) { g.gen.P(args...) }


// Init will initalize the godin plugin
func (g *godin) Init(gen *generator.Generator)  {
	g.gen = gen
}

func (g *godin) Generate(file *generator.FileDescriptor)  {
	g.P("// Godin: Why not Zoidberg?")
}

func (g *godin) GenerateImports(file *generator.FileDescriptor)  {
	
}
