package ecdsa

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"io"
	"math/big"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func ecdsa.GenerateKey(c elliptic.Curve, rand io.Reader) (*ecdsa.PrivateKey, error)
func execGenerateKey(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := ecdsa.GenerateKey(args[0].(elliptic.Curve), args[1].(io.Reader))
	p.Ret(2, ret, ret1)
}

// func (*ecdsa.PrivateKey).Public() crypto.PublicKey
func execmPrivateKeyPublic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ecdsa.PrivateKey).Public()
	p.Ret(1, ret)
}

// func (*ecdsa.PrivateKey).Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
func execmPrivateKeySign(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*ecdsa.PrivateKey).Sign(args[1].(io.Reader), args[2].([]byte), args[3].(crypto.SignerOpts))
	p.Ret(4, ret, ret1)
}

// func ecdsa.Sign(rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) (r *big.Int, s *big.Int, err error)
func execSign(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := ecdsa.Sign(args[0].(io.Reader), args[1].(*ecdsa.PrivateKey), args[2].([]byte))
	p.Ret(3, ret, ret1, ret2)
}

// func ecdsa.Verify(pub *ecdsa.PublicKey, hash []byte, r *big.Int, s *big.Int) bool
func execVerify(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := ecdsa.Verify(args[0].(*ecdsa.PublicKey), args[1].([]byte), args[2].(*big.Int), args[3].(*big.Int))
	p.Ret(4, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/ecdsa")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*ecdsa.PrivateKey)(nil))),
		I.Rtype(reflect.TypeOf((*ecdsa.PublicKey)(nil))),
	)
	I.RegisterFuncs(
		I.Func("GenerateKey", ecdsa.GenerateKey, execGenerateKey),
		I.Func("(*PrivateKey).Public", (*ecdsa.PrivateKey).Public, execmPrivateKeyPublic),
		I.Func("(*PrivateKey).Sign", (*ecdsa.PrivateKey).Sign, execmPrivateKeySign),
		I.Func("Sign", ecdsa.Sign, execSign),
		I.Func("Verify", ecdsa.Verify, execVerify),
	)
}
