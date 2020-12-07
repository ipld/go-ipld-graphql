package gen

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
	"github.com/warpfork/go-wish"
)

// EmitFileHeader templates the start of the schema.go generated content.
func EmitFileHeader(w io.Writer, pkg, tsPkg string, c *config) {
	commonInfo := map[string]string{
		"Package":   pkg,
		"TSPackage": tsPkg,
	}

	writeTemplate("common.tmpl", w, commonInfo, c)
}

// EmitFileCompletion templates deferred definitions run at `init()` time.
func EmitFileCompletion(w io.Writer, ts schema.TypeSystem, c *config) error {
	if _, err := w.Write([]byte("\nfunc init() {\n")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(wish.Dedent(c.initDirectives.String()))); err != nil {
		return err
	}
	_, err := w.Write([]byte("\n}\n"))
	return err
}
