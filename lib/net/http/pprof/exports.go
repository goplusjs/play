package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/qiniu/goplus/gop"
)

// func pprof.Cmdline(w http.ResponseWriter, r *http.Request)
func execCmdline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	pprof.Cmdline(args[0].(http.ResponseWriter), args[1].(*http.Request))
}

// func pprof.Handler(name string) http.Handler
func execHandler(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := pprof.Handler(args[0].(string))
	p.Ret(1, ret)
}

// func pprof.Index(w http.ResponseWriter, r *http.Request)
func execIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	pprof.Index(args[0].(http.ResponseWriter), args[1].(*http.Request))
}

// func pprof.Profile(w http.ResponseWriter, r *http.Request)
func execProfile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	pprof.Profile(args[0].(http.ResponseWriter), args[1].(*http.Request))
}

// func pprof.Symbol(w http.ResponseWriter, r *http.Request)
func execSymbol(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	pprof.Symbol(args[0].(http.ResponseWriter), args[1].(*http.Request))
}

// func pprof.Trace(w http.ResponseWriter, r *http.Request)
func execTrace(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	pprof.Trace(args[0].(http.ResponseWriter), args[1].(*http.Request))
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http/pprof")

func init() {
	I.RegisterFuncs(
		I.Func("Cmdline", pprof.Cmdline, execCmdline),
		I.Func("Handler", pprof.Handler, execHandler),
		I.Func("Index", pprof.Index, execIndex),
		I.Func("Profile", pprof.Profile, execProfile),
		I.Func("Symbol", pprof.Symbol, execSymbol),
		I.Func("Trace", pprof.Trace, execTrace),
	)
}
