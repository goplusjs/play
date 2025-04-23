// export by github.com/goplus/igop/cmd/qexp

package stdio

import (
	q "github.com/goplus/mcp/server/stdio"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "stdio",
		Path: "github.com/goplus/mcp/server/stdio",
		Deps: map[string]string{
			"github.com/goplus/mcp/server/svx":   "svx",
			"github.com/mark3labs/mcp-go/server": "server",
			"log":                                "log",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ListenAndServe": reflect.ValueOf(q.ListenAndServe),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Scheme": {"untyped string", constant.MakeString(string(q.Scheme))},
		},
		Source: source,
	})
}

var source = "package stdio\n\nimport (\n\t\"log\"\n\t\"github.com/goplus/mcp/server/svx\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\nconst (\n\tScheme = \"stdio\"\n)\n\nfunc init() {\n\tsvx.Register(Scheme, ListenAndServe)\n}\n\nfunc ListenAndServe(addr string, svr *server.MCPServer) error {\n\tlog.Println(\"Serving MCP server with stdio ...\")\n\treturn server.ServeStdio(svr)\n}\n"
