package gen

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
)

// EmitStruct generates the definition of a struct type.
func EmitStruct(t *schema.TypeStruct, w io.Writer, c *config) {
	if len(t.Fields()) == 0 {
		writeTemplate("struct_empty.tmpl", w, t, c)

		return
	}

	for _, field := range t.Fields() {
		writeMethod(t.Name()+`__`+field.Name()+`__resolve`, "struct_field.tmpl", w, map[string]interface{}{
			"t":     t,
			"field": field,
		}, c)
	}

	writeTemplate("struct.tmpl", w, t, c)
}
