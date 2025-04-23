// export by github.com/goplus/igop/cmd/qexp

package svx

import (
	q "github.com/goplus/mcp/server/svx"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "svx",
		Path: "github.com/goplus/mcp/server/svx",
		Deps: map[string]string{
			"errors":                             "errors",
			"github.com/mark3labs/mcp-go/server": "server",
			"io/fs":                              "fs",
			"strings":                            "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"LAS": reflect.TypeOf((*q.LAS)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrUnknownScheme": reflect.ValueOf(&q.ErrUnknownScheme),
		},
		Funcs: map[string]reflect.Value{
			"ListenAndServe": reflect.ValueOf(q.ListenAndServe),
			"Register":       reflect.ValueOf(q.Register),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
		Source:        source,
	})
}

var source = "package svx\n\nimport (\n\t\"errors\"\n\t\"io/fs\"\n\t\"strings\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\nvar (\n\tErrUnknownScheme = errors.New(\"unknown scheme\")\n)\n\ntype LAS = func(addr string, svr *server.MCPServer) error\n\nvar (\n\tsvxs = make(map[string]LAS, 4)\n)\n\nfunc Register(scheme string, las LAS) {\n\tsvxs[scheme] = las\n}\n\nfunc ListenAndServe(addr string, svr *server.MCPServer) error {\n\tscheme := schemeOf(addr)\n\tif las, ok := svxs[scheme]; ok {\n\t\treturn las(addr, svr)\n\t}\n\treturn &fs.PathError{Op: \"svx.ListenAndServe\", Err: ErrUnknownScheme, Path: addr}\n}\n\nfunc schemeOf(url string) (scheme string) {\n\tpos := strings.IndexAny(url, \":/\")\n\tif pos > 0 {\n\t\tif url[pos] == ':' {\n\t\t\treturn url[:pos]\n\t\t}\n\t}\n\treturn \"\"\n}\n"
