package sha256

import (
	"crypto/sha256"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func sha256.New() hash.Hash
func execNew(_ int, p *gop.Context) {
	ret := sha256.New()
	p.Ret(0, ret)
}

// func sha256.New224() hash.Hash
func execNew224(_ int, p *gop.Context) {
	ret := sha256.New224()
	p.Ret(0, ret)
}

// func sha256.Sum224(data []byte) (sum224 [28]byte)
func execSum224(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sha256.Sum224(args[0].([]byte))
	p.Ret(1, ret)
}

// func sha256.Sum256(data []byte) [32]byte
func execSum256(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sha256.Sum256(args[0].([]byte))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/sha256")

func init() {
	I.RegisterConsts(
		I.Const("BlockSize", qspec.ConstUnboundInt, sha256.BlockSize),
		I.Const("Size", qspec.ConstUnboundInt, sha256.Size),
		I.Const("Size224", qspec.ConstUnboundInt, sha256.Size224),
	)
	I.RegisterFuncs(
		I.Func("New", sha256.New, execNew),
		I.Func("New224", sha256.New224, execNew224),
		I.Func("Sum224", sha256.Sum224, execSum224),
		I.Func("Sum256", sha256.Sum256, execSum256),
	)
}
