// export by github.com/goplus/igop/cmd/qexp

package rtx

import (
	q "github.com/goplus/mcp/mtest/rtx"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
		Source:        source,
	})
}

var source = "package rtx\n\nimport (\n\t\"context\"\n\t\"encoding/json\"\n\t\"errors\"\n\t\"fmt\"\n\t\"maps\"\n\t\"sync/atomic\"\n\t\"github.com/mark3labs/mcp-go/client/transport\"\n\t\"github.com/mark3labs/mcp-go/mcp\"\n)\n\nvar (\n\tErrUnknownMethod = errors.New(\"unknown method\")\n)\n\ntype M = map[string]any\n\ntype RoundTripper interface {\n\tRoundTrip(ctx context.Context, method string, params M) (resp M, err error)\n\tOnNotify(notify func(method string, params M))\n\tClose() error\n}\n\ntype Transport struct {\n\tt\t\ttransport.Interface\n\trequestID\tatomic.Int64\n}\n\nfunc New(t transport.Interface) *Transport {\n\treturn &Transport{\n\t\tt: t,\n\t}\n}\n\nfunc (p *Transport) Close() error {\n\treturn p.t.Close()\n}\n\nfunc (p *Transport) OnNotify(notify func(method string, params M)) {\n\tp.t.SetNotificationHandler(func(in mcp.JSONRPCNotification) {\n\t\tnotify(in.Method, Params(in.Params.Meta, in.Params.AdditionalFields))\n\t})\n}\n\nfunc Params(meta, addition M) M {\n\tif meta == nil {\n\t\treturn addition\n\t}\n\tm := make(M, len(addition)+1)\n\tmaps.Copy(m, addition)\n\tm[\"_meta\"] = meta\n\treturn m\n}\n\nfunc (p *Transport) RoundTrip(ctx context.Context, method string, params M) (ret M, err error) {\n\tid := p.requestID.Add(1)\n\trequest := transport.JSONRPCRequest{\n\t\tJSONRPC:\tmcp.JSONRPC_VERSION,\n\t\tID:\t\tid,\n\t\tMethod:\t\tmethod,\n\t\tParams:\t\tparams,\n\t}\n\tresp, err := p.t.SendRequest(ctx, request)\n\tif err != nil {\n\t\treturn\n\t}\n\tif e := resp.Error; e != nil {\n\t\terr = &Error{\n\t\t\tCode:\t\te.Code,\n\t\t\tMessage:\te.Message,\n\t\t}\n\t} else {\n\t\terr = json.Unmarshal(resp.Result, &ret)\n\t}\n\treturn\n}\n\ntype Error struct {\n\tCode\tint\t`json:\"code\"`\n\tMessage\tstring\t`json:\"message\"`\n}\n\nfunc (p *Error) Error() string {\n\treturn fmt.Sprintf(\"code: %d, message: %s\", p.Code, p.Message)\n}\n"
