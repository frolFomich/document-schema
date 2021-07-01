package document_schema

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
)

type SchemaBase struct {
	*doc.AbstractDocument
}

type SchemaOption func(schema *SchemaBase)

var (
	ErrorInvalidArgument = errors.New("invalid argument")
)

func New(schemaType SchemaType, options ...SchemaOption) *SchemaBase {
	sb := &SchemaBase{
		doc.New(),
	}
	sb.Put(TypeSchemaKeyword, schemaType.String())
	for _, opt := range options {
		opt(sb)
	}
	return sb
}

func NewRef(s string) *SchemaBase {
	if s == "" {
		panic(ErrorInvalidArgument)
	}
	sb := &SchemaBase{
		doc.New(),
	}
	sb.Put(RefSchemaKeyword, s)
	return sb
}

func UnmarshalJSON(bytes []byte) (*SchemaBase, error) {
	d, err := doc.UnmarshalJson(bytes)
	if err != nil {
		return nil, err
	}
	return &SchemaBase{doc.FromOther(d)}, nil
}

func FromOther(schema Schema) *SchemaBase {
	if schema == nil {
		panic(ErrorInvalidArgument)
	}
	return &SchemaBase{doc.FromOther(schema)}
}

func SchemaOf(data map[string]interface{}) (*SchemaBase, error) {
	if data == nil || len(data) <= 0 {
		return nil, ErrorInvalidArgument
	}
	//if _,found := data[TypeSchemaKeyword]; !found {
	//
	//	panic(ErrorInvalidSchemaType)
	//}
	return &SchemaBase{doc.Of(data)}, nil
}

func (s SchemaBase) Type() SchemaType {
	t, err := s.String(TypeSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return SchemaTypeOf(t)
}

func (s SchemaBase) IsRef() bool {
	return s.IsExist(RefSchemaKeyword)
}

func (s SchemaBase) IsNullable() bool {
	b, err := s.Boolean(NullableSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return b
}

func (s SchemaBase) ReadOnly() bool {
	b, err := s.Boolean(ReadOnlySchemaKeyword)
	if err != nil {
		return false
	}
	return b
}

func (s SchemaBase) WriteOnly() bool {
	b, err := s.Boolean(WriteOnlySchemaKeyword)
	if err != nil {
		return false
	}
	return b
}

func (s SchemaBase) Title() string {
	t, err := s.String(TitleSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return t
}

func (s SchemaBase) Description() string {
	d, err := s.String(DescriptionSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return d
}

func (s SchemaBase) Example() interface{} {
	return s.Get(ExampleSchemaKeyword)
}

func (s SchemaBase) AsRef() string {
	r, err := s.String(RefSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return r
}
