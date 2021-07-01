package document_schema

import doc "github.com/frolFomich/abstract-document"

func WithNullable(val bool) SchemaOption {
	return func(schema *SchemaBase) {
		if schema != nil {
			schema.Put(NullableSchemaKeyword, val)
		}
	}
}

func WithReadOnly(val bool) SchemaOption {
	return func(schema *SchemaBase) {
		if schema != nil {
			schema.Put(ReadOnlySchemaKeyword, val)
		}
	}
}

func WithWriteOnly(val bool) SchemaOption {
	return func(schema *SchemaBase) {
		if schema != nil {
			schema.Put(WriteOnlySchemaKeyword, val)
		}
	}
}

func WithTitle(val string) SchemaOption {
	return func(schema *SchemaBase) {
		if schema != nil {
			schema.Put(TitleSchemaKeyword, val)
		}
	}
}

func WithDescription(val string) SchemaOption {
	return func(schema *SchemaBase) {
		if schema != nil {
			schema.Put(DescriptionSchemaKeyword, val)
		}
	}
}

func WithExample(e interface{}) SchemaOption {
	return func(sch *SchemaBase) {
		if sch != nil {
			sch.Put(ExampleSchemaKeyword, e)
		}
	}
}

func WithDefault(d doc.Document) SchemaOption {
	return func(sch *SchemaBase) {
		if sch != nil && d != nil && d.Size() > 0 {
			sch.Put(DefaultSchemaKeyword, d)
		}
	}
}
