package sha1

import (
	"crypto/sha1"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func sha1.New() hash.Hash
func execNew(_ int, p *gop.Context) {
	ret := sha1.New()
	p.Ret(0, ret)
}

// func sha1.Sum(data []byte) [20]byte
func execSum(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sha1.Sum(args[0].([]byte))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/sha1")

func init() {
	I.RegisterConsts(
		I.Const("BlockSize", qspec.ConstUnboundInt, sha1.BlockSize),
		I.Const("Size", qspec.ConstUnboundInt, sha1.Size),
	)
	I.RegisterFuncs(
		I.Func("New", sha1.New, execNew),
		I.Func("Sum", sha1.Sum, execSum),
	)
}
