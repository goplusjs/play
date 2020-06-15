package constant

import (
	"go/constant"
	"go/token"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func constant.BinaryOp(x_ constant.Value, op token.Token, y_ constant.Value) constant.Value
func execBinaryOp(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := constant.BinaryOp(args[0].(constant.Value), token.Token(args[1].(int)), args[2].(constant.Value))
	p.Ret(3, ret)
}

// func constant.BitLen(x constant.Value) int
func execBitLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.BitLen(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.BoolVal(x constant.Value) bool
func execBoolVal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.BoolVal(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.Bytes(x constant.Value) []byte
func execBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Bytes(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.Compare(x_ constant.Value, op token.Token, y_ constant.Value) bool
func execCompare(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := constant.Compare(args[0].(constant.Value), token.Token(args[1].(int)), args[2].(constant.Value))
	p.Ret(3, ret)
}

// func constant.Denom(x constant.Value) constant.Value
func execDenom(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Denom(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.Float32Val(x constant.Value) (float32, bool)
func execFloat32Val(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := constant.Float32Val(args[0].(constant.Value))
	p.Ret(1, ret, ret1)
}

// func constant.Float64Val(x constant.Value) (float64, bool)
func execFloat64Val(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := constant.Float64Val(args[0].(constant.Value))
	p.Ret(1, ret, ret1)
}

// func constant.Imag(x constant.Value) constant.Value
func execImag(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Imag(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.Int64Val(x constant.Value) (int64, bool)
func execInt64Val(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := constant.Int64Val(args[0].(constant.Value))
	p.Ret(1, ret, ret1)
}

// func constant.Make(x interface{}) constant.Value
func execMake(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Make(args[0].(interface{}))
	p.Ret(1, ret)
}

// func constant.MakeBool(b bool) constant.Value
func execMakeBool(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.MakeBool(args[0].(bool))
	p.Ret(1, ret)
}

// func constant.MakeFloat64(x float64) constant.Value
func execMakeFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.MakeFloat64(args[0].(float64))
	p.Ret(1, ret)
}

// func constant.MakeFromBytes(bytes []byte) constant.Value
func execMakeFromBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.MakeFromBytes(args[0].([]byte))
	p.Ret(1, ret)
}

// func constant.MakeFromLiteral(lit string, tok token.Token, zero uint) constant.Value
func execMakeFromLiteral(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := constant.MakeFromLiteral(args[0].(string), token.Token(args[1].(int)), args[2].(uint))
	p.Ret(3, ret)
}

// func constant.MakeImag(x constant.Value) constant.Value
func execMakeImag(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.MakeImag(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.MakeInt64(x int64) constant.Value
func execMakeInt64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.MakeInt64(args[0].(int64))
	p.Ret(1, ret)
}

// func constant.MakeString(s string) constant.Value
func execMakeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.MakeString(args[0].(string))
	p.Ret(1, ret)
}

// func constant.MakeUint64(x uint64) constant.Value
func execMakeUint64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.MakeUint64(args[0].(uint64))
	p.Ret(1, ret)
}

// func constant.MakeUnknown() constant.Value
func execMakeUnknown(_ int, p *gop.Context) {
	ret := constant.MakeUnknown()
	p.Ret(0, ret)
}

// func constant.Num(x constant.Value) constant.Value
func execNum(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Num(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.Real(x constant.Value) constant.Value
func execReal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Real(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.Shift(x constant.Value, op token.Token, s uint) constant.Value
func execShift(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := constant.Shift(args[0].(constant.Value), token.Token(args[1].(int)), args[2].(uint))
	p.Ret(3, ret)
}

// func constant.Sign(x constant.Value) int
func execSign(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Sign(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.StringVal(x constant.Value) string
func execStringVal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.StringVal(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.ToComplex(x constant.Value) constant.Value
func execToComplex(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.ToComplex(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.ToFloat(x constant.Value) constant.Value
func execToFloat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.ToFloat(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.ToInt(x constant.Value) constant.Value
func execToInt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.ToInt(args[0].(constant.Value))
	p.Ret(1, ret)
}

// func constant.Uint64Val(x constant.Value) (uint64, bool)
func execUint64Val(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := constant.Uint64Val(args[0].(constant.Value))
	p.Ret(1, ret, ret1)
}

// func constant.UnaryOp(op token.Token, y constant.Value, prec uint) constant.Value
func execUnaryOp(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := constant.UnaryOp(token.Token(args[0].(int)), args[1].(constant.Value), args[2].(uint))
	p.Ret(3, ret)
}

// func constant.Val(x constant.Value) interface{}
func execVal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := constant.Val(args[0].(constant.Value))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/constant")

func init() {
	I.RegisterConsts(
		I.Const("Bool", reflect.Int, constant.Bool),
		I.Const("Complex", reflect.Int, constant.Complex),
		I.Const("Float", reflect.Int, constant.Float),
		I.Const("Int", reflect.Int, constant.Int),
		I.Const("String", reflect.Int, constant.String),
		I.Const("Unknown", reflect.Int, constant.Unknown),
	)
	I.RegisterTypes(
		I.Type("Kind", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("BinaryOp", constant.BinaryOp, execBinaryOp),
		I.Func("BitLen", constant.BitLen, execBitLen),
		I.Func("BoolVal", constant.BoolVal, execBoolVal),
		I.Func("Bytes", constant.Bytes, execBytes),
		I.Func("Compare", constant.Compare, execCompare),
		I.Func("Denom", constant.Denom, execDenom),
		I.Func("Float32Val", constant.Float32Val, execFloat32Val),
		I.Func("Float64Val", constant.Float64Val, execFloat64Val),
		I.Func("Imag", constant.Imag, execImag),
		I.Func("Int64Val", constant.Int64Val, execInt64Val),
		I.Func("Make", constant.Make, execMake),
		I.Func("MakeBool", constant.MakeBool, execMakeBool),
		I.Func("MakeFloat64", constant.MakeFloat64, execMakeFloat64),
		I.Func("MakeFromBytes", constant.MakeFromBytes, execMakeFromBytes),
		I.Func("MakeFromLiteral", constant.MakeFromLiteral, execMakeFromLiteral),
		I.Func("MakeImag", constant.MakeImag, execMakeImag),
		I.Func("MakeInt64", constant.MakeInt64, execMakeInt64),
		I.Func("MakeString", constant.MakeString, execMakeString),
		I.Func("MakeUint64", constant.MakeUint64, execMakeUint64),
		I.Func("MakeUnknown", constant.MakeUnknown, execMakeUnknown),
		I.Func("Num", constant.Num, execNum),
		I.Func("Real", constant.Real, execReal),
		I.Func("Shift", constant.Shift, execShift),
		I.Func("Sign", constant.Sign, execSign),
		I.Func("StringVal", constant.StringVal, execStringVal),
		I.Func("ToComplex", constant.ToComplex, execToComplex),
		I.Func("ToFloat", constant.ToFloat, execToFloat),
		I.Func("ToInt", constant.ToInt, execToInt),
		I.Func("Uint64Val", constant.Uint64Val, execUint64Val),
		I.Func("UnaryOp", constant.UnaryOp, execUnaryOp),
		I.Func("Val", constant.Val, execVal),
	)
}
