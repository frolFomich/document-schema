package object

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
	schema "github.com/frolFomich/document-schema"
)

type ObjectSchema interface {
	schema.Schema
	IsComposed() bool
	AsComposed() ComposedObjectSchema
	AsPlain() PlainObjectSchema
}

type PlainObjectSchema interface {
	schema.Schema
	Required() []string
	MaxProperties() int64
	MinProperties() int64
	PropertyNames() []string
	PropertySchema(name string) schema.Schema
	IsAdditionalPropertiesAllowed() bool
	AdditionalPropertiesSchema() schema.Schema
	Default() doc.Document
}

type CompositionType int

type ComposedObjectSchema interface {
	schema.Schema
	CompositionType() CompositionType
	Schemas() []schema.Schema
}

type DocumentSchema interface {
	ObjectSchema
	Name() string
}

const (
	UnknownCompositionType CompositionType = iota + 1
	AllOfCompositionType
	AnyOfCompositionType
	OneOfCompositionType
)

const (
	DocumentNameKeyword = "name"
)

var (
	compositionTypeNames = []string{"Unknown", schema.AllOfSchemaKeyword, schema.AnyOfSchemaKeyword,
		schema.OneOfSchemaKeyword}

	ErrorInvalidCompositionType = errors.New("invalid composition types")
)

func (ct CompositionType) String() string {
	if ct <= 0 || int(ct) > len(compositionTypeNames) {
		panic(ErrorInvalidCompositionType)
	}
	return compositionTypeNames[ct-1]
}

func CompositionTypeOf(s string) CompositionType {
	if s == "" {
		panic(schema.ErrorInvalidArgument)
	}
	for i, v := range compositionTypeNames {
		if v == s {
			return CompositionType(i + 1)
		}
	}
	panic(ErrorInvalidCompositionType)
}
