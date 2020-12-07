package gen

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
)

// EmitList generates the code for a list type
func EmitList(t *schema.TypeList, w io.Writer, c *config) {
	writeTemplate("list.tmpl", w, t, c)
}
