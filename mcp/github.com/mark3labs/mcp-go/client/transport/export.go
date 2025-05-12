// export by github.com/goplus/igop/cmd/qexp

package transport

import (
	q "github.com/mark3labs/mcp-go/client/transport"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "transport",
		Path: "github.com/mark3labs/mcp-go/client/transport",
		Deps: map[string]string{
			"bufio":                              "bufio",
			"bytes":                              "bytes",
			"context":                            "context",
			"encoding/json":                      "json",
			"fmt":                                "fmt",
			"github.com/mark3labs/mcp-go/mcp":    "mcp",
			"github.com/mark3labs/mcp-go/server": "server",
			"io":                                 "io",
			"net/http":                           "http",
			"net/url":                            "url",
			"os":                                 "os",
			"os/exec":                            "exec",
			"strings":                            "strings",
			"sync":                               "sync",
			"sync/atomic":                        "atomic",
			"time":                               "time",
		},
		Interfaces: map[string]reflect.Type{
			"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"ClientOption":          reflect.TypeOf((*q.ClientOption)(nil)).Elem(),
			"InProcessTransport":    reflect.TypeOf((*q.InProcessTransport)(nil)).Elem(),
			"JSONRPCRequest":        reflect.TypeOf((*q.JSONRPCRequest)(nil)).Elem(),
			"JSONRPCResponse":       reflect.TypeOf((*q.JSONRPCResponse)(nil)).Elem(),
			"SSE":                   reflect.TypeOf((*q.SSE)(nil)).Elem(),
			"Stdio":                 reflect.TypeOf((*q.Stdio)(nil)).Elem(),
			"StreamableHTTP":        reflect.TypeOf((*q.StreamableHTTP)(nil)).Elem(),
			"StreamableHTTPCOption": reflect.TypeOf((*q.StreamableHTTPCOption)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewInProcessTransport": reflect.ValueOf(q.NewInProcessTransport),
			"NewSSE":                reflect.ValueOf(q.NewSSE),
			"NewStdio":              reflect.ValueOf(q.NewStdio),
			"NewStreamableHTTP":     reflect.ValueOf(q.NewStreamableHTTP),
			"WithHTTPHeaders":       reflect.ValueOf(q.WithHTTPHeaders),
			"WithHTTPTimeout":       reflect.ValueOf(q.WithHTTPTimeout),
			"WithHeaders":           reflect.ValueOf(q.WithHeaders),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
