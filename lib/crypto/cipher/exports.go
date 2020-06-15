package cipher

import (
	"crypto/cipher"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func cipher.NewCBCDecrypter(b cipher.Block, iv []byte) cipher.BlockMode
func execNewCBCDecrypter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cipher.NewCBCDecrypter(args[0].(cipher.Block), args[1].([]byte))
	p.Ret(2, ret)
}

// func cipher.NewCBCEncrypter(b cipher.Block, iv []byte) cipher.BlockMode
func execNewCBCEncrypter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cipher.NewCBCEncrypter(args[0].(cipher.Block), args[1].([]byte))
	p.Ret(2, ret)
}

// func cipher.NewCFBDecrypter(block cipher.Block, iv []byte) cipher.Stream
func execNewCFBDecrypter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cipher.NewCFBDecrypter(args[0].(cipher.Block), args[1].([]byte))
	p.Ret(2, ret)
}

// func cipher.NewCFBEncrypter(block cipher.Block, iv []byte) cipher.Stream
func execNewCFBEncrypter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cipher.NewCFBEncrypter(args[0].(cipher.Block), args[1].([]byte))
	p.Ret(2, ret)
}

// func cipher.NewCTR(block cipher.Block, iv []byte) cipher.Stream
func execNewCTR(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cipher.NewCTR(args[0].(cipher.Block), args[1].([]byte))
	p.Ret(2, ret)
}

// func cipher.NewGCM(cipher cipher.Block) (cipher.AEAD, error)
func execNewGCM(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := cipher.NewGCM(args[0].(cipher.Block))
	p.Ret(1, ret, ret1)
}

// func cipher.NewGCMWithNonceSize(cipher cipher.Block, size int) (cipher.AEAD, error)
func execNewGCMWithNonceSize(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := cipher.NewGCMWithNonceSize(args[0].(cipher.Block), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func cipher.NewGCMWithTagSize(cipher cipher.Block, tagSize int) (cipher.AEAD, error)
func execNewGCMWithTagSize(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := cipher.NewGCMWithTagSize(args[0].(cipher.Block), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func cipher.NewOFB(b cipher.Block, iv []byte) cipher.Stream
func execNewOFB(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := cipher.NewOFB(args[0].(cipher.Block), args[1].([]byte))
	p.Ret(2, ret)
}

// func (cipher.StreamReader).Read(dst []byte) (n int, err error)
func execmStreamReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(cipher.StreamReader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (cipher.StreamWriter).Close() error
func execmStreamWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(cipher.StreamWriter).Close()
	p.Ret(1, ret)
}

// func (cipher.StreamWriter).Write(src []byte) (n int, err error)
func execmStreamWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(cipher.StreamWriter).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/cipher")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*cipher.StreamReader)(nil))),
		I.Rtype(reflect.TypeOf((*cipher.StreamWriter)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewCBCDecrypter", cipher.NewCBCDecrypter, execNewCBCDecrypter),
		I.Func("NewCBCEncrypter", cipher.NewCBCEncrypter, execNewCBCEncrypter),
		I.Func("NewCFBDecrypter", cipher.NewCFBDecrypter, execNewCFBDecrypter),
		I.Func("NewCFBEncrypter", cipher.NewCFBEncrypter, execNewCFBEncrypter),
		I.Func("NewCTR", cipher.NewCTR, execNewCTR),
		I.Func("NewGCM", cipher.NewGCM, execNewGCM),
		I.Func("NewGCMWithNonceSize", cipher.NewGCMWithNonceSize, execNewGCMWithNonceSize),
		I.Func("NewGCMWithTagSize", cipher.NewGCMWithTagSize, execNewGCMWithTagSize),
		I.Func("NewOFB", cipher.NewOFB, execNewOFB),
		I.Func("(StreamReader).Read", (cipher.StreamReader).Read, execmStreamReaderRead),
		I.Func("(StreamWriter).Close", (cipher.StreamWriter).Close, execmStreamWriterClose),
		I.Func("(StreamWriter).Write", (cipher.StreamWriter).Write, execmStreamWriterWrite),
	)
}
