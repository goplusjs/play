package httptest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func httptest.NewRecorder() *httptest.ResponseRecorder
func execNewRecorder(_ int, p *gop.Context) {
	ret := httptest.NewRecorder()
	p.Ret(0, ret)
}

// func httptest.NewRequest(method string, target string, body io.Reader) *http.Request
func execNewRequest(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := httptest.NewRequest(args[0].(string), args[1].(string), args[2].(io.Reader))
	p.Ret(3, ret)
}

// func httptest.NewServer(handler http.Handler) *httptest.Server
func execNewServer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := httptest.NewServer(args[0].(http.Handler))
	p.Ret(1, ret)
}

// func httptest.NewTLSServer(handler http.Handler) *httptest.Server
func execNewTLSServer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := httptest.NewTLSServer(args[0].(http.Handler))
	p.Ret(1, ret)
}

// func httptest.NewUnstartedServer(handler http.Handler) *httptest.Server
func execNewUnstartedServer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := httptest.NewUnstartedServer(args[0].(http.Handler))
	p.Ret(1, ret)
}

// func (*httptest.ResponseRecorder).Flush()
func execmResponseRecorderFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*httptest.ResponseRecorder).Flush()
}

// func (*httptest.ResponseRecorder).Header() http.Header
func execmResponseRecorderHeader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httptest.ResponseRecorder).Header()
	p.Ret(1, ret)
}

// func (*httptest.ResponseRecorder).Result() *http.Response
func execmResponseRecorderResult(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httptest.ResponseRecorder).Result()
	p.Ret(1, ret)
}

// func (*httptest.ResponseRecorder).Write(buf []byte) (int, error)
func execmResponseRecorderWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*httptest.ResponseRecorder).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*httptest.ResponseRecorder).WriteHeader(code int)
func execmResponseRecorderWriteHeader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*httptest.ResponseRecorder).WriteHeader(args[1].(int))
}

// func (*httptest.ResponseRecorder).WriteString(str string) (int, error)
func execmResponseRecorderWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*httptest.ResponseRecorder).WriteString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*httptest.Server).Certificate() *x509.Certificate
func execmServerCertificate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httptest.Server).Certificate()
	p.Ret(1, ret)
}

// func (*httptest.Server).Client() *http.Client
func execmServerClient(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*httptest.Server).Client()
	p.Ret(1, ret)
}

// func (*httptest.Server).Close()
func execmServerClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*httptest.Server).Close()
}

// func (*httptest.Server).CloseClientConnections()
func execmServerCloseClientConnections(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*httptest.Server).CloseClientConnections()
}

// func (*httptest.Server).Start()
func execmServerStart(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*httptest.Server).Start()
}

// func (*httptest.Server).StartTLS()
func execmServerStartTLS(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*httptest.Server).StartTLS()
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http/httptest")

func init() {
	I.RegisterConsts(
		I.Const("DefaultRemoteAddr", qspec.ConstBoundString, httptest.DefaultRemoteAddr),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*httptest.ResponseRecorder)(nil))),
		I.Rtype(reflect.TypeOf((*httptest.Server)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewRecorder", httptest.NewRecorder, execNewRecorder),
		I.Func("NewRequest", httptest.NewRequest, execNewRequest),
		I.Func("NewServer", httptest.NewServer, execNewServer),
		I.Func("NewTLSServer", httptest.NewTLSServer, execNewTLSServer),
		I.Func("NewUnstartedServer", httptest.NewUnstartedServer, execNewUnstartedServer),
		I.Func("(*ResponseRecorder).Flush", (*httptest.ResponseRecorder).Flush, execmResponseRecorderFlush),
		I.Func("(*ResponseRecorder).Header", (*httptest.ResponseRecorder).Header, execmResponseRecorderHeader),
		I.Func("(*ResponseRecorder).Result", (*httptest.ResponseRecorder).Result, execmResponseRecorderResult),
		I.Func("(*ResponseRecorder).Write", (*httptest.ResponseRecorder).Write, execmResponseRecorderWrite),
		I.Func("(*ResponseRecorder).WriteHeader", (*httptest.ResponseRecorder).WriteHeader, execmResponseRecorderWriteHeader),
		I.Func("(*ResponseRecorder).WriteString", (*httptest.ResponseRecorder).WriteString, execmResponseRecorderWriteString),
		I.Func("(*Server).Certificate", (*httptest.Server).Certificate, execmServerCertificate),
		I.Func("(*Server).Client", (*httptest.Server).Client, execmServerClient),
		I.Func("(*Server).Close", (*httptest.Server).Close, execmServerClose),
		I.Func("(*Server).CloseClientConnections", (*httptest.Server).CloseClientConnections, execmServerCloseClientConnections),
		I.Func("(*Server).Start", (*httptest.Server).Start, execmServerStart),
		I.Func("(*Server).StartTLS", (*httptest.Server).StartTLS, execmServerStartTLS),
	)
}
