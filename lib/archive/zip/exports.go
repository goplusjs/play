package zip

import (
	"archive/zip"
	"io"
	"os"
	"reflect"
	"time"

	"github.com/qiniu/goplus/gop"
)

// func (*zip.File).DataOffset() (offset int64, err error)
func execmFileDataOffset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*zip.File).DataOffset()
	p.Ret(1, ret, ret1)
}

// func (*zip.File).Open() (io.ReadCloser, error)
func execmFileOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*zip.File).Open()
	p.Ret(1, ret, ret1)
}

// func (*zip.FileHeader).FileInfo() os.FileInfo
func execmFileHeaderFileInfo(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zip.FileHeader).FileInfo()
	p.Ret(1, ret)
}

// func (*zip.FileHeader).ModTime() time.Time
func execmFileHeaderModTime(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zip.FileHeader).ModTime()
	p.Ret(1, ret)
}

// func (*zip.FileHeader).Mode() (mode os.FileMode)
func execmFileHeaderMode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zip.FileHeader).Mode()
	p.Ret(1, ret)
}

// func (*zip.FileHeader).SetModTime(t time.Time)
func execmFileHeaderSetModTime(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*zip.FileHeader).SetModTime(args[1].(time.Time))
}

// func (*zip.FileHeader).SetMode(mode os.FileMode)
func execmFileHeaderSetMode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*zip.FileHeader).SetMode(os.FileMode(args[1].(uint32)))
}

// func zip.FileInfoHeader(fi os.FileInfo) (*zip.FileHeader, error)
func execFileInfoHeader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := zip.FileInfoHeader(args[0].(os.FileInfo))
	p.Ret(1, ret, ret1)
}

// func zip.NewReader(r io.ReaderAt, size int64) (*zip.Reader, error)
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := zip.NewReader(args[0].(io.ReaderAt), args[1].(int64))
	p.Ret(2, ret, ret1)
}

// func zip.NewWriter(w io.Writer) *zip.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := zip.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func zip.OpenReader(name string) (*zip.ReadCloser, error)
func execOpenReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := zip.OpenReader(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*zip.ReadCloser).Close() error
func execmReadCloserClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zip.ReadCloser).Close()
	p.Ret(1, ret)
}

// func (*zip.Reader).RegisterDecompressor(method uint16, dcomp zip.Decompressor)
func execmReaderRegisterDecompressor(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*zip.Reader).RegisterDecompressor(args[1].(uint16), args[2].(zip.Decompressor))
}

// func zip.RegisterCompressor(method uint16, comp zip.Compressor)
func execRegisterCompressor(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	zip.RegisterCompressor(args[0].(uint16), args[1].(zip.Compressor))
}

// func zip.RegisterDecompressor(method uint16, dcomp zip.Decompressor)
func execRegisterDecompressor(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	zip.RegisterDecompressor(args[0].(uint16), args[1].(zip.Decompressor))
}

// func (*zip.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zip.Writer).Close()
	p.Ret(1, ret)
}

// func (*zip.Writer).Create(name string) (io.Writer, error)
func execmWriterCreate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*zip.Writer).Create(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*zip.Writer).CreateHeader(fh *zip.FileHeader) (io.Writer, error)
func execmWriterCreateHeader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*zip.Writer).CreateHeader(args[1].(*zip.FileHeader))
	p.Ret(2, ret, ret1)
}

// func (*zip.Writer).Flush() error
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*zip.Writer).Flush()
	p.Ret(1, ret)
}

// func (*zip.Writer).RegisterCompressor(method uint16, comp zip.Compressor)
func execmWriterRegisterCompressor(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*zip.Writer).RegisterCompressor(args[1].(uint16), args[2].(zip.Compressor))
}

// func (*zip.Writer).SetComment(comment string) error
func execmWriterSetComment(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*zip.Writer).SetComment(args[1].(string))
	p.Ret(2, ret)
}

// func (*zip.Writer).SetOffset(n int64)
func execmWriterSetOffset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*zip.Writer).SetOffset(args[1].(int64))
}

// I is a Go package instance.
var I = gop.NewGoPackage("archive/zip")

func init() {
	I.RegisterConsts(
		I.Const("Deflate", reflect.Uint16, zip.Deflate),
		I.Const("Store", reflect.Uint16, zip.Store),
	)
	I.RegisterVars(
		I.Var("ErrAlgorithm", &zip.ErrAlgorithm),
		I.Var("ErrChecksum", &zip.ErrChecksum),
		I.Var("ErrFormat", &zip.ErrFormat),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*zip.File)(nil))),
		I.Rtype(reflect.TypeOf((*zip.FileHeader)(nil))),
		I.Rtype(reflect.TypeOf((*zip.ReadCloser)(nil))),
		I.Rtype(reflect.TypeOf((*zip.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*zip.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*File).DataOffset", (*zip.File).DataOffset, execmFileDataOffset),
		I.Func("(*File).Open", (*zip.File).Open, execmFileOpen),
		I.Func("(*FileHeader).FileInfo", (*zip.FileHeader).FileInfo, execmFileHeaderFileInfo),
		I.Func("(*FileHeader).ModTime", (*zip.FileHeader).ModTime, execmFileHeaderModTime),
		I.Func("(*FileHeader).Mode", (*zip.FileHeader).Mode, execmFileHeaderMode),
		I.Func("(*FileHeader).SetModTime", (*zip.FileHeader).SetModTime, execmFileHeaderSetModTime),
		I.Func("(*FileHeader).SetMode", (*zip.FileHeader).SetMode, execmFileHeaderSetMode),
		I.Func("FileInfoHeader", zip.FileInfoHeader, execFileInfoHeader),
		I.Func("NewReader", zip.NewReader, execNewReader),
		I.Func("NewWriter", zip.NewWriter, execNewWriter),
		I.Func("OpenReader", zip.OpenReader, execOpenReader),
		I.Func("(*ReadCloser).Close", (*zip.ReadCloser).Close, execmReadCloserClose),
		I.Func("(*Reader).RegisterDecompressor", (*zip.Reader).RegisterDecompressor, execmReaderRegisterDecompressor),
		I.Func("RegisterCompressor", zip.RegisterCompressor, execRegisterCompressor),
		I.Func("RegisterDecompressor", zip.RegisterDecompressor, execRegisterDecompressor),
		I.Func("(*Writer).Close", (*zip.Writer).Close, execmWriterClose),
		I.Func("(*Writer).Create", (*zip.Writer).Create, execmWriterCreate),
		I.Func("(*Writer).CreateHeader", (*zip.Writer).CreateHeader, execmWriterCreateHeader),
		I.Func("(*Writer).Flush", (*zip.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).RegisterCompressor", (*zip.Writer).RegisterCompressor, execmWriterRegisterCompressor),
		I.Func("(*Writer).SetComment", (*zip.Writer).SetComment, execmWriterSetComment),
		I.Func("(*Writer).SetOffset", (*zip.Writer).SetOffset, execmWriterSetOffset),
	)
}
