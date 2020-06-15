package multipart

import (
	"io"
	"mime/multipart"
	"net/textproto"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*multipart.FileHeader).Open() (multipart.File, error)
func execmFileHeaderOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*multipart.FileHeader).Open()
	p.Ret(1, ret, ret1)
}

// func (*multipart.Form).RemoveAll() error
func execmFormRemoveAll(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*multipart.Form).RemoveAll()
	p.Ret(1, ret)
}

// func multipart.NewReader(r io.Reader, boundary string) *multipart.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := multipart.NewReader(args[0].(io.Reader), args[1].(string))
	p.Ret(2, ret)
}

// func multipart.NewWriter(w io.Writer) *multipart.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := multipart.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func (*multipart.Part).Close() error
func execmPartClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*multipart.Part).Close()
	p.Ret(1, ret)
}

// func (*multipart.Part).FileName() string
func execmPartFileName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*multipart.Part).FileName()
	p.Ret(1, ret)
}

// func (*multipart.Part).FormName() string
func execmPartFormName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*multipart.Part).FormName()
	p.Ret(1, ret)
}

// func (*multipart.Part).Read(d []byte) (n int, err error)
func execmPartRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*multipart.Part).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*multipart.Reader).NextPart() (*multipart.Part, error)
func execmReaderNextPart(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*multipart.Reader).NextPart()
	p.Ret(1, ret, ret1)
}

// func (*multipart.Reader).NextRawPart() (*multipart.Part, error)
func execmReaderNextRawPart(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*multipart.Reader).NextRawPart()
	p.Ret(1, ret, ret1)
}

// func (*multipart.Reader).ReadForm(maxMemory int64) (*multipart.Form, error)
func execmReaderReadForm(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*multipart.Reader).ReadForm(args[1].(int64))
	p.Ret(2, ret, ret1)
}

// func (*multipart.Writer).Boundary() string
func execmWriterBoundary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*multipart.Writer).Boundary()
	p.Ret(1, ret)
}

// func (*multipart.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*multipart.Writer).Close()
	p.Ret(1, ret)
}

// func (*multipart.Writer).CreateFormField(fieldname string) (io.Writer, error)
func execmWriterCreateFormField(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*multipart.Writer).CreateFormField(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*multipart.Writer).CreateFormFile(fieldname string, filename string) (io.Writer, error)
func execmWriterCreateFormFile(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*multipart.Writer).CreateFormFile(args[1].(string), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*multipart.Writer).CreatePart(header textproto.MIMEHeader) (io.Writer, error)
func execmWriterCreatePart(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*multipart.Writer).CreatePart(args[1].(textproto.MIMEHeader))
	p.Ret(2, ret, ret1)
}

// func (*multipart.Writer).FormDataContentType() string
func execmWriterFormDataContentType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*multipart.Writer).FormDataContentType()
	p.Ret(1, ret)
}

// func (*multipart.Writer).SetBoundary(boundary string) error
func execmWriterSetBoundary(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*multipart.Writer).SetBoundary(args[1].(string))
	p.Ret(2, ret)
}

// func (*multipart.Writer).WriteField(fieldname string, value string) error
func execmWriterWriteField(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*multipart.Writer).WriteField(args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("mime/multipart")

func init() {
	I.RegisterVars(
		I.Var("ErrMessageTooLarge", &multipart.ErrMessageTooLarge),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*multipart.FileHeader)(nil))),
		I.Rtype(reflect.TypeOf((*multipart.Form)(nil))),
		I.Rtype(reflect.TypeOf((*multipart.Part)(nil))),
		I.Rtype(reflect.TypeOf((*multipart.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*multipart.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*FileHeader).Open", (*multipart.FileHeader).Open, execmFileHeaderOpen),
		I.Func("(*Form).RemoveAll", (*multipart.Form).RemoveAll, execmFormRemoveAll),
		I.Func("NewReader", multipart.NewReader, execNewReader),
		I.Func("NewWriter", multipart.NewWriter, execNewWriter),
		I.Func("(*Part).Close", (*multipart.Part).Close, execmPartClose),
		I.Func("(*Part).FileName", (*multipart.Part).FileName, execmPartFileName),
		I.Func("(*Part).FormName", (*multipart.Part).FormName, execmPartFormName),
		I.Func("(*Part).Read", (*multipart.Part).Read, execmPartRead),
		I.Func("(*Reader).NextPart", (*multipart.Reader).NextPart, execmReaderNextPart),
		I.Func("(*Reader).NextRawPart", (*multipart.Reader).NextRawPart, execmReaderNextRawPart),
		I.Func("(*Reader).ReadForm", (*multipart.Reader).ReadForm, execmReaderReadForm),
		I.Func("(*Writer).Boundary", (*multipart.Writer).Boundary, execmWriterBoundary),
		I.Func("(*Writer).Close", (*multipart.Writer).Close, execmWriterClose),
		I.Func("(*Writer).CreateFormField", (*multipart.Writer).CreateFormField, execmWriterCreateFormField),
		I.Func("(*Writer).CreateFormFile", (*multipart.Writer).CreateFormFile, execmWriterCreateFormFile),
		I.Func("(*Writer).CreatePart", (*multipart.Writer).CreatePart, execmWriterCreatePart),
		I.Func("(*Writer).FormDataContentType", (*multipart.Writer).FormDataContentType, execmWriterFormDataContentType),
		I.Func("(*Writer).SetBoundary", (*multipart.Writer).SetBoundary, execmWriterSetBoundary),
		I.Func("(*Writer).WriteField", (*multipart.Writer).WriteField, execmWriterWriteField),
	)
}
