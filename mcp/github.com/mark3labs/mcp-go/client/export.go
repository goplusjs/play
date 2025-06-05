// export by github.com/goplus/ixgo/cmd/qexp

package client

import (
	q "github.com/mark3labs/mcp-go/client"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "client",
		Path: "github.com/mark3labs/mcp-go/client",
		Deps: map[string]string{
			"context":       "context",
			"encoding/json": "json",
			"errors":        "errors",
			"fmt":           "fmt",
			"github.com/mark3labs/mcp-go/client/transport": "transport",
			"github.com/mark3labs/mcp-go/mcp":              "mcp",
			"github.com/mark3labs/mcp-go/server":           "server",
			"io":                                           "io",
			"net/http":                                     "http",
			"net/url":                                      "url",
			"sync":                                         "sync",
			"sync/atomic":                                  "atomic",
		},
		Interfaces: map[string]reflect.Type{
			"MCPClient": reflect.TypeOf((*q.MCPClient)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Client":       reflect.TypeOf((*q.Client)(nil)).Elem(),
			"ClientOption": reflect.TypeOf((*q.ClientOption)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GetEndpoint":             reflect.ValueOf(q.GetEndpoint),
			"GetStderr":               reflect.ValueOf(q.GetStderr),
			"NewClient":               reflect.ValueOf(q.NewClient),
			"NewInProcessClient":      reflect.ValueOf(q.NewInProcessClient),
			"NewSSEMCPClient":         reflect.ValueOf(q.NewSSEMCPClient),
			"NewStdioMCPClient":       reflect.ValueOf(q.NewStdioMCPClient),
			"NewStreamableHttpClient": reflect.ValueOf(q.NewStreamableHttpClient),
			"WithClientCapabilities":  reflect.ValueOf(q.WithClientCapabilities),
			"WithHTTPClient":          reflect.ValueOf(q.WithHTTPClient),
			"WithHeaders":             reflect.ValueOf(q.WithHeaders),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
