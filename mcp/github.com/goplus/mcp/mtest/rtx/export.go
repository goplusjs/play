// export by github.com/goplus/ixgo/cmd/qexp

package rtx

import (
	q "github.com/goplus/mcp/mtest/rtx"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "rtx",
		Path: "github.com/goplus/mcp/mtest/rtx",
		Deps: map[string]string{
			"context":       "context",
			"encoding/json": "json",
			"errors":        "errors",
			"fmt":           "fmt",
			"github.com/mark3labs/mcp-go/client/transport": "transport",
			"github.com/mark3labs/mcp-go/mcp":              "mcp",
			"maps":                                         "maps",
			"sync/atomic":                                  "atomic",
		},
		Interfaces: map[string]reflect.Type{
			"RoundTripper": reflect.TypeOf((*q.RoundTripper)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
			"Transport": reflect.TypeOf((*q.Transport)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"M": reflect.TypeOf((*q.M)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrUnknownMethod": reflect.ValueOf(&q.ErrUnknownMethod),
		},
		Funcs: map[string]reflect.Value{
			"New":    reflect.ValueOf(q.New),
			"Params": reflect.ValueOf(q.Params),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
