// export by github.com/goplus/ixgo/cmd/qexp

package sse

import (
	q "github.com/goplus/mcp/server/sse"

	"go/constant"
	"reflect"

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
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"Scheme": {"untyped string", constant.MakeString(string(q.Scheme))},
		},
	})
}
