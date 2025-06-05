// export by github.com/goplus/ixgo/cmd/qexp

package mtest

import (
	q "github.com/goplus/mcp/mtest"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "mtest",
		Path: "github.com/goplus/mcp/mtest",
		Deps: map[string]string{
			"context":                          "context",
			"encoding/json":                    "json",
			"fmt":                              "fmt",
			"github.com/goplus/mcp/mtest/mock": "mock",
			"github.com/goplus/mcp/mtest/rtx":  "rtx",
			"github.com/mark3labs/mcp-go/client/transport": "transport",
			"github.com/mark3labs/mcp-go/mcp":              "mcp",
			"github.com/mark3labs/mcp-go/server":           "server",
			"github.com/qiniu/x/test":                      "test",
			"log":                                          "log",
			"maps":                                         "maps",
			"os":                                           "os",
			"testing":                                      "testing",
		},
		Interfaces: map[string]reflect.Type{
			"MCPAppType": reflect.TypeOf((*q.MCPAppType)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"App":     reflect.TypeOf((*q.App)(nil)).Elem(),
			"CaseApp": reflect.TypeOf((*q.CaseApp)(nil)).Elem(),
			"MainApp": reflect.TypeOf((*q.MainApp)(nil)).Elem(),
			"Request": reflect.TypeOf((*q.Request)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"CaseT": reflect.TypeOf((*q.CaseT)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Dump":                  reflect.ValueOf(q.Dump),
			"Gopt_CaseApp_TestMain": reflect.ValueOf(q.Gopt_CaseApp_TestMain),
			"Gopt_MainApp_TestMain": reflect.ValueOf(q.Gopt_MainApp_TestMain),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"GopPackage":       {"untyped string", constant.MakeString(string(q.GopPackage))},
			"GopTestClass":     {"untyped bool", constant.MakeBool(bool(q.GopTestClass))},
			"Gopo_Request_Ret": {"untyped string", constant.MakeString(string(q.Gopo_Request_Ret))},
		},
	})
}
