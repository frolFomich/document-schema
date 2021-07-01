package array

import (
	doc "github.com/frolFomich/abstract-document"
	schema "github.com/frolFomich/document-schema"
)

func WithItems(itemSchema ...schema.Schema) schema.SchemaOption {
	if itemSchema == nil || len(itemSchema) <= 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ArraySchemaType == sch.Type() {
			existed := sch.Children(schema.ItemsSchemaKeyword)
			updated := make([]doc.Document, 0)
			for _, d := range existed {
				updated = append(updated, d)
			}
			sch.Put(schema.ItemsSchemaKeyword, updated)
		}
	}
}

func WithMinItems(cnt int64) schema.SchemaOption {
	if cnt < 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ArraySchemaType == sch.Type() {
			sch.Put(schema.MinItemsSchemaKeyword, cnt)
		}
	}
}

func WithMaxItems(cnt int64) schema.SchemaOption {
	if cnt < 0 {
		panic(schema.ErrorInvalidArgument)
	}
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ArraySchemaType == sch.Type() {
			sch.Put(schema.MaxItemsSchemaKeyword, cnt)
		}
	}
}

func WithUniqueItems(unique bool) schema.SchemaOption {
	return func(sch *schema.SchemaBase) {
		if sch != nil && schema.ArraySchemaType == sch.Type() {
			sch.Put(schema.UniqueItemsSchemaKeyword, unique)
		}
	}
}
