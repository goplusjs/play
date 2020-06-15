package cgi

import (
	"net/http"
	"net/http/cgi"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*cgi.Handler).ServeHTTP(rw http.ResponseWriter, req *http.Request)
func execmHandlerServeHTTP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*cgi.Handler).ServeHTTP(args[1].(http.ResponseWriter), args[2].(*http.Request))
}

// func cgi.Request() (*http.Request, error)
func execRequest(_ int, p *gop.Context) {
	ret, ret1 := cgi.Request()
	p.Ret(0, ret, ret1)
}

// func cgi.RequestFromMap(params map[string]string) (*http.Request, error)
func execRequestFromMap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := cgi.RequestFromMap(args[0].(map[string]string))
	p.Ret(1, ret, ret1)
}

// func cgi.Serve(handler http.Handler) error
func execServe(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cgi.Serve(args[0].(http.Handler))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http/cgi")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*cgi.Handler)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Handler).ServeHTTP", (*cgi.Handler).ServeHTTP, execmHandlerServeHTTP),
		I.Func("Request", cgi.Request, execRequest),
		I.Func("RequestFromMap", cgi.RequestFromMap, execRequestFromMap),
		I.Func("Serve", cgi.Serve, execServe),
	)
}
