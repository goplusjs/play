package heap

import (
	"container/heap"

	"github.com/qiniu/goplus/gop"
)

// func heap.Fix(h heap.Interface, i int)
func execFix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	heap.Fix(args[0].(heap.Interface), args[1].(int))
}

// func heap.Init(h heap.Interface)
func execInit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	heap.Init(args[0].(heap.Interface))
}

// func heap.Pop(h heap.Interface) interface{}
func execPop(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := heap.Pop(args[0].(heap.Interface))
	p.Ret(1, ret)
}

// func heap.Push(h heap.Interface, x interface{})
func execPush(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	heap.Push(args[0].(heap.Interface), args[1].(interface{}))
}

// func heap.Remove(h heap.Interface, i int) interface{}
func execRemove(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := heap.Remove(args[0].(heap.Interface), args[1].(int))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("container/heap")

func init() {
	I.RegisterFuncs(
		I.Func("Fix", heap.Fix, execFix),
		I.Func("Init", heap.Init, execInit),
		I.Func("Pop", heap.Pop, execPop),
		I.Func("Push", heap.Push, execPush),
		I.Func("Remove", heap.Remove, execRemove),
	)
}
