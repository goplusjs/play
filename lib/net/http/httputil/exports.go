package httputil

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*httputil.ClientConn).Close() error
func execmClientConnClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httputil.ClientConn).Close()
	p.Ret(1, ret)
}

// func (*httputil.ClientConn).Do(req *http.Request) (*http.Response, error)
func execmClientConnDo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*httputil.ClientConn).Do(args[1].(*http.Request))
	p.Ret(2, ret, ret1)
}

// func (*httputil.ClientConn).Hijack() (c net.Conn, r *bufio.Reader)
func execmClientConnHijack(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*httputil.ClientConn).Hijack()
	p.Ret(1, ret, ret1)
}

// func (*httputil.ClientConn).Pending() int
func execmClientConnPending(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httputil.ClientConn).Pending()
	p.Ret(1, ret)
}

// func (*httputil.ClientConn).Read(req *http.Request) (resp *http.Response, err error)
func execmClientConnRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*httputil.ClientConn).Read(args[1].(*http.Request))
	p.Ret(2, ret, ret1)
}

// func (*httputil.ClientConn).Write(req *http.Request) error
func execmClientConnWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*httputil.ClientConn).Write(args[1].(*http.Request))
	p.Ret(2, ret)
}

// func httputil.DumpRequest(req *http.Request, body bool) ([]byte, error)
func execDumpRequest(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := httputil.DumpRequest(args[0].(*http.Request), args[1].(bool))
	p.Ret(2, ret, ret1)
}

// func httputil.DumpRequestOut(req *http.Request, body bool) ([]byte, error)
func execDumpRequestOut(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := httputil.DumpRequestOut(args[0].(*http.Request), args[1].(bool))
	p.Ret(2, ret, ret1)
}

// func httputil.DumpResponse(resp *http.Response, body bool) ([]byte, error)
func execDumpResponse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := httputil.DumpResponse(args[0].(*http.Response), args[1].(bool))
	p.Ret(2, ret, ret1)
}

// func httputil.NewChunkedReader(r io.Reader) io.Reader
func execNewChunkedReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := httputil.NewChunkedReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func httputil.NewChunkedWriter(w io.Writer) io.WriteCloser
func execNewChunkedWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := httputil.NewChunkedWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func httputil.NewClientConn(c net.Conn, r *bufio.Reader) *httputil.ClientConn
func execNewClientConn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := httputil.NewClientConn(args[0].(net.Conn), args[1].(*bufio.Reader))
	p.Ret(2, ret)
}

// func httputil.NewProxyClientConn(c net.Conn, r *bufio.Reader) *httputil.ClientConn
func execNewProxyClientConn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := httputil.NewProxyClientConn(args[0].(net.Conn), args[1].(*bufio.Reader))
	p.Ret(2, ret)
}

// func httputil.NewServerConn(c net.Conn, r *bufio.Reader) *httputil.ServerConn
func execNewServerConn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := httputil.NewServerConn(args[0].(net.Conn), args[1].(*bufio.Reader))
	p.Ret(2, ret)
}

// func httputil.NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy
func execNewSingleHostReverseProxy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := httputil.NewSingleHostReverseProxy(args[0].(*url.URL))
	p.Ret(1, ret)
}

// func (*httputil.ReverseProxy).ServeHTTP(rw http.ResponseWriter, req *http.Request)
func execmReverseProxyServeHTTP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*httputil.ReverseProxy).ServeHTTP(args[1].(http.ResponseWriter), args[2].(*http.Request))
}

// func (*httputil.ServerConn).Close() error
func execmServerConnClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httputil.ServerConn).Close()
	p.Ret(1, ret)
}

// func (*httputil.ServerConn).Hijack() (net.Conn, *bufio.Reader)
func execmServerConnHijack(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*httputil.ServerConn).Hijack()
	p.Ret(1, ret, ret1)
}

// func (*httputil.ServerConn).Pending() int
func execmServerConnPending(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httputil.ServerConn).Pending()
	p.Ret(1, ret)
}

// func (*httputil.ServerConn).Read() (*http.Request, error)
func execmServerConnRead(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*httputil.ServerConn).Read()
	p.Ret(1, ret, ret1)
}

// func (*httputil.ServerConn).Write(req *http.Request, resp *http.Response) error
func execmServerConnWrite(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*httputil.ServerConn).Write(args[1].(*http.Request), args[2].(*http.Response))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http/httputil")

func init() {
	I.RegisterVars(
		I.Var("ErrClosed", &httputil.ErrClosed),
		I.Var("ErrLineTooLong", &httputil.ErrLineTooLong),
		I.Var("ErrPersistEOF", &httputil.ErrPersistEOF),
		I.Var("ErrPipeline", &httputil.ErrPipeline),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*httputil.ClientConn)(nil))),
		I.Rtype(reflect.TypeOf((*httputil.ReverseProxy)(nil))),
		I.Rtype(reflect.TypeOf((*httputil.ServerConn)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*ClientConn).Close", (*httputil.ClientConn).Close, execmClientConnClose),
		I.Func("(*ClientConn).Do", (*httputil.ClientConn).Do, execmClientConnDo),
		I.Func("(*ClientConn).Hijack", (*httputil.ClientConn).Hijack, execmClientConnHijack),
		I.Func("(*ClientConn).Pending", (*httputil.ClientConn).Pending, execmClientConnPending),
		I.Func("(*ClientConn).Read", (*httputil.ClientConn).Read, execmClientConnRead),
		I.Func("(*ClientConn).Write", (*httputil.ClientConn).Write, execmClientConnWrite),
		I.Func("DumpRequest", httputil.DumpRequest, execDumpRequest),
		I.Func("DumpRequestOut", httputil.DumpRequestOut, execDumpRequestOut),
		I.Func("DumpResponse", httputil.DumpResponse, execDumpResponse),
		I.Func("NewChunkedReader", httputil.NewChunkedReader, execNewChunkedReader),
		I.Func("NewChunkedWriter", httputil.NewChunkedWriter, execNewChunkedWriter),
		I.Func("NewClientConn", httputil.NewClientConn, execNewClientConn),
		I.Func("NewProxyClientConn", httputil.NewProxyClientConn, execNewProxyClientConn),
		I.Func("NewServerConn", httputil.NewServerConn, execNewServerConn),
		I.Func("NewSingleHostReverseProxy", httputil.NewSingleHostReverseProxy, execNewSingleHostReverseProxy),
		I.Func("(*ReverseProxy).ServeHTTP", (*httputil.ReverseProxy).ServeHTTP, execmReverseProxyServeHTTP),
		I.Func("(*ServerConn).Close", (*httputil.ServerConn).Close, execmServerConnClose),
		I.Func("(*ServerConn).Hijack", (*httputil.ServerConn).Hijack, execmServerConnHijack),
		I.Func("(*ServerConn).Pending", (*httputil.ServerConn).Pending, execmServerConnPending),
		I.Func("(*ServerConn).Read", (*httputil.ServerConn).Read, execmServerConnRead),
		I.Func("(*ServerConn).Write", (*httputil.ServerConn).Write, execmServerConnWrite),
	)
}
