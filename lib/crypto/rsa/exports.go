package rsa

import (
	"crypto"
	"crypto/rsa"
	"hash"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func rsa.DecryptOAEP(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
func execDecryptOAEP(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := rsa.DecryptOAEP(args[0].(hash.Hash), args[1].(io.Reader), args[2].(*rsa.PrivateKey), args[3].([]byte), args[4].([]byte))
	p.Ret(5, ret, ret1)
}

// func rsa.DecryptPKCS1v15(rand io.Reader, priv *rsa.PrivateKey, ciphertext []byte) ([]byte, error)
func execDecryptPKCS1v15(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := rsa.DecryptPKCS1v15(args[0].(io.Reader), args[1].(*rsa.PrivateKey), args[2].([]byte))
	p.Ret(3, ret, ret1)
}

// func rsa.DecryptPKCS1v15SessionKey(rand io.Reader, priv *rsa.PrivateKey, ciphertext []byte, key []byte) error
func execDecryptPKCS1v15SessionKey(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := rsa.DecryptPKCS1v15SessionKey(args[0].(io.Reader), args[1].(*rsa.PrivateKey), args[2].([]byte), args[3].([]byte))
	p.Ret(4, ret)
}

// func rsa.EncryptOAEP(hash hash.Hash, random io.Reader, pub *rsa.PublicKey, msg []byte, label []byte) ([]byte, error)
func execEncryptOAEP(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := rsa.EncryptOAEP(args[0].(hash.Hash), args[1].(io.Reader), args[2].(*rsa.PublicKey), args[3].([]byte), args[4].([]byte))
	p.Ret(5, ret, ret1)
}

// func rsa.EncryptPKCS1v15(rand io.Reader, pub *rsa.PublicKey, msg []byte) ([]byte, error)
func execEncryptPKCS1v15(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := rsa.EncryptPKCS1v15(args[0].(io.Reader), args[1].(*rsa.PublicKey), args[2].([]byte))
	p.Ret(3, ret, ret1)
}

// func rsa.GenerateKey(random io.Reader, bits int) (*rsa.PrivateKey, error)
func execGenerateKey(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := rsa.GenerateKey(args[0].(io.Reader), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func rsa.GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*rsa.PrivateKey, error)
func execGenerateMultiPrimeKey(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := rsa.GenerateMultiPrimeKey(args[0].(io.Reader), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func (*rsa.PSSOptions).HashFunc() crypto.Hash
func execmPSSOptionsHashFunc(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rsa.PSSOptions).HashFunc()
	p.Ret(1, ret)
}

// func (*rsa.PrivateKey).Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
func execmPrivateKeyDecrypt(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*rsa.PrivateKey).Decrypt(args[1].(io.Reader), args[2].([]byte), args[3].(crypto.DecrypterOpts))
	p.Ret(4, ret, ret1)
}

// func (*rsa.PrivateKey).Precompute()
func execmPrivateKeyPrecompute(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*rsa.PrivateKey).Precompute()
}

// func (*rsa.PrivateKey).Public() crypto.PublicKey
func execmPrivateKeyPublic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rsa.PrivateKey).Public()
	p.Ret(1, ret)
}

// func (*rsa.PrivateKey).Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
func execmPrivateKeySign(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*rsa.PrivateKey).Sign(args[1].(io.Reader), args[2].([]byte), args[3].(crypto.SignerOpts))
	p.Ret(4, ret, ret1)
}

// func (*rsa.PrivateKey).Validate() error
func execmPrivateKeyValidate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rsa.PrivateKey).Validate()
	p.Ret(1, ret)
}

// func (*rsa.PublicKey).Size() int
func execmPublicKeySize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rsa.PublicKey).Size()
	p.Ret(1, ret)
}

// func rsa.SignPKCS1v15(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
func execSignPKCS1v15(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := rsa.SignPKCS1v15(args[0].(io.Reader), args[1].(*rsa.PrivateKey), crypto.Hash(args[2].(uint)), args[3].([]byte))
	p.Ret(4, ret, ret1)
}

// func rsa.SignPSS(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions) ([]byte, error)
func execSignPSS(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := rsa.SignPSS(args[0].(io.Reader), args[1].(*rsa.PrivateKey), crypto.Hash(args[2].(uint)), args[3].([]byte), args[4].(*rsa.PSSOptions))
	p.Ret(5, ret, ret1)
}

// func rsa.VerifyPKCS1v15(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
func execVerifyPKCS1v15(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := rsa.VerifyPKCS1v15(args[0].(*rsa.PublicKey), crypto.Hash(args[1].(uint)), args[2].([]byte), args[3].([]byte))
	p.Ret(4, ret)
}

// func rsa.VerifyPSS(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *rsa.PSSOptions) error
func execVerifyPSS(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := rsa.VerifyPSS(args[0].(*rsa.PublicKey), crypto.Hash(args[1].(uint)), args[2].([]byte), args[3].([]byte), args[4].(*rsa.PSSOptions))
	p.Ret(5, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/rsa")

func init() {
	I.RegisterConsts(
		I.Const("PSSSaltLengthAuto", qspec.ConstUnboundInt, rsa.PSSSaltLengthAuto),
		I.Const("PSSSaltLengthEqualsHash", qspec.ConstUnboundInt, rsa.PSSSaltLengthEqualsHash),
	)
	I.RegisterVars(
		I.Var("ErrDecryption", &rsa.ErrDecryption),
		I.Var("ErrMessageTooLong", &rsa.ErrMessageTooLong),
		I.Var("ErrVerification", &rsa.ErrVerification),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*rsa.CRTValue)(nil))),
		I.Rtype(reflect.TypeOf((*rsa.OAEPOptions)(nil))),
		I.Rtype(reflect.TypeOf((*rsa.PKCS1v15DecryptOptions)(nil))),
		I.Rtype(reflect.TypeOf((*rsa.PSSOptions)(nil))),
		I.Rtype(reflect.TypeOf((*rsa.PrecomputedValues)(nil))),
		I.Rtype(reflect.TypeOf((*rsa.PrivateKey)(nil))),
		I.Rtype(reflect.TypeOf((*rsa.PublicKey)(nil))),
	)
	I.RegisterFuncs(
		I.Func("DecryptOAEP", rsa.DecryptOAEP, execDecryptOAEP),
		I.Func("DecryptPKCS1v15", rsa.DecryptPKCS1v15, execDecryptPKCS1v15),
		I.Func("DecryptPKCS1v15SessionKey", rsa.DecryptPKCS1v15SessionKey, execDecryptPKCS1v15SessionKey),
		I.Func("EncryptOAEP", rsa.EncryptOAEP, execEncryptOAEP),
		I.Func("EncryptPKCS1v15", rsa.EncryptPKCS1v15, execEncryptPKCS1v15),
		I.Func("GenerateKey", rsa.GenerateKey, execGenerateKey),
		I.Func("GenerateMultiPrimeKey", rsa.GenerateMultiPrimeKey, execGenerateMultiPrimeKey),
		I.Func("(*PSSOptions).HashFunc", (*rsa.PSSOptions).HashFunc, execmPSSOptionsHashFunc),
		I.Func("(*PrivateKey).Decrypt", (*rsa.PrivateKey).Decrypt, execmPrivateKeyDecrypt),
		I.Func("(*PrivateKey).Precompute", (*rsa.PrivateKey).Precompute, execmPrivateKeyPrecompute),
		I.Func("(*PrivateKey).Public", (*rsa.PrivateKey).Public, execmPrivateKeyPublic),
		I.Func("(*PrivateKey).Sign", (*rsa.PrivateKey).Sign, execmPrivateKeySign),
		I.Func("(*PrivateKey).Validate", (*rsa.PrivateKey).Validate, execmPrivateKeyValidate),
		I.Func("(*PublicKey).Size", (*rsa.PublicKey).Size, execmPublicKeySize),
		I.Func("SignPKCS1v15", rsa.SignPKCS1v15, execSignPKCS1v15),
		I.Func("SignPSS", rsa.SignPSS, execSignPSS),
		I.Func("VerifyPKCS1v15", rsa.VerifyPKCS1v15, execVerifyPKCS1v15),
		I.Func("VerifyPSS", rsa.VerifyPSS, execVerifyPSS),
	)
}
