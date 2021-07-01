package string

import schema "github.com/frolFomich/document-schema"

type StringSchema interface {
	schema.Schema
	MaxLength() int64
	MinLength() int64
	Format() schema.FormatType
	Pattern() string
	Enum() []string
	Default() string
}
