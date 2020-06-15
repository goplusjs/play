package flate

import (
	"compress/flate"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (flate.CorruptInputError).Error() string
func execmCorruptInputErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(flate.CorruptInputError).Error()
	p.Ret(1, ret)
}

// func (flate.InternalError).Error() string
func execmInternalErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(flate.InternalError).Error()
	p.Ret(1, ret)
}

// func flate.NewReader(r io.Reader) io.ReadCloser
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := flate.NewReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func flate.NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
func execNewReaderDict(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := flate.NewReaderDict(args[0].(io.Reader), args[1].([]byte))
	p.Ret(2, ret)
}

// func flate.NewWriter(w io.Writer, level int) (*flate.Writer, error)
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := flate.NewWriter(args[0].(io.Writer), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func flate.NewWriterDict(w io.Writer, level int, dict []byte) (*flate.Writer, error)
func execNewWriterDict(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := flate.NewWriterDict(args[0].(io.Writer), args[1].(int), args[2].([]byte))
	p.Ret(3, ret, ret1)
}

// func (*flate.ReadError).Error() string
func execmReadErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flate.ReadError).Error()
	p.Ret(1, ret)
}

// func (*flate.WriteError).Error() string
func execmWriteErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flate.WriteError).Error()
	p.Ret(1, ret)
}

// func (*flate.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flate.Writer).Close()
	p.Ret(1, ret)
}

// func (*flate.Writer).Flush() error
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flate.Writer).Flush()
	p.Ret(1, ret)
}

// func (*flate.Writer).Reset(dst io.Writer)
func execmWriterReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*flate.Writer).Reset(args[1].(io.Writer))
}

// func (*flate.Writer).Write(data []byte) (n int, err error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*flate.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("compress/flate")

func init() {
	I.RegisterConsts(
		I.Const("BestCompression", qspec.ConstUnboundInt, flate.BestCompression),
		I.Const("BestSpeed", qspec.ConstUnboundInt, flate.BestSpeed),
		I.Const("DefaultCompression", qspec.ConstUnboundInt, flate.DefaultCompression),
		I.Const("HuffmanOnly", qspec.ConstUnboundInt, flate.HuffmanOnly),
		I.Const("NoCompression", qspec.ConstUnboundInt, flate.NoCompression),
	)
	I.RegisterTypes(
		I.Type("CorruptInputError", qspec.TyInt64),
		I.Type("InternalError", qspec.TyString),
		I.Rtype(reflect.TypeOf((*flate.ReadError)(nil))),
		I.Rtype(reflect.TypeOf((*flate.WriteError)(nil))),
		I.Rtype(reflect.TypeOf((*flate.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(CorruptInputError).Error", (flate.CorruptInputError).Error, execmCorruptInputErrorError),
		I.Func("(InternalError).Error", (flate.InternalError).Error, execmInternalErrorError),
		I.Func("NewReader", flate.NewReader, execNewReader),
		I.Func("NewReaderDict", flate.NewReaderDict, execNewReaderDict),
		I.Func("NewWriter", flate.NewWriter, execNewWriter),
		I.Func("NewWriterDict", flate.NewWriterDict, execNewWriterDict),
		I.Func("(*ReadError).Error", (*flate.ReadError).Error, execmReadErrorError),
		I.Func("(*WriteError).Error", (*flate.WriteError).Error, execmWriteErrorError),
		I.Func("(*Writer).Close", (*flate.Writer).Close, execmWriterClose),
		I.Func("(*Writer).Flush", (*flate.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).Reset", (*flate.Writer).Reset, execmWriterReset),
		I.Func("(*Writer).Write", (*flate.Writer).Write, execmWriterWrite),
	)
}
