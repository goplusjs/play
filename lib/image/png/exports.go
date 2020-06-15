package png

import (
	"image"
	"image/png"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func png.Decode(r io.Reader) (image.Image, error)
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := png.Decode(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func png.DecodeConfig(r io.Reader) (image.Config, error)
func execDecodeConfig(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := png.DecodeConfig(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func png.Encode(w io.Writer, m image.Image) error
func execEncode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := png.Encode(args[0].(io.Writer), args[1].(image.Image))
	p.Ret(2, ret)
}

// func (*png.Encoder).Encode(w io.Writer, m image.Image) error
func execmEncoderEncode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*png.Encoder).Encode(args[1].(io.Writer), args[2].(image.Image))
	p.Ret(3, ret)
}

// func (png.FormatError).Error() string
func execmFormatErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(png.FormatError).Error()
	p.Ret(1, ret)
}

// func (png.UnsupportedError).Error() string
func execmUnsupportedErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(png.UnsupportedError).Error()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("image/png")

func init() {
	I.RegisterConsts(
		I.Const("BestCompression", reflect.Int, png.BestCompression),
		I.Const("BestSpeed", reflect.Int, png.BestSpeed),
		I.Const("DefaultCompression", reflect.Int, png.DefaultCompression),
		I.Const("NoCompression", reflect.Int, png.NoCompression),
	)
	I.RegisterTypes(
		I.Type("CompressionLevel", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*png.Encoder)(nil))),
		I.Rtype(reflect.TypeOf((*png.EncoderBuffer)(nil))),
		I.Type("FormatError", qspec.TyString),
		I.Type("UnsupportedError", qspec.TyString),
	)
	I.RegisterFuncs(
		I.Func("Decode", png.Decode, execDecode),
		I.Func("DecodeConfig", png.DecodeConfig, execDecodeConfig),
		I.Func("Encode", png.Encode, execEncode),
		I.Func("(*Encoder).Encode", (*png.Encoder).Encode, execmEncoderEncode),
		I.Func("(FormatError).Error", (png.FormatError).Error, execmFormatErrorError),
		I.Func("(UnsupportedError).Error", (png.UnsupportedError).Error, execmUnsupportedErrorError),
	)
}
