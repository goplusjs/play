package io

import (
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func io.Copy(dst io.Writer, src io.Reader) (written int64, err error)
func execCopy(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := io.Copy(args[0].(io.Writer), args[1].(io.Reader))
	p.Ret(2, ret, ret1)
}

// func io.CopyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error)
func execCopyBuffer(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := io.CopyBuffer(args[0].(io.Writer), args[1].(io.Reader), args[2].([]byte))
	p.Ret(3, ret, ret1)
}

// func io.CopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error)
func execCopyN(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := io.CopyN(args[0].(io.Writer), args[1].(io.Reader), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func io.LimitReader(r io.Reader, n int64) io.Reader
func execLimitReader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := io.LimitReader(args[0].(io.Reader), args[1].(int64))
	p.Ret(2, ret)
}

// func (*io.LimitedReader).Read(p []byte) (n int, err error)
func execmLimitedReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*io.LimitedReader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func io.MultiReader(readers ..io.Reader) io.Reader
func execMultiReader(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []io.Reader {
		ret := make([]io.Reader, len(args))
		for i, arg := range args {
			ret[i] = arg.(io.Reader)
		}
		return ret
	}
	ret := io.MultiReader(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func io.MultiWriter(writers ..io.Writer) io.Writer
func execMultiWriter(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []io.Writer {
		ret := make([]io.Writer, len(args))
		for i, arg := range args {
			ret[i] = arg.(io.Writer)
		}
		return ret
	}
	ret := io.MultiWriter(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func io.NewSectionReader(r io.ReaderAt, off int64, n int64) *io.SectionReader
func execNewSectionReader(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := io.NewSectionReader(args[0].(io.ReaderAt), args[1].(int64), args[2].(int64))
	p.Ret(3, ret)
}

// func io.Pipe() (*io.PipeReader, *io.PipeWriter)
func execPipe(_ int, p *gop.Context) {
	ret, ret1 := io.Pipe()
	p.Ret(0, ret, ret1)
}

// func (*io.PipeReader).Close() error
func execmPipeReaderClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*io.PipeReader).Close()
	p.Ret(1, ret)
}

// func (*io.PipeReader).CloseWithError(err error) error
func execmPipeReaderCloseWithError(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*io.PipeReader).CloseWithError(args[1].(error))
	p.Ret(2, ret)
}

// func (*io.PipeReader).Read(data []byte) (n int, err error)
func execmPipeReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*io.PipeReader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*io.PipeWriter).Close() error
func execmPipeWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*io.PipeWriter).Close()
	p.Ret(1, ret)
}

// func (*io.PipeWriter).CloseWithError(err error) error
func execmPipeWriterCloseWithError(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*io.PipeWriter).CloseWithError(args[1].(error))
	p.Ret(2, ret)
}

// func (*io.PipeWriter).Write(data []byte) (n int, err error)
func execmPipeWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*io.PipeWriter).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func io.ReadAtLeast(r io.Reader, buf []byte, min int) (n int, err error)
func execReadAtLeast(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := io.ReadAtLeast(args[0].(io.Reader), args[1].([]byte), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func io.ReadFull(r io.Reader, buf []byte) (n int, err error)
func execReadFull(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := io.ReadFull(args[0].(io.Reader), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*io.SectionReader).Read(p []byte) (n int, err error)
func execmSectionReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*io.SectionReader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*io.SectionReader).ReadAt(p []byte, off int64) (n int, err error)
func execmSectionReaderReadAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*io.SectionReader).ReadAt(args[1].([]byte), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func (*io.SectionReader).Seek(offset int64, whence int) (int64, error)
func execmSectionReaderSeek(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*io.SectionReader).Seek(args[1].(int64), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func (*io.SectionReader).Size() int64
func execmSectionReaderSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*io.SectionReader).Size()
	p.Ret(1, ret)
}

// func io.TeeReader(r io.Reader, w io.Writer) io.Reader
func execTeeReader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := io.TeeReader(args[0].(io.Reader), args[1].(io.Writer))
	p.Ret(2, ret)
}

// func io.WriteString(w io.Writer, s string) (n int, err error)
func execWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := io.WriteString(args[0].(io.Writer), args[1].(string))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("io")

func init() {
	I.RegisterConsts(
		I.Const("SeekCurrent", qspec.ConstUnboundInt, io.SeekCurrent),
		I.Const("SeekEnd", qspec.ConstUnboundInt, io.SeekEnd),
		I.Const("SeekStart", qspec.ConstUnboundInt, io.SeekStart),
	)
	I.RegisterVars(
		I.Var("EOF", &io.EOF),
		I.Var("ErrClosedPipe", &io.ErrClosedPipe),
		I.Var("ErrNoProgress", &io.ErrNoProgress),
		I.Var("ErrShortBuffer", &io.ErrShortBuffer),
		I.Var("ErrShortWrite", &io.ErrShortWrite),
		I.Var("ErrUnexpectedEOF", &io.ErrUnexpectedEOF),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*io.LimitedReader)(nil))),
		I.Rtype(reflect.TypeOf((*io.PipeReader)(nil))),
		I.Rtype(reflect.TypeOf((*io.PipeWriter)(nil))),
		I.Rtype(reflect.TypeOf((*io.SectionReader)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Copy", io.Copy, execCopy),
		I.Func("CopyBuffer", io.CopyBuffer, execCopyBuffer),
		I.Func("CopyN", io.CopyN, execCopyN),
		I.Func("LimitReader", io.LimitReader, execLimitReader),
		I.Func("(*LimitedReader).Read", (*io.LimitedReader).Read, execmLimitedReaderRead),
		I.Func("NewSectionReader", io.NewSectionReader, execNewSectionReader),
		I.Func("Pipe", io.Pipe, execPipe),
		I.Func("(*PipeReader).Close", (*io.PipeReader).Close, execmPipeReaderClose),
		I.Func("(*PipeReader).CloseWithError", (*io.PipeReader).CloseWithError, execmPipeReaderCloseWithError),
		I.Func("(*PipeReader).Read", (*io.PipeReader).Read, execmPipeReaderRead),
		I.Func("(*PipeWriter).Close", (*io.PipeWriter).Close, execmPipeWriterClose),
		I.Func("(*PipeWriter).CloseWithError", (*io.PipeWriter).CloseWithError, execmPipeWriterCloseWithError),
		I.Func("(*PipeWriter).Write", (*io.PipeWriter).Write, execmPipeWriterWrite),
		I.Func("ReadAtLeast", io.ReadAtLeast, execReadAtLeast),
		I.Func("ReadFull", io.ReadFull, execReadFull),
		I.Func("(*SectionReader).Read", (*io.SectionReader).Read, execmSectionReaderRead),
		I.Func("(*SectionReader).ReadAt", (*io.SectionReader).ReadAt, execmSectionReaderReadAt),
		I.Func("(*SectionReader).Seek", (*io.SectionReader).Seek, execmSectionReaderSeek),
		I.Func("(*SectionReader).Size", (*io.SectionReader).Size, execmSectionReaderSize),
		I.Func("TeeReader", io.TeeReader, execTeeReader),
		I.Func("WriteString", io.WriteString, execWriteString),
	)
	I.RegisterFuncvs(
		I.Funcv("MultiReader", io.MultiReader, execMultiReader),
		I.Funcv("MultiWriter", io.MultiWriter, execMultiWriter),
	)
}
