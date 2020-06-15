package ed25519

import (
	"crypto"
	"crypto/ed25519"
	"io"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func ed25519.GenerateKey(rand io.Reader) (ed25519.PublicKey, ed25519.PrivateKey, error)
func execGenerateKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := ed25519.GenerateKey(args[0].(io.Reader))
	p.Ret(1, ret, ret1, ret2)
}

// func ed25519.NewKeyFromSeed(seed []byte) ed25519.PrivateKey
func execNewKeyFromSeed(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ed25519.NewKeyFromSeed(args[0].([]byte))
	p.Ret(1, ret)
}

// func (ed25519.PrivateKey).Public() crypto.PublicKey
func execmPrivateKeyPublic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(ed25519.PrivateKey).Public()
	p.Ret(1, ret)
}

// func (ed25519.PrivateKey).Seed() []byte
func execmPrivateKeySeed(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(ed25519.PrivateKey).Seed()
	p.Ret(1, ret)
}

// func (ed25519.PrivateKey).Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)
func execmPrivateKeySign(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(ed25519.PrivateKey).Sign(args[1].(io.Reader), args[2].([]byte), args[3].(crypto.SignerOpts))
	p.Ret(4, ret, ret1)
}

// func ed25519.Sign(privateKey ed25519.PrivateKey, message []byte) []byte
func execSign(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ed25519.Sign(args[0].(ed25519.PrivateKey), args[1].([]byte))
	p.Ret(2, ret)
}

// func ed25519.Verify(publicKey ed25519.PublicKey, message []byte, sig []byte) bool
func execVerify(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := ed25519.Verify(args[0].(ed25519.PublicKey), args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/ed25519")

func init() {
	I.RegisterConsts(
		I.Const("PrivateKeySize", qspec.ConstUnboundInt, ed25519.PrivateKeySize),
		I.Const("PublicKeySize", qspec.ConstUnboundInt, ed25519.PublicKeySize),
		I.Const("SeedSize", qspec.ConstUnboundInt, ed25519.SeedSize),
		I.Const("SignatureSize", qspec.ConstUnboundInt, ed25519.SignatureSize),
	)
	I.RegisterFuncs(
		I.Func("GenerateKey", ed25519.GenerateKey, execGenerateKey),
		I.Func("NewKeyFromSeed", ed25519.NewKeyFromSeed, execNewKeyFromSeed),
		I.Func("(PrivateKey).Public", (ed25519.PrivateKey).Public, execmPrivateKeyPublic),
		I.Func("(PrivateKey).Seed", (ed25519.PrivateKey).Seed, execmPrivateKeySeed),
		I.Func("(PrivateKey).Sign", (ed25519.PrivateKey).Sign, execmPrivateKeySign),
		I.Func("Sign", ed25519.Sign, execSign),
		I.Func("Verify", ed25519.Verify, execVerify),
	)
}
