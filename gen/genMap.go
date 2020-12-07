package gen

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
)

// EmitMap generates a map type schema
func EmitMap(t *schema.TypeMap, w io.Writer, c *config) {
	writeTemplate("map.tmpl", w, t, c)
}
