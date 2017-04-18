package types

import (
	"fmt"
	"github.com/alanctgardner/gogen-avro/generator"
)

/*
  A named Reference to a user-defined type (fixed, enum, record). Just a wrapper with a name around a Definition.
*/

type Reference struct {
	typeName QualifiedName
	def      Definition
}

func NewReference(typeName QualifiedName) *Reference {
	return &Reference{
		typeName: typeName,
	}
}

func (s *Reference) Name() string {
	return s.def.Name()
}

func (s *Reference) GoType() string {
	return s.def.GoType()
}

func (s *Reference) SerializerMethod() string {
	return s.def.SerializerMethod()
}

func (s *Reference) DeserializerMethod() string {
	return s.def.DeserializerMethod()
}

func (s *Reference) AddStruct(p *generator.Package) {
	s.def.AddStruct(p)
}

func (s *Reference) AddSerializer(p *generator.Package) {
	s.def.AddSerializer(p)
}

func (s *Reference) AddDeserializer(p *generator.Package) {
	s.def.AddDeserializer(p)
}

func (s *Reference) ResolveReferences(n *Namespace) error {
	if s.def == nil {
		var ok bool
		if s.def, ok = n.Definitions[s.typeName]; !ok {
			return fmt.Errorf("Unable to resolve definition of type %v", s.typeName)
		}
		return s.def.ResolveReferences(n)
	}
	return nil
}

func (s *Reference) Definition(scope map[QualifiedName]interface{}) interface{} {
	return s.def.Definition(scope)
}

func (s *Reference) DefaultValue(lvalue string, rvalue interface{}) string {
	return s.def.DefaultValue(lvalue, rvalue)
}
