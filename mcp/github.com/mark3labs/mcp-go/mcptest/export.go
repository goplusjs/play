// export by github.com/goplus/ixgo/cmd/qexp

package mcptest

import (
	q "github.com/mark3labs/mcp-go/mcptest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "mcptest",
		Path: "github.com/mark3labs/mcp-go/mcptest",
		Deps: map[string]string{
			"bytes":                              "bytes",
			"context":                            "context",
			"fmt":                                "fmt",
			"github.com/mark3labs/mcp-go/client": "client",
			"github.com/mark3labs/mcp-go/client/transport": "transport",
			"github.com/mark3labs/mcp-go/mcp":              "mcp",
			"github.com/mark3labs/mcp-go/server":           "server",
			"io":                                           "io",
			"log":                                          "log",
			"sync":                                         "sync",
			"testing":                                      "testing",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Server": reflect.TypeOf((*q.Server)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewServer":          reflect.ValueOf(q.NewServer),
			"NewUnstartedServer": reflect.ValueOf(q.NewUnstartedServer),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
