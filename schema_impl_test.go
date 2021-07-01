package document_schema

import (
	doc "github.com/frolFomich/abstract-document"
	"reflect"
	"testing"
)

func TestFromOther(t *testing.T) {
	type args struct {
		schema Schema
	}

	m := map[string]interface{}{
		TypeSchemaKeyword:    StringSchemaType,
		DefaultSchemaKeyword: "ABBA",
	}
	given := &SchemaBase{AbstractDocument: doc.Of(m)}
	expected, err := SchemaOf(m)
	if err != nil {
		t.Errorf("Error FromOther() - %v", err)
		return
	}

	tests := []struct {
		name string
		args args
		want *SchemaBase
	}{
		{name: "FromOther able to create schema from other schema",
			args: args{schema: given},
			want: expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromOther(tt.args.schema); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromOther() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		schemaType SchemaType
		options    []SchemaOption
	}

	const (
		givenTitle       = "Some title"
		givenDescription = "dedicated description"
	)

	tests := []struct {
		name string
		args args
		want *SchemaBase
	}{
		{name: "New() without options creates valid schema",
			args: args{schemaType: StringSchemaType},
			want: &SchemaBase{AbstractDocument: doc.Of(map[string]interface{}{
				TypeSchemaKeyword: TypeStringSchemaKeyword,
			})}},
		{name: "New() with Title and Description options creates valid schema",
			args: args{schemaType: StringSchemaType, options: []SchemaOption{
				WithTitle(givenTitle),
				WithDescription(givenDescription),
			}},
			want: &SchemaBase{AbstractDocument: doc.Of(map[string]interface{}{
				TypeSchemaKeyword:        TypeStringSchemaKeyword,
				TitleSchemaKeyword:       givenTitle,
				DescriptionSchemaKeyword: givenDescription,
			})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.schemaType, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRef(t *testing.T) {
	type args struct {
		s string
	}

	const givenRef = "/a/b/c/d/schema"

	tests := []struct {
		name string
		args args
		want *SchemaBase
	}{
		{name: "NewRef() creates schema reference successfully",
			args: args{s: givenRef},
			want: &SchemaBase{AbstractDocument: doc.Of(map[string]interface{}{
				RefSchemaKeyword: givenRef,
			})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRef(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_AsRef(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}

	const givenRef = "/reference/to/schema"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "AsRef should return correct reference value",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				RefSchemaKeyword: givenRef})},
			want: givenRef},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.AsRef(); got != tt.want {
				t.Errorf("AsRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_Description(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}

	const givenDescription = "this is the description"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Description should return correct value",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				DescriptionSchemaKeyword: givenDescription})},
			want: givenDescription},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.Description(); got != tt.want {
				t.Errorf("Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_Example(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}

	example := map[string]interface{}{
		"A": "B",
		"C": 100,
		"D": true,
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{name: "Example should return correct value",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				ExampleSchemaKeyword: example})},
			want: example},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.Example(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Example() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_IsNullable(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "IsNullable should return correct value",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				NullableSchemaKeyword: true})},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.IsNullable(); got != tt.want {
				t.Errorf("IsNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_IsRef(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "IsRef should return true for reference",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				RefSchemaKeyword: "A"})},
			want: true},
		{name: "IsRef should return false for schema",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				TypeSchemaKeyword: TypeStringSchemaKeyword})},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.IsRef(); got != tt.want {
				t.Errorf("IsRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_ReadOnly(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "ReadOnly should return true",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				ReadOnlySchemaKeyword: true})},
			want: true},
		{name: "ReadOnly should return false",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				ReadOnlySchemaKeyword: false})},
			want: false},
		{name: "ReadOnly should return false",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{})},
			want:   false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.ReadOnly(); got != tt.want {
				t.Errorf("ReadOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_Title(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}

	const givenTitle = "Test title"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Title should return correct value",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				TitleSchemaKeyword: givenTitle})},
			want: givenTitle},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.Title(); got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_Type(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}
	tests := []struct {
		name   string
		fields fields
		want   SchemaType
	}{
		{name: "Type should return correct value",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				TypeSchemaKeyword: TypeStringSchemaKeyword})},
			want: StringSchemaType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaBase_WriteOnly(t *testing.T) {
	type fields struct {
		AbstractDocument *doc.AbstractDocument
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "WriteOnly should return true",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				WriteOnlySchemaKeyword: true})},
			want: true},
		{name: "WriteOnly should return false",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{
				WriteOnlySchemaKeyword: false})},
			want: false},
		{name: "WriteOnly should return false",
			fields: fields{AbstractDocument: doc.Of(map[string]interface{}{})},
			want:   false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaBase{
				AbstractDocument: tt.fields.AbstractDocument,
			}
			if got := s.WriteOnly(); got != tt.want {
				t.Errorf("WriteOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOf(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}

	givenData := map[string]interface{}{
		TypeSchemaKeyword:    TypeStringSchemaKeyword,
		DefaultSchemaKeyword: "ABBA"}

	tests := []struct {
		name    string
		args    args
		want    *SchemaBase
		wantErr bool
	}{
		{name: "SchemaOf should return correct schema",
			args: args{data: givenData},
			want: &SchemaBase{AbstractDocument: doc.Of(givenData)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SchemaOf(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaOf() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}

	givenJson := `{"type" : "string", "minLength": 1, "maxLength": 50, "default": "Country"}`

	tests := []struct {
		name    string
		args    args
		want    *SchemaBase
		wantErr bool
	}{
		{name: "Unmarshal should return valid schema",
			args: args{[]byte(givenJson)},
			want: &SchemaBase{AbstractDocument: doc.Of(map[string]interface{}{
				TypeSchemaKeyword:      TypeStringSchemaKeyword,
				MinLengthSchemaKeyword: float64(1),
				MaxLengthSchemaKeyword: float64(50),
				DefaultSchemaKeyword:   "Country",
			})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalJSON(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
