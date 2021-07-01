package object

import (
	doc "github.com/frolFomich/abstract-document"
	schema "github.com/frolFomich/document-schema"
)

func WithRequired(fields ...string) schema.SchemaOption {
	if fields == nil || len(fields) <= 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ObjectSchemaType == sch.Type() {
			existed := sch.Array(schema.RequiredSchemaKeyword)
			updated := make([]string, 0)
			for _, d := range existed {
				if v, ok := d.(string); ok {
					updated = append(updated, v)
				} else {
					panic(ErrorInvalidRequiredFieldName)
				}
			}
			updated = append(updated, fields...)
			sch.Put(schema.RequiredSchemaKeyword, updated)
		}
	}
}

func WithMaxProperties(val int64) schema.SchemaOption {
	if val < 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ObjectSchemaType == sch.Type() {
			sch.Put(schema.MaxPropertiesSchemaKeyword, val)
		}
	}
}

func WithMinProperties(val int64) schema.SchemaOption {
	if val < 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ObjectSchemaType == sch.Type() {
			sch.Put(schema.MinPropertiesSchemaKeyword, val)
		}
	}
}

func WithProperty(name string, prop schema.Schema) schema.SchemaOption {
	if name == "" || prop == nil {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ObjectSchemaType == sch.Type() {
			dc := sch.Document(schema.PropertiesSchemaKeyword)
			if dc == nil {
				dc = doc.New()
				sch.Put(schema.PropertiesSchemaKeyword, dc)
			}
			dc.Put(name, prop)
		}
	}
}

func WithAdditionalPropertiesAllowed(val bool) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ObjectSchemaType == sch.Type() {
			sch.Put(schema.AdditionalPropertiesSchemaKeyword, val)
		}
	}
}

func WithAdditionalPropertiesSchema(val schema.Schema) schema.SchemaOption {
	if val == nil {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ObjectSchemaType == sch.Type() {
			sch.Put(schema.AdditionalPropertiesSchemaKeyword, val)
		}
	}
}

func WithDefault(dc doc.Document) schema.SchemaOption {
	if dc == nil {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ObjectSchemaType == sch.Type() {
			sch.Put(schema.DefaultSchemaKeyword, dc)
		}
	}
}
