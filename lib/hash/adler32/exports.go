package adler32

import (
	"hash/adler32"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func adler32.Checksum(data []byte) uint32
func execChecksum(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := adler32.Checksum(args[0].([]byte))
	p.Ret(1, ret)
}

// func adler32.New() hash.Hash32
func execNew(_ int, p *gop.Context) {
	ret := adler32.New()
	p.Ret(0, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("hash/adler32")

func init() {
	I.RegisterConsts(
		I.Const("Size", qspec.ConstUnboundInt, adler32.Size),
	)
	I.RegisterFuncs(
		I.Func("Checksum", adler32.Checksum, execChecksum),
		I.Func("New", adler32.New, execNew),
	)
}
