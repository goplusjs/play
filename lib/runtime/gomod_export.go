// Package runtime provide Go+ "runtime" package, as "runtime" package in Go.
package runtime

import (
	runtime "runtime"

	gop "github.com/goplus/gop"
	qspec "github.com/goplus/gop/exec.spec"
)

func execGOMAXPROCS(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0 := runtime.GOMAXPROCS(args[0].(int))
	p.Ret(1, ret0)
}

func execGOROOT(_ int, p *gop.Context) {
	ret0 := runtime.GOROOT()
	p.Ret(0, ret0)
}

func execGoexit(_ int, p *gop.Context) {
	runtime.Goexit()
	p.PopN(0)
}

func execGosched(_ int, p *gop.Context) {
	runtime.Gosched()
	p.PopN(0)
}

func execKeepAlive(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	runtime.KeepAlive(args[0])
	p.PopN(1)
}

func execLockOSThread(_ int, p *gop.Context) {
	runtime.LockOSThread()
	p.PopN(0)
}

func execUnlockOSThread(_ int, p *gop.Context) {
	runtime.UnlockOSThread()
	p.PopN(0)
}

func execVersion(_ int, p *gop.Context) {
	ret0 := runtime.Version()
	p.Ret(0, ret0)
}

// I is a Go package instance.
var I = gop.NewGoPackage("runtime")

func init() {
	I.RegisterFuncs(
		I.Func("GOMAXPROCS", runtime.GOMAXPROCS, execGOMAXPROCS),
		I.Func("GOROOT", runtime.GOROOT, execGOROOT),
		I.Func("Goexit", runtime.Goexit, execGoexit),
		I.Func("Gosched", runtime.Gosched, execGosched),
		I.Func("KeepAlive", runtime.KeepAlive, execKeepAlive),
		I.Func("LockOSThread", runtime.LockOSThread, execLockOSThread),
		I.Func("UnlockOSThread", runtime.UnlockOSThread, execUnlockOSThread),
		I.Func("Version", runtime.Version, execVersion),
	)
	I.RegisterConsts(
		I.Const("Compiler", qspec.ConstBoundString, runtime.Compiler),
		I.Const("GOARCH", qspec.String, runtime.GOARCH),
		I.Const("GOOS", qspec.String, runtime.GOOS),
	)
}
