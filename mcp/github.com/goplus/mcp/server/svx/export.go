// export by github.com/goplus/ixgo/cmd/qexp

package svx

import (
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "svx",
		Path: "github.com/goplus/mcp/server/svx",
		Deps: map[string]string{
			"errors":                             "errors",
			"github.com/mark3labs/mcp-go/server": "server",
			"io/fs":                              "fs",
			"strings":                            "strings",
		},
		Source: source,
	})
}

var source = "package svx\n\nimport (\n\t\"errors\"\n\t\"io/fs\"\n\t\"strings\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\nvar (\n\tErrUnknownScheme = errors.New(\"unknown scheme\")\n)\n\ntype LAS = func(addr string, svr *server.MCPServer) error\n\nvar (\n\tsvxs = make(map[string]LAS, 4)\n)\n\nfunc Register(scheme string, las LAS) {\n\tsvxs[scheme] = las\n}\n\nfunc ListenAndServe(addr string, svr *server.MCPServer) error {\n\tscheme := schemeOf(addr)\n\tif las, ok := svxs[scheme]; ok {\n\t\treturn las(addr, svr)\n\t}\n\treturn &fs.PathError{Op: \"svx.ListenAndServe\", Err: ErrUnknownScheme, Path: addr}\n}\n\nfunc schemeOf(url string) (scheme string) {\n\tpos := strings.IndexAny(url, \":/\")\n\tif pos > 0 {\n\t\tif url[pos] == ':' {\n\t\t\treturn url[:pos]\n\t\t}\n\t}\n\treturn \"\"\n}\n"
