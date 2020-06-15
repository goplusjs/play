package rc4

import (
	"crypto/rc4"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*rc4.Cipher).Reset()
func execmCipherReset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*rc4.Cipher).Reset()
}

// func (*rc4.Cipher).XORKeyStream(dst []byte, src []byte)
func execmCipherXORKeyStream(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*rc4.Cipher).XORKeyStream(args[1].([]byte), args[2].([]byte))
}

// func (rc4.KeySizeError).Error() string
func execmKeySizeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(rc4.KeySizeError).Error()
	p.Ret(1, ret)
}

// func rc4.NewCipher(key []byte) (*rc4.Cipher, error)
func execNewCipher(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := rc4.NewCipher(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/rc4")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*rc4.Cipher)(nil))),
		I.Type("KeySizeError", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("(*Cipher).Reset", (*rc4.Cipher).Reset, execmCipherReset),
		I.Func("(*Cipher).XORKeyStream", (*rc4.Cipher).XORKeyStream, execmCipherXORKeyStream),
		I.Func("(KeySizeError).Error", (rc4.KeySizeError).Error, execmKeySizeErrorError),
		I.Func("NewCipher", rc4.NewCipher, execNewCipher),
	)
}
