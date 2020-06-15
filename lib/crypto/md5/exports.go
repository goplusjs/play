package md5

import (
	"crypto/md5"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func md5.New() hash.Hash
func execNew(_ int, p *gop.Context) {
	ret := md5.New()
	p.Ret(0, ret)
}

// func md5.Sum(data []byte) [16]byte
func execSum(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := md5.Sum(args[0].([]byte))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/md5")

func init() {
	I.RegisterConsts(
		I.Const("BlockSize", qspec.ConstUnboundInt, md5.BlockSize),
		I.Const("Size", qspec.ConstUnboundInt, md5.Size),
	)
	I.RegisterFuncs(
		I.Func("New", md5.New, execNew),
		I.Func("Sum", md5.Sum, execSum),
	)
}
