package fcgi

import (
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/qiniu/goplus/gop"
)

// func fcgi.ProcessEnv(r *http.Request) map[string]string
func execProcessEnv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := fcgi.ProcessEnv(args[0].(*http.Request))
	p.Ret(1, ret)
}

// func fcgi.Serve(l net.Listener, handler http.Handler) error
func execServe(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := fcgi.Serve(args[0].(net.Listener), args[1].(http.Handler))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http/fcgi")

func init() {
	I.RegisterVars(
		I.Var("ErrConnClosed", &fcgi.ErrConnClosed),
		I.Var("ErrRequestAborted", &fcgi.ErrRequestAborted),
	)
	I.RegisterFuncs(
		I.Func("ProcessEnv", fcgi.ProcessEnv, execProcessEnv),
		I.Func("Serve", fcgi.Serve, execServe),
	)
}
