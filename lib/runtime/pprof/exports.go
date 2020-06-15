package pprof

import (
	"context"
	"io"
	"reflect"
	"runtime/pprof"

	"github.com/qiniu/goplus/gop"
)

// func pprof.Do(ctx context.Context, labels pprof.LabelSet, f func(context.Context))
func execDo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	pprof.Do(args[0].(context.Context), args[1].(pprof.LabelSet), args[2].(func(context.Context)))
}

// func pprof.ForLabels(ctx context.Context, f func(key string, value string) bool)
func execForLabels(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	pprof.ForLabels(args[0].(context.Context), args[1].(func(key string, value string) bool))
}

// func pprof.Label(ctx context.Context, key string) (string, bool)
func execLabel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := pprof.Label(args[0].(context.Context), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func pprof.Labels(args ..string) pprof.LabelSet
func execLabels(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret := pprof.Labels(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func pprof.Lookup(name string) *pprof.Profile
func execLookup(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := pprof.Lookup(args[0].(string))
	p.Ret(1, ret)
}

// func pprof.NewProfile(name string) *pprof.Profile
func execNewProfile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := pprof.NewProfile(args[0].(string))
	p.Ret(1, ret)
}

// func (*pprof.Profile).Add(value interface{}, skip int)
func execmProfileAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*pprof.Profile).Add(args[1].(interface{}), args[2].(int))
}

// func (*pprof.Profile).Count() int
func execmProfileCount(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*pprof.Profile).Count()
	p.Ret(1, ret)
}

// func (*pprof.Profile).Name() string
func execmProfileName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*pprof.Profile).Name()
	p.Ret(1, ret)
}

// func (*pprof.Profile).Remove(value interface{})
func execmProfileRemove(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*pprof.Profile).Remove(args[1].(interface{}))
}

// func (*pprof.Profile).WriteTo(w io.Writer, debug int) error
func execmProfileWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*pprof.Profile).WriteTo(args[1].(io.Writer), args[2].(int))
	p.Ret(3, ret)
}

// func pprof.Profiles() []*pprof.Profile
func execProfiles(_ int, p *gop.Context) {
	ret := pprof.Profiles()
	p.Ret(0, ret)
}

// func pprof.SetGoroutineLabels(ctx context.Context)
func execSetGoroutineLabels(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	pprof.SetGoroutineLabels(args[0].(context.Context))
}

// func pprof.StartCPUProfile(w io.Writer) error
func execStartCPUProfile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := pprof.StartCPUProfile(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func pprof.StopCPUProfile()
func execStopCPUProfile(_ int, p *gop.Context) {
	pprof.StopCPUProfile()
}

// func pprof.WithLabels(ctx context.Context, labels pprof.LabelSet) context.Context
func execWithLabels(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := pprof.WithLabels(args[0].(context.Context), args[1].(pprof.LabelSet))
	p.Ret(2, ret)
}

// func pprof.WriteHeapProfile(w io.Writer) error
func execWriteHeapProfile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := pprof.WriteHeapProfile(args[0].(io.Writer))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("runtime/pprof")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*pprof.LabelSet)(nil))),
		I.Rtype(reflect.TypeOf((*pprof.Profile)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Do", pprof.Do, execDo),
		I.Func("ForLabels", pprof.ForLabels, execForLabels),
		I.Func("Label", pprof.Label, execLabel),
		I.Func("Lookup", pprof.Lookup, execLookup),
		I.Func("NewProfile", pprof.NewProfile, execNewProfile),
		I.Func("(*Profile).Add", (*pprof.Profile).Add, execmProfileAdd),
		I.Func("(*Profile).Count", (*pprof.Profile).Count, execmProfileCount),
		I.Func("(*Profile).Name", (*pprof.Profile).Name, execmProfileName),
		I.Func("(*Profile).Remove", (*pprof.Profile).Remove, execmProfileRemove),
		I.Func("(*Profile).WriteTo", (*pprof.Profile).WriteTo, execmProfileWriteTo),
		I.Func("Profiles", pprof.Profiles, execProfiles),
		I.Func("SetGoroutineLabels", pprof.SetGoroutineLabels, execSetGoroutineLabels),
		I.Func("StartCPUProfile", pprof.StartCPUProfile, execStartCPUProfile),
		I.Func("StopCPUProfile", pprof.StopCPUProfile, execStopCPUProfile),
		I.Func("WithLabels", pprof.WithLabels, execWithLabels),
		I.Func("WriteHeapProfile", pprof.WriteHeapProfile, execWriteHeapProfile),
	)
	I.RegisterFuncvs(
		I.Funcv("Labels", pprof.Labels, execLabels),
	)
}
