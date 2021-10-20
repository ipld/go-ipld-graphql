package gen

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
)

var graphQLScalars = []string{"Int", "Float", "Boolean", "String", "ID"}

func isBuiltInScalar(t schema.Type) bool {
	for _, bi := range graphQLScalars {
		if bi == t.Name() {
			return true
		}
	}
	return t.Name() == "Bool"
}

// EmitScalar defines a scalar type for custom scalars in the type system.
func EmitScalar(t schema.Type, w io.Writer, c *config) {
	if isBuiltInScalar(t) {
		return
	}
	writeMethod(graphQLName(t)+`__serialize`, "scalar_serialize.tmpl", w, t, c)

	writeMethod(graphQLName(t)+`__parse`, "scalar_parse.tmpl", w, t, c)

	writeMethod(graphQLName(t)+`__parseLiteral`, "scalar_parseliteral.tmpl", w, t, c)

	writeTemplate("scalar.tmpl", w, t, c)
}
