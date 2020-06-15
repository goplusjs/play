package des

import (
	"crypto/des"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (des.KeySizeError).Error() string
func execmKeySizeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(des.KeySizeError).Error()
	p.Ret(1, ret)
}

// func des.NewCipher(key []byte) (cipher.Block, error)
func execNewCipher(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := des.NewCipher(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func des.NewTripleDESCipher(key []byte) (cipher.Block, error)
func execNewTripleDESCipher(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := des.NewTripleDESCipher(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/des")

func init() {
	I.RegisterConsts(
		I.Const("BlockSize", qspec.ConstUnboundInt, des.BlockSize),
	)
	I.RegisterTypes(
		I.Type("KeySizeError", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("(KeySizeError).Error", (des.KeySizeError).Error, execmKeySizeErrorError),
		I.Func("NewCipher", des.NewCipher, execNewCipher),
		I.Func("NewTripleDESCipher", des.NewTripleDESCipher, execNewTripleDESCipher),
	)
}
