package string

import (
	schema "github.com/frolFomich/document-schema"
)

func WithMaxLength(l int64) schema.SchemaOption {
	if l < 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.StringSchemaType == sch.Type() {
			sch.Put(schema.MaxLengthSchemaKeyword, l)
		}
	}
}

func WithMinLength(l int64) schema.SchemaOption {
	if l < 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.StringSchemaType == sch.Type() {
			sch.Put(schema.MinLengthSchemaKeyword, l)
		}
	}
}

func WithFormat(format schema.FormatType) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.StringSchemaType == sch.Type() {
			sch.Put(schema.FormatSchemaKeyword, format.String())
		}
	}
}

func WithPattern(p string) schema.SchemaOption {
	if p == "" {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.StringSchemaType == sch.Type() {
			sch.Put(schema.PatternSchemaKeyword, p)
		}
	}
}

func WithEnum(values ...string) schema.SchemaOption {
	if values == nil || len(values) <= 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.StringSchemaType == sch.Type() {
			existed := sch.Array(schema.EnumSchemaKeyword)
			updated := make([]interface{}, 0)
			for _, d := range existed {
				if v, ok := d.(string); ok {
					updated = append(updated, v)
				} else {
					panic(ErrorInvalidEnumValue)
				}
			}
			for _, i := range values {
				updated = append(updated, interface{}(i))
			}
			sch.Put(schema.EnumSchemaKeyword, updated)
		}
	}
}

func WithDefault(d string) schema.SchemaOption {
	if d == "" {
		panic(ErrorInvalidEnumValue)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.StringSchemaType == sch.Type() {
			sch.Put(schema.DefaultSchemaKeyword, d)
		}

	}
}
