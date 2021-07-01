package document_schema

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
)

type SchemaType int
type FormatType int

type Schema interface {
	doc.Document

	Type() SchemaType

	IsRef() bool
	AsRef() string
	IsNullable() bool
	ReadOnly() bool
	WriteOnly() bool
	Title() string
	Description() string

	Example() interface{}
}

const (
	ArraySchemaType SchemaType = iota + 1
	BooleanSchemaType
	IntegerSchemaType
	NumberSchemaType
	ObjectSchemaType
	StringSchemaType
)
const (
	BinaryFormatType FormatType = iota + 1
	ByteFormatType
	FloatFormatType
	DateFormatType
	DateTimeFormatType
	DoubleFormatType
	EmailFormatType
	HostnameFormatType
	Int32FormatType
	Int64FormatType
	IpV4FormatType
	IpV6FormatType
	PasswordFormatType
	URIFormatType
	URIREFFormatType
	URIReferenceFormatType
)

var (
	schemaTypeNames = [...]string{TypeArraySchemaKeyword, TypeBooleanSchemaKeyword, TypeIntegerSchemaKeyword,
		TypeNumberSchemaKeyword, TypeObjectSchemaKeyword, TypeStringSchemaKeyword}

	formatTypeNames = [...]string{FormatBinarySchemaKeyword, FormatByteSchemaKeyword, FormatFloatSchemaKeyword,
		FormatDateSchemaKeyword, FormatDateTimeSchemaKeyword, FormatDoubleSchemaKeyword, FormatEmailSchemaKeyword,
		FormatHostnameSchemaKeyword, FormatInt32SchemaKeyword, FormatInt64SchemaKeyword, FormatIpV4SchemaKeyword,
		FormatIpV6SchemaKeyword, FormatPasswordSchemaKeyword, FormatURISchemaKeyword, FormatURIREFSchemaKeyword,
		FormatURIReferenceSchemaKeyword}

	ErrorInvalidSchemaType = errors.New("invalid schema type")
	ErrorInvalidFormatType = errors.New("invalid format type")
)

func (st SchemaType) String() string {
	if st <= 0 || int(st) > len(schemaTypeNames) {
		panic(ErrorInvalidSchemaType)
	}
	return schemaTypeNames[st-1]
}

func SchemaTypeOf(s string) SchemaType {
	for i, t := range schemaTypeNames {
		if t == s {
			return SchemaType(i + 1)
		}
	}
	panic(ErrorInvalidSchemaType)
}

func (ft FormatType) String() string {
	if ft <= 0 || int(ft) > len(formatTypeNames) {
		panic(ErrorInvalidFormatType)
	}
	return formatTypeNames[ft-1]
}

func FormatTypeOf(s string) FormatType {
	for i, t := range formatTypeNames {
		if t == s {
			return FormatType(i + 1)
		}
	}
	panic(ErrorInvalidFormatType)
}
