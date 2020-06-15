package rpc

import (
	"io"
	"net"
	"net/http"
	"net/rpc"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func rpc.Accept(lis net.Listener)
func execAccept(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	rpc.Accept(args[0].(net.Listener))
}

// func (*rpc.Client).Call(serviceMethod string, args interface{}, reply interface{}) error
func execmClientCall(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*rpc.Client).Call(args[1].(string), args[2].(interface{}), args[3].(interface{}))
	p.Ret(4, ret)
}

// func (*rpc.Client).Close() error
func execmClientClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rpc.Client).Close()
	p.Ret(1, ret)
}

// func (*rpc.Client).Go(serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call) *rpc.Call
func execmClientGo(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := args[0].(*rpc.Client).Go(args[1].(string), args[2].(interface{}), args[3].(interface{}), args[4].(chan *rpc.Call))
	p.Ret(5, ret)
}

// func rpc.Dial(network string, address string) (*rpc.Client, error)
func execDial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := rpc.Dial(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func rpc.DialHTTP(network string, address string) (*rpc.Client, error)
func execDialHTTP(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := rpc.DialHTTP(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func rpc.DialHTTPPath(network string, address string, path string) (*rpc.Client, error)
func execDialHTTPPath(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := rpc.DialHTTPPath(args[0].(string), args[1].(string), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func rpc.HandleHTTP()
func execHandleHTTP(_ int, p *gop.Context) {
	rpc.HandleHTTP()
}

// func rpc.NewClient(conn io.ReadWriteCloser) *rpc.Client
func execNewClient(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rpc.NewClient(args[0].(io.ReadWriteCloser))
	p.Ret(1, ret)
}

// func rpc.NewClientWithCodec(codec rpc.ClientCodec) *rpc.Client
func execNewClientWithCodec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rpc.NewClientWithCodec(args[0].(rpc.ClientCodec))
	p.Ret(1, ret)
}

// func rpc.NewServer() *rpc.Server
func execNewServer(_ int, p *gop.Context) {
	ret := rpc.NewServer()
	p.Ret(0, ret)
}

// func rpc.Register(rcvr interface{}) error
func execRegister(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rpc.Register(args[0].(interface{}))
	p.Ret(1, ret)
}

// func rpc.RegisterName(name string, rcvr interface{}) error
func execRegisterName(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := rpc.RegisterName(args[0].(string), args[1].(interface{}))
	p.Ret(2, ret)
}

// func rpc.ServeCodec(codec rpc.ServerCodec)
func execServeCodec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	rpc.ServeCodec(args[0].(rpc.ServerCodec))
}

// func rpc.ServeConn(conn io.ReadWriteCloser)
func execServeConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	rpc.ServeConn(args[0].(io.ReadWriteCloser))
}

// func rpc.ServeRequest(codec rpc.ServerCodec) error
func execServeRequest(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rpc.ServeRequest(args[0].(rpc.ServerCodec))
	p.Ret(1, ret)
}

// func (*rpc.Server).Accept(lis net.Listener)
func execmServerAccept(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*rpc.Server).Accept(args[1].(net.Listener))
}

// func (*rpc.Server).HandleHTTP(rpcPath string, debugPath string)
func execmServerHandleHTTP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*rpc.Server).HandleHTTP(args[1].(string), args[2].(string))
}

// func (*rpc.Server).Register(rcvr interface{}) error
func execmServerRegister(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*rpc.Server).Register(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*rpc.Server).RegisterName(name string, rcvr interface{}) error
func execmServerRegisterName(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*rpc.Server).RegisterName(args[1].(string), args[2].(interface{}))
	p.Ret(3, ret)
}

// func (*rpc.Server).ServeCodec(codec rpc.ServerCodec)
func execmServerServeCodec(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*rpc.Server).ServeCodec(args[1].(rpc.ServerCodec))
}

// func (*rpc.Server).ServeConn(conn io.ReadWriteCloser)
func execmServerServeConn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*rpc.Server).ServeConn(args[1].(io.ReadWriteCloser))
}

// func (*rpc.Server).ServeHTTP(w http.ResponseWriter, req *http.Request)
func execmServerServeHTTP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*rpc.Server).ServeHTTP(args[1].(http.ResponseWriter), args[2].(*http.Request))
}

// func (*rpc.Server).ServeRequest(codec rpc.ServerCodec) error
func execmServerServeRequest(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*rpc.Server).ServeRequest(args[1].(rpc.ServerCodec))
	p.Ret(2, ret)
}

// func (rpc.ServerError).Error() string
func execmServerErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(rpc.ServerError).Error()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/rpc")

func init() {
	I.RegisterConsts(
		I.Const("DefaultDebugPath", qspec.ConstBoundString, rpc.DefaultDebugPath),
		I.Const("DefaultRPCPath", qspec.ConstBoundString, rpc.DefaultRPCPath),
	)
	I.RegisterVars(
		I.Var("DefaultServer", &rpc.DefaultServer),
		I.Var("ErrShutdown", &rpc.ErrShutdown),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*rpc.Call)(nil))),
		I.Rtype(reflect.TypeOf((*rpc.Client)(nil))),
		I.Rtype(reflect.TypeOf((*rpc.Request)(nil))),
		I.Rtype(reflect.TypeOf((*rpc.Response)(nil))),
		I.Rtype(reflect.TypeOf((*rpc.Server)(nil))),
		I.Type("ServerError", qspec.TyString),
	)
	I.RegisterFuncs(
		I.Func("Accept", rpc.Accept, execAccept),
		I.Func("(*Client).Call", (*rpc.Client).Call, execmClientCall),
		I.Func("(*Client).Close", (*rpc.Client).Close, execmClientClose),
		I.Func("(*Client).Go", (*rpc.Client).Go, execmClientGo),
		I.Func("Dial", rpc.Dial, execDial),
		I.Func("DialHTTP", rpc.DialHTTP, execDialHTTP),
		I.Func("DialHTTPPath", rpc.DialHTTPPath, execDialHTTPPath),
		I.Func("HandleHTTP", rpc.HandleHTTP, execHandleHTTP),
		I.Func("NewClient", rpc.NewClient, execNewClient),
		I.Func("NewClientWithCodec", rpc.NewClientWithCodec, execNewClientWithCodec),
		I.Func("NewServer", rpc.NewServer, execNewServer),
		I.Func("Register", rpc.Register, execRegister),
		I.Func("RegisterName", rpc.RegisterName, execRegisterName),
		I.Func("ServeCodec", rpc.ServeCodec, execServeCodec),
		I.Func("ServeConn", rpc.ServeConn, execServeConn),
		I.Func("ServeRequest", rpc.ServeRequest, execServeRequest),
		I.Func("(*Server).Accept", (*rpc.Server).Accept, execmServerAccept),
		I.Func("(*Server).HandleHTTP", (*rpc.Server).HandleHTTP, execmServerHandleHTTP),
		I.Func("(*Server).Register", (*rpc.Server).Register, execmServerRegister),
		I.Func("(*Server).RegisterName", (*rpc.Server).RegisterName, execmServerRegisterName),
		I.Func("(*Server).ServeCodec", (*rpc.Server).ServeCodec, execmServerServeCodec),
		I.Func("(*Server).ServeConn", (*rpc.Server).ServeConn, execmServerServeConn),
		I.Func("(*Server).ServeHTTP", (*rpc.Server).ServeHTTP, execmServerServeHTTP),
		I.Func("(*Server).ServeRequest", (*rpc.Server).ServeRequest, execmServerServeRequest),
		I.Func("(ServerError).Error", (rpc.ServerError).Error, execmServerErrorError),
	)
}
