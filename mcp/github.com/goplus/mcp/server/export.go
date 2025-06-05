// export by github.com/goplus/ixgo/cmd/qexp

package server

import (
	q "github.com/goplus/mcp/server"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "server",
		Path: "github.com/goplus/mcp/server",
		Deps: map[string]string{
			"context":                            "context",
			"encoding/json":                      "json",
			"errors":                             "errors",
			"github.com/goplus/mcp/server/stdio": "stdio",
			"github.com/goplus/mcp/server/svx":   "svx",
			"github.com/mark3labs/mcp-go/mcp":    "mcp",
			"github.com/mark3labs/mcp-go/server": "server",
			"github.com/yosida95/uritemplate/v3": "uritemplate",
			"log":                                "log",
			"reflect":                            "reflect",
			"strconv":                            "strconv",
			"strings":                            "strings",
		},
		Interfaces: map[string]reflect.Type{
			"PromptProto":   reflect.TypeOf((*q.PromptProto)(nil)).Elem(),
			"ResourceProto": reflect.TypeOf((*q.ResourceProto)(nil)).Elem(),
			"ToolProto":     reflect.TypeOf((*q.ToolProto)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"JsonContent":              reflect.TypeOf((*q.JsonContent)(nil)).Elem(),
			"JsonResourceContents":     reflect.TypeOf((*q.JsonResourceContents)(nil)).Elem(),
			"MCPApp":                   reflect.TypeOf((*q.MCPApp)(nil)).Elem(),
			"PromptApp":                reflect.TypeOf((*q.PromptApp)(nil)).Elem(),
			"PromptAppProto":           reflect.TypeOf((*q.PromptAppProto)(nil)).Elem(),
			"ResourceApp":              reflect.TypeOf((*q.ResourceApp)(nil)).Elem(),
			"ResourceAppProto":         reflect.TypeOf((*q.ResourceAppProto)(nil)).Elem(),
			"TextResourceByteContents": reflect.TypeOf((*q.TextResourceByteContents)(nil)).Elem(),
			"ToolApp":                  reflect.TypeOf((*q.ToolApp)(nil)).Elem(),
			"ToolAppProto":             reflect.TypeOf((*q.ToolAppProto)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Role": reflect.TypeOf((*q.Role)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Content__0":       reflect.ValueOf(q.Content__0),
			"Content__1":       reflect.ValueOf(q.Content__1),
			"Content__2":       reflect.ValueOf(q.Content__2),
			"Content__3":       reflect.ValueOf(q.Content__3),
			"Embedded__0":      reflect.ValueOf(q.Embedded__0),
			"Embedded__1":      reflect.ValueOf(q.Embedded__1),
			"Embedded__2":      reflect.ValueOf(q.Embedded__2),
			"Embedded__3":      reflect.ValueOf(q.Embedded__3),
			"Gopt_MCPApp_Main": reflect.ValueOf(q.Gopt_MCPApp_Main),
			"Image__0":         reflect.ValueOf(q.Image__0),
			"Image__1":         reflect.ValueOf(q.Image__1),
			"Multiple":         reflect.ValueOf(q.Multiple),
			"NewError__0":      reflect.ValueOf(q.NewError__0),
			"NewError__1":      reflect.ValueOf(q.NewError__1),
			"Number__0":        reflect.ValueOf(q.Number__0),
			"Number__1":        reflect.ValueOf(q.Number__1),
			"Number__2":        reflect.ValueOf(q.Number__2),
			"Text__0":          reflect.ValueOf(q.Text__0),
			"Text__1":          reflect.ValueOf(q.Text__1),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"RoleAssistant": {reflect.TypeOf(q.RoleAssistant), constant.MakeString(string(q.RoleAssistant))},
			"RoleUser":      {reflect.TypeOf(q.RoleUser), constant.MakeString(string(q.RoleUser))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"GopPackage": {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
		},
	})
}
