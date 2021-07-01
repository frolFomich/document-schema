package string

import (
	document "github.com/frolFomich/abstract-document"
	schema "github.com/frolFomich/document-schema"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		options []schema.SchemaOption
	}

	expectedEmpty, err := schema.SchemaOf(map[string]interface{}{
		schema.TypeSchemaKeyword: schema.TypeStringSchemaKeyword,
	})
	if err != nil {
		t.Errorf("Error: New() - %v", err)
		return
	}
	expectedMaxLength, err := schema.SchemaOf(map[string]interface{}{
		schema.TypeSchemaKeyword:      schema.TypeStringSchemaKeyword,
		schema.MaxLengthSchemaKeyword: float64(1234),
	})
	if err != nil {
		t.Errorf("Error: New() - %v", err)
		return
	}
	expectedMinLength, err := schema.SchemaOf(map[string]interface{}{
		schema.TypeSchemaKeyword:      schema.TypeStringSchemaKeyword,
		schema.MinLengthSchemaKeyword: float64(4321),
	})
	if err != nil {
		t.Errorf("Error: New() - %v", err)
		return
	}
	expectedFormat, err := schema.SchemaOf(map[string]interface{}{
		schema.TypeSchemaKeyword:   schema.TypeStringSchemaKeyword,
		schema.FormatSchemaKeyword: schema.FormatPasswordSchemaKeyword,
	})
	if err != nil {
		t.Errorf("Error: New() - %v", err)
		return
	}
	expectedDefault, err := schema.SchemaOf(map[string]interface{}{
		schema.TypeSchemaKeyword:    schema.TypeStringSchemaKeyword,
		schema.DefaultSchemaKeyword: "ABBA",
	})
	if err != nil {
		t.Errorf("Error: New() - %v", err)
		return
	}
	expectedEnum, err := schema.SchemaOf(map[string]interface{}{
		schema.TypeSchemaKeyword: schema.TypeStringSchemaKeyword,
		schema.EnumSchemaKeyword: []interface{}{interface{}("A"), interface{}("B"), interface{}("B"), interface{}("A")},
	})
	if err != nil {
		t.Errorf("Error: New() - %v", err)
		return
	}

	tests := []struct {
		name string
		args args
		want StringSchema
	}{
		{name: "Create String schema without options",
			args: args{options: []schema.SchemaOption{}},
			want: &stringSchemaImpl{
				expectedEmpty,
			}},
		{name: "Create String schema with MaxLength option",
			args: args{options: []schema.SchemaOption{WithMaxLength(1234)}},
			want: &stringSchemaImpl{
				expectedMaxLength,
			}},
		{name: "Create String schema with MinLength option",
			args: args{options: []schema.SchemaOption{WithMinLength(4321)}},
			want: &stringSchemaImpl{
				expectedMinLength,
			}},
		{name: "Create String schema with Format option",
			args: args{options: []schema.SchemaOption{WithFormat(schema.PasswordFormatType)}},
			want: &stringSchemaImpl{
				expectedFormat,
			}},
		{name: "Create String schema with Default option",
			args: args{options: []schema.SchemaOption{WithDefault("ABBA")}},
			want: &stringSchemaImpl{
				expectedDefault,
			}},
		{name: "Create String schema with Enum option",
			args: args{options: []schema.SchemaOption{WithEnum("A", "B", "B", "A")}},
			want: &stringSchemaImpl{
				expectedEnum,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSchemaImpl_Default(t *testing.T) {
	type fields struct {
		SchemaBase *schema.SchemaBase
	}

	givenDefault := "theDefaultString"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Default returns correct value",
			fields: fields{SchemaBase: &schema.SchemaBase{AbstractDocument: document.Of(map[string]interface{}{
				schema.TypeSchemaKeyword:    schema.StringSchemaType,
				schema.DefaultSchemaKeyword: givenDefault,
			})}},
			want: givenDefault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringSchemaImpl{
				SchemaBase: tt.fields.SchemaBase,
			}
			if got := s.Default(); got != tt.want {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSchemaImpl_Enum(t *testing.T) {
	type fields struct {
		SchemaBase *schema.SchemaBase
	}

	givenEnum := []interface{}{interface{}("A"), interface{}("B"), interface{}("C")}
	expectedEnum := []string{"A", "B", "C"}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{name: "Enum returns correct value",
			fields: fields{SchemaBase: &schema.SchemaBase{AbstractDocument: document.Of(map[string]interface{}{
				schema.TypeSchemaKeyword: schema.StringSchemaType,
				schema.EnumSchemaKeyword: givenEnum,
			})}},
			want: expectedEnum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringSchemaImpl{
				SchemaBase: tt.fields.SchemaBase,
			}
			if got := s.Enum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSchemaImpl_Format(t *testing.T) {
	type fields struct {
		SchemaBase *schema.SchemaBase
	}
	tests := []struct {
		name   string
		fields fields
		want   schema.FormatType
	}{
		{name: "Format returns correct value",
			fields: fields{SchemaBase: &schema.SchemaBase{AbstractDocument: document.Of(map[string]interface{}{
				schema.TypeSchemaKeyword:   schema.StringSchemaType,
				schema.FormatSchemaKeyword: schema.FormatInt64SchemaKeyword,
			})}},
			want: schema.Int64FormatType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringSchemaImpl{
				SchemaBase: tt.fields.SchemaBase,
			}
			if got := s.Format(); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSchemaImpl_MaxLength(t *testing.T) {
	type fields struct {
		SchemaBase *schema.SchemaBase
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "MaxLength returns correct value",
			fields: fields{SchemaBase: &schema.SchemaBase{AbstractDocument: document.Of(map[string]interface{}{
				schema.TypeSchemaKeyword:      schema.StringSchemaType,
				schema.MaxLengthSchemaKeyword: float64(100),
			})}},
			want: int64(100),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringSchemaImpl{
				SchemaBase: tt.fields.SchemaBase,
			}
			if got := s.MaxLength(); got != tt.want {
				t.Errorf("MaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSchemaImpl_MinLength(t *testing.T) {
	type fields struct {
		SchemaBase *schema.SchemaBase
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "MinLength returns correct value",
			fields: fields{SchemaBase: &schema.SchemaBase{AbstractDocument: document.Of(map[string]interface{}{
				schema.TypeSchemaKeyword:      schema.StringSchemaType,
				schema.MaxLengthSchemaKeyword: float64(300),
			})}},
			want: int64(300),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringSchemaImpl{
				SchemaBase: tt.fields.SchemaBase,
			}
			if got := s.MinLength(); got != tt.want {
				t.Errorf("MinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSchemaImpl_Pattern(t *testing.T) {
	type fields struct {
		SchemaBase *schema.SchemaBase
	}

	const givenPattern = "^[A|a]bba$"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Pattern returns correct value",
			fields: fields{SchemaBase: &schema.SchemaBase{AbstractDocument: document.Of(map[string]interface{}{
				schema.TypeSchemaKeyword:    schema.StringSchemaType,
				schema.PatternSchemaKeyword: givenPattern,
			})}},
			want: givenPattern,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringSchemaImpl{
				SchemaBase: tt.fields.SchemaBase,
			}
			if got := s.Pattern(); got != tt.want {
				t.Errorf("Pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
