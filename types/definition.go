package types

import (
	"github.com/alanctgardner/gogen-avro/generator"
)

/*
  The definition of a record, fixed or enum satisfies this interface.
*/

type Definition interface {
	AvroName() QualifiedName
	Aliases() []QualifiedName

	// A user-friendly name that can be built into a Go string (for unions, mostly)
	Name() string

	GoType() string

	SerializerMethod() string
	DeserializerMethod() string

	// Add the imports and struct for the definition of this type to the generator.Package
	AddStruct(*generator.Package)
	AddSerializer(*generator.Package)
	AddDeserializer(*generator.Package)

	// Resolve references to user-defined types
	ResolveReferences(*Namespace) error

	// A JSON object defining this object, for writing the schema back out
	Definition(scope map[QualifiedName]interface{}) interface{}
	DefaultValue(lvalue string, rvalue interface{}) string
}
