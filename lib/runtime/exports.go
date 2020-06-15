package runtime

import (
	"reflect"
	"runtime"
	"unsafe"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func runtime.BlockProfile(p []runtime.BlockProfileRecord) (n int, ok bool)
func execBlockProfile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := runtime.BlockProfile(args[0].([]runtime.BlockProfileRecord))
	p.Ret(1, ret, ret1)
}

// func runtime.Breakpoint()
func execBreakpoint(_ int, p *gop.Context) {
	runtime.Breakpoint()
}

// func runtime.CPUProfile() []byte
func execCPUProfile(_ int, p *gop.Context) {
	ret := runtime.CPUProfile()
	p.Ret(0, ret)
}

// func runtime.Caller(skip int) (pc uintptr, file string, line int, ok bool)
func execCaller(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2, ret3 := runtime.Caller(args[0].(int))
	p.Ret(1, ret, ret1, ret2, ret3)
}

// func runtime.Callers(skip int, pc []uintptr) int
func execCallers(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := runtime.Callers(args[0].(int), args[1].([]uintptr))
	p.Ret(2, ret)
}

// func runtime.CallersFrames(callers []uintptr) *runtime.Frames
func execCallersFrames(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := runtime.CallersFrames(args[0].([]uintptr))
	p.Ret(1, ret)
}

// func (*runtime.Frames).Next() (frame runtime.Frame, more bool)
func execmFramesNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*runtime.Frames).Next()
	p.Ret(1, ret, ret1)
}

// func (*runtime.Func).Entry() uintptr
func execmFuncEntry(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*runtime.Func).Entry()
	p.Ret(1, ret)
}

// func (*runtime.Func).FileLine(pc uintptr) (file string, line int)
func execmFuncFileLine(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*runtime.Func).FileLine(args[1].(uintptr))
	p.Ret(2, ret, ret1)
}

// func (*runtime.Func).Name() string
func execmFuncName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*runtime.Func).Name()
	p.Ret(1, ret)
}

// func runtime.FuncForPC(pc uintptr) *runtime.Func
func execFuncForPC(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := runtime.FuncForPC(args[0].(uintptr))
	p.Ret(1, ret)
}

// func runtime.GC()
func execGC(_ int, p *gop.Context) {
	runtime.GC()
}

// func runtime.GOMAXPROCS(n int) int
func execGOMAXPROCS(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := runtime.GOMAXPROCS(args[0].(int))
	p.Ret(1, ret)
}

// func runtime.GOROOT() string
func execGOROOT(_ int, p *gop.Context) {
	ret := runtime.GOROOT()
	p.Ret(0, ret)
}

// func runtime.Goexit()
func execGoexit(_ int, p *gop.Context) {
	runtime.Goexit()
}

// func runtime.GoroutineProfile(p []runtime.StackRecord) (n int, ok bool)
func execGoroutineProfile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := runtime.GoroutineProfile(args[0].([]runtime.StackRecord))
	p.Ret(1, ret, ret1)
}

// func runtime.Gosched()
func execGosched(_ int, p *gop.Context) {
	runtime.Gosched()
}

// func runtime.KeepAlive(x interface{})
func execKeepAlive(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	runtime.KeepAlive(args[0].(interface{}))
}

// func runtime.LockOSThread()
func execLockOSThread(_ int, p *gop.Context) {
	runtime.LockOSThread()
}

// func runtime.MemProfile(p []runtime.MemProfileRecord, inuseZero bool) (n int, ok bool)
func execMemProfile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := runtime.MemProfile(args[0].([]runtime.MemProfileRecord), args[1].(bool))
	p.Ret(2, ret, ret1)
}

// func (*runtime.MemProfileRecord).InUseBytes() int64
func execmMemProfileRecordInUseBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*runtime.MemProfileRecord).InUseBytes()
	p.Ret(1, ret)
}

// func (*runtime.MemProfileRecord).InUseObjects() int64
func execmMemProfileRecordInUseObjects(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*runtime.MemProfileRecord).InUseObjects()
	p.Ret(1, ret)
}

// func (*runtime.MemProfileRecord).Stack() []uintptr
func execmMemProfileRecordStack(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*runtime.MemProfileRecord).Stack()
	p.Ret(1, ret)
}

// func runtime.MutexProfile(p []runtime.BlockProfileRecord) (n int, ok bool)
func execMutexProfile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := runtime.MutexProfile(args[0].([]runtime.BlockProfileRecord))
	p.Ret(1, ret, ret1)
}

// func runtime.NumCPU() int
func execNumCPU(_ int, p *gop.Context) {
	ret := runtime.NumCPU()
	p.Ret(0, ret)
}

// func runtime.NumCgoCall() int64
func execNumCgoCall(_ int, p *gop.Context) {
	ret := runtime.NumCgoCall()
	p.Ret(0, ret)
}

// func runtime.NumGoroutine() int
func execNumGoroutine(_ int, p *gop.Context) {
	ret := runtime.NumGoroutine()
	p.Ret(0, ret)
}

// func runtime.ReadMemStats(m *runtime.MemStats)
func execReadMemStats(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	runtime.ReadMemStats(args[0].(*runtime.MemStats))
}

// func runtime.ReadTrace() []byte
func execReadTrace(_ int, p *gop.Context) {
	ret := runtime.ReadTrace()
	p.Ret(0, ret)
}

// func runtime.SetBlockProfileRate(rate int)
func execSetBlockProfileRate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	runtime.SetBlockProfileRate(args[0].(int))
}

// func runtime.SetCPUProfileRate(hz int)
func execSetCPUProfileRate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	runtime.SetCPUProfileRate(args[0].(int))
}

// func runtime.SetCgoTraceback(version int, traceback unsafe.Pointer, context unsafe.Pointer, symbolizer unsafe.Pointer)
func execSetCgoTraceback(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	runtime.SetCgoTraceback(args[0].(int), args[1].(unsafe.Pointer), args[2].(unsafe.Pointer), args[3].(unsafe.Pointer))
}

// func runtime.SetFinalizer(obj interface{}, finalizer interface{})
func execSetFinalizer(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	runtime.SetFinalizer(args[0].(interface{}), args[1].(interface{}))
}

// func runtime.SetMutexProfileFraction(rate int) int
func execSetMutexProfileFraction(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := runtime.SetMutexProfileFraction(args[0].(int))
	p.Ret(1, ret)
}

// func runtime.Stack(buf []byte, all bool) int
func execStack(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := runtime.Stack(args[0].([]byte), args[1].(bool))
	p.Ret(2, ret)
}

// func (*runtime.StackRecord).Stack() []uintptr
func execmStackRecordStack(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*runtime.StackRecord).Stack()
	p.Ret(1, ret)
}

// func runtime.StartTrace() error
func execStartTrace(_ int, p *gop.Context) {
	ret := runtime.StartTrace()
	p.Ret(0, ret)
}

// func runtime.StopTrace()
func execStopTrace(_ int, p *gop.Context) {
	runtime.StopTrace()
}

// func runtime.ThreadCreateProfile(p []runtime.StackRecord) (n int, ok bool)
func execThreadCreateProfile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := runtime.ThreadCreateProfile(args[0].([]runtime.StackRecord))
	p.Ret(1, ret, ret1)
}

// func (*runtime.TypeAssertionError).Error() string
func execmTypeAssertionErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*runtime.TypeAssertionError).Error()
	p.Ret(1, ret)
}

// func (*runtime.TypeAssertionError).RuntimeError()
func execmTypeAssertionErrorRuntimeError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*runtime.TypeAssertionError).RuntimeError()
}

// func runtime.UnlockOSThread()
func execUnlockOSThread(_ int, p *gop.Context) {
	runtime.UnlockOSThread()
}

// func runtime.Version() string
func execVersion(_ int, p *gop.Context) {
	ret := runtime.Version()
	p.Ret(0, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("runtime")

func init() {
	I.RegisterConsts(
		I.Const("Compiler", qspec.ConstBoundString, runtime.Compiler),
		I.Const("GOARCH", reflect.String, runtime.GOARCH),
		I.Const("GOOS", reflect.String, runtime.GOOS),
	)
	I.RegisterVars(
		I.Var("MemProfileRate", &runtime.MemProfileRate),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*runtime.BlockProfileRecord)(nil))),
		I.Rtype(reflect.TypeOf((*runtime.Frame)(nil))),
		I.Rtype(reflect.TypeOf((*runtime.Frames)(nil))),
		I.Rtype(reflect.TypeOf((*runtime.Func)(nil))),
		I.Rtype(reflect.TypeOf((*runtime.MemProfileRecord)(nil))),
		I.Rtype(reflect.TypeOf((*runtime.MemStats)(nil))),
		I.Rtype(reflect.TypeOf((*runtime.StackRecord)(nil))),
		I.Rtype(reflect.TypeOf((*runtime.TypeAssertionError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("BlockProfile", runtime.BlockProfile, execBlockProfile),
		I.Func("Breakpoint", runtime.Breakpoint, execBreakpoint),
		I.Func("CPUProfile", runtime.CPUProfile, execCPUProfile),
		I.Func("Caller", runtime.Caller, execCaller),
		I.Func("Callers", runtime.Callers, execCallers),
		I.Func("CallersFrames", runtime.CallersFrames, execCallersFrames),
		I.Func("(*Frames).Next", (*runtime.Frames).Next, execmFramesNext),
		I.Func("(*Func).Entry", (*runtime.Func).Entry, execmFuncEntry),
		I.Func("(*Func).FileLine", (*runtime.Func).FileLine, execmFuncFileLine),
		I.Func("(*Func).Name", (*runtime.Func).Name, execmFuncName),
		I.Func("FuncForPC", runtime.FuncForPC, execFuncForPC),
		I.Func("GC", runtime.GC, execGC),
		I.Func("GOMAXPROCS", runtime.GOMAXPROCS, execGOMAXPROCS),
		I.Func("GOROOT", runtime.GOROOT, execGOROOT),
		I.Func("Goexit", runtime.Goexit, execGoexit),
		I.Func("GoroutineProfile", runtime.GoroutineProfile, execGoroutineProfile),
		I.Func("Gosched", runtime.Gosched, execGosched),
		I.Func("KeepAlive", runtime.KeepAlive, execKeepAlive),
		I.Func("LockOSThread", runtime.LockOSThread, execLockOSThread),
		I.Func("MemProfile", runtime.MemProfile, execMemProfile),
		I.Func("(*MemProfileRecord).InUseBytes", (*runtime.MemProfileRecord).InUseBytes, execmMemProfileRecordInUseBytes),
		I.Func("(*MemProfileRecord).InUseObjects", (*runtime.MemProfileRecord).InUseObjects, execmMemProfileRecordInUseObjects),
		I.Func("(*MemProfileRecord).Stack", (*runtime.MemProfileRecord).Stack, execmMemProfileRecordStack),
		I.Func("MutexProfile", runtime.MutexProfile, execMutexProfile),
		I.Func("NumCPU", runtime.NumCPU, execNumCPU),
		I.Func("NumCgoCall", runtime.NumCgoCall, execNumCgoCall),
		I.Func("NumGoroutine", runtime.NumGoroutine, execNumGoroutine),
		I.Func("ReadMemStats", runtime.ReadMemStats, execReadMemStats),
		I.Func("ReadTrace", runtime.ReadTrace, execReadTrace),
		I.Func("SetBlockProfileRate", runtime.SetBlockProfileRate, execSetBlockProfileRate),
		I.Func("SetCPUProfileRate", runtime.SetCPUProfileRate, execSetCPUProfileRate),
		I.Func("SetCgoTraceback", runtime.SetCgoTraceback, execSetCgoTraceback),
		I.Func("SetFinalizer", runtime.SetFinalizer, execSetFinalizer),
		I.Func("SetMutexProfileFraction", runtime.SetMutexProfileFraction, execSetMutexProfileFraction),
		I.Func("Stack", runtime.Stack, execStack),
		I.Func("(*StackRecord).Stack", (*runtime.StackRecord).Stack, execmStackRecordStack),
		I.Func("StartTrace", runtime.StartTrace, execStartTrace),
		I.Func("StopTrace", runtime.StopTrace, execStopTrace),
		I.Func("ThreadCreateProfile", runtime.ThreadCreateProfile, execThreadCreateProfile),
		I.Func("(*TypeAssertionError).Error", (*runtime.TypeAssertionError).Error, execmTypeAssertionErrorError),
		I.Func("(*TypeAssertionError).RuntimeError", (*runtime.TypeAssertionError).RuntimeError, execmTypeAssertionErrorRuntimeError),
		I.Func("UnlockOSThread", runtime.UnlockOSThread, execUnlockOSThread),
		I.Func("Version", runtime.Version, execVersion),
	)
}
