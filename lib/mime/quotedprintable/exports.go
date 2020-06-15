package quotedprintable

import (
	"io"
	"mime/quotedprintable"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func quotedprintable.NewReader(r io.Reader) *quotedprintable.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := quotedprintable.NewReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func quotedprintable.NewWriter(w io.Writer) *quotedprintable.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := quotedprintable.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func (*quotedprintable.Reader).Read(p []byte) (n int, err error)
func execmReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*quotedprintable.Reader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*quotedprintable.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*quotedprintable.Writer).Close()
	p.Ret(1, ret)
}

// func (*quotedprintable.Writer).Write(p []byte) (n int, err error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*quotedprintable.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("mime/quotedprintable")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*quotedprintable.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*quotedprintable.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewReader", quotedprintable.NewReader, execNewReader),
		I.Func("NewWriter", quotedprintable.NewWriter, execNewWriter),
		I.Func("(*Reader).Read", (*quotedprintable.Reader).Read, execmReaderRead),
		I.Func("(*Writer).Close", (*quotedprintable.Writer).Close, execmWriterClose),
		I.Func("(*Writer).Write", (*quotedprintable.Writer).Write, execmWriterWrite),
	)
}
