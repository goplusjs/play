package ioutil

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/qiniu/goplus/gop"
)

// func ioutil.NopCloser(r io.Reader) io.ReadCloser
func execNopCloser(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ioutil.NopCloser(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func ioutil.ReadAll(r io.Reader) ([]byte, error)
func execReadAll(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := ioutil.ReadAll(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// func ioutil.ReadDir(dirname string) ([]os.FileInfo, error)
func execReadDir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := ioutil.ReadDir(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func ioutil.ReadFile(filename string) ([]byte, error)
func execReadFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := ioutil.ReadFile(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func ioutil.TempDir(dir string, pattern string) (name string, err error)
func execTempDir(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := ioutil.TempDir(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func ioutil.TempFile(dir string, pattern string) (f *os.File, err error)
func execTempFile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := ioutil.TempFile(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func ioutil.WriteFile(filename string, data []byte, perm os.FileMode) error
func execWriteFile(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := ioutil.WriteFile(args[0].(string), args[1].([]byte), os.FileMode(args[2].(uint32)))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("io/ioutil")

func init() {
	I.RegisterVars(
		I.Var("Discard", &ioutil.Discard),
	)
	I.RegisterFuncs(
		I.Func("NopCloser", ioutil.NopCloser, execNopCloser),
		I.Func("ReadAll", ioutil.ReadAll, execReadAll),
		I.Func("ReadDir", ioutil.ReadDir, execReadDir),
		I.Func("ReadFile", ioutil.ReadFile, execReadFile),
		I.Func("TempDir", ioutil.TempDir, execTempDir),
		I.Func("TempFile", ioutil.TempFile, execTempFile),
		I.Func("WriteFile", ioutil.WriteFile, execWriteFile),
	)
}
