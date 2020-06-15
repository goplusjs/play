package gif

import (
	"image"
	"image/gif"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func gif.Decode(r io.Reader) (image.Image, error)
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := gif.Decode(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func gif.DecodeAll(r io.Reader) (*gif.GIF, error)
func execDecodeAll(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := gif.DecodeAll(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func gif.DecodeConfig(r io.Reader) (image.Config, error)
func execDecodeConfig(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := gif.DecodeConfig(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func gif.Encode(w io.Writer, m image.Image, o *gif.Options) error
func execEncode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := gif.Encode(args[0].(io.Writer), args[1].(image.Image), args[2].(*gif.Options))
	p.Ret(3, ret)
}

// func gif.EncodeAll(w io.Writer, g *gif.GIF) error
func execEncodeAll(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := gif.EncodeAll(args[0].(io.Writer), args[1].(*gif.GIF))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("image/gif")

func init() {
	I.RegisterConsts(
		I.Const("DisposalBackground", qspec.ConstUnboundInt, gif.DisposalBackground),
		I.Const("DisposalNone", qspec.ConstUnboundInt, gif.DisposalNone),
		I.Const("DisposalPrevious", qspec.ConstUnboundInt, gif.DisposalPrevious),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*gif.GIF)(nil))),
		I.Rtype(reflect.TypeOf((*gif.Options)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Decode", gif.Decode, execDecode),
		I.Func("DecodeAll", gif.DecodeAll, execDecodeAll),
		I.Func("DecodeConfig", gif.DecodeConfig, execDecodeConfig),
		I.Func("Encode", gif.Encode, execEncode),
		I.Func("EncodeAll", gif.EncodeAll, execEncodeAll),
	)
}
