package rand

import (
	"crypto/rand"
	"io"
	"math/big"

	"github.com/qiniu/goplus/gop"
)

// func rand.Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
func execInt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := rand.Int(args[0].(io.Reader), args[1].(*big.Int))
	p.Ret(2, ret, ret1)
}

// func rand.Prime(rand io.Reader, bits int) (p *big.Int, err error)
func execPrime(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := rand.Prime(args[0].(io.Reader), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func rand.Read(b []byte) (n int, err error)
func execRead(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := rand.Read(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/rand")

func init() {
	I.RegisterVars(
		I.Var("Reader", &rand.Reader),
	)
	I.RegisterFuncs(
		I.Func("Int", rand.Int, execInt),
		I.Func("Prime", rand.Prime, execPrime),
		I.Func("Read", rand.Read, execRead),
	)
}
