package number

import schema "github.com/frolFomich/document-schema"

type numberSchemaBase interface {
	schema.Schema
	Maximum() float64
	Minimum() float64
	ExclusiveMaximum() bool
	ExclusiveMinimum() bool
	MultipleOf() float64
	Format() schema.FormatType
}

type NumberSchema interface {
	numberSchemaBase
	Default() float64
}

type IntegerSchema interface {
	numberSchemaBase
	Default() int64
}
