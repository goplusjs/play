package math

import (
	"math"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func math.Abs(x float64) float64
func execAbs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Abs(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Acos(x float64) float64
func execAcos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Acos(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Acosh(x float64) float64
func execAcosh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Acosh(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Asin(x float64) float64
func execAsin(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Asin(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Asinh(x float64) float64
func execAsinh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Asinh(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Atan(x float64) float64
func execAtan(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Atan(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Atan2(y float64, x float64) float64
func execAtan2(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Atan2(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Atanh(x float64) float64
func execAtanh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Atanh(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Cbrt(x float64) float64
func execCbrt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Cbrt(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Ceil(x float64) float64
func execCeil(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Ceil(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Copysign(x float64, y float64) float64
func execCopysign(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Copysign(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Cos(x float64) float64
func execCos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Cos(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Cosh(x float64) float64
func execCosh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Cosh(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Dim(x float64, y float64) float64
func execDim(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Dim(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Erf(x float64) float64
func execErf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Erf(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Erfc(x float64) float64
func execErfc(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Erfc(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Erfcinv(x float64) float64
func execErfcinv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Erfcinv(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Erfinv(x float64) float64
func execErfinv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Erfinv(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Exp(x float64) float64
func execExp(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Exp(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Exp2(x float64) float64
func execExp2(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Exp2(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Expm1(x float64) float64
func execExpm1(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Expm1(args[0].(float64))
	p.Ret(1, ret)
}

// func math.FMA(x float64, y float64, z float64) float64
func execFMA(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := math.FMA(args[0].(float64), args[1].(float64), args[2].(float64))
	p.Ret(3, ret)
}

// func math.Float32bits(f float32) uint32
func execFloat32bits(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Float32bits(args[0].(float32))
	p.Ret(1, ret)
}

// func math.Float32frombits(b uint32) float32
func execFloat32frombits(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Float32frombits(args[0].(uint32))
	p.Ret(1, ret)
}

// func math.Float64bits(f float64) uint64
func execFloat64bits(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Float64bits(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Float64frombits(b uint64) float64
func execFloat64frombits(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Float64frombits(args[0].(uint64))
	p.Ret(1, ret)
}

// func math.Floor(x float64) float64
func execFloor(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Floor(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Frexp(f float64) (frac float64, exp int)
func execFrexp(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := math.Frexp(args[0].(float64))
	p.Ret(1, ret, ret1)
}

// func math.Gamma(x float64) float64
func execGamma(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Gamma(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Hypot(p float64, q float64) float64
func execHypot(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Hypot(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Ilogb(x float64) int
func execIlogb(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Ilogb(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Inf(sign int) float64
func execInf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Inf(args[0].(int))
	p.Ret(1, ret)
}

// func math.IsInf(f float64, sign int) bool
func execIsInf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.IsInf(args[0].(float64), args[1].(int))
	p.Ret(2, ret)
}

// func math.IsNaN(f float64) (is bool)
func execIsNaN(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.IsNaN(args[0].(float64))
	p.Ret(1, ret)
}

// func math.J0(x float64) float64
func execJ0(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.J0(args[0].(float64))
	p.Ret(1, ret)
}

// func math.J1(x float64) float64
func execJ1(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.J1(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Jn(n int, x float64) float64
func execJn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Jn(args[0].(int), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Ldexp(frac float64, exp int) float64
func execLdexp(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Ldexp(args[0].(float64), args[1].(int))
	p.Ret(2, ret)
}

// func math.Lgamma(x float64) (lgamma float64, sign int)
func execLgamma(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := math.Lgamma(args[0].(float64))
	p.Ret(1, ret, ret1)
}

// func math.Log(x float64) float64
func execLog(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Log(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Log10(x float64) float64
func execLog10(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Log10(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Log1p(x float64) float64
func execLog1p(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Log1p(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Log2(x float64) float64
func execLog2(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Log2(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Logb(x float64) float64
func execLogb(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Logb(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Max(x float64, y float64) float64
func execMax(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Max(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Min(x float64, y float64) float64
func execMin(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Min(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Mod(x float64, y float64) float64
func execMod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Mod(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Modf(f float64) (int float64, frac float64)
func execModf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := math.Modf(args[0].(float64))
	p.Ret(1, ret, ret1)
}

// func math.NaN() float64
func execNaN(_ int, p *gop.Context) {
	ret := math.NaN()
	p.Ret(0, ret)
}

// func math.Nextafter(x float64, y float64) (r float64)
func execNextafter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Nextafter(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Nextafter32(x float32, y float32) (r float32)
func execNextafter32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Nextafter32(args[0].(float32), args[1].(float32))
	p.Ret(2, ret)
}

// func math.Pow(x float64, y float64) float64
func execPow(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Pow(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Pow10(n int) float64
func execPow10(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Pow10(args[0].(int))
	p.Ret(1, ret)
}

// func math.Remainder(x float64, y float64) float64
func execRemainder(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Remainder(args[0].(float64), args[1].(float64))
	p.Ret(2, ret)
}

// func math.Round(x float64) float64
func execRound(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Round(args[0].(float64))
	p.Ret(1, ret)
}

// func math.RoundToEven(x float64) float64
func execRoundToEven(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.RoundToEven(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Signbit(x float64) bool
func execSignbit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Signbit(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Sin(x float64) float64
func execSin(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Sin(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Sincos(x float64) (sin float64, cos float64)
func execSincos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := math.Sincos(args[0].(float64))
	p.Ret(1, ret, ret1)
}

// func math.Sinh(x float64) float64
func execSinh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Sinh(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Sqrt(x float64) float64
func execSqrt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Sqrt(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Tan(x float64) float64
func execTan(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Tan(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Tanh(x float64) float64
func execTanh(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Tanh(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Trunc(x float64) float64
func execTrunc(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Trunc(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Y0(x float64) float64
func execY0(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Y0(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Y1(x float64) float64
func execY1(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := math.Y1(args[0].(float64))
	p.Ret(1, ret)
}

// func math.Yn(n int, x float64) float64
func execYn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := math.Yn(args[0].(int), args[1].(float64))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("math")

func init() {
	I.RegisterConsts(
		I.Const("E", qspec.ConstUnboundFloat, math.E),
		I.Const("Ln10", qspec.ConstUnboundFloat, math.Ln10),
		I.Const("Ln2", qspec.ConstUnboundFloat, math.Ln2),
		I.Const("Log10E", qspec.ConstUnboundFloat, math.Log10E),
		I.Const("Log2E", qspec.ConstUnboundFloat, math.Log2E),
		I.Const("MaxFloat32", qspec.ConstUnboundFloat, math.MaxFloat32),
		I.Const("MaxFloat64", qspec.ConstUnboundFloat, math.MaxFloat64),
		I.Const("MaxInt16", qspec.ConstUnboundInt, math.MaxInt16),
		I.Const("MaxInt32", qspec.ConstUnboundInt, math.MaxInt32),
		I.Const("MaxInt64", qspec.Uint64, uint64(math.MaxInt64)),
		I.Const("MaxInt8", qspec.ConstUnboundInt, math.MaxInt8),
		I.Const("MaxUint16", qspec.ConstUnboundInt, math.MaxUint16),
		I.Const("MaxUint32", qspec.Uint64, uint64(math.MaxUint32)),
		I.Const("MaxUint64", qspec.Uint64, uint64(math.MaxUint64)),
		I.Const("MaxUint8", qspec.ConstUnboundInt, math.MaxUint8),
		I.Const("MinInt16", qspec.ConstUnboundInt, math.MinInt16),
		I.Const("MinInt32", qspec.ConstUnboundInt, math.MinInt32),
		I.Const("MinInt64", qspec.Int64, int64(math.MinInt64)),
		I.Const("MinInt8", qspec.ConstUnboundInt, math.MinInt8),
		I.Const("Phi", qspec.ConstUnboundFloat, math.Phi),
		I.Const("Pi", qspec.ConstUnboundFloat, math.Pi),
		I.Const("SmallestNonzeroFloat32", qspec.ConstUnboundFloat, math.SmallestNonzeroFloat32),
		I.Const("SmallestNonzeroFloat64", qspec.ConstUnboundFloat, math.SmallestNonzeroFloat64),
		I.Const("Sqrt2", qspec.ConstUnboundFloat, math.Sqrt2),
		I.Const("SqrtE", qspec.ConstUnboundFloat, math.SqrtE),
		I.Const("SqrtPhi", qspec.ConstUnboundFloat, math.SqrtPhi),
		I.Const("SqrtPi", qspec.ConstUnboundFloat, math.SqrtPi),
	)
	I.RegisterFuncs(
		I.Func("Abs", math.Abs, execAbs),
		I.Func("Acos", math.Acos, execAcos),
		I.Func("Acosh", math.Acosh, execAcosh),
		I.Func("Asin", math.Asin, execAsin),
		I.Func("Asinh", math.Asinh, execAsinh),
		I.Func("Atan", math.Atan, execAtan),
		I.Func("Atan2", math.Atan2, execAtan2),
		I.Func("Atanh", math.Atanh, execAtanh),
		I.Func("Cbrt", math.Cbrt, execCbrt),
		I.Func("Ceil", math.Ceil, execCeil),
		I.Func("Copysign", math.Copysign, execCopysign),
		I.Func("Cos", math.Cos, execCos),
		I.Func("Cosh", math.Cosh, execCosh),
		I.Func("Dim", math.Dim, execDim),
		I.Func("Erf", math.Erf, execErf),
		I.Func("Erfc", math.Erfc, execErfc),
		I.Func("Erfcinv", math.Erfcinv, execErfcinv),
		I.Func("Erfinv", math.Erfinv, execErfinv),
		I.Func("Exp", math.Exp, execExp),
		I.Func("Exp2", math.Exp2, execExp2),
		I.Func("Expm1", math.Expm1, execExpm1),
		I.Func("FMA", math.FMA, execFMA),
		I.Func("Float32bits", math.Float32bits, execFloat32bits),
		I.Func("Float32frombits", math.Float32frombits, execFloat32frombits),
		I.Func("Float64bits", math.Float64bits, execFloat64bits),
		I.Func("Float64frombits", math.Float64frombits, execFloat64frombits),
		I.Func("Floor", math.Floor, execFloor),
		I.Func("Frexp", math.Frexp, execFrexp),
		I.Func("Gamma", math.Gamma, execGamma),
		I.Func("Hypot", math.Hypot, execHypot),
		I.Func("Ilogb", math.Ilogb, execIlogb),
		I.Func("Inf", math.Inf, execInf),
		I.Func("IsInf", math.IsInf, execIsInf),
		I.Func("IsNaN", math.IsNaN, execIsNaN),
		I.Func("J0", math.J0, execJ0),
		I.Func("J1", math.J1, execJ1),
		I.Func("Jn", math.Jn, execJn),
		I.Func("Ldexp", math.Ldexp, execLdexp),
		I.Func("Lgamma", math.Lgamma, execLgamma),
		I.Func("Log", math.Log, execLog),
		I.Func("Log10", math.Log10, execLog10),
		I.Func("Log1p", math.Log1p, execLog1p),
		I.Func("Log2", math.Log2, execLog2),
		I.Func("Logb", math.Logb, execLogb),
		I.Func("Max", math.Max, execMax),
		I.Func("Min", math.Min, execMin),
		I.Func("Mod", math.Mod, execMod),
		I.Func("Modf", math.Modf, execModf),
		I.Func("NaN", math.NaN, execNaN),
		I.Func("Nextafter", math.Nextafter, execNextafter),
		I.Func("Nextafter32", math.Nextafter32, execNextafter32),
		I.Func("Pow", math.Pow, execPow),
		I.Func("Pow10", math.Pow10, execPow10),
		I.Func("Remainder", math.Remainder, execRemainder),
		I.Func("Round", math.Round, execRound),
		I.Func("RoundToEven", math.RoundToEven, execRoundToEven),
		I.Func("Signbit", math.Signbit, execSignbit),
		I.Func("Sin", math.Sin, execSin),
		I.Func("Sincos", math.Sincos, execSincos),
		I.Func("Sinh", math.Sinh, execSinh),
		I.Func("Sqrt", math.Sqrt, execSqrt),
		I.Func("Tan", math.Tan, execTan),
		I.Func("Tanh", math.Tanh, execTanh),
		I.Func("Trunc", math.Trunc, execTrunc),
		I.Func("Y0", math.Y0, execY0),
		I.Func("Y1", math.Y1, execY1),
		I.Func("Yn", math.Yn, execYn),
	)
}
