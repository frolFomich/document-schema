package boolean

import schema "github.com/frolFomich/document-schema"

func WithDefault(b bool) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.BooleanSchemaType == sch.Type() {
			sch.Put(schema.DefaultSchemaKeyword, b)
		}
	}
}
