package jsonrpc

import (
	"io"
	"net/rpc/jsonrpc"

	"github.com/qiniu/goplus/gop"
)

// func jsonrpc.Dial(network string, address string) (*rpc.Client, error)
func execDial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := jsonrpc.Dial(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func jsonrpc.NewClient(conn io.ReadWriteCloser) *rpc.Client
func execNewClient(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := jsonrpc.NewClient(args[0].(io.ReadWriteCloser))
	p.Ret(1, ret)
}

// func jsonrpc.NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec
func execNewClientCodec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := jsonrpc.NewClientCodec(args[0].(io.ReadWriteCloser))
	p.Ret(1, ret)
}

// func jsonrpc.NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec
func execNewServerCodec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := jsonrpc.NewServerCodec(args[0].(io.ReadWriteCloser))
	p.Ret(1, ret)
}

// func jsonrpc.ServeConn(conn io.ReadWriteCloser)
func execServeConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	jsonrpc.ServeConn(args[0].(io.ReadWriteCloser))
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/rpc/jsonrpc")

func init() {
	I.RegisterFuncs(
		I.Func("Dial", jsonrpc.Dial, execDial),
		I.Func("NewClient", jsonrpc.NewClient, execNewClient),
		I.Func("NewClientCodec", jsonrpc.NewClientCodec, execNewClientCodec),
		I.Func("NewServerCodec", jsonrpc.NewServerCodec, execNewServerCodec),
		I.Func("ServeConn", jsonrpc.ServeConn, execServeConn),
	)
}
