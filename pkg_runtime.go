// export by github.com/goplus/igop/cmd/qexp

package main

import (
	q "runtime"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "runtime",
		Path: "runtime",
		Deps: map[string]string{
			"internal/bytealg":        "bytealg",
			"internal/cpu":            "cpu",
			"runtime/internal/atomic": "atomic",
			"runtime/internal/math":   "math",
			"runtime/internal/sys":    "sys",
			"unsafe":                  "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Error": reflect.TypeOf((*q.Error)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Frame":              reflect.TypeOf((*q.Frame)(nil)).Elem(),
			"Frames":             reflect.TypeOf((*q.Frames)(nil)).Elem(),
			"Func":               reflect.TypeOf((*q.Func)(nil)).Elem(),
			"TypeAssertionError": reflect.TypeOf((*q.TypeAssertionError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Funcs: map[string]reflect.Value{
			"Breakpoint":              reflect.ValueOf(q.Breakpoint),
			"Caller":                  reflect.ValueOf(q.Caller),
			"Callers":                 reflect.ValueOf(q.Callers),
			"CallersFrames":           reflect.ValueOf(q.CallersFrames),
			"FuncForPC":               reflect.ValueOf(q.FuncForPC),
			"GC":                      reflect.ValueOf(q.GC),
			"GOMAXPROCS":              reflect.ValueOf(q.GOMAXPROCS),
			"GOROOT":                  reflect.ValueOf(q.GOROOT),
			"Goexit":                  reflect.ValueOf(q.Goexit),
			"Gosched":                 reflect.ValueOf(q.Gosched),
			"KeepAlive":               reflect.ValueOf(q.KeepAlive),
			"LockOSThread":            reflect.ValueOf(q.LockOSThread),
			"NumCPU":                  reflect.ValueOf(q.NumCPU),
			"NumCgoCall":              reflect.ValueOf(q.NumCgoCall),
			"NumGoroutine":            reflect.ValueOf(q.NumGoroutine),
			"ReadMemStats":            reflect.ValueOf(q.ReadMemStats),
			"ReadTrace":               reflect.ValueOf(q.ReadTrace),
			"SetBlockProfileRate":     reflect.ValueOf(q.SetBlockProfileRate),
			"SetFinalizer":            reflect.ValueOf(q.SetFinalizer),
			"SetMutexProfileFraction": reflect.ValueOf(q.SetMutexProfileFraction),
			"Stack":                   reflect.ValueOf(q.Stack),
			"StartTrace":              reflect.ValueOf(q.StartTrace),
			"StopTrace":               reflect.ValueOf(q.StopTrace),
			"UnlockOSThread":          reflect.ValueOf(q.UnlockOSThread),
			"Version":                 reflect.ValueOf(q.Version),
		},
		TypedConsts: map[string]igop.TypedConst{
			"GOARCH": {reflect.TypeOf(q.GOARCH), constant.MakeString(q.GOOS)},
			"GOOS":   {reflect.TypeOf(q.GOOS), constant.MakeString(q.GOARCH)},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"Compiler": {"untyped string", constant.MakeString(q.Compiler)},
		},
	})
}
