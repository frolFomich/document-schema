package number

import schema "github.com/frolFomich/document-schema"

func WithMaximum(m float64) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			sch.Put(schema.MaximumSchemaKeyword, m)
		}
	}
}

func WithMinimum(m float64) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			sch.Put(schema.MinimumSchemaKeyword, m)
		}
	}
}

func WithExclusiveMaximum(m bool) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			sch.Put(schema.ExclusiveMaximumSchemaKeyword, m)
		}
	}
}

func WithExclusiveMinimum(m bool) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			sch.Put(schema.ExclusiveMinimumSchemaKeyword, m)
		}
	}
}

func WithMultipleOf(m float64) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			sch.Put(schema.MultipleOfSchemaKeyword, m)
		}
	}
}

func WithFormat(m schema.FormatType) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			// TODO check m for int* and float* formats
			sch.Put(schema.ExclusiveMaximumSchemaKeyword, m)
		}
	}
}

func WithDefaultInt(m int64) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			sch.Put(schema.DefaultSchemaKeyword, m)
		}
	}
}

func WithDefaultFloat(m float64) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && (schema.NumberSchemaType == sch.Type() || schema.IntegerSchemaType == sch.Type()) {
			sch.Put(schema.DefaultSchemaKeyword, m)
		}
	}
}
