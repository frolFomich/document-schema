package boolean

import "github.com/frolFomich/document-schema"

type BooleanSchema interface {
	document_schema.Schema
	Default() bool
}
