package signal

import (
	"os"
	"os/signal"

	"github.com/qiniu/goplus/gop"
)

// func signal.Ignore(sig ..os.Signal)
func execIgnore(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []os.Signal {
		ret := make([]os.Signal, len(args))
		for i, arg := range args {
			ret[i] = arg.(os.Signal)
		}
		return ret
	}
	signal.Ignore(conv(args[0:])...)
}

// func signal.Ignored(sig os.Signal) bool
func execIgnored(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := signal.Ignored(args[0].(os.Signal))
	p.Ret(1, ret)
}

// func signal.Notify(c chan<- os.Signal, sig ..os.Signal)
func execNotify(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []os.Signal {
		ret := make([]os.Signal, len(args))
		for i, arg := range args {
			ret[i] = arg.(os.Signal)
		}
		return ret
	}
	signal.Notify(args[0].(chan<- os.Signal), conv(args[1:])...)
}

// func signal.Reset(sig ..os.Signal)
func execReset(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []os.Signal {
		ret := make([]os.Signal, len(args))
		for i, arg := range args {
			ret[i] = arg.(os.Signal)
		}
		return ret
	}
	signal.Reset(conv(args[0:])...)
}

// func signal.Stop(c chan<- os.Signal)
func execStop(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	signal.Stop(args[0].(chan<- os.Signal))
}

// I is a Go package instance.
var I = gop.NewGoPackage("os/signal")

func init() {
	I.RegisterFuncs(
		I.Func("Ignored", signal.Ignored, execIgnored),
		I.Func("Stop", signal.Stop, execStop),
	)
	I.RegisterFuncvs(
		I.Funcv("Ignore", signal.Ignore, execIgnore),
		I.Funcv("Notify", signal.Notify, execNotify),
		I.Funcv("Reset", signal.Reset, execReset),
	)
}
