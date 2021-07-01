package array

import schema "github.com/frolFomich/document-schema"

type ArraySchema interface {
	schema.Schema
	Items() schema.Schema
	MinItems() int64
	MaxItems() int64
	IsUniqueItems() bool
}
