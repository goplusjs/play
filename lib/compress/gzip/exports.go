package gzip

import (
	"compress/gzip"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func gzip.NewReader(r io.Reader) (*gzip.Reader, error)
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := gzip.NewReader(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func gzip.NewWriter(w io.Writer) *gzip.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := gzip.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func gzip.NewWriterLevel(w io.Writer, level int) (*gzip.Writer, error)
func execNewWriterLevel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := gzip.NewWriterLevel(args[0].(io.Writer), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func (*gzip.Reader).Close() error
func execmReaderClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gzip.Reader).Close()
	p.Ret(1, ret)
}

// func (*gzip.Reader).Multistream(ok bool)
func execmReaderMultistream(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*gzip.Reader).Multistream(args[1].(bool))
}

// func (*gzip.Reader).Read(p []byte) (n int, err error)
func execmReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*gzip.Reader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*gzip.Reader).Reset(r io.Reader) error
func execmReaderReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gzip.Reader).Reset(args[1].(io.Reader))
	p.Ret(2, ret)
}

// func (*gzip.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gzip.Writer).Close()
	p.Ret(1, ret)
}

// func (*gzip.Writer).Flush() error
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gzip.Writer).Flush()
	p.Ret(1, ret)
}

// func (*gzip.Writer).Reset(w io.Writer)
func execmWriterReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*gzip.Writer).Reset(args[1].(io.Writer))
}

// func (*gzip.Writer).Write(p []byte) (int, error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*gzip.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("compress/gzip")

func init() {
	I.RegisterConsts(
		I.Const("BestCompression", qspec.ConstUnboundInt, gzip.BestCompression),
		I.Const("BestSpeed", qspec.ConstUnboundInt, gzip.BestSpeed),
		I.Const("DefaultCompression", qspec.ConstUnboundInt, gzip.DefaultCompression),
		I.Const("HuffmanOnly", qspec.ConstUnboundInt, gzip.HuffmanOnly),
		I.Const("NoCompression", qspec.ConstUnboundInt, gzip.NoCompression),
	)
	I.RegisterVars(
		I.Var("ErrChecksum", &gzip.ErrChecksum),
		I.Var("ErrHeader", &gzip.ErrHeader),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*gzip.Header)(nil))),
		I.Rtype(reflect.TypeOf((*gzip.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*gzip.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewReader", gzip.NewReader, execNewReader),
		I.Func("NewWriter", gzip.NewWriter, execNewWriter),
		I.Func("NewWriterLevel", gzip.NewWriterLevel, execNewWriterLevel),
		I.Func("(*Reader).Close", (*gzip.Reader).Close, execmReaderClose),
		I.Func("(*Reader).Multistream", (*gzip.Reader).Multistream, execmReaderMultistream),
		I.Func("(*Reader).Read", (*gzip.Reader).Read, execmReaderRead),
		I.Func("(*Reader).Reset", (*gzip.Reader).Reset, execmReaderReset),
		I.Func("(*Writer).Close", (*gzip.Writer).Close, execmWriterClose),
		I.Func("(*Writer).Flush", (*gzip.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).Reset", (*gzip.Writer).Reset, execmWriterReset),
		I.Func("(*Writer).Write", (*gzip.Writer).Write, execmWriterWrite),
	)
}
