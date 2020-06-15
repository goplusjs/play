package debug

import (
	"reflect"
	"runtime/debug"

	"github.com/qiniu/goplus/gop"
)

// func debug.FreeOSMemory()
func execFreeOSMemory(_ int, p *gop.Context) {
	debug.FreeOSMemory()
}

// func debug.PrintStack()
func execPrintStack(_ int, p *gop.Context) {
	debug.PrintStack()
}

// func debug.ReadBuildInfo() (info *debug.BuildInfo, ok bool)
func execReadBuildInfo(_ int, p *gop.Context) {
	ret, ret1 := debug.ReadBuildInfo()
	p.Ret(0, ret, ret1)
}

// func debug.ReadGCStats(stats *debug.GCStats)
func execReadGCStats(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	debug.ReadGCStats(args[0].(*debug.GCStats))
}

// func debug.SetGCPercent(percent int) int
func execSetGCPercent(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := debug.SetGCPercent(args[0].(int))
	p.Ret(1, ret)
}

// func debug.SetMaxStack(bytes int) int
func execSetMaxStack(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := debug.SetMaxStack(args[0].(int))
	p.Ret(1, ret)
}

// func debug.SetMaxThreads(threads int) int
func execSetMaxThreads(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := debug.SetMaxThreads(args[0].(int))
	p.Ret(1, ret)
}

// func debug.SetPanicOnFault(enabled bool) bool
func execSetPanicOnFault(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := debug.SetPanicOnFault(args[0].(bool))
	p.Ret(1, ret)
}

// func debug.SetTraceback(level string)
func execSetTraceback(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	debug.SetTraceback(args[0].(string))
}

// func debug.Stack() []byte
func execStack(_ int, p *gop.Context) {
	ret := debug.Stack()
	p.Ret(0, ret)
}

// func debug.WriteHeapDump(fd uintptr)
func execWriteHeapDump(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	debug.WriteHeapDump(args[0].(uintptr))
}

// I is a Go package instance.
var I = gop.NewGoPackage("runtime/debug")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*debug.BuildInfo)(nil))),
		I.Rtype(reflect.TypeOf((*debug.GCStats)(nil))),
		I.Rtype(reflect.TypeOf((*debug.Module)(nil))),
	)
	I.RegisterFuncs(
		I.Func("FreeOSMemory", debug.FreeOSMemory, execFreeOSMemory),
		I.Func("PrintStack", debug.PrintStack, execPrintStack),
		I.Func("ReadBuildInfo", debug.ReadBuildInfo, execReadBuildInfo),
		I.Func("ReadGCStats", debug.ReadGCStats, execReadGCStats),
		I.Func("SetGCPercent", debug.SetGCPercent, execSetGCPercent),
		I.Func("SetMaxStack", debug.SetMaxStack, execSetMaxStack),
		I.Func("SetMaxThreads", debug.SetMaxThreads, execSetMaxThreads),
		I.Func("SetPanicOnFault", debug.SetPanicOnFault, execSetPanicOnFault),
		I.Func("SetTraceback", debug.SetTraceback, execSetTraceback),
		I.Func("Stack", debug.Stack, execStack),
		I.Func("WriteHeapDump", debug.WriteHeapDump, execWriteHeapDump),
	)
}
