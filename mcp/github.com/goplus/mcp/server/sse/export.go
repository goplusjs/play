// export by github.com/goplus/igop/cmd/qexp

package sse

import (
	q "github.com/goplus/mcp/server/sse"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "sse",
		Path: "github.com/goplus/mcp/server/sse",
		Deps: map[string]string{
			"github.com/goplus/mcp/server/svx":   "svx",
			"github.com/mark3labs/mcp-go/server": "server",
			"log":                                "log",
			"net/http":                           "http",
			"net/url":                            "url",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Options": reflect.TypeOf((*q.Options)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ListenAndServe":    reflect.ValueOf(q.ListenAndServe),
			"ListenAndServeTLS": reflect.ValueOf(q.ListenAndServeTLS),
			"NewServer":         reflect.ValueOf(q.NewServer),
			"ParseAddr":         reflect.ValueOf(q.ParseAddr),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Scheme": {"untyped string", constant.MakeString(string(q.Scheme))},
		},
		Source: source,
	})
}

var source = "package sse\n\nimport (\n\t\"log\"\n\t\"net/http\"\n\t\"net/url\"\n\t\"github.com/goplus/mcp/server/svx\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\nconst (\n\tScheme = \"sse\"\n)\n\nfunc init() {\n\tsvx.Register(Scheme, ListenAndServe)\n}\n\nfunc ListenAndServe(addr string, svr *server.MCPServer) (err error) {\n\tserver, err := NewServer(addr, svr)\n\tif err != nil {\n\t\treturn\n\t}\n\tlog.Println(\"Serving MCP server at\", addr)\n\treturn server.ListenAndServe()\n}\n\nfunc ListenAndServeTLS(addr, certFile, keyFile string, svr *server.MCPServer) (err error) {\n\tserver, err := NewServer(addr, svr)\n\tif err != nil {\n\t\treturn\n\t}\n\tlog.Println(\"Serving TLS MCP server at\", addr)\n\treturn server.ListenAndServeTLS(certFile, keyFile)\n}\n\nfunc NewServer(addr string, svr *server.MCPServer) (ret *http.Server, err error) {\n\topts, err := ParseAddr(addr)\n\tif err != nil {\n\t\treturn\n\t}\n\tsse := server.NewSSEServer(svr)\n\tif opts.Path != \"\" {\n\t\tserver.WithBasePath(opts.Path)(sse)\n\t}\n\tif opts.Endpoint != \"\" {\n\t\tserver.WithSSEEndpoint(opts.Endpoint)(sse)\n\t}\n\tif opts.MsgEndpoint != \"\" {\n\t\tserver.WithMessageEndpoint(opts.MsgEndpoint)(sse)\n\t}\n\treturn &http.Server{Addr: opts.Host, Handler: sse}, nil\n}\n\ntype Options struct {\n\tHost\t\tstring\n\tPath\t\tstring\n\tEndpoint\tstring\n\tMsgEndpoint\tstring\n}\n\nfunc ParseAddr(addr string) (ret Options, err error) {\n\tu, err := url.Parse(addr)\n\tif err != nil {\n\t\treturn\n\t}\n\tif u.Scheme != Scheme {\n\t\terr = svx.ErrUnknownScheme\n\t\treturn\n\t}\n\tret.Host = u.Host\n\tret.Path = u.Path\n\tif u.RawQuery == \"\" {\n\t\treturn\n\t}\n\tparams, err := url.ParseQuery(u.RawQuery)\n\tif err != nil {\n\t\treturn\n\t}\n\tret.Endpoint = params.Get(\"sse\")\n\tret.MsgEndpoint = params.Get(\"msg\")\n\treturn\n}\n"
