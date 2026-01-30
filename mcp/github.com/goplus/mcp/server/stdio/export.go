// export by github.com/goplus/ixgo/cmd/qexp

package stdio

import (
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "stdio",
		Path: "github.com/goplus/mcp/server/stdio",
		Deps: map[string]string{
			"github.com/goplus/mcp/server/svx":   "svx",
			"github.com/mark3labs/mcp-go/server": "server",
			"log":                                "log",
		},
		Source: source,
	})
}

var source = "package stdio\n\nimport (\n\t\"log\"\n\t\"github.com/goplus/mcp/server/svx\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\nconst (\n\tScheme = \"stdio\"\n)\n\nfunc init() {\n\tsvx.Register(Scheme, ListenAndServe)\n}\n\nfunc ListenAndServe(addr string, svr *server.MCPServer) error {\n\tlog.Println(\"Serving MCP server with stdio ...\")\n\treturn server.ServeStdio(svr)\n}\n"
