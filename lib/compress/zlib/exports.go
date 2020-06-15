package zlib

import (
	"compress/zlib"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func zlib.NewReader(r io.Reader) (io.ReadCloser, error)
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := zlib.NewReader(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func zlib.NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)
func execNewReaderDict(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := zlib.NewReaderDict(args[0].(io.Reader), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func zlib.NewWriter(w io.Writer) *zlib.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := zlib.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func zlib.NewWriterLevel(w io.Writer, level int) (*zlib.Writer, error)
func execNewWriterLevel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := zlib.NewWriterLevel(args[0].(io.Writer), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func zlib.NewWriterLevelDict(w io.Writer, level int, dict []byte) (*zlib.Writer, error)
func execNewWriterLevelDict(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := zlib.NewWriterLevelDict(args[0].(io.Writer), args[1].(int), args[2].([]byte))
	p.Ret(3, ret, ret1)
}

// func (*zlib.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zlib.Writer).Close()
	p.Ret(1, ret)
}

// func (*zlib.Writer).Flush() error
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zlib.Writer).Flush()
	p.Ret(1, ret)
}

// func (*zlib.Writer).Reset(w io.Writer)
func execmWriterReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*zlib.Writer).Reset(args[1].(io.Writer))
}

// func (*zlib.Writer).Write(p []byte) (n int, err error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*zlib.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("compress/zlib")

func init() {
	I.RegisterConsts(
		I.Const("BestCompression", qspec.ConstUnboundInt, zlib.BestCompression),
		I.Const("BestSpeed", qspec.ConstUnboundInt, zlib.BestSpeed),
		I.Const("DefaultCompression", qspec.ConstUnboundInt, zlib.DefaultCompression),
		I.Const("HuffmanOnly", qspec.ConstUnboundInt, zlib.HuffmanOnly),
		I.Const("NoCompression", qspec.ConstUnboundInt, zlib.NoCompression),
	)
	I.RegisterVars(
		I.Var("ErrChecksum", &zlib.ErrChecksum),
		I.Var("ErrDictionary", &zlib.ErrDictionary),
		I.Var("ErrHeader", &zlib.ErrHeader),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*zlib.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewReader", zlib.NewReader, execNewReader),
		I.Func("NewReaderDict", zlib.NewReaderDict, execNewReaderDict),
		I.Func("NewWriter", zlib.NewWriter, execNewWriter),
		I.Func("NewWriterLevel", zlib.NewWriterLevel, execNewWriterLevel),
		I.Func("NewWriterLevelDict", zlib.NewWriterLevelDict, execNewWriterLevelDict),
		I.Func("(*Writer).Close", (*zlib.Writer).Close, execmWriterClose),
		I.Func("(*Writer).Flush", (*zlib.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).Reset", (*zlib.Writer).Reset, execmWriterReset),
		I.Func("(*Writer).Write", (*zlib.Writer).Write, execmWriterWrite),
	)
}
