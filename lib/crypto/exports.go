package crypto

import (
	"crypto"
	"hash"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (crypto.Hash).Available() bool
func execmHashAvailable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(crypto.Hash).Available()
	p.Ret(1, ret)
}

// func (crypto.Hash).HashFunc() crypto.Hash
func execmHashHashFunc(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(crypto.Hash).HashFunc()
	p.Ret(1, ret)
}

// func (crypto.Hash).New() hash.Hash
func execmHashNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(crypto.Hash).New()
	p.Ret(1, ret)
}

// func (crypto.Hash).Size() int
func execmHashSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(crypto.Hash).Size()
	p.Ret(1, ret)
}

// func crypto.RegisterHash(h crypto.Hash, f func() hash.Hash)
func execRegisterHash(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	crypto.RegisterHash(crypto.Hash(args[0].(uint)), args[1].(func() hash.Hash))
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto")

func init() {
	I.RegisterConsts(
		I.Const("BLAKE2b_256", reflect.Uint, crypto.BLAKE2b_256),
		I.Const("BLAKE2b_384", reflect.Uint, crypto.BLAKE2b_384),
		I.Const("BLAKE2b_512", reflect.Uint, crypto.BLAKE2b_512),
		I.Const("BLAKE2s_256", reflect.Uint, crypto.BLAKE2s_256),
		I.Const("MD4", reflect.Uint, crypto.MD4),
		I.Const("MD5", reflect.Uint, crypto.MD5),
		I.Const("MD5SHA1", reflect.Uint, crypto.MD5SHA1),
		I.Const("RIPEMD160", reflect.Uint, crypto.RIPEMD160),
		I.Const("SHA1", reflect.Uint, crypto.SHA1),
		I.Const("SHA224", reflect.Uint, crypto.SHA224),
		I.Const("SHA256", reflect.Uint, crypto.SHA256),
		I.Const("SHA384", reflect.Uint, crypto.SHA384),
		I.Const("SHA3_224", reflect.Uint, crypto.SHA3_224),
		I.Const("SHA3_256", reflect.Uint, crypto.SHA3_256),
		I.Const("SHA3_384", reflect.Uint, crypto.SHA3_384),
		I.Const("SHA3_512", reflect.Uint, crypto.SHA3_512),
		I.Const("SHA512", reflect.Uint, crypto.SHA512),
		I.Const("SHA512_224", reflect.Uint, crypto.SHA512_224),
		I.Const("SHA512_256", reflect.Uint, crypto.SHA512_256),
	)
	I.RegisterTypes(
		I.Type("Hash", qspec.TyUint),
	)
	I.RegisterFuncs(
		I.Func("(Hash).Available", (crypto.Hash).Available, execmHashAvailable),
		I.Func("(Hash).HashFunc", (crypto.Hash).HashFunc, execmHashHashFunc),
		I.Func("(Hash).New", (crypto.Hash).New, execmHashNew),
		I.Func("(Hash).Size", (crypto.Hash).Size, execmHashSize),
		I.Func("RegisterHash", crypto.RegisterHash, execRegisterHash),
	)
}
