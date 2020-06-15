package sha512

import (
	"crypto/sha512"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func sha512.New() hash.Hash
func execNew(_ int, p *gop.Context) {
	ret := sha512.New()
	p.Ret(0, ret)
}

// func sha512.New384() hash.Hash
func execNew384(_ int, p *gop.Context) {
	ret := sha512.New384()
	p.Ret(0, ret)
}

// func sha512.New512_224() hash.Hash
func execNew512_224(_ int, p *gop.Context) {
	ret := sha512.New512_224()
	p.Ret(0, ret)
}

// func sha512.New512_256() hash.Hash
func execNew512_256(_ int, p *gop.Context) {
	ret := sha512.New512_256()
	p.Ret(0, ret)
}

// func sha512.Sum384(data []byte) (sum384 [48]byte)
func execSum384(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sha512.Sum384(args[0].([]byte))
	p.Ret(1, ret)
}

// func sha512.Sum512(data []byte) [64]byte
func execSum512(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sha512.Sum512(args[0].([]byte))
	p.Ret(1, ret)
}

// func sha512.Sum512_224(data []byte) (sum224 [28]byte)
func execSum512_224(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sha512.Sum512_224(args[0].([]byte))
	p.Ret(1, ret)
}

// func sha512.Sum512_256(data []byte) (sum256 [32]byte)
func execSum512_256(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sha512.Sum512_256(args[0].([]byte))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/sha512")

func init() {
	I.RegisterConsts(
		I.Const("BlockSize", qspec.ConstUnboundInt, sha512.BlockSize),
		I.Const("Size", qspec.ConstUnboundInt, sha512.Size),
		I.Const("Size224", qspec.ConstUnboundInt, sha512.Size224),
		I.Const("Size256", qspec.ConstUnboundInt, sha512.Size256),
		I.Const("Size384", qspec.ConstUnboundInt, sha512.Size384),
	)
	I.RegisterFuncs(
		I.Func("New", sha512.New, execNew),
		I.Func("New384", sha512.New384, execNew384),
		I.Func("New512_224", sha512.New512_224, execNew512_224),
		I.Func("New512_256", sha512.New512_256, execNew512_256),
		I.Func("Sum384", sha512.Sum384, execSum384),
		I.Func("Sum512", sha512.Sum512, execSum512),
		I.Func("Sum512_224", sha512.Sum512_224, execSum512_224),
		I.Func("Sum512_256", sha512.Sum512_256, execSum512_256),
	)
}
