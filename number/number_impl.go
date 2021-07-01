package number

import schema "github.com/frolFomich/document-schema"

type numberSchemaBaseImpl struct {
	*schema.SchemaBase
}

type numberSchemaImpl struct {
	*numberSchemaBaseImpl
}

type integerSchemaImpl struct {
	*numberSchemaBaseImpl
}

func newNumberSchema(typeNum schema.SchemaType, options ...schema.SchemaOption) *numberSchemaBaseImpl {
	return &numberSchemaBaseImpl{
		schema.New(typeNum, options...),
	}
}

func NewNumberSchema(options ...schema.SchemaOption) NumberSchema {
	return &numberSchemaImpl{
		newNumberSchema(schema.NumberSchemaType, options...),
	}
}

func NewIntegerSchema(options ...schema.SchemaOption) IntegerSchema {
	return &integerSchemaImpl{
		newNumberSchema(schema.IntegerSchemaType, options...),
	}
}

func (n *numberSchemaBaseImpl) Maximum() float64 {
	f, err := n.Number(schema.MaximumSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return f
}

func (n *numberSchemaBaseImpl) Minimum() float64 {
	f, err := n.Number(schema.MinimumSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return f
}

func (n *numberSchemaBaseImpl) ExclusiveMaximum() bool {
	b, err := n.Boolean(schema.ExclusiveMaximumSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return b
}

func (n *numberSchemaBaseImpl) ExclusiveMinimum() bool {
	b, err := n.Boolean(schema.ExclusiveMinimumSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return b
}

func (n *numberSchemaBaseImpl) MultipleOf() float64 {
	f, err := n.Number(schema.MultipleOfSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return f
}

func (n *numberSchemaBaseImpl) Format() schema.FormatType {
	s, err := n.String(schema.FormatSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return schema.FormatTypeOf(s)
}

func (n *numberSchemaImpl) Default() float64 {
	v, err := n.Number(schema.DefaultSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return v
}

func (i *integerSchemaImpl) Default() int64 {
	v, err := i.Integer(schema.DefaultSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return v
}
