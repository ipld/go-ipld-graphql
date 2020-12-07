package gen

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
)

// EmitUnion emits the template for a union type
func EmitUnion(t *schema.TypeUnion, w io.Writer, c *config) {
	writeTemplate("union.tmpl", w, t, c)

	// types which may involve type reference cycles defer to init block
	// to make golang compiler happy.
	writeTemplate("union_degenerate.tmpl", c.initDirectives, t, c)
}
