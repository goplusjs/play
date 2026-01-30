// export by github.com/goplus/ixgo/cmd/qexp

package sse

import (
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "sse",
		Path: "github.com/goplus/mcp/server/sse",
		Deps: map[string]string{
			"github.com/goplus/mcp/server/svx":   "svx",
			"github.com/mark3labs/mcp-go/server": "server",
			"log":                                "log",
			"net/http":                           "http",
			"net/url":                            "url",
		},
		Source: source,
	})
}

var source = "package sse\n\nimport (\n\t\"log\"\n\t\"net/http\"\n\t\"net/url\"\n\t\"github.com/goplus/mcp/server/svx\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\nconst (\n\tScheme = \"sse\"\n)\n\nfunc init() {\n\tsvx.Register(Scheme, ListenAndServe)\n}\n\nfunc ListenAndServe(addr string, svr *server.MCPServer) (err error) {\n\tserver, err := NewServer(addr, svr)\n\tif err != nil {\n\t\treturn\n\t}\n\tlog.Println(\"Serving MCP server at\", addr)\n\treturn server.ListenAndServe()\n}\n\nfunc ListenAndServeTLS(addr, certFile, keyFile string, svr *server.MCPServer) (err error) {\n\tserver, err := NewServer(addr, svr)\n\tif err != nil {\n\t\treturn\n\t}\n\tlog.Println(\"Serving TLS MCP server at\", addr)\n\treturn server.ListenAndServeTLS(certFile, keyFile)\n}\n\nfunc NewServer(addr string, svr *server.MCPServer) (ret *http.Server, err error) {\n\topts, err := ParseAddr(addr)\n\tif err != nil {\n\t\treturn\n\t}\n\toptions := make([]server.SSEOption, 0, 3)\n\tif opts.Path != \"\" {\n\t\toptions = append(options, server.WithStaticBasePath(opts.Path))\n\t}\n\tif opts.Endpoint != \"\" {\n\t\toptions = append(options, server.WithSSEEndpoint(opts.Endpoint))\n\t}\n\tif opts.MsgEndpoint != \"\" {\n\t\toptions = append(options, server.WithMessageEndpoint(opts.MsgEndpoint))\n\t}\n\tsse := server.NewSSEServer(svr, options...)\n\treturn &http.Server{Addr: opts.Host, Handler: sse}, nil\n}\n\ntype Options struct {\n\tHost\t\tstring\n\tPath\t\tstring\n\tEndpoint\tstring\n\tMsgEndpoint\tstring\n}\n\nfunc ParseAddr(addr string) (ret Options, err error) {\n\tu, err := url.Parse(addr)\n\tif err != nil {\n\t\treturn\n\t}\n\tif u.Scheme != Scheme {\n\t\terr = svx.ErrUnknownScheme\n\t\treturn\n\t}\n\tret.Host = u.Host\n\tret.Path = u.Path\n\tif u.RawQuery == \"\" {\n\t\treturn\n\t}\n\tparams, err := url.ParseQuery(u.RawQuery)\n\tif err != nil {\n\t\treturn\n\t}\n\tret.Endpoint = params.Get(\"sse\")\n\tret.MsgEndpoint = params.Get(\"msg\")\n\treturn\n}\n"
