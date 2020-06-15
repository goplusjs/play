package color

import (
	"image/color"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (color.Alpha).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmAlphaRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.Alpha).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (color.Alpha16).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmAlpha16RGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.Alpha16).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (color.CMYK).RGBA() (uint32, uint32, uint32, uint32)
func execmCMYKRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.CMYK).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func color.CMYKToRGB(c uint8, m uint8, y uint8, k uint8) (uint8, uint8, uint8)
func execCMYKToRGB(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := color.CMYKToRGB(args[0].(uint8), args[1].(uint8), args[2].(uint8), args[3].(uint8))
	p.Ret(4, ret, ret1, ret2)
}

// func (color.Gray).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmGrayRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.Gray).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (color.Gray16).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmGray16RGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.Gray16).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func color.ModelFunc(f func(color.Color) color.Color) color.Model
func execModelFunc(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := color.ModelFunc(args[0].(func(color.Color) color.Color))
	p.Ret(1, ret)
}

// func (color.NRGBA).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmNRGBARGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.NRGBA).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (color.NRGBA64).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmNRGBA64RGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.NRGBA64).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (color.NYCbCrA).RGBA() (uint32, uint32, uint32, uint32)
func execmNYCbCrARGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.NYCbCrA).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (color.Palette).Convert(c color.Color) color.Color
func execmPaletteConvert(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(color.Palette).Convert(args[1].(color.Color))
	p.Ret(2, ret)
}

// func (color.Palette).Index(c color.Color) int
func execmPaletteIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(color.Palette).Index(args[1].(color.Color))
	p.Ret(2, ret)
}

// func (color.RGBA).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmRGBARGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.RGBA).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (color.RGBA64).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmRGBA64RGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.RGBA64).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func color.RGBToCMYK(r uint8, g uint8, b uint8) (uint8, uint8, uint8, uint8)
func execRGBToCMYK(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2, ret3 := color.RGBToCMYK(args[0].(uint8), args[1].(uint8), args[2].(uint8))
	p.Ret(3, ret, ret1, ret2, ret3)
}

// func color.RGBToYCbCr(r uint8, g uint8, b uint8) (uint8, uint8, uint8)
func execRGBToYCbCr(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := color.RGBToYCbCr(args[0].(uint8), args[1].(uint8), args[2].(uint8))
	p.Ret(3, ret, ret1, ret2)
}

// func (color.YCbCr).RGBA() (uint32, uint32, uint32, uint32)
func execmYCbCrRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(color.YCbCr).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func color.YCbCrToRGB(y uint8, cb uint8, cr uint8) (uint8, uint8, uint8)
func execYCbCrToRGB(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := color.YCbCrToRGB(args[0].(uint8), args[1].(uint8), args[2].(uint8))
	p.Ret(3, ret, ret1, ret2)
}

// I is a Go package instance.
var I = gop.NewGoPackage("image/color")

func init() {
	I.RegisterVars(
		I.Var("Alpha16Model", &color.Alpha16Model),
		I.Var("AlphaModel", &color.AlphaModel),
		I.Var("Black", &color.Black),
		I.Var("CMYKModel", &color.CMYKModel),
		I.Var("Gray16Model", &color.Gray16Model),
		I.Var("GrayModel", &color.GrayModel),
		I.Var("NRGBA64Model", &color.NRGBA64Model),
		I.Var("NRGBAModel", &color.NRGBAModel),
		I.Var("NYCbCrAModel", &color.NYCbCrAModel),
		I.Var("Opaque", &color.Opaque),
		I.Var("RGBA64Model", &color.RGBA64Model),
		I.Var("RGBAModel", &color.RGBAModel),
		I.Var("Transparent", &color.Transparent),
		I.Var("White", &color.White),
		I.Var("YCbCrModel", &color.YCbCrModel),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*color.Alpha)(nil))),
		I.Rtype(reflect.TypeOf((*color.Alpha16)(nil))),
		I.Rtype(reflect.TypeOf((*color.CMYK)(nil))),
		I.Rtype(reflect.TypeOf((*color.Gray)(nil))),
		I.Rtype(reflect.TypeOf((*color.Gray16)(nil))),
		I.Rtype(reflect.TypeOf((*color.NRGBA)(nil))),
		I.Rtype(reflect.TypeOf((*color.NRGBA64)(nil))),
		I.Rtype(reflect.TypeOf((*color.NYCbCrA)(nil))),
		I.Rtype(reflect.TypeOf((*color.RGBA)(nil))),
		I.Rtype(reflect.TypeOf((*color.RGBA64)(nil))),
		I.Rtype(reflect.TypeOf((*color.YCbCr)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(Alpha).RGBA", (color.Alpha).RGBA, execmAlphaRGBA),
		I.Func("(Alpha16).RGBA", (color.Alpha16).RGBA, execmAlpha16RGBA),
		I.Func("(CMYK).RGBA", (color.CMYK).RGBA, execmCMYKRGBA),
		I.Func("CMYKToRGB", color.CMYKToRGB, execCMYKToRGB),
		I.Func("(Gray).RGBA", (color.Gray).RGBA, execmGrayRGBA),
		I.Func("(Gray16).RGBA", (color.Gray16).RGBA, execmGray16RGBA),
		I.Func("ModelFunc", color.ModelFunc, execModelFunc),
		I.Func("(NRGBA).RGBA", (color.NRGBA).RGBA, execmNRGBARGBA),
		I.Func("(NRGBA64).RGBA", (color.NRGBA64).RGBA, execmNRGBA64RGBA),
		I.Func("(NYCbCrA).RGBA", (color.NYCbCrA).RGBA, execmNYCbCrARGBA),
		I.Func("(Palette).Convert", (color.Palette).Convert, execmPaletteConvert),
		I.Func("(Palette).Index", (color.Palette).Index, execmPaletteIndex),
		I.Func("(RGBA).RGBA", (color.RGBA).RGBA, execmRGBARGBA),
		I.Func("(RGBA64).RGBA", (color.RGBA64).RGBA, execmRGBA64RGBA),
		I.Func("RGBToCMYK", color.RGBToCMYK, execRGBToCMYK),
		I.Func("RGBToYCbCr", color.RGBToYCbCr, execRGBToYCbCr),
		I.Func("(YCbCr).RGBA", (color.YCbCr).RGBA, execmYCbCrRGBA),
		I.Func("YCbCrToRGB", color.YCbCrToRGB, execYCbCrToRGB),
	)
}
