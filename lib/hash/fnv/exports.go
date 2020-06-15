package fnv

import (
	"hash/fnv"

	"github.com/qiniu/goplus/gop"
)

// func fnv.New128() hash.Hash
func execNew128(_ int, p *gop.Context) {
	ret := fnv.New128()
	p.Ret(0, ret)
}

// func fnv.New128a() hash.Hash
func execNew128a(_ int, p *gop.Context) {
	ret := fnv.New128a()
	p.Ret(0, ret)
}

// func fnv.New32() hash.Hash32
func execNew32(_ int, p *gop.Context) {
	ret := fnv.New32()
	p.Ret(0, ret)
}

// func fnv.New32a() hash.Hash32
func execNew32a(_ int, p *gop.Context) {
	ret := fnv.New32a()
	p.Ret(0, ret)
}

// func fnv.New64() hash.Hash64
func execNew64(_ int, p *gop.Context) {
	ret := fnv.New64()
	p.Ret(0, ret)
}

// func fnv.New64a() hash.Hash64
func execNew64a(_ int, p *gop.Context) {
	ret := fnv.New64a()
	p.Ret(0, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("hash/fnv")

func init() {
	I.RegisterFuncs(
		I.Func("New128", fnv.New128, execNew128),
		I.Func("New128a", fnv.New128a, execNew128a),
		I.Func("New32", fnv.New32, execNew32),
		I.Func("New32a", fnv.New32a, execNew32a),
		I.Func("New64", fnv.New64, execNew64),
		I.Func("New64a", fnv.New64a, execNew64a),
	)
}
