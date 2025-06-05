// export by github.com/goplus/ixgo/cmd/qexp

package svx

import (
	q "github.com/goplus/mcp/server/svx"

	"reflect"

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
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"LAS": reflect.TypeOf((*q.LAS)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrUnknownScheme": reflect.ValueOf(&q.ErrUnknownScheme),
		},
		Funcs: map[string]reflect.Value{
			"ListenAndServe": reflect.ValueOf(q.ListenAndServe),
			"Register":       reflect.ValueOf(q.Register),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
