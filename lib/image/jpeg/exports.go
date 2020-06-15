package jpeg

import (
	"image"
	"image/jpeg"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func jpeg.Decode(r io.Reader) (image.Image, error)
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := jpeg.Decode(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func jpeg.DecodeConfig(r io.Reader) (image.Config, error)
func execDecodeConfig(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := jpeg.DecodeConfig(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func jpeg.Encode(w io.Writer, m image.Image, o *jpeg.Options) error
func execEncode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := jpeg.Encode(args[0].(io.Writer), args[1].(image.Image), args[2].(*jpeg.Options))
	p.Ret(3, ret)
}

// func (jpeg.FormatError).Error() string
func execmFormatErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(jpeg.FormatError).Error()
	p.Ret(1, ret)
}

// func (jpeg.UnsupportedError).Error() string
func execmUnsupportedErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(jpeg.UnsupportedError).Error()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("image/jpeg")

func init() {
	I.RegisterConsts(
		I.Const("DefaultQuality", qspec.ConstUnboundInt, jpeg.DefaultQuality),
	)
	I.RegisterTypes(
		I.Type("FormatError", qspec.TyString),
		I.Rtype(reflect.TypeOf((*jpeg.Options)(nil))),
		I.Type("UnsupportedError", qspec.TyString),
	)
	I.RegisterFuncs(
		I.Func("Decode", jpeg.Decode, execDecode),
		I.Func("DecodeConfig", jpeg.DecodeConfig, execDecodeConfig),
		I.Func("Encode", jpeg.Encode, execEncode),
		I.Func("(FormatError).Error", (jpeg.FormatError).Error, execmFormatErrorError),
		I.Func("(UnsupportedError).Error", (jpeg.UnsupportedError).Error, execmUnsupportedErrorError),
	)
}
