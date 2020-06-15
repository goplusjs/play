package trace

import (
	"context"
	"io"
	"reflect"
	"runtime/trace"

	"github.com/qiniu/goplus/gop"
)

// func trace.IsEnabled() bool
func execIsEnabled(_ int, p *gop.Context) {
	ret := trace.IsEnabled()
	p.Ret(0, ret)
}

// func trace.Log(ctx context.Context, category string, message string)
func execLog(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	trace.Log(args[0].(context.Context), args[1].(string), args[2].(string))
}

// func trace.Logf(ctx context.Context, category string, format string, args ..interface{})
func execLogf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	trace.Logf(args[0].(context.Context), args[1].(string), args[2].(string), args[3:]...)
}

// func trace.NewTask(pctx context.Context, taskType string) (ctx context.Context, task *trace.Task)
func execNewTask(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := trace.NewTask(args[0].(context.Context), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*trace.Region).End()
func execmRegionEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*trace.Region).End()
}

// func trace.Start(w io.Writer) error
func execStart(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := trace.Start(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func trace.StartRegion(ctx context.Context, regionType string) *trace.Region
func execStartRegion(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := trace.StartRegion(args[0].(context.Context), args[1].(string))
	p.Ret(2, ret)
}

// func trace.Stop()
func execStop(_ int, p *gop.Context) {
	trace.Stop()
}

// func (*trace.Task).End()
func execmTaskEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*trace.Task).End()
}

// func trace.WithRegion(ctx context.Context, regionType string, fn func())
func execWithRegion(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	trace.WithRegion(args[0].(context.Context), args[1].(string), args[2].(func()))
}

// I is a Go package instance.
var I = gop.NewGoPackage("runtime/trace")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*trace.Region)(nil))),
		I.Rtype(reflect.TypeOf((*trace.Task)(nil))),
	)
	I.RegisterFuncs(
		I.Func("IsEnabled", trace.IsEnabled, execIsEnabled),
		I.Func("Log", trace.Log, execLog),
		I.Func("NewTask", trace.NewTask, execNewTask),
		I.Func("(*Region).End", (*trace.Region).End, execmRegionEnd),
		I.Func("Start", trace.Start, execStart),
		I.Func("StartRegion", trace.StartRegion, execStartRegion),
		I.Func("Stop", trace.Stop, execStop),
		I.Func("(*Task).End", (*trace.Task).End, execmTaskEnd),
		I.Func("WithRegion", trace.WithRegion, execWithRegion),
	)
	I.RegisterFuncvs(
		I.Funcv("Logf", trace.Logf, execLogf),
	)
}
