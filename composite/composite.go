package composite

import schema "github.com/frolFomich/document-schema"

type AllOfSchema interface {
	schema.Schema
	Schemas() []schema.Schema
}

type AnyOfSchema interface {
	schema.Schema
	Schemas() []schema.Schema
}

type OneOfSchema interface {
	schema.Schema
	Schemas() []schema.Schema
}

//type NotSchema interface {
//	schema.Schema
//	Schema() schema.Schema
//}
