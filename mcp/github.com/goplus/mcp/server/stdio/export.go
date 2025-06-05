// export by github.com/goplus/ixgo/cmd/qexp

package stdio

import (
	q "github.com/goplus/mcp/server/stdio"

	"go/constant"
	"reflect"

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
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ListenAndServe": reflect.ValueOf(q.ListenAndServe),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"Scheme": {"untyped string", constant.MakeString(string(q.Scheme))},
		},
	})
}
