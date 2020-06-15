package elliptic

import (
	"crypto/elliptic"
	"io"
	"math/big"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*elliptic.CurveParams).Add(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (*big.Int, *big.Int)
func execmCurveParamsAdd(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := args[0].(*elliptic.CurveParams).Add(args[1].(*big.Int), args[2].(*big.Int), args[3].(*big.Int), args[4].(*big.Int))
	p.Ret(5, ret, ret1)
}

// func (*elliptic.CurveParams).Double(x1 *big.Int, y1 *big.Int) (*big.Int, *big.Int)
func execmCurveParamsDouble(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*elliptic.CurveParams).Double(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret, ret1)
}

// func (*elliptic.CurveParams).IsOnCurve(x *big.Int, y *big.Int) bool
func execmCurveParamsIsOnCurve(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*elliptic.CurveParams).IsOnCurve(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*elliptic.CurveParams).Params() *elliptic.CurveParams
func execmCurveParamsParams(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*elliptic.CurveParams).Params()
	p.Ret(1, ret)
}

// func (*elliptic.CurveParams).ScalarBaseMult(k []byte) (*big.Int, *big.Int)
func execmCurveParamsScalarBaseMult(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*elliptic.CurveParams).ScalarBaseMult(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*elliptic.CurveParams).ScalarMult(Bx *big.Int, By *big.Int, k []byte) (*big.Int, *big.Int)
func execmCurveParamsScalarMult(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*elliptic.CurveParams).ScalarMult(args[1].(*big.Int), args[2].(*big.Int), args[3].([]byte))
	p.Ret(4, ret, ret1)
}

// func elliptic.GenerateKey(curve elliptic.Curve, rand io.Reader) (priv []byte, x *big.Int, y *big.Int, err error)
func execGenerateKey(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2, ret3 := elliptic.GenerateKey(args[0].(elliptic.Curve), args[1].(io.Reader))
	p.Ret(2, ret, ret1, ret2, ret3)
}

// func elliptic.Marshal(curve elliptic.Curve, x *big.Int, y *big.Int) []byte
func execMarshal(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := elliptic.Marshal(args[0].(elliptic.Curve), args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func elliptic.P224() elliptic.Curve
func execP224(_ int, p *gop.Context) {
	ret := elliptic.P224()
	p.Ret(0, ret)
}

// func elliptic.P256() elliptic.Curve
func execP256(_ int, p *gop.Context) {
	ret := elliptic.P256()
	p.Ret(0, ret)
}

// func elliptic.P384() elliptic.Curve
func execP384(_ int, p *gop.Context) {
	ret := elliptic.P384()
	p.Ret(0, ret)
}

// func elliptic.P521() elliptic.Curve
func execP521(_ int, p *gop.Context) {
	ret := elliptic.P521()
	p.Ret(0, ret)
}

// func elliptic.Unmarshal(curve elliptic.Curve, data []byte) (x *big.Int, y *big.Int)
func execUnmarshal(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := elliptic.Unmarshal(args[0].(elliptic.Curve), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/elliptic")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*elliptic.CurveParams)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*CurveParams).Add", (*elliptic.CurveParams).Add, execmCurveParamsAdd),
		I.Func("(*CurveParams).Double", (*elliptic.CurveParams).Double, execmCurveParamsDouble),
		I.Func("(*CurveParams).IsOnCurve", (*elliptic.CurveParams).IsOnCurve, execmCurveParamsIsOnCurve),
		I.Func("(*CurveParams).Params", (*elliptic.CurveParams).Params, execmCurveParamsParams),
		I.Func("(*CurveParams).ScalarBaseMult", (*elliptic.CurveParams).ScalarBaseMult, execmCurveParamsScalarBaseMult),
		I.Func("(*CurveParams).ScalarMult", (*elliptic.CurveParams).ScalarMult, execmCurveParamsScalarMult),
		I.Func("GenerateKey", elliptic.GenerateKey, execGenerateKey),
		I.Func("Marshal", elliptic.Marshal, execMarshal),
		I.Func("P224", elliptic.P224, execP224),
		I.Func("P256", elliptic.P256, execP256),
		I.Func("P384", elliptic.P384, execP384),
		I.Func("P521", elliptic.P521, execP521),
		I.Func("Unmarshal", elliptic.Unmarshal, execUnmarshal),
	)
}
