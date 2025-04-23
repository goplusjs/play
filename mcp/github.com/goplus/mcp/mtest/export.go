// export by github.com/goplus/igop/cmd/qexp

package mtest

import (
	q "github.com/goplus/mcp/mtest"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"GopPackage":       {"untyped string", constant.MakeString(string(q.GopPackage))},
			"GopTestClass":     {"untyped bool", constant.MakeBool(bool(q.GopTestClass))},
			"Gopo_Request_Ret": {"untyped string", constant.MakeString(string(q.Gopo_Request_Ret))},
		},
		Source: source,
	})
}

var source = "package mtest\n\nimport (\n\t\"context\"\n\t\"maps\"\n\t\"testing\"\n\t\"github.com/goplus/mcp/mtest/rtx\"\n\t\"github.com/mark3labs/mcp-go/mcp\"\n\t\"github.com/qiniu/x/test\"\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"log\"\n\t\"os\"\n\t\"github.com/goplus/mcp/mtest/mock\"\n\t\"github.com/mark3labs/mcp-go/client/transport\"\n\t\"github.com/mark3labs/mcp-go/server\"\n)\n\ntype CaseT = test.CaseT\n\ntype CaseApp struct {\n\t*Request\n\t*App\n\ttest.Case\n\tctx\tcontext.Context\n}\n\nfunc (p *CaseApp) Req__0(method string) *Request {\n\tp.Request = &Request{method: method, ctx: p}\n\treturn p.Request\n}\n\nfunc (p *CaseApp) Req__1() *Request {\n\treturn p.Request\n}\n\nfunc (p *CaseApp) Initialize(params rtx.M) *Request {\n\treturn p.Req__0(\"initialize\").Params(mapJoin(rtx.M{\n\t\t\"protocolVersion\":\tmcp.LATEST_PROTOCOL_VERSION,\n\t\t\"capabilities\": map[string]any{\n\t\t\t\"roots\":\tmap[string]any{\"listChanged\": true},\n\t\t\t\"sampling\":\tmap[string]any{},\n\t\t},\n\t\t\"clientInfo\": map[string]any{\n\t\t\t\"name\":\t\t\"github.com/goplus/mcp/mtest\",\n\t\t\t\"version\":\t\"0.3.0\",\n\t\t},\n\t}, params))\n}\n\nfunc (p *CaseApp) Ping() *Request {\n\treturn p.Req__0(\"ping\")\n}\n\nfunc (p *CaseApp) List__0(what string) *Request {\n\treturn p.Req__0(what + \"/list\")\n}\n\nfunc (p *CaseApp) List__1(what string, params rtx.M) *Request {\n\treturn p.Req__0(what + \"/list\").Params(params)\n}\n\nfunc (p *CaseApp) Read__0(uri string) *Request {\n\treturn p.Read__1(uri, nil)\n}\n\nfunc (p *CaseApp) Read__1(uri string, args map[string]any) *Request {\n\treturn p.Req__0(\"resources/read\").Params(map[string]any{\n\t\t\"uri\":\t\turi,\n\t\t\"arguments\":\targs,\n\t})\n}\n\nfunc (p *CaseApp) Subscribe(uri string) *Request {\n\treturn p.Req__0(\"resources/subscribe\").Params(map[string]any{\n\t\t\"uri\": uri,\n\t})\n}\n\nfunc (p *CaseApp) Unsubscribe(uri string) *Request {\n\treturn p.Req__0(\"resources/unsubscribe\").Params(map[string]any{\n\t\t\"uri\": uri,\n\t})\n}\n\nfunc (p *CaseApp) Prompt(name string, args map[string]any) *Request {\n\treturn p.Req__0(\"prompts/get\").Params(map[string]any{\n\t\t\"name\":\t\tname,\n\t\t\"arguments\":\targs,\n\t})\n}\n\nfunc (p *CaseApp) Call__0(name string, args map[string]any) *Request {\n\treturn p.Req__0(\"tools/call\").Params(map[string]any{\n\t\t\"name\":\t\tname,\n\t\t\"arguments\":\targs,\n\t})\n}\n\nfunc (p *CaseApp) Call__1(name string, args, meta map[string]any) *Request {\n\treturn p.Req__0(\"tools/call\").Params(map[string]any{\n\t\t\"name\":\t\tname,\n\t\t\"arguments\":\targs,\n\t\t\"_meta\":\tmeta,\n\t})\n}\n\nfunc (p *CaseApp) SetLevel(level mcp.LoggingLevel) *Request {\n\treturn p.Req__0(\"logging/setLevel\").Params(map[string]any{\n\t\t\"level\": level,\n\t})\n}\n\nfunc (p *CaseApp) Complete(args rtx.M) *Request {\n\treturn p.Req__0(\"completion/complete\").Params(args)\n}\n\nfunc (p *CaseApp) OnNotify__0(method string, notify func(params rtx.M)) {\n\tp.rt.OnNotify(func(methodIn string, params rtx.M) {\n\t\tif method == methodIn {\n\t\t\tnotify(params)\n\t\t}\n\t})\n}\n\nfunc (p *CaseApp) OnNotify__1(notify func(method string, params rtx.M)) {\n\tp.rt.OnNotify(notify)\n}\n\nfunc mapJoin(a, b map[string]any) map[string]any {\n\tif a == nil {\n\t\treturn b\n\t}\n\tif b == nil {\n\t\treturn a\n\t}\n\tm := make(map[string]any, len(a)+len(b))\n\tmaps.Copy(m, a)\n\tmaps.Copy(m, b)\n\treturn m\n}\n\ntype iCaseProto interface {\n\tinitCaseApp(*App, CaseT)\n\tMain()\n}\n\nfunc Gopt_CaseApp_TestMain(c iCaseProto, t *testing.T) {\n\tapp := new(App).initApp()\n\tc.initCaseApp(app, test.NewT(t))\n\tc.Main()\n\tapp.shutdown()\n}\n\nfunc (p *CaseApp) initCaseApp(app *App, t CaseT) {\n\tp.App = app\n\tp.CaseT = t\n\tp.ctx = context.TODO()\n}\n\nconst (\n\tGopPackage\t= \"github.com/qiniu/x/test\"\n\tGopTestClass\t= true\n)\n\nfunc Dump(args ...any) {\n\tin := make([]any, len(args))\n\tfor i, arg := range args {\n\t\tif _, ok := arg.(error); ok {\n\t\t\tin[i] = arg\n\t\t} else if b, e := json.MarshalIndent(arg, \"\", \"  \"); e == nil {\n\t\t\tin[i] = string(b)\n\t\t} else {\n\t\t\tin[i] = arg\n\t\t}\n\t}\n\tfmt.Println(in...)\n}\n\ntype MainApp struct {\n}\n\nfunc Gopt_MainApp_TestMain(app any, m *testing.M) {\n\tif me, ok := app.(interface{ MainEntry() }); ok {\n\t\tme.MainEntry()\n\t}\n\tos.Exit(m.Run())\n}\n\ntype App struct {\n\trt rtx.RoundTripper\n}\n\nfunc (p *App) initApp() *App {\n\treturn p\n}\n\nfunc (p *App) shutdown() {\n\tif rt := p.rt; rt != nil {\n\t\tp.rt = nil\n\t\trt.Close()\n\t}\n}\n\ntype MCPAppType interface {\n\tSetLAS(las func(addr string, svr *server.MCPServer) error)\n\tMain()\n}\n\nfunc (p *App) Mock(app MCPAppType) {\n\tapp.SetLAS(func(addr string, svr *server.MCPServer) (err error) {\n\t\tp.rt = mock.New(svr)\n\t\tlog.Println(\"Mocking MCP server\")\n\t\treturn nil\n\t})\n\tapp.Main()\n}\n\nfunc (p *App) TestServer__0(app MCPAppType) {\n\tp.TestServer__1(\"/sse\", app)\n}\n\nfunc (p *App) TestServer__1(path string, app MCPAppType) {\n\tapp.SetLAS(func(addr string, svr *server.MCPServer) (err error) {\n\t\tts := server.NewTestServer(svr)\n\t\tlog.Println(\"Serving MCP server at\", ts.URL)\n\t\tbaseURL := ts.URL + path\n\t\tclient, err := transport.NewSSE(baseURL)\n\t\tif err != nil {\n\t\t\tlog.Println(\"NewSSEMCPClient:\", err)\n\t\t\treturn\n\t\t}\n\t\terr = client.Start(context.Background())\n\t\tif err != nil {\n\t\t\tlog.Println(\"SSEMCPClient.Start:\", err)\n\t\t}\n\t\tp.rt = rtx.New(client)\n\t\treturn\n\t})\n\tapp.Main()\n}\n\ntype Request struct {\n\tmethod\tstring\n\tparams\trtx.M\n\tctx\t*CaseApp\n\tresp\trtx.M\n\trerr\terror\n}\n\nfunc (p *Request) t() CaseT {\n\treturn p.ctx.CaseT\n}\n\nfunc (p *Request) Params(params rtx.M) *Request {\n\tp.params = params\n\treturn p\n}\n\nfunc (p *Request) Resp() rtx.M {\n\treturn p.resp\n}\n\nfunc (p *Request) LastErr() error {\n\treturn p.rerr\n}\n\nconst (\n\tGopo_Request_Ret = \".Send,.RetWith\"\n)\n\nfunc (p *Request) RetWith(resp any) *Request {\n\tt := p.t()\n\tt.Helper()\n\tp.Send()\n\tif p.rerr != nil {\n\t\tt.Fatal(p.rerr)\n\t}\n\ttest.Gopt_Case_MatchAny(t, resp, p.resp, \"resp\")\n\treturn p\n}\n\nfunc (p *Request) Send() *Request {\n\tapp := p.ctx\n\tp.resp, p.rerr = app.rt.RoundTrip(app.ctx, p.method, p.params)\n\treturn p\n}\n"
