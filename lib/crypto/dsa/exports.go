package dsa

import (
	"crypto/dsa"
	"io"
	"math/big"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func dsa.GenerateKey(priv *dsa.PrivateKey, rand io.Reader) error
func execGenerateKey(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := dsa.GenerateKey(args[0].(*dsa.PrivateKey), args[1].(io.Reader))
	p.Ret(2, ret)
}

// func dsa.GenerateParameters(params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) error
func execGenerateParameters(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := dsa.GenerateParameters(args[0].(*dsa.Parameters), args[1].(io.Reader), dsa.ParameterSizes(args[2].(int)))
	p.Ret(3, ret)
}

// func dsa.Sign(rand io.Reader, priv *dsa.PrivateKey, hash []byte) (r *big.Int, s *big.Int, err error)
func execSign(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := dsa.Sign(args[0].(io.Reader), args[1].(*dsa.PrivateKey), args[2].([]byte))
	p.Ret(3, ret, ret1, ret2)
}

// func dsa.Verify(pub *dsa.PublicKey, hash []byte, r *big.Int, s *big.Int) bool
func execVerify(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := dsa.Verify(args[0].(*dsa.PublicKey), args[1].([]byte), args[2].(*big.Int), args[3].(*big.Int))
	p.Ret(4, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/dsa")

func init() {
	I.RegisterConsts(
		I.Const("L1024N160", reflect.Int, dsa.L1024N160),
		I.Const("L2048N224", reflect.Int, dsa.L2048N224),
		I.Const("L2048N256", reflect.Int, dsa.L2048N256),
		I.Const("L3072N256", reflect.Int, dsa.L3072N256),
	)
	I.RegisterVars(
		I.Var("ErrInvalidPublicKey", &dsa.ErrInvalidPublicKey),
	)
	I.RegisterTypes(
		I.Type("ParameterSizes", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*dsa.Parameters)(nil))),
		I.Rtype(reflect.TypeOf((*dsa.PrivateKey)(nil))),
		I.Rtype(reflect.TypeOf((*dsa.PublicKey)(nil))),
	)
	I.RegisterFuncs(
		I.Func("GenerateKey", dsa.GenerateKey, execGenerateKey),
		I.Func("GenerateParameters", dsa.GenerateParameters, execGenerateParameters),
		I.Func("Sign", dsa.Sign, execSign),
		I.Func("Verify", dsa.Verify, execVerify),
	)
}
