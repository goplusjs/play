package tar

import (
	"archive/tar"
	"io"
	"os"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func tar.FileInfoHeader(fi os.FileInfo, link string) (*tar.Header, error)
func execFileInfoHeader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := tar.FileInfoHeader(args[0].(os.FileInfo), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (tar.Format).String() string
func execmFormatString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(tar.Format).String()
	p.Ret(1, ret)
}

// func (*tar.Header).FileInfo() os.FileInfo
func execmHeaderFileInfo(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tar.Header).FileInfo()
	p.Ret(1, ret)
}

// func tar.NewReader(r io.Reader) *tar.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := tar.NewReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func tar.NewWriter(w io.Writer) *tar.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := tar.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func (*tar.Reader).Next() (*tar.Header, error)
func execmReaderNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*tar.Reader).Next()
	p.Ret(1, ret, ret1)
}

// func (*tar.Reader).Read(b []byte) (int, error)
func execmReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*tar.Reader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*tar.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tar.Writer).Close()
	p.Ret(1, ret)
}

// func (*tar.Writer).Flush() error
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tar.Writer).Flush()
	p.Ret(1, ret)
}

// func (*tar.Writer).Write(b []byte) (int, error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*tar.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*tar.Writer).WriteHeader(hdr *tar.Header) error
func execmWriterWriteHeader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*tar.Writer).WriteHeader(args[1].(*tar.Header))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("archive/tar")

func init() {
	I.RegisterConsts(
		I.Const("FormatGNU", reflect.Int, tar.FormatGNU),
		I.Const("FormatPAX", reflect.Int, tar.FormatPAX),
		I.Const("FormatUSTAR", reflect.Int, tar.FormatUSTAR),
		I.Const("FormatUnknown", reflect.Int, tar.FormatUnknown),
		I.Const("TypeBlock", qspec.ConstBoundRune, tar.TypeBlock),
		I.Const("TypeChar", qspec.ConstBoundRune, tar.TypeChar),
		I.Const("TypeCont", qspec.ConstBoundRune, tar.TypeCont),
		I.Const("TypeDir", qspec.ConstBoundRune, tar.TypeDir),
		I.Const("TypeFifo", qspec.ConstBoundRune, tar.TypeFifo),
		I.Const("TypeGNULongLink", qspec.ConstBoundRune, tar.TypeGNULongLink),
		I.Const("TypeGNULongName", qspec.ConstBoundRune, tar.TypeGNULongName),
		I.Const("TypeGNUSparse", qspec.ConstBoundRune, tar.TypeGNUSparse),
		I.Const("TypeLink", qspec.ConstBoundRune, tar.TypeLink),
		I.Const("TypeReg", qspec.ConstBoundRune, tar.TypeReg),
		I.Const("TypeRegA", qspec.ConstBoundRune, tar.TypeRegA),
		I.Const("TypeSymlink", qspec.ConstBoundRune, tar.TypeSymlink),
		I.Const("TypeXGlobalHeader", qspec.ConstBoundRune, tar.TypeXGlobalHeader),
		I.Const("TypeXHeader", qspec.ConstBoundRune, tar.TypeXHeader),
	)
	I.RegisterVars(
		I.Var("ErrFieldTooLong", &tar.ErrFieldTooLong),
		I.Var("ErrHeader", &tar.ErrHeader),
		I.Var("ErrWriteAfterClose", &tar.ErrWriteAfterClose),
		I.Var("ErrWriteTooLong", &tar.ErrWriteTooLong),
	)
	I.RegisterTypes(
		I.Type("Format", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*tar.Header)(nil))),
		I.Rtype(reflect.TypeOf((*tar.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*tar.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("FileInfoHeader", tar.FileInfoHeader, execFileInfoHeader),
		I.Func("(Format).String", (tar.Format).String, execmFormatString),
		I.Func("(*Header).FileInfo", (*tar.Header).FileInfo, execmHeaderFileInfo),
		I.Func("NewReader", tar.NewReader, execNewReader),
		I.Func("NewWriter", tar.NewWriter, execNewWriter),
		I.Func("(*Reader).Next", (*tar.Reader).Next, execmReaderNext),
		I.Func("(*Reader).Read", (*tar.Reader).Read, execmReaderRead),
		I.Func("(*Writer).Close", (*tar.Writer).Close, execmWriterClose),
		I.Func("(*Writer).Flush", (*tar.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).Write", (*tar.Writer).Write, execmWriterWrite),
		I.Func("(*Writer).WriteHeader", (*tar.Writer).WriteHeader, execmWriterWriteHeader),
	)
}
