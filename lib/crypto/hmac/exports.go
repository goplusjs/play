package hmac

import (
	"crypto/hmac"
	"hash"

	"github.com/qiniu/goplus/gop"
)

// func hmac.Equal(mac1 []byte, mac2 []byte) bool
func execEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := hmac.Equal(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func hmac.New(h func() hash.Hash, key []byte) hash.Hash
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := hmac.New(args[0].(func() hash.Hash), args[1].([]byte))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/hmac")

func init() {
	I.RegisterFuncs(
		I.Func("Equal", hmac.Equal, execEqual),
		I.Func("New", hmac.New, execNew),
	)
}
