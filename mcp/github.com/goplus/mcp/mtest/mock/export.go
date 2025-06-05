// export by github.com/goplus/ixgo/cmd/qexp

package mock

import (
	q "github.com/goplus/mcp/mtest/mock"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "mock",
		Path: "github.com/goplus/mcp/mtest/mock",
		Deps: map[string]string{
			"context":                            "context",
			"encoding/json":                      "json",
			"fmt":                                "fmt",
			"github.com/goplus/mcp/mtest/rtx":    "rtx",
			"github.com/mark3labs/mcp-go/mcp":    "mcp",
			"github.com/mark3labs/mcp-go/server": "server",
			"sync/atomic":                        "atomic",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Transport": reflect.TypeOf((*q.Transport)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
