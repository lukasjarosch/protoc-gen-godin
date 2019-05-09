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
	for _, service := range file.Service {
		for _, meth := range service.Method {
			inputType := g.typeName(meth.GetInputType())
			g.P("// Encode", inputType, " is responsible of encoding a domain-layer ", inputType, " into a protobuf message")
			g.P("func Encode", inputType, "(request interface{}) (interface{}, error) {")
			g.P("if request == nil {")
			g.P("return nil, errors.New(", "\"nil ", inputType, "\")")
			g.P("}")
			g.P("req := pbRequest.(*", inputType, ")")
			g.P("}")
		}
	}
}

func (g *godin) GenerateImports(file *generator.FileDescriptor)  {
	
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *godin) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Return the printable name of a type (e.g. InputType)
func (g *godin) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}