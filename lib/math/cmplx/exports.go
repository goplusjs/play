package cmplx

import (
	"math/cmplx"

	"github.com/qiniu/goplus/gop"
)

// func cmplx.Abs(x complex128) float64
func execAbs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Abs(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Acos(x complex128) complex128
func execAcos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Acos(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Acosh(x complex128) complex128
func execAcosh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Acosh(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Asin(x complex128) complex128
func execAsin(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Asin(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Asinh(x complex128) complex128
func execAsinh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Asinh(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Atan(x complex128) complex128
func execAtan(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Atan(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Atanh(x complex128) complex128
func execAtanh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Atanh(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Conj(x complex128) complex128
func execConj(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Conj(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Cos(x complex128) complex128
func execCos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Cos(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Cosh(x complex128) complex128
func execCosh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Cosh(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Cot(x complex128) complex128
func execCot(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Cot(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Exp(x complex128) complex128
func execExp(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Exp(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Inf() complex128
func execInf(_ int, p *gop.Context) {
	ret := cmplx.Inf()
	p.Ret(0, ret)
}

// func cmplx.IsInf(x complex128) bool
func execIsInf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.IsInf(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.IsNaN(x complex128) bool
func execIsNaN(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.IsNaN(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Log(x complex128) complex128
func execLog(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Log(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Log10(x complex128) complex128
func execLog10(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Log10(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.NaN() complex128
func execNaN(_ int, p *gop.Context) {
	ret := cmplx.NaN()
	p.Ret(0, ret)
}

// func cmplx.Phase(x complex128) float64
func execPhase(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Phase(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Polar(x complex128) (r float64, θ float64)
func execPolar(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := cmplx.Polar(args[0].(complex128))
	p.Ret(1, ret, ret1)
}

// func cmplx.Pow(x complex128, y complex128) complex128
func execPow(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cmplx.Pow(args[0].(complex128), args[1].(complex128))
	p.Ret(2, ret)
}

// func cmplx.Rect(r float64, θ float64) complex128
func execRect(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cmplx.Rect(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func cmplx.Sin(x complex128) complex128
func execSin(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Sin(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Sinh(x complex128) complex128
func execSinh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Sinh(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Sqrt(x complex128) complex128
func execSqrt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Sqrt(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Tan(x complex128) complex128
func execTan(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Tan(args[0].(complex128))
	p.Ret(1, ret)
}

// func cmplx.Tanh(x complex128) complex128
func execTanh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := cmplx.Tanh(args[0].(complex128))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("math/cmplx")

func init() {
	I.RegisterFuncs(
		I.Func("Abs", cmplx.Abs, execAbs),
		I.Func("Acos", cmplx.Acos, execAcos),
		I.Func("Acosh", cmplx.Acosh, execAcosh),
		I.Func("Asin", cmplx.Asin, execAsin),
		I.Func("Asinh", cmplx.Asinh, execAsinh),
		I.Func("Atan", cmplx.Atan, execAtan),
		I.Func("Atanh", cmplx.Atanh, execAtanh),
		I.Func("Conj", cmplx.Conj, execConj),
		I.Func("Cos", cmplx.Cos, execCos),
		I.Func("Cosh", cmplx.Cosh, execCosh),
		I.Func("Cot", cmplx.Cot, execCot),
		I.Func("Exp", cmplx.Exp, execExp),
		I.Func("Inf", cmplx.Inf, execInf),
		I.Func("IsInf", cmplx.IsInf, execIsInf),
		I.Func("IsNaN", cmplx.IsNaN, execIsNaN),
		I.Func("Log", cmplx.Log, execLog),
		I.Func("Log10", cmplx.Log10, execLog10),
		I.Func("NaN", cmplx.NaN, execNaN),
		I.Func("Phase", cmplx.Phase, execPhase),
		I.Func("Polar", cmplx.Polar, execPolar),
		I.Func("Pow", cmplx.Pow, execPow),
		I.Func("Rect", cmplx.Rect, execRect),
		I.Func("Sin", cmplx.Sin, execSin),
		I.Func("Sinh", cmplx.Sinh, execSinh),
		I.Func("Sqrt", cmplx.Sqrt, execSqrt),
		I.Func("Tan", cmplx.Tan, execTan),
		I.Func("Tanh", cmplx.Tanh, execTanh),
	)
}
