package array

import (
	schema "github.com/frolFomich/document-schema"
)

type arraySchemaImpl struct {
	*schema.SchemaBase
}

func NewArraySchema(options ...schema.SchemaOption) ArraySchema {
	return &arraySchemaImpl{
		schema.New(schema.ArraySchemaType, options...),
	}
}

func (a arraySchemaImpl) Items() schema.Schema {
	d := a.Document(schema.ItemsSchemaKeyword)

	s, err := schema.SchemaOf(d.AsPlainMap())
	if err != nil {
		panic(err)
	}
	return s
}

func (a arraySchemaImpl) MinItems() int64 {
	i, err := a.Integer(schema.MinItemsSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return i
}

func (a arraySchemaImpl) MaxItems() int64 {
	i, err := a.Integer(schema.MaxItemsSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return i
}

func (a arraySchemaImpl) IsUniqueItems() bool {
	i, err := a.Boolean(schema.UniqueItemsSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return i
}
