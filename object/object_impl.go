package object

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
	schema "github.com/frolFomich/document-schema"
)

type objectSchemaImpl struct {
	*schema.SchemaBase
}

type plainObjectSchemaImpl struct {
	*objectSchemaImpl
}

type composedObjectSchemaImpl struct {
	*objectSchemaImpl
}

type documentSchemaImpl struct {
	*objectSchemaImpl
}

var (
	ErrorInvalidRequiredFieldName = errors.New("one of required field names is invalid")
)

func NewPlainObjectSchema(options ...schema.SchemaOption) ObjectSchema {
	pos := &plainObjectSchemaImpl{
		&objectSchemaImpl{
			SchemaBase: &schema.SchemaBase{
				AbstractDocument: doc.Of(map[string]interface{}{
					schema.TypeSchemaKeyword:       schema.TypeObjectSchemaKeyword,
					schema.PropertiesSchemaKeyword: doc.Of(map[string]interface{}{}),
				})}}}
	for _, opt := range options {
		opt(pos.SchemaBase)
	}
	return pos
}

func NewComposedObjectSchema(compositionType CompositionType, schemas ...schema.Schema) ObjectSchema {
	return &composedObjectSchemaImpl{
		&objectSchemaImpl{
			SchemaBase: &schema.SchemaBase{
				AbstractDocument: doc.Of(map[string]interface{}{
					schema.TypeSchemaKeyword:       schema.TypeObjectSchemaKeyword,
					schema.PropertiesSchemaKeyword: doc.Of(map[string]interface{}{}),
				})}}}
}

func (o *plainObjectSchemaImpl) Required() []string {
	is := o.Array(schema.RequiredSchemaKeyword)
	if is == nil {
		return nil
	}
	res := make([]string, len(is))
	for i, v := range is {
		if s, ok := v.(string); ok {
			res[i] = s
		} else {
			panic(ErrorInvalidRequiredFieldName)
		}
	}
	return res
}

func (o *plainObjectSchemaImpl) MaxProperties() int64 {
	i, err := o.Integer(schema.MaxPropertiesSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return i
}

func (o *plainObjectSchemaImpl) MinProperties() int64 {
	i, err := o.Integer(schema.MinPropertiesSchemaKeyword)
	if err != nil {
		panic(err)
	}
	return i
}

func (o *plainObjectSchemaImpl) PropertySchema(name string) schema.Schema {
	if name == "" {
		panic(schema.ErrorInvalidArgument)
	}
	dc := o.Document(schema.PropertiesSchemaKeyword)
	if dc == nil || dc.Size() <= 0 {
		return nil
	}

	s := dc.Document(name)
	if s == nil {
		return nil
	}

	sch, err := schema.SchemaOf(s.AsPlainMap())
	if err != nil {
		panic(err)
	}
	return sch
}

func (o *plainObjectSchemaImpl) PropertyNames() []string {
	dc := o.Document(schema.PropertiesSchemaKeyword)
	if dc == nil || dc.Size() <= 0 {
		return nil
	}
	return dc.Keys()
}

func (o *plainObjectSchemaImpl) IsAdditionalPropertiesAllowed() bool {
	if o.IsNull(schema.AdditionalPropertiesSchemaKeyword) {
		return true
	}
	i := o.Get(schema.AdditionalPropertiesSchemaKeyword)
	if b, ok := i.(bool); ok {
		return b
	}
	return true
}

func (o *plainObjectSchemaImpl) AdditionalPropertiesSchema() schema.Schema {
	i := o.Get(schema.AdditionalPropertiesSchemaKeyword)
	if i == nil {
		return nil
	}
	if m, ok := i.(map[string]interface{}); ok {
		s, err := schema.SchemaOf(m)
		if err != nil {
			panic(err)
		}
		return s
	}
	return nil
}

func (o *plainObjectSchemaImpl) Default() doc.Document {
	return o.Document(schema.DefaultSchemaKeyword)
}

func (o *objectSchemaImpl) AsComposed() ComposedObjectSchema {
	if !o.IsComposed() {
		return nil
	}
	return &composedObjectSchemaImpl{
		objectSchemaImpl: o,
	}
}

func (o *objectSchemaImpl) AsPlain() PlainObjectSchema {
	if o.IsComposed() {
		return nil
	}
	return &plainObjectSchemaImpl{
		objectSchemaImpl: o,
	}
}

func (o *objectSchemaImpl) IsComposed() bool {
	return !o.IsExist(schema.PropertiesSchemaKeyword)
}

func (c *composedObjectSchemaImpl) CompositionType() CompositionType {
	res := UnknownCompositionType
	for i, v := range compositionTypeNames {
		if c.IsExist(v) {
			res = CompositionType(i + 1)
		}
	}
	return res
}

func (c *composedObjectSchemaImpl) Schemas() []schema.Schema {
	key := c.CompositionType().String()
	dcs := c.Children(key)
	res := make([]schema.Schema, len(dcs))
	for i, v := range dcs {
		sch, err := schema.SchemaOf(v.AsPlainMap())
		if err != nil {
			panic(err)
		}
		res[i] = sch
	}
	return res
}

func (d *documentSchemaImpl) Name() string {
	s, err := d.String(DocumentNameKeyword)
	if err != nil {
		panic(err)
	}
	return s
}
