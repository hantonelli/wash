package syntax

import "go/ast"

// Field represents a field
type Field struct {
	fieldName string
	typeName  string
}

func NewIdent(name string) *ast.Ident {
	if name == "" {
		return nil
	}
	return &ast.Ident{
		Name: name,
	}
}

func newIdentList(name string) []*ast.Ident {
	if name == "" {
		return make([]*ast.Ident, 0)
	}
	return []*ast.Ident{NewIdent(name)}
}

// NewBasicLit creates a new basic literal from the given text
func NewBasicLit(value string) *ast.BasicLit {
	return &ast.BasicLit{
		Value: value,
	}
}

// NewField creates a new field
func NewField(name string, typeName string) Field {
	return Field{
		fieldName: name,
		typeName:  typeName,
	}
}

// NewFieldList creates a new list of fields
func NewFieldList(fields []Field) *ast.FieldList {
	l := &ast.FieldList{}
	l.List = make([]*ast.Field, len(fields))
	for i, p := range fields {
		l.List[i] = &ast.Field{
			Names: newIdentList(p.fieldName),
			Type:  NewBasicLit(p.typeName),
		}
	}
	return l
}
