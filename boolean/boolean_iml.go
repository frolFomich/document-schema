package boolean

import schema "github.com/frolFomich/document-schema"

type booleanSchemaImpl struct {
	*schema.SchemaBase
}

func NewBoolean(options ...schema.SchemaOption) BooleanSchema {
	return &booleanSchemaImpl{
		schema.New(schema.BooleanSchemaType, options...),
	}
}

func (b booleanSchemaImpl) Default() bool {
	d, err := b.Boolean(schema.DefaultSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return d
}
