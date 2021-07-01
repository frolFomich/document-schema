package string

import (
	"errors"
	schema "github.com/frolFomich/document-schema"
)

type stringSchemaImpl struct {
	*schema.SchemaBase
}

var (
	ErrorInvalidEnumValue = errors.New("invalid enum value")
)

func New(options ...schema.SchemaOption) StringSchema {
	return &stringSchemaImpl{
		schema.New(schema.StringSchemaType, options...),
	}
}

func (s *stringSchemaImpl) MaxLength() int64 {
	l, err := s.Integer(schema.MaxLengthSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return l
}

func (s *stringSchemaImpl) MinLength() int64 {
	l, err := s.Integer(schema.MaxLengthSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return l
}

func (s *stringSchemaImpl) Format() schema.FormatType {
	f, err := s.String(schema.FormatSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return schema.FormatTypeOf(f)
}

func (s *stringSchemaImpl) Pattern() string {
	p, err := s.String(schema.PatternSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return p
}

func (s *stringSchemaImpl) Enum() []string {
	is := s.Array(schema.EnumSchemaKeyword)
	if is == nil {
		return nil
	}
	res := make([]string, len(is))
	for i, v := range is {
		if s, ok := v.(string); ok {
			res[i] = s
		} else {
			panic(ErrorInvalidEnumValue)
		}
	}
	return res
}

func (s *stringSchemaImpl) Default() string {
	d, err := s.String(schema.DefaultSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return d
}
