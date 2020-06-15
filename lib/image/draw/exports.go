package draw

import (
	"image"
	"image/draw"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func draw.Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, op draw.Op)
func execDraw(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	draw.Draw(args[0].(draw.Image), args[1].(image.Rectangle), args[2].(image.Image), args[3].(image.Point), draw.Op(args[4].(int)))
}

// func draw.DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op draw.Op)
func execDrawMask(_ int, p *gop.Context) {
	args := p.GetArgs(7)
	draw.DrawMask(args[0].(draw.Image), args[1].(image.Rectangle), args[2].(image.Image), args[3].(image.Point), args[4].(image.Image), args[5].(image.Point), draw.Op(args[6].(int)))
}

// func (draw.Op).Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)
func execmOpDraw(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(draw.Op).Draw(args[1].(draw.Image), args[2].(image.Rectangle), args[3].(image.Image), args[4].(image.Point))
}

// I is a Go package instance.
var I = gop.NewGoPackage("image/draw")

func init() {
	I.RegisterConsts(
		I.Const("Over", reflect.Int, draw.Over),
		I.Const("Src", reflect.Int, draw.Src),
	)
	I.RegisterVars(
		I.Var("FloydSteinberg", &draw.FloydSteinberg),
	)
	I.RegisterTypes(
		I.Type("Op", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("Draw", draw.Draw, execDraw),
		I.Func("DrawMask", draw.DrawMask, execDrawMask),
		I.Func("(Op).Draw", (draw.Op).Draw, execmOpDraw),
	)
}
