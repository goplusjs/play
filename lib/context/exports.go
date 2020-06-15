package context

import (
	"context"
	"time"

	"github.com/qiniu/goplus/gop"
)

// func context.Background() context.Context
func execBackground(_ int, p *gop.Context) {
	ret := context.Background()
	p.Ret(0, ret)
}

// func context.TODO() context.Context
func execTODO(_ int, p *gop.Context) {
	ret := context.TODO()
	p.Ret(0, ret)
}

// func context.WithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc)
func execWithCancel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := context.WithCancel(args[0].(context.Context))
	p.Ret(1, ret, ret1)
}

// func context.WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc)
func execWithDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := context.WithDeadline(args[0].(context.Context), args[1].(time.Time))
	p.Ret(2, ret, ret1)
}

// func context.WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
func execWithTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := context.WithTimeout(args[0].(context.Context), time.Duration(args[1].(int64)))
	p.Ret(2, ret, ret1)
}

// func context.WithValue(parent context.Context, key interface{}, val interface{}) context.Context
func execWithValue(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := context.WithValue(args[0].(context.Context), args[1].(interface{}), args[2].(interface{}))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("context")

func init() {
	I.RegisterVars(
		I.Var("Canceled", &context.Canceled),
		I.Var("DeadlineExceeded", &context.DeadlineExceeded),
	)
	I.RegisterFuncs(
		I.Func("Background", context.Background, execBackground),
		I.Func("TODO", context.TODO, execTODO),
		I.Func("WithCancel", context.WithCancel, execWithCancel),
		I.Func("WithDeadline", context.WithDeadline, execWithDeadline),
		I.Func("WithTimeout", context.WithTimeout, execWithTimeout),
		I.Func("WithValue", context.WithValue, execWithValue),
	)
}
