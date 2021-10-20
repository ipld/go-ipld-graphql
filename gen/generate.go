package gen

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o ./tmplgen/bindata.go -pkg tmplgen ./tmpl

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/ipld/go-ipld-graphql/gen/tmplgen"
	"github.com/ipld/go-ipld-prime/schema"
)

type config struct {
	tmpl           *template.Template
	schemaPkg      string
	initDirectives *bytes.Buffer
	overrides      map[string]struct{}
}

// Generate creates `schema.go` at `pth` (with go package `pkg`) that
// represents a graphql schema of the `ts` schema. The generated code will expect the
// schema also has been turned into a golang schema using the go-ipld-prime code generation,
// and that the generated golang schema can be referenced as `tsPkgName` at `tsPkgPath`.
func Generate(pth string, pkg string, ts schema.TypeSystem, tsPkgName, tsPkgPath string) error {
	c := config{
		schemaPkg:      tsPkgName,
		initDirectives: bytes.NewBuffer(nil),
		overrides:      GetPreExistingMethods(pth),
	}
	setupTemplate(&c)
	return withFile(filepath.Join(pth, "schema.go"), func(f io.Writer) error {
		EmitFileHeader(f, pkg, tsPkgPath, &c)
		types := ts.GetTypes()
		typeKeys := make([]schema.TypeName, 0, len(types))
		for k := range types {
			typeKeys = append(typeKeys, k)
		}
		sort.Slice(typeKeys, func(i, j int) bool {
			return typeKeys[i] < typeKeys[j]
		})
		for _, tk := range typeKeys {
			switch t2 := types[tk].(type) {
			case *schema.TypeBool:
				EmitScalar(t2, f, &c)
			case *schema.TypeInt:
				EmitScalar(t2, f, &c)
			case *schema.TypeFloat:
				EmitScalar(t2, f, &c)
			case *schema.TypeString:
				EmitScalar(t2, f, &c)
			case *schema.TypeBytes:
				EmitScalar(t2, f, &c)
			case *schema.TypeLink:
			case *schema.TypeStruct:
				EmitStruct(t2, f, &c)
			case *schema.TypeMap:
				EmitMap(t2, f, &c)
			case *schema.TypeList:
				EmitList(t2, f, &c)
			case *schema.TypeUnion:
				EmitUnion(t2, f, &c)
			default:
				panic("unknown type" + t2.Name())
			}
		}
		return EmitFileCompletion(f, ts, &c)
	})
}

func withFile(filename string, fn func(io.Writer) error) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return fn(f)
}

func graphQLType(t schema.Type, allowRecurse bool) string {
	switch t2 := t.(type) {
	case *schema.TypeList:
		if !allowRecurse {
			// multiple layers of lists can't be directly expressed.
			return string(t.Name())
		}
		if t2.ValueIsNullable() {
			return "[" + graphQLType(t2.ValueType(), false) + "]"
		}
		return "[" + graphQLType(t2.ValueType(), false) + "!]"
	case *schema.TypeLink:
		if t2.HasReferencedType() {
			return graphQLType(t2.ReferencedType(), allowRecurse)
		}
		return "ID"
	}

	return string(t.Name())
}

func graphQLName(t schema.Type) string {
	n := t.Name()
	if isBuiltInScalar(t) {
		if n == "Bool" {
			n = "Boolean"
		}
		return "graphql." + n
	}
	if t.TypeKind() == schema.TypeKind_Link {
		sl, ok := t.(*schema.TypeLink)
		if !ok {
			return fmt.Sprintf("t is link but err: %v\n", t)
		}
		if sl.HasReferencedType() {
			return graphQLName(sl.ReferencedType())
		}
		return "graphql.ID"
	}
	return n + "__type"
}

func followingLinks(t schema.Type, linkName, into string, c *config) string {
	target := linkTarget(t)
	return fmt.Sprintf(`
	if cl, ok := %s.(cidlink.Link); ok {
		v := p.Context.Value(nodeLoaderCtxKey)
		if v == nil {
			return cl.Cid, nil
		}
		loader, ok := v.(func(context.Context, cidlink.Link, ipld.NodeBuilder) (ipld.Node, error))
		if !ok {
			return nil, errInvalidLoader
		}

		builder := %s.Type.%s__Repr.NewBuilder()
		n, err := loader(p.Context, cl, builder);
		if err != nil {
			return nil, err
		}
		%s = n
	} else {
		return nil, errInvalidLink
	}
	`, linkName, c.schemaPkg, target, into)
}

func linkTarget(t schema.Type) string {
	switch t2 := t.(type) {
	case *schema.TypeLink:
		if t2.HasReferencedType() {
			return t2.ReferencedType().Name()
		}
	}
	return "graphql.ID"
}

func setupTemplate(c *config) {
	f := template.FuncMap{
		"TypePackage":         func() string { return c.schemaPkg },
		"TypeSymbol":          func(t schema.Type) string { return graphQLType(t, true) },
		"TypeKind":            func(t schema.Type) string { return t.TypeKind().String() },
		"LocalName":           graphQLName,
		"TypeSymbolNoRecurse": func(t schema.Type) string { return graphQLType(t, false) },
		"IsBuiltIn":           isBuiltInScalar,
		"LinkTarget":          linkTarget,
		"ReturnTarget": func(name string, t schema.Type) string {
			return `
			var node ipld.Node
			` + followingLinks(t, name, "node", c) + `
			return node, nil
			`
		},
		"IntoTarget": func(name, into string, t schema.Type) string { return followingLinks(t, name, into, c) },
		"IsComplex": func(t schema.Type) bool {
			switch t2 := t.(type) {
			case *schema.TypeStruct:
				return true
			case *schema.TypeUnion:
				return true
			case *schema.TypeMap:
				return true
			case *schema.TypeList:
				return !t2.IsAnonymous()
			}
			return false
		},
		"IsStruct": func(t schema.Type) bool {
			switch t.(type) {
			case *schema.TypeStruct:
				return true
			}
			return false
		},
	}
	tmpl := template.New("").Funcs(f)
	for _, f := range tmplgen.AssetNames() {
		fd, _ := tmplgen.Asset(f)
		_, err := tmpl.New(path.Base(f)).Parse(string(fd))
		if err != nil {
			panic(err)
		}
	}

	c.tmpl = tmpl
}

func writeTemplate(tmpl string, w io.Writer, data interface{}, c *config) {
	if err := c.tmpl.ExecuteTemplate(w, tmpl, data); err != nil {
		panic(err)
	}
}

func writeMethod(name string, tmpl string, w io.Writer, data interface{}, c *config) {
	if _, ok := c.overrides[name]; ok {
		return
	}
	writeTemplate(tmpl, w, data, c)
}
