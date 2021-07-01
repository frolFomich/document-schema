package composite

import (
	schema "github.com/frolFomich/document-schema"
)

type allOfSchemaImpl struct {
	*schema.SchemaBase
}
type anyOfSchemaImpl struct {
	*schema.SchemaBase
}
type oneOfSchemaImpl struct {
	*schema.SchemaBase
}

func AllOf(schemas ...schema.Schema) AllOfSchema {
	if schemas == nil || len(schemas) <= 0 {
		panic(schema.ErrorInvalidArgument)
	}
	s := &allOfSchemaImpl{}
	s.Put(schema.AllOfSchemaKeyword, schemas)
	return s
}

func AnyOf(schemas ...schema.Schema) AllOfSchema {
	if schemas == nil || len(schemas) <= 0 {
		panic(schema.ErrorInvalidArgument)
	}
	s := &anyOfSchemaImpl{}
	s.Put(schema.AnyOfSchemaKeyword, schemas)
	return s
}

func OneOf(schemas ...schema.Schema) AllOfSchema {
	if schemas == nil || len(schemas) <= 0 {
		panic(schema.ErrorInvalidArgument)
	}
	s := &oneOfSchemaImpl{}
	s.Put(schema.OneOfSchemaKeyword, schemas)
	return s
}

func (a *allOfSchemaImpl) Schemas() []schema.Schema {
	return getSchemasFrom(schema.AllOfSchemaKeyword, a.SchemaBase)
}

func (a *anyOfSchemaImpl) Schemas() []schema.Schema {
	return getSchemasFrom(schema.OneOfSchemaKeyword, a.SchemaBase)
}

func (a *oneOfSchemaImpl) Schemas() []schema.Schema {
	return getSchemasFrom(schema.OneOfSchemaKeyword, a.SchemaBase)
}

func getSchemasFrom(kind string, sch *schema.SchemaBase) []schema.Schema {
	if kind == "" || sch == nil {
		panic(schema.ErrorInvalidArgument)
	}
	docs := sch.Children(kind)
	res := make([]schema.Schema, len(docs))
	for i, d := range docs {
		r, err := schema.SchemaOf(d.AsPlainMap())
		if err != nil {
			panic(err)
		}
		res[i] = r
	}
	return res
}
