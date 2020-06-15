package httptrace

import (
	"context"
	"net/http/httptrace"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func httptrace.ContextClientTrace(ctx context.Context) *httptrace.ClientTrace
func execContextClientTrace(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := httptrace.ContextClientTrace(args[0].(context.Context))
	p.Ret(1, ret)
}

// func httptrace.WithClientTrace(ctx context.Context, trace *httptrace.ClientTrace) context.Context
func execWithClientTrace(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := httptrace.WithClientTrace(args[0].(context.Context), args[1].(*httptrace.ClientTrace))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http/httptrace")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*httptrace.ClientTrace)(nil))),
		I.Rtype(reflect.TypeOf((*httptrace.DNSDoneInfo)(nil))),
		I.Rtype(reflect.TypeOf((*httptrace.DNSStartInfo)(nil))),
		I.Rtype(reflect.TypeOf((*httptrace.GotConnInfo)(nil))),
		I.Rtype(reflect.TypeOf((*httptrace.WroteRequestInfo)(nil))),
	)
	I.RegisterFuncs(
		I.Func("ContextClientTrace", httptrace.ContextClientTrace, execContextClientTrace),
		I.Func("WithClientTrace", httptrace.WithClientTrace, execWithClientTrace),
	)
}
