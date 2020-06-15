package aes

import (
	"crypto/aes"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (aes.KeySizeError).Error() string
func execmKeySizeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(aes.KeySizeError).Error()
	p.Ret(1, ret)
}

// func aes.NewCipher(key []byte) (cipher.Block, error)
func execNewCipher(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := aes.NewCipher(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/aes")

func init() {
	I.RegisterConsts(
		I.Const("BlockSize", qspec.ConstUnboundInt, aes.BlockSize),
	)
	I.RegisterTypes(
		I.Type("KeySizeError", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("(KeySizeError).Error", (aes.KeySizeError).Error, execmKeySizeErrorError),
		I.Func("NewCipher", aes.NewCipher, execNewCipher),
	)
}
