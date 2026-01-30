// export by github.com/goplus/ixgo/cmd/qexp

package client

import (
	q "github.com/mark3labs/mcp-go/client"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "client",
		Path: "github.com/mark3labs/mcp-go/client",
		Deps: map[string]string{
			"context":       "context",
			"encoding/json": "json",
			"errors":        "errors",
			"fmt":           "fmt",
			"github.com/mark3labs/mcp-go/client/transport": "transport",
			"github.com/mark3labs/mcp-go/mcp":              "mcp",
			"github.com/mark3labs/mcp-go/server":           "server",
			"io":                                           "io",
			"net/http":                                     "http",
			"net/url":                                      "url",
			"slices":                                       "slices",
			"sync":                                         "sync",
			"sync/atomic":                                  "atomic",
		},
		Interfaces: map[string]reflect.Type{
			"ElicitationHandler": reflect.TypeOf((*q.ElicitationHandler)(nil)).Elem(),
			"MCPClient":          reflect.TypeOf((*q.MCPClient)(nil)).Elem(),
			"RootsHandler":       reflect.TypeOf((*q.RootsHandler)(nil)).Elem(),
			"SamplingHandler":    reflect.TypeOf((*q.SamplingHandler)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Client":       reflect.TypeOf((*q.Client)(nil)).Elem(),
			"ClientOption": reflect.TypeOf((*q.ClientOption)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"MemoryTokenStore":                reflect.TypeOf((*q.MemoryTokenStore)(nil)).Elem(),
			"OAuthAuthorizationRequiredError": reflect.TypeOf((*q.OAuthAuthorizationRequiredError)(nil)).Elem(),
			"OAuthConfig":                     reflect.TypeOf((*q.OAuthConfig)(nil)).Elem(),
			"Token":                           reflect.TypeOf((*q.Token)(nil)).Elem(),
			"TokenStore":                      reflect.TypeOf((*q.TokenStore)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"GenerateCodeChallenge": reflect.ValueOf(&q.GenerateCodeChallenge),
			"GenerateCodeVerifier":  reflect.ValueOf(&q.GenerateCodeVerifier),
			"GenerateState":         reflect.ValueOf(&q.GenerateState),
			"NewMemoryTokenStore":   reflect.ValueOf(&q.NewMemoryTokenStore),
		},
		Funcs: map[string]reflect.Value{
			"GetEndpoint":                           reflect.ValueOf(q.GetEndpoint),
			"GetOAuthHandler":                       reflect.ValueOf(q.GetOAuthHandler),
			"GetStderr":                             reflect.ValueOf(q.GetStderr),
			"IsOAuthAuthorizationRequiredError":     reflect.ValueOf(q.IsOAuthAuthorizationRequiredError),
			"NewClient":                             reflect.ValueOf(q.NewClient),
			"NewInProcessClient":                    reflect.ValueOf(q.NewInProcessClient),
			"NewInProcessClientWithSamplingHandler": reflect.ValueOf(q.NewInProcessClientWithSamplingHandler),
			"NewOAuthSSEClient":                     reflect.ValueOf(q.NewOAuthSSEClient),
			"NewOAuthStreamableHttpClient":          reflect.ValueOf(q.NewOAuthStreamableHttpClient),
			"NewSSEMCPClient":                       reflect.ValueOf(q.NewSSEMCPClient),
			"NewStdioMCPClient":                     reflect.ValueOf(q.NewStdioMCPClient),
			"NewStdioMCPClientWithOptions":          reflect.ValueOf(q.NewStdioMCPClientWithOptions),
			"NewStreamableHttpClient":               reflect.ValueOf(q.NewStreamableHttpClient),
			"WithClientCapabilities":                reflect.ValueOf(q.WithClientCapabilities),
			"WithElicitationHandler":                reflect.ValueOf(q.WithElicitationHandler),
			"WithHTTPClient":                        reflect.ValueOf(q.WithHTTPClient),
			"WithHeaderFunc":                        reflect.ValueOf(q.WithHeaderFunc),
			"WithHeaders":                           reflect.ValueOf(q.WithHeaders),
			"WithRootsHandler":                      reflect.ValueOf(q.WithRootsHandler),
			"WithSamplingHandler":                   reflect.ValueOf(q.WithSamplingHandler),
			"WithSession":                           reflect.ValueOf(q.WithSession),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
