package iotest

import (
	"io"
	"testing/iotest"

	"github.com/qiniu/goplus/gop"
)

// func iotest.DataErrReader(r io.Reader) io.Reader
func execDataErrReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := iotest.DataErrReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func iotest.HalfReader(r io.Reader) io.Reader
func execHalfReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := iotest.HalfReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func iotest.NewReadLogger(prefix string, r io.Reader) io.Reader
func execNewReadLogger(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := iotest.NewReadLogger(args[0].(string), args[1].(io.Reader))
	p.Ret(2, ret)
}

// func iotest.NewWriteLogger(prefix string, w io.Writer) io.Writer
func execNewWriteLogger(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := iotest.NewWriteLogger(args[0].(string), args[1].(io.Writer))
	p.Ret(2, ret)
}

// func iotest.OneByteReader(r io.Reader) io.Reader
func execOneByteReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := iotest.OneByteReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func iotest.TimeoutReader(r io.Reader) io.Reader
func execTimeoutReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := iotest.TimeoutReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func iotest.TruncateWriter(w io.Writer, n int64) io.Writer
func execTruncateWriter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := iotest.TruncateWriter(args[0].(io.Writer), args[1].(int64))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("testing/iotest")

func init() {
	I.RegisterVars(
		I.Var("ErrTimeout", &iotest.ErrTimeout),
	)
	I.RegisterFuncs(
		I.Func("DataErrReader", iotest.DataErrReader, execDataErrReader),
		I.Func("HalfReader", iotest.HalfReader, execHalfReader),
		I.Func("NewReadLogger", iotest.NewReadLogger, execNewReadLogger),
		I.Func("NewWriteLogger", iotest.NewWriteLogger, execNewWriteLogger),
		I.Func("OneByteReader", iotest.OneByteReader, execOneByteReader),
		I.Func("TimeoutReader", iotest.TimeoutReader, execTimeoutReader),
		I.Func("TruncateWriter", iotest.TruncateWriter, execTruncateWriter),
	)
}
