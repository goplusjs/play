// export by github.com/goplus/igop/cmd/qexp

package mock

import (
	q "github.com/goplus/mcp/mtest/mock"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
		Source:        source,
	})
}

var source = "package mock\n\nimport (\n\t\"context\"\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"sync/atomic\"\n\t\"github.com/goplus/mcp/mtest/rtx\"\n\t\"github.com/mark3labs/mcp-go/mcp\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\nfunc (p *Transport) Initialize() {\n}\n\nfunc (p *Transport) Initialized() bool {\n\treturn true\n}\n\nfunc (p *Transport) NotificationChannel() chan<- mcp.JSONRPCNotification {\n\treturn p.ch\n}\n\nfunc (p *Transport) SessionID() string {\n\tpanic(\"unreachable\")\n}\n\ntype notifyNode struct {\n\tnotify\tfunc(method string, params rtx.M)\n\tprev\t*notifyNode\n}\n\ntype Transport struct {\n\tsvr\t\t*server.MCPServer\n\tch\t\tchan<- mcp.JSONRPCNotification\n\tnotify\t\tatomic.Pointer[notifyNode]\n\trequestID\tatomic.Int64\n}\n\nfunc New(svr *server.MCPServer) *Transport {\n\tch := make(chan mcp.JSONRPCNotification, 8)\n\tret := &Transport{\n\t\tsvr:\tsvr,\n\t\tch:\tch,\n\t}\n\tgo func(ch chan mcp.JSONRPCNotification, t *Transport) {\n\t\tfor in := range ch {\n\t\t\tfor lst := ret.notify.Load(); lst != nil; lst = lst.prev {\n\t\t\t\tlst.notify(in.Method, rtx.Params(in.Params.Meta, in.Params.AdditionalFields))\n\t\t\t}\n\t\t}\n\t}(ch, ret)\n\treturn ret\n}\n\nfunc (p *Transport) Close() error {\n\tif ch := p.ch; ch != nil {\n\t\tp.ch = nil\n\t\tclose(ch)\n\t}\n\treturn nil\n}\n\nfunc (p *Transport) OnNotify(notify func(method string, params rtx.M)) {\n\tn := &notifyNode{\n\t\tnotify: notify,\n\t}\n\tfor {\n\t\tprev := p.notify.Load()\n\t\tn.prev = prev\n\t\tif p.notify.CompareAndSwap(prev, n) {\n\t\t\tbreak\n\t\t}\n\t}\n}\n\nfunc (p *Transport) RoundTrip(ctx context.Context, method string, params rtx.M) (ret rtx.M, err error) {\n\tid := p.requestID.Add(1)\n\trequest := mcp.JSONRPCRequest{\n\t\tJSONRPC:\tmcp.JSONRPC_VERSION,\n\t\tID:\t\tmcp.NewRequestId(id),\n\t\tRequest: mcp.Request{\n\t\t\tMethod: method,\n\t\t},\n\t\tParams:\tparams,\n\t}\n\trequestBytes, err := json.Marshal(request)\n\tif err != nil {\n\t\treturn nil, fmt.Errorf(\"failed to marshal request: %w\", err)\n\t}\n\tsvr := p.svr\n\tctx = svr.WithContext(ctx, p)\n\tresp := svr.HandleMessage(ctx, requestBytes)\n\tswitch resp := resp.(type) {\n\tcase mcp.JSONRPCResponse:\n\t\tb, e := json.Marshal(resp.Result)\n\t\tif e != nil {\n\t\t\treturn nil, e\n\t\t}\n\t\terr = json.Unmarshal(b, &ret)\n\tcase mcp.JSONRPCError:\n\t\terr = &rtx.Error{\n\t\t\tCode:\t\tresp.Error.Code,\n\t\t\tMessage:\tresp.Error.Message,\n\t\t}\n\tdefault:\n\t\tpanic(\"unexpected response type\")\n\t}\n\treturn\n}\n"
