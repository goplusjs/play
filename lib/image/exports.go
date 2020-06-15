package image

import (
	"image"
	"image/color"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*image.Alpha).AlphaAt(x int, y int) color.Alpha
func execmAlphaAlphaAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Alpha).AlphaAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Alpha).At(x int, y int) color.Color
func execmAlphaAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Alpha).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Alpha).Bounds() image.Rectangle
func execmAlphaBounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Alpha).Bounds()
	p.Ret(1, ret)
}

// func (*image.Alpha).ColorModel() color.Model
func execmAlphaColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Alpha).ColorModel()
	p.Ret(1, ret)
}

// func (*image.Alpha).Opaque() bool
func execmAlphaOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Alpha).Opaque()
	p.Ret(1, ret)
}

// func (*image.Alpha).PixOffset(x int, y int) int
func execmAlphaPixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Alpha).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Alpha).Set(x int, y int, c color.Color)
func execmAlphaSet(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Alpha).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.Alpha).SetAlpha(x int, y int, c color.Alpha)
func execmAlphaSetAlpha(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Alpha).SetAlpha(args[1].(int), args[2].(int), args[3].(color.Alpha))
}

// func (*image.Alpha).SubImage(r image.Rectangle) image.Image
func execmAlphaSubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.Alpha).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.Alpha16).Alpha16At(x int, y int) color.Alpha16
func execmAlpha16Alpha16At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Alpha16).Alpha16At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Alpha16).At(x int, y int) color.Color
func execmAlpha16At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Alpha16).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Alpha16).Bounds() image.Rectangle
func execmAlpha16Bounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Alpha16).Bounds()
	p.Ret(1, ret)
}

// func (*image.Alpha16).ColorModel() color.Model
func execmAlpha16ColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Alpha16).ColorModel()
	p.Ret(1, ret)
}

// func (*image.Alpha16).Opaque() bool
func execmAlpha16Opaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Alpha16).Opaque()
	p.Ret(1, ret)
}

// func (*image.Alpha16).PixOffset(x int, y int) int
func execmAlpha16PixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Alpha16).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Alpha16).Set(x int, y int, c color.Color)
func execmAlpha16Set(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Alpha16).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.Alpha16).SetAlpha16(x int, y int, c color.Alpha16)
func execmAlpha16SetAlpha16(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Alpha16).SetAlpha16(args[1].(int), args[2].(int), args[3].(color.Alpha16))
}

// func (*image.Alpha16).SubImage(r image.Rectangle) image.Image
func execmAlpha16SubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.Alpha16).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.CMYK).At(x int, y int) color.Color
func execmCMYKAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.CMYK).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.CMYK).Bounds() image.Rectangle
func execmCMYKBounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.CMYK).Bounds()
	p.Ret(1, ret)
}

// func (*image.CMYK).CMYKAt(x int, y int) color.CMYK
func execmCMYKCMYKAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.CMYK).CMYKAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.CMYK).ColorModel() color.Model
func execmCMYKColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.CMYK).ColorModel()
	p.Ret(1, ret)
}

// func (*image.CMYK).Opaque() bool
func execmCMYKOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.CMYK).Opaque()
	p.Ret(1, ret)
}

// func (*image.CMYK).PixOffset(x int, y int) int
func execmCMYKPixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.CMYK).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.CMYK).Set(x int, y int, c color.Color)
func execmCMYKSet(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.CMYK).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.CMYK).SetCMYK(x int, y int, c color.CMYK)
func execmCMYKSetCMYK(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.CMYK).SetCMYK(args[1].(int), args[2].(int), args[3].(color.CMYK))
}

// func (*image.CMYK).SubImage(r image.Rectangle) image.Image
func execmCMYKSubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.CMYK).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func image.Decode(r io.Reader) (image.Image, string, error)
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := image.Decode(args[0].(io.Reader))
	p.Ret(1, ret, ret1, ret2)
}

// func image.DecodeConfig(r io.Reader) (image.Config, string, error)
func execDecodeConfig(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := image.DecodeConfig(args[0].(io.Reader))
	p.Ret(1, ret, ret1, ret2)
}

// func (*image.Gray).At(x int, y int) color.Color
func execmGrayAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Gray).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Gray).Bounds() image.Rectangle
func execmGrayBounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Gray).Bounds()
	p.Ret(1, ret)
}

// func (*image.Gray).ColorModel() color.Model
func execmGrayColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Gray).ColorModel()
	p.Ret(1, ret)
}

// func (*image.Gray).GrayAt(x int, y int) color.Gray
func execmGrayGrayAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Gray).GrayAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Gray).Opaque() bool
func execmGrayOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Gray).Opaque()
	p.Ret(1, ret)
}

// func (*image.Gray).PixOffset(x int, y int) int
func execmGrayPixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Gray).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Gray).Set(x int, y int, c color.Color)
func execmGraySet(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Gray).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.Gray).SetGray(x int, y int, c color.Gray)
func execmGraySetGray(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Gray).SetGray(args[1].(int), args[2].(int), args[3].(color.Gray))
}

// func (*image.Gray).SubImage(r image.Rectangle) image.Image
func execmGraySubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.Gray).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.Gray16).At(x int, y int) color.Color
func execmGray16At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Gray16).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Gray16).Bounds() image.Rectangle
func execmGray16Bounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Gray16).Bounds()
	p.Ret(1, ret)
}

// func (*image.Gray16).ColorModel() color.Model
func execmGray16ColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Gray16).ColorModel()
	p.Ret(1, ret)
}

// func (*image.Gray16).Gray16At(x int, y int) color.Gray16
func execmGray16Gray16At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Gray16).Gray16At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Gray16).Opaque() bool
func execmGray16Opaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Gray16).Opaque()
	p.Ret(1, ret)
}

// func (*image.Gray16).PixOffset(x int, y int) int
func execmGray16PixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Gray16).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Gray16).Set(x int, y int, c color.Color)
func execmGray16Set(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Gray16).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.Gray16).SetGray16(x int, y int, c color.Gray16)
func execmGray16SetGray16(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Gray16).SetGray16(args[1].(int), args[2].(int), args[3].(color.Gray16))
}

// func (*image.Gray16).SubImage(r image.Rectangle) image.Image
func execmGray16SubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.Gray16).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.NRGBA).At(x int, y int) color.Color
func execmNRGBAAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NRGBA).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NRGBA).Bounds() image.Rectangle
func execmNRGBABounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NRGBA).Bounds()
	p.Ret(1, ret)
}

// func (*image.NRGBA).ColorModel() color.Model
func execmNRGBAColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NRGBA).ColorModel()
	p.Ret(1, ret)
}

// func (*image.NRGBA).NRGBAAt(x int, y int) color.NRGBA
func execmNRGBANRGBAAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NRGBA).NRGBAAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NRGBA).Opaque() bool
func execmNRGBAOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NRGBA).Opaque()
	p.Ret(1, ret)
}

// func (*image.NRGBA).PixOffset(x int, y int) int
func execmNRGBAPixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NRGBA).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NRGBA).Set(x int, y int, c color.Color)
func execmNRGBASet(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.NRGBA).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.NRGBA).SetNRGBA(x int, y int, c color.NRGBA)
func execmNRGBASetNRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.NRGBA).SetNRGBA(args[1].(int), args[2].(int), args[3].(color.NRGBA))
}

// func (*image.NRGBA).SubImage(r image.Rectangle) image.Image
func execmNRGBASubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.NRGBA).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.NRGBA64).At(x int, y int) color.Color
func execmNRGBA64At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NRGBA64).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NRGBA64).Bounds() image.Rectangle
func execmNRGBA64Bounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NRGBA64).Bounds()
	p.Ret(1, ret)
}

// func (*image.NRGBA64).ColorModel() color.Model
func execmNRGBA64ColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NRGBA64).ColorModel()
	p.Ret(1, ret)
}

// func (*image.NRGBA64).NRGBA64At(x int, y int) color.NRGBA64
func execmNRGBA64NRGBA64At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NRGBA64).NRGBA64At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NRGBA64).Opaque() bool
func execmNRGBA64Opaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NRGBA64).Opaque()
	p.Ret(1, ret)
}

// func (*image.NRGBA64).PixOffset(x int, y int) int
func execmNRGBA64PixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NRGBA64).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NRGBA64).Set(x int, y int, c color.Color)
func execmNRGBA64Set(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.NRGBA64).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.NRGBA64).SetNRGBA64(x int, y int, c color.NRGBA64)
func execmNRGBA64SetNRGBA64(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.NRGBA64).SetNRGBA64(args[1].(int), args[2].(int), args[3].(color.NRGBA64))
}

// func (*image.NRGBA64).SubImage(r image.Rectangle) image.Image
func execmNRGBA64SubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.NRGBA64).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.NYCbCrA).AOffset(x int, y int) int
func execmNYCbCrAAOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NYCbCrA).AOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NYCbCrA).At(x int, y int) color.Color
func execmNYCbCrAAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NYCbCrA).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NYCbCrA).ColorModel() color.Model
func execmNYCbCrAColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NYCbCrA).ColorModel()
	p.Ret(1, ret)
}

// func (*image.NYCbCrA).NYCbCrAAt(x int, y int) color.NYCbCrA
func execmNYCbCrANYCbCrAAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.NYCbCrA).NYCbCrAAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.NYCbCrA).Opaque() bool
func execmNYCbCrAOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.NYCbCrA).Opaque()
	p.Ret(1, ret)
}

// func (*image.NYCbCrA).SubImage(r image.Rectangle) image.Image
func execmNYCbCrASubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.NYCbCrA).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func image.NewAlpha(r image.Rectangle) *image.Alpha
func execNewAlpha(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewAlpha(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewAlpha16(r image.Rectangle) *image.Alpha16
func execNewAlpha16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewAlpha16(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewCMYK(r image.Rectangle) *image.CMYK
func execNewCMYK(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewCMYK(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewGray(r image.Rectangle) *image.Gray
func execNewGray(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewGray(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewGray16(r image.Rectangle) *image.Gray16
func execNewGray16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewGray16(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewNRGBA(r image.Rectangle) *image.NRGBA
func execNewNRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewNRGBA(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewNRGBA64(r image.Rectangle) *image.NRGBA64
func execNewNRGBA64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewNRGBA64(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewNYCbCrA(r image.Rectangle, subsampleRatio image.YCbCrSubsampleRatio) *image.NYCbCrA
func execNewNYCbCrA(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := image.NewNYCbCrA(args[0].(image.Rectangle), image.YCbCrSubsampleRatio(args[1].(int)))
	p.Ret(2, ret)
}

// func image.NewPaletted(r image.Rectangle, p color.Palette) *image.Paletted
func execNewPaletted(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := image.NewPaletted(args[0].(image.Rectangle), args[1].(color.Palette))
	p.Ret(2, ret)
}

// func image.NewRGBA(r image.Rectangle) *image.RGBA
func execNewRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewRGBA(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewRGBA64(r image.Rectangle) *image.RGBA64
func execNewRGBA64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewRGBA64(args[0].(image.Rectangle))
	p.Ret(1, ret)
}

// func image.NewUniform(c color.Color) *image.Uniform
func execNewUniform(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := image.NewUniform(args[0].(color.Color))
	p.Ret(1, ret)
}

// func image.NewYCbCr(r image.Rectangle, subsampleRatio image.YCbCrSubsampleRatio) *image.YCbCr
func execNewYCbCr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := image.NewYCbCr(args[0].(image.Rectangle), image.YCbCrSubsampleRatio(args[1].(int)))
	p.Ret(2, ret)
}

// func (*image.Paletted).At(x int, y int) color.Color
func execmPalettedAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Paletted).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Paletted).Bounds() image.Rectangle
func execmPalettedBounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Paletted).Bounds()
	p.Ret(1, ret)
}

// func (*image.Paletted).ColorIndexAt(x int, y int) uint8
func execmPalettedColorIndexAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Paletted).ColorIndexAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Paletted).ColorModel() color.Model
func execmPalettedColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Paletted).ColorModel()
	p.Ret(1, ret)
}

// func (*image.Paletted).Opaque() bool
func execmPalettedOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Paletted).Opaque()
	p.Ret(1, ret)
}

// func (*image.Paletted).PixOffset(x int, y int) int
func execmPalettedPixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Paletted).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Paletted).Set(x int, y int, c color.Color)
func execmPalettedSet(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Paletted).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.Paletted).SetColorIndex(x int, y int, index uint8)
func execmPalettedSetColorIndex(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.Paletted).SetColorIndex(args[1].(int), args[2].(int), args[3].(uint8))
}

// func (*image.Paletted).SubImage(r image.Rectangle) image.Image
func execmPalettedSubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.Paletted).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (image.Point).Add(q image.Point) image.Point
func execmPointAdd(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Point).Add(args[1].(image.Point))
	p.Ret(2, ret)
}

// func (image.Point).Div(k int) image.Point
func execmPointDiv(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Point).Div(args[1].(int))
	p.Ret(2, ret)
}

// func (image.Point).Eq(q image.Point) bool
func execmPointEq(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Point).Eq(args[1].(image.Point))
	p.Ret(2, ret)
}

// func (image.Point).In(r image.Rectangle) bool
func execmPointIn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Point).In(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (image.Point).Mod(r image.Rectangle) image.Point
func execmPointMod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Point).Mod(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (image.Point).Mul(k int) image.Point
func execmPointMul(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Point).Mul(args[1].(int))
	p.Ret(2, ret)
}

// func (image.Point).String() string
func execmPointString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Point).String()
	p.Ret(1, ret)
}

// func (image.Point).Sub(q image.Point) image.Point
func execmPointSub(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Point).Sub(args[1].(image.Point))
	p.Ret(2, ret)
}

// func image.Pt(X int, Y int) image.Point
func execPt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := image.Pt(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func (*image.RGBA).At(x int, y int) color.Color
func execmRGBAAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.RGBA).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.RGBA).Bounds() image.Rectangle
func execmRGBABounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.RGBA).Bounds()
	p.Ret(1, ret)
}

// func (*image.RGBA).ColorModel() color.Model
func execmRGBAColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.RGBA).ColorModel()
	p.Ret(1, ret)
}

// func (*image.RGBA).Opaque() bool
func execmRGBAOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.RGBA).Opaque()
	p.Ret(1, ret)
}

// func (*image.RGBA).PixOffset(x int, y int) int
func execmRGBAPixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.RGBA).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.RGBA).RGBAAt(x int, y int) color.RGBA
func execmRGBARGBAAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.RGBA).RGBAAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.RGBA).Set(x int, y int, c color.Color)
func execmRGBASet(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.RGBA).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.RGBA).SetRGBA(x int, y int, c color.RGBA)
func execmRGBASetRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.RGBA).SetRGBA(args[1].(int), args[2].(int), args[3].(color.RGBA))
}

// func (*image.RGBA).SubImage(r image.Rectangle) image.Image
func execmRGBASubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.RGBA).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.RGBA64).At(x int, y int) color.Color
func execmRGBA64At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.RGBA64).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.RGBA64).Bounds() image.Rectangle
func execmRGBA64Bounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.RGBA64).Bounds()
	p.Ret(1, ret)
}

// func (*image.RGBA64).ColorModel() color.Model
func execmRGBA64ColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.RGBA64).ColorModel()
	p.Ret(1, ret)
}

// func (*image.RGBA64).Opaque() bool
func execmRGBA64Opaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.RGBA64).Opaque()
	p.Ret(1, ret)
}

// func (*image.RGBA64).PixOffset(x int, y int) int
func execmRGBA64PixOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.RGBA64).PixOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.RGBA64).RGBA64At(x int, y int) color.RGBA64
func execmRGBA64RGBA64At(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.RGBA64).RGBA64At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.RGBA64).Set(x int, y int, c color.Color)
func execmRGBA64Set(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.RGBA64).Set(args[1].(int), args[2].(int), args[3].(color.Color))
}

// func (*image.RGBA64).SetRGBA64(x int, y int, c color.RGBA64)
func execmRGBA64SetRGBA64(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*image.RGBA64).SetRGBA64(args[1].(int), args[2].(int), args[3].(color.RGBA64))
}

// func (*image.RGBA64).SubImage(r image.Rectangle) image.Image
func execmRGBA64SubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.RGBA64).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func image.Rect(x0 int, y0 int, x1 int, y1 int) image.Rectangle
func execRect(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := image.Rect(args[0].(int), args[1].(int), args[2].(int), args[3].(int))
	p.Ret(4, ret)
}

// func (image.Rectangle).Add(p image.Point) image.Rectangle
func execmRectangleAdd(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).Add(args[1].(image.Point))
	p.Ret(2, ret)
}

// func (image.Rectangle).At(x int, y int) color.Color
func execmRectangleAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(image.Rectangle).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (image.Rectangle).Bounds() image.Rectangle
func execmRectangleBounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).Bounds()
	p.Ret(1, ret)
}

// func (image.Rectangle).Canon() image.Rectangle
func execmRectangleCanon(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).Canon()
	p.Ret(1, ret)
}

// func (image.Rectangle).ColorModel() color.Model
func execmRectangleColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).ColorModel()
	p.Ret(1, ret)
}

// func (image.Rectangle).Dx() int
func execmRectangleDx(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).Dx()
	p.Ret(1, ret)
}

// func (image.Rectangle).Dy() int
func execmRectangleDy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).Dy()
	p.Ret(1, ret)
}

// func (image.Rectangle).Empty() bool
func execmRectangleEmpty(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).Empty()
	p.Ret(1, ret)
}

// func (image.Rectangle).Eq(s image.Rectangle) bool
func execmRectangleEq(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).Eq(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (image.Rectangle).In(s image.Rectangle) bool
func execmRectangleIn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).In(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (image.Rectangle).Inset(n int) image.Rectangle
func execmRectangleInset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).Inset(args[1].(int))
	p.Ret(2, ret)
}

// func (image.Rectangle).Intersect(s image.Rectangle) image.Rectangle
func execmRectangleIntersect(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).Intersect(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (image.Rectangle).Overlaps(s image.Rectangle) bool
func execmRectangleOverlaps(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).Overlaps(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (image.Rectangle).Size() image.Point
func execmRectangleSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).Size()
	p.Ret(1, ret)
}

// func (image.Rectangle).String() string
func execmRectangleString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.Rectangle).String()
	p.Ret(1, ret)
}

// func (image.Rectangle).Sub(p image.Point) image.Rectangle
func execmRectangleSub(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).Sub(args[1].(image.Point))
	p.Ret(2, ret)
}

// func (image.Rectangle).Union(s image.Rectangle) image.Rectangle
func execmRectangleUnion(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(image.Rectangle).Union(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func image.RegisterFormat(name string, magic string, decode func(io.Reader) (image.Image, error), decodeConfig func(io.Reader) (image.Config, error))
func execRegisterFormat(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	image.RegisterFormat(args[0].(string), args[1].(string), args[2].(func(io.Reader) (image.Image, error)), args[3].(func(io.Reader) (image.Config, error)))
}

// func (*image.Uniform).At(x int, y int) color.Color
func execmUniformAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.Uniform).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.Uniform).Bounds() image.Rectangle
func execmUniformBounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Uniform).Bounds()
	p.Ret(1, ret)
}

// func (*image.Uniform).ColorModel() color.Model
func execmUniformColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Uniform).ColorModel()
	p.Ret(1, ret)
}

// func (*image.Uniform).Convert(color.Color) color.Color
func execmUniformConvert(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.Uniform).Convert(args[1].(color.Color))
	p.Ret(2, ret)
}

// func (*image.Uniform).Opaque() bool
func execmUniformOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.Uniform).Opaque()
	p.Ret(1, ret)
}

// func (*image.Uniform).RGBA() (r uint32, g uint32, b uint32, a uint32)
func execmUniformRGBA(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := args[0].(*image.Uniform).RGBA()
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func (*image.YCbCr).At(x int, y int) color.Color
func execmYCbCrAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.YCbCr).At(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.YCbCr).Bounds() image.Rectangle
func execmYCbCrBounds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.YCbCr).Bounds()
	p.Ret(1, ret)
}

// func (*image.YCbCr).COffset(x int, y int) int
func execmYCbCrCOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.YCbCr).COffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.YCbCr).ColorModel() color.Model
func execmYCbCrColorModel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.YCbCr).ColorModel()
	p.Ret(1, ret)
}

// func (*image.YCbCr).Opaque() bool
func execmYCbCrOpaque(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*image.YCbCr).Opaque()
	p.Ret(1, ret)
}

// func (*image.YCbCr).SubImage(r image.Rectangle) image.Image
func execmYCbCrSubImage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*image.YCbCr).SubImage(args[1].(image.Rectangle))
	p.Ret(2, ret)
}

// func (*image.YCbCr).YCbCrAt(x int, y int) color.YCbCr
func execmYCbCrYCbCrAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.YCbCr).YCbCrAt(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*image.YCbCr).YOffset(x int, y int) int
func execmYCbCrYOffset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*image.YCbCr).YOffset(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (image.YCbCrSubsampleRatio).String() string
func execmYCbCrSubsampleRatioString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(image.YCbCrSubsampleRatio).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("image")

func init() {
	I.RegisterConsts(
		I.Const("YCbCrSubsampleRatio410", reflect.Int, image.YCbCrSubsampleRatio410),
		I.Const("YCbCrSubsampleRatio411", reflect.Int, image.YCbCrSubsampleRatio411),
		I.Const("YCbCrSubsampleRatio420", reflect.Int, image.YCbCrSubsampleRatio420),
		I.Const("YCbCrSubsampleRatio422", reflect.Int, image.YCbCrSubsampleRatio422),
		I.Const("YCbCrSubsampleRatio440", reflect.Int, image.YCbCrSubsampleRatio440),
		I.Const("YCbCrSubsampleRatio444", reflect.Int, image.YCbCrSubsampleRatio444),
	)
	I.RegisterVars(
		I.Var("Black", &image.Black),
		I.Var("ErrFormat", &image.ErrFormat),
		I.Var("Opaque", &image.Opaque),
		I.Var("Transparent", &image.Transparent),
		I.Var("White", &image.White),
		I.Var("ZP", &image.ZP),
		I.Var("ZR", &image.ZR),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*image.Alpha)(nil))),
		I.Rtype(reflect.TypeOf((*image.Alpha16)(nil))),
		I.Rtype(reflect.TypeOf((*image.CMYK)(nil))),
		I.Rtype(reflect.TypeOf((*image.Config)(nil))),
		I.Rtype(reflect.TypeOf((*image.Gray)(nil))),
		I.Rtype(reflect.TypeOf((*image.Gray16)(nil))),
		I.Rtype(reflect.TypeOf((*image.NRGBA)(nil))),
		I.Rtype(reflect.TypeOf((*image.NRGBA64)(nil))),
		I.Rtype(reflect.TypeOf((*image.NYCbCrA)(nil))),
		I.Rtype(reflect.TypeOf((*image.Paletted)(nil))),
		I.Rtype(reflect.TypeOf((*image.Point)(nil))),
		I.Rtype(reflect.TypeOf((*image.RGBA)(nil))),
		I.Rtype(reflect.TypeOf((*image.RGBA64)(nil))),
		I.Rtype(reflect.TypeOf((*image.Rectangle)(nil))),
		I.Rtype(reflect.TypeOf((*image.Uniform)(nil))),
		I.Rtype(reflect.TypeOf((*image.YCbCr)(nil))),
		I.Type("YCbCrSubsampleRatio", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("(*Alpha).AlphaAt", (*image.Alpha).AlphaAt, execmAlphaAlphaAt),
		I.Func("(*Alpha).At", (*image.Alpha).At, execmAlphaAt),
		I.Func("(*Alpha).Bounds", (*image.Alpha).Bounds, execmAlphaBounds),
		I.Func("(*Alpha).ColorModel", (*image.Alpha).ColorModel, execmAlphaColorModel),
		I.Func("(*Alpha).Opaque", (*image.Alpha).Opaque, execmAlphaOpaque),
		I.Func("(*Alpha).PixOffset", (*image.Alpha).PixOffset, execmAlphaPixOffset),
		I.Func("(*Alpha).Set", (*image.Alpha).Set, execmAlphaSet),
		I.Func("(*Alpha).SetAlpha", (*image.Alpha).SetAlpha, execmAlphaSetAlpha),
		I.Func("(*Alpha).SubImage", (*image.Alpha).SubImage, execmAlphaSubImage),
		I.Func("(*Alpha16).Alpha16At", (*image.Alpha16).Alpha16At, execmAlpha16Alpha16At),
		I.Func("(*Alpha16).At", (*image.Alpha16).At, execmAlpha16At),
		I.Func("(*Alpha16).Bounds", (*image.Alpha16).Bounds, execmAlpha16Bounds),
		I.Func("(*Alpha16).ColorModel", (*image.Alpha16).ColorModel, execmAlpha16ColorModel),
		I.Func("(*Alpha16).Opaque", (*image.Alpha16).Opaque, execmAlpha16Opaque),
		I.Func("(*Alpha16).PixOffset", (*image.Alpha16).PixOffset, execmAlpha16PixOffset),
		I.Func("(*Alpha16).Set", (*image.Alpha16).Set, execmAlpha16Set),
		I.Func("(*Alpha16).SetAlpha16", (*image.Alpha16).SetAlpha16, execmAlpha16SetAlpha16),
		I.Func("(*Alpha16).SubImage", (*image.Alpha16).SubImage, execmAlpha16SubImage),
		I.Func("(*CMYK).At", (*image.CMYK).At, execmCMYKAt),
		I.Func("(*CMYK).Bounds", (*image.CMYK).Bounds, execmCMYKBounds),
		I.Func("(*CMYK).CMYKAt", (*image.CMYK).CMYKAt, execmCMYKCMYKAt),
		I.Func("(*CMYK).ColorModel", (*image.CMYK).ColorModel, execmCMYKColorModel),
		I.Func("(*CMYK).Opaque", (*image.CMYK).Opaque, execmCMYKOpaque),
		I.Func("(*CMYK).PixOffset", (*image.CMYK).PixOffset, execmCMYKPixOffset),
		I.Func("(*CMYK).Set", (*image.CMYK).Set, execmCMYKSet),
		I.Func("(*CMYK).SetCMYK", (*image.CMYK).SetCMYK, execmCMYKSetCMYK),
		I.Func("(*CMYK).SubImage", (*image.CMYK).SubImage, execmCMYKSubImage),
		I.Func("Decode", image.Decode, execDecode),
		I.Func("DecodeConfig", image.DecodeConfig, execDecodeConfig),
		I.Func("(*Gray).At", (*image.Gray).At, execmGrayAt),
		I.Func("(*Gray).Bounds", (*image.Gray).Bounds, execmGrayBounds),
		I.Func("(*Gray).ColorModel", (*image.Gray).ColorModel, execmGrayColorModel),
		I.Func("(*Gray).GrayAt", (*image.Gray).GrayAt, execmGrayGrayAt),
		I.Func("(*Gray).Opaque", (*image.Gray).Opaque, execmGrayOpaque),
		I.Func("(*Gray).PixOffset", (*image.Gray).PixOffset, execmGrayPixOffset),
		I.Func("(*Gray).Set", (*image.Gray).Set, execmGraySet),
		I.Func("(*Gray).SetGray", (*image.Gray).SetGray, execmGraySetGray),
		I.Func("(*Gray).SubImage", (*image.Gray).SubImage, execmGraySubImage),
		I.Func("(*Gray16).At", (*image.Gray16).At, execmGray16At),
		I.Func("(*Gray16).Bounds", (*image.Gray16).Bounds, execmGray16Bounds),
		I.Func("(*Gray16).ColorModel", (*image.Gray16).ColorModel, execmGray16ColorModel),
		I.Func("(*Gray16).Gray16At", (*image.Gray16).Gray16At, execmGray16Gray16At),
		I.Func("(*Gray16).Opaque", (*image.Gray16).Opaque, execmGray16Opaque),
		I.Func("(*Gray16).PixOffset", (*image.Gray16).PixOffset, execmGray16PixOffset),
		I.Func("(*Gray16).Set", (*image.Gray16).Set, execmGray16Set),
		I.Func("(*Gray16).SetGray16", (*image.Gray16).SetGray16, execmGray16SetGray16),
		I.Func("(*Gray16).SubImage", (*image.Gray16).SubImage, execmGray16SubImage),
		I.Func("(*NRGBA).At", (*image.NRGBA).At, execmNRGBAAt),
		I.Func("(*NRGBA).Bounds", (*image.NRGBA).Bounds, execmNRGBABounds),
		I.Func("(*NRGBA).ColorModel", (*image.NRGBA).ColorModel, execmNRGBAColorModel),
		I.Func("(*NRGBA).NRGBAAt", (*image.NRGBA).NRGBAAt, execmNRGBANRGBAAt),
		I.Func("(*NRGBA).Opaque", (*image.NRGBA).Opaque, execmNRGBAOpaque),
		I.Func("(*NRGBA).PixOffset", (*image.NRGBA).PixOffset, execmNRGBAPixOffset),
		I.Func("(*NRGBA).Set", (*image.NRGBA).Set, execmNRGBASet),
		I.Func("(*NRGBA).SetNRGBA", (*image.NRGBA).SetNRGBA, execmNRGBASetNRGBA),
		I.Func("(*NRGBA).SubImage", (*image.NRGBA).SubImage, execmNRGBASubImage),
		I.Func("(*NRGBA64).At", (*image.NRGBA64).At, execmNRGBA64At),
		I.Func("(*NRGBA64).Bounds", (*image.NRGBA64).Bounds, execmNRGBA64Bounds),
		I.Func("(*NRGBA64).ColorModel", (*image.NRGBA64).ColorModel, execmNRGBA64ColorModel),
		I.Func("(*NRGBA64).NRGBA64At", (*image.NRGBA64).NRGBA64At, execmNRGBA64NRGBA64At),
		I.Func("(*NRGBA64).Opaque", (*image.NRGBA64).Opaque, execmNRGBA64Opaque),
		I.Func("(*NRGBA64).PixOffset", (*image.NRGBA64).PixOffset, execmNRGBA64PixOffset),
		I.Func("(*NRGBA64).Set", (*image.NRGBA64).Set, execmNRGBA64Set),
		I.Func("(*NRGBA64).SetNRGBA64", (*image.NRGBA64).SetNRGBA64, execmNRGBA64SetNRGBA64),
		I.Func("(*NRGBA64).SubImage", (*image.NRGBA64).SubImage, execmNRGBA64SubImage),
		I.Func("(*NYCbCrA).AOffset", (*image.NYCbCrA).AOffset, execmNYCbCrAAOffset),
		I.Func("(*NYCbCrA).At", (*image.NYCbCrA).At, execmNYCbCrAAt),
		I.Func("(*NYCbCrA).ColorModel", (*image.NYCbCrA).ColorModel, execmNYCbCrAColorModel),
		I.Func("(*NYCbCrA).NYCbCrAAt", (*image.NYCbCrA).NYCbCrAAt, execmNYCbCrANYCbCrAAt),
		I.Func("(*NYCbCrA).Opaque", (*image.NYCbCrA).Opaque, execmNYCbCrAOpaque),
		I.Func("(*NYCbCrA).SubImage", (*image.NYCbCrA).SubImage, execmNYCbCrASubImage),
		I.Func("NewAlpha", image.NewAlpha, execNewAlpha),
		I.Func("NewAlpha16", image.NewAlpha16, execNewAlpha16),
		I.Func("NewCMYK", image.NewCMYK, execNewCMYK),
		I.Func("NewGray", image.NewGray, execNewGray),
		I.Func("NewGray16", image.NewGray16, execNewGray16),
		I.Func("NewNRGBA", image.NewNRGBA, execNewNRGBA),
		I.Func("NewNRGBA64", image.NewNRGBA64, execNewNRGBA64),
		I.Func("NewNYCbCrA", image.NewNYCbCrA, execNewNYCbCrA),
		I.Func("NewPaletted", image.NewPaletted, execNewPaletted),
		I.Func("NewRGBA", image.NewRGBA, execNewRGBA),
		I.Func("NewRGBA64", image.NewRGBA64, execNewRGBA64),
		I.Func("NewUniform", image.NewUniform, execNewUniform),
		I.Func("NewYCbCr", image.NewYCbCr, execNewYCbCr),
		I.Func("(*Paletted).At", (*image.Paletted).At, execmPalettedAt),
		I.Func("(*Paletted).Bounds", (*image.Paletted).Bounds, execmPalettedBounds),
		I.Func("(*Paletted).ColorIndexAt", (*image.Paletted).ColorIndexAt, execmPalettedColorIndexAt),
		I.Func("(*Paletted).ColorModel", (*image.Paletted).ColorModel, execmPalettedColorModel),
		I.Func("(*Paletted).Opaque", (*image.Paletted).Opaque, execmPalettedOpaque),
		I.Func("(*Paletted).PixOffset", (*image.Paletted).PixOffset, execmPalettedPixOffset),
		I.Func("(*Paletted).Set", (*image.Paletted).Set, execmPalettedSet),
		I.Func("(*Paletted).SetColorIndex", (*image.Paletted).SetColorIndex, execmPalettedSetColorIndex),
		I.Func("(*Paletted).SubImage", (*image.Paletted).SubImage, execmPalettedSubImage),
		I.Func("(Point).Add", (image.Point).Add, execmPointAdd),
		I.Func("(Point).Div", (image.Point).Div, execmPointDiv),
		I.Func("(Point).Eq", (image.Point).Eq, execmPointEq),
		I.Func("(Point).In", (image.Point).In, execmPointIn),
		I.Func("(Point).Mod", (image.Point).Mod, execmPointMod),
		I.Func("(Point).Mul", (image.Point).Mul, execmPointMul),
		I.Func("(Point).String", (image.Point).String, execmPointString),
		I.Func("(Point).Sub", (image.Point).Sub, execmPointSub),
		I.Func("Pt", image.Pt, execPt),
		I.Func("(*RGBA).At", (*image.RGBA).At, execmRGBAAt),
		I.Func("(*RGBA).Bounds", (*image.RGBA).Bounds, execmRGBABounds),
		I.Func("(*RGBA).ColorModel", (*image.RGBA).ColorModel, execmRGBAColorModel),
		I.Func("(*RGBA).Opaque", (*image.RGBA).Opaque, execmRGBAOpaque),
		I.Func("(*RGBA).PixOffset", (*image.RGBA).PixOffset, execmRGBAPixOffset),
		I.Func("(*RGBA).RGBAAt", (*image.RGBA).RGBAAt, execmRGBARGBAAt),
		I.Func("(*RGBA).Set", (*image.RGBA).Set, execmRGBASet),
		I.Func("(*RGBA).SetRGBA", (*image.RGBA).SetRGBA, execmRGBASetRGBA),
		I.Func("(*RGBA).SubImage", (*image.RGBA).SubImage, execmRGBASubImage),
		I.Func("(*RGBA64).At", (*image.RGBA64).At, execmRGBA64At),
		I.Func("(*RGBA64).Bounds", (*image.RGBA64).Bounds, execmRGBA64Bounds),
		I.Func("(*RGBA64).ColorModel", (*image.RGBA64).ColorModel, execmRGBA64ColorModel),
		I.Func("(*RGBA64).Opaque", (*image.RGBA64).Opaque, execmRGBA64Opaque),
		I.Func("(*RGBA64).PixOffset", (*image.RGBA64).PixOffset, execmRGBA64PixOffset),
		I.Func("(*RGBA64).RGBA64At", (*image.RGBA64).RGBA64At, execmRGBA64RGBA64At),
		I.Func("(*RGBA64).Set", (*image.RGBA64).Set, execmRGBA64Set),
		I.Func("(*RGBA64).SetRGBA64", (*image.RGBA64).SetRGBA64, execmRGBA64SetRGBA64),
		I.Func("(*RGBA64).SubImage", (*image.RGBA64).SubImage, execmRGBA64SubImage),
		I.Func("Rect", image.Rect, execRect),
		I.Func("(Rectangle).Add", (image.Rectangle).Add, execmRectangleAdd),
		I.Func("(Rectangle).At", (image.Rectangle).At, execmRectangleAt),
		I.Func("(Rectangle).Bounds", (image.Rectangle).Bounds, execmRectangleBounds),
		I.Func("(Rectangle).Canon", (image.Rectangle).Canon, execmRectangleCanon),
		I.Func("(Rectangle).ColorModel", (image.Rectangle).ColorModel, execmRectangleColorModel),
		I.Func("(Rectangle).Dx", (image.Rectangle).Dx, execmRectangleDx),
		I.Func("(Rectangle).Dy", (image.Rectangle).Dy, execmRectangleDy),
		I.Func("(Rectangle).Empty", (image.Rectangle).Empty, execmRectangleEmpty),
		I.Func("(Rectangle).Eq", (image.Rectangle).Eq, execmRectangleEq),
		I.Func("(Rectangle).In", (image.Rectangle).In, execmRectangleIn),
		I.Func("(Rectangle).Inset", (image.Rectangle).Inset, execmRectangleInset),
		I.Func("(Rectangle).Intersect", (image.Rectangle).Intersect, execmRectangleIntersect),
		I.Func("(Rectangle).Overlaps", (image.Rectangle).Overlaps, execmRectangleOverlaps),
		I.Func("(Rectangle).Size", (image.Rectangle).Size, execmRectangleSize),
		I.Func("(Rectangle).String", (image.Rectangle).String, execmRectangleString),
		I.Func("(Rectangle).Sub", (image.Rectangle).Sub, execmRectangleSub),
		I.Func("(Rectangle).Union", (image.Rectangle).Union, execmRectangleUnion),
		I.Func("RegisterFormat", image.RegisterFormat, execRegisterFormat),
		I.Func("(*Uniform).At", (*image.Uniform).At, execmUniformAt),
		I.Func("(*Uniform).Bounds", (*image.Uniform).Bounds, execmUniformBounds),
		I.Func("(*Uniform).ColorModel", (*image.Uniform).ColorModel, execmUniformColorModel),
		I.Func("(*Uniform).Convert", (*image.Uniform).Convert, execmUniformConvert),
		I.Func("(*Uniform).Opaque", (*image.Uniform).Opaque, execmUniformOpaque),
		I.Func("(*Uniform).RGBA", (*image.Uniform).RGBA, execmUniformRGBA),
		I.Func("(*YCbCr).At", (*image.YCbCr).At, execmYCbCrAt),
		I.Func("(*YCbCr).Bounds", (*image.YCbCr).Bounds, execmYCbCrBounds),
		I.Func("(*YCbCr).COffset", (*image.YCbCr).COffset, execmYCbCrCOffset),
		I.Func("(*YCbCr).ColorModel", (*image.YCbCr).ColorModel, execmYCbCrColorModel),
		I.Func("(*YCbCr).Opaque", (*image.YCbCr).Opaque, execmYCbCrOpaque),
		I.Func("(*YCbCr).SubImage", (*image.YCbCr).SubImage, execmYCbCrSubImage),
		I.Func("(*YCbCr).YCbCrAt", (*image.YCbCr).YCbCrAt, execmYCbCrYCbCrAt),
		I.Func("(*YCbCr).YOffset", (*image.YCbCr).YOffset, execmYCbCrYOffset),
		I.Func("(YCbCrSubsampleRatio).String", (image.YCbCrSubsampleRatio).String, execmYCbCrSubsampleRatioString),
	)
}
