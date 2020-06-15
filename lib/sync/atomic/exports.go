package atomic

import (
	"reflect"
	"sync/atomic"
	"unsafe"

	"github.com/qiniu/goplus/gop"
)

// func atomic.AddInt32(addr *int32, delta int32) (new int32)
func execAddInt32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.AddInt32(args[0].(*int32), args[1].(int32))
	p.Ret(2, ret)
}

// func atomic.AddInt64(addr *int64, delta int64) (new int64)
func execAddInt64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.AddInt64(args[0].(*int64), args[1].(int64))
	p.Ret(2, ret)
}

// func atomic.AddUint32(addr *uint32, delta uint32) (new uint32)
func execAddUint32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.AddUint32(args[0].(*uint32), args[1].(uint32))
	p.Ret(2, ret)
}

// func atomic.AddUint64(addr *uint64, delta uint64) (new uint64)
func execAddUint64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.AddUint64(args[0].(*uint64), args[1].(uint64))
	p.Ret(2, ret)
}

// func atomic.AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
func execAddUintptr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.AddUintptr(args[0].(*uintptr), args[1].(uintptr))
	p.Ret(2, ret)
}

// func atomic.CompareAndSwapInt32(addr *int32, old int32, new int32) (swapped bool)
func execCompareAndSwapInt32(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := atomic.CompareAndSwapInt32(args[0].(*int32), args[1].(int32), args[2].(int32))
	p.Ret(3, ret)
}

// func atomic.CompareAndSwapInt64(addr *int64, old int64, new int64) (swapped bool)
func execCompareAndSwapInt64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := atomic.CompareAndSwapInt64(args[0].(*int64), args[1].(int64), args[2].(int64))
	p.Ret(3, ret)
}

// func atomic.CompareAndSwapPointer(addr *unsafe.Pointer, old unsafe.Pointer, new unsafe.Pointer) (swapped bool)
func execCompareAndSwapPointer(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := atomic.CompareAndSwapPointer(args[0].(*unsafe.Pointer), args[1].(unsafe.Pointer), args[2].(unsafe.Pointer))
	p.Ret(3, ret)
}

// func atomic.CompareAndSwapUint32(addr *uint32, old uint32, new uint32) (swapped bool)
func execCompareAndSwapUint32(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := atomic.CompareAndSwapUint32(args[0].(*uint32), args[1].(uint32), args[2].(uint32))
	p.Ret(3, ret)
}

// func atomic.CompareAndSwapUint64(addr *uint64, old uint64, new uint64) (swapped bool)
func execCompareAndSwapUint64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := atomic.CompareAndSwapUint64(args[0].(*uint64), args[1].(uint64), args[2].(uint64))
	p.Ret(3, ret)
}

// func atomic.CompareAndSwapUintptr(addr *uintptr, old uintptr, new uintptr) (swapped bool)
func execCompareAndSwapUintptr(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := atomic.CompareAndSwapUintptr(args[0].(*uintptr), args[1].(uintptr), args[2].(uintptr))
	p.Ret(3, ret)
}

// func atomic.LoadInt32(addr *int32) (val int32)
func execLoadInt32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := atomic.LoadInt32(args[0].(*int32))
	p.Ret(1, ret)
}

// func atomic.LoadInt64(addr *int64) (val int64)
func execLoadInt64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := atomic.LoadInt64(args[0].(*int64))
	p.Ret(1, ret)
}

// func atomic.LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func execLoadPointer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := atomic.LoadPointer(args[0].(*unsafe.Pointer))
	p.Ret(1, ret)
}

// func atomic.LoadUint32(addr *uint32) (val uint32)
func execLoadUint32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := atomic.LoadUint32(args[0].(*uint32))
	p.Ret(1, ret)
}

// func atomic.LoadUint64(addr *uint64) (val uint64)
func execLoadUint64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := atomic.LoadUint64(args[0].(*uint64))
	p.Ret(1, ret)
}

// func atomic.LoadUintptr(addr *uintptr) (val uintptr)
func execLoadUintptr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := atomic.LoadUintptr(args[0].(*uintptr))
	p.Ret(1, ret)
}

// func atomic.StoreInt32(addr *int32, val int32)
func execStoreInt32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	atomic.StoreInt32(args[0].(*int32), args[1].(int32))
}

// func atomic.StoreInt64(addr *int64, val int64)
func execStoreInt64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	atomic.StoreInt64(args[0].(*int64), args[1].(int64))
}

// func atomic.StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
func execStorePointer(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	atomic.StorePointer(args[0].(*unsafe.Pointer), args[1].(unsafe.Pointer))
}

// func atomic.StoreUint32(addr *uint32, val uint32)
func execStoreUint32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	atomic.StoreUint32(args[0].(*uint32), args[1].(uint32))
}

// func atomic.StoreUint64(addr *uint64, val uint64)
func execStoreUint64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	atomic.StoreUint64(args[0].(*uint64), args[1].(uint64))
}

// func atomic.StoreUintptr(addr *uintptr, val uintptr)
func execStoreUintptr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	atomic.StoreUintptr(args[0].(*uintptr), args[1].(uintptr))
}

// func atomic.SwapInt32(addr *int32, new int32) (old int32)
func execSwapInt32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.SwapInt32(args[0].(*int32), args[1].(int32))
	p.Ret(2, ret)
}

// func atomic.SwapInt64(addr *int64, new int64) (old int64)
func execSwapInt64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.SwapInt64(args[0].(*int64), args[1].(int64))
	p.Ret(2, ret)
}

// func atomic.SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
func execSwapPointer(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.SwapPointer(args[0].(*unsafe.Pointer), args[1].(unsafe.Pointer))
	p.Ret(2, ret)
}

// func atomic.SwapUint32(addr *uint32, new uint32) (old uint32)
func execSwapUint32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.SwapUint32(args[0].(*uint32), args[1].(uint32))
	p.Ret(2, ret)
}

// func atomic.SwapUint64(addr *uint64, new uint64) (old uint64)
func execSwapUint64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.SwapUint64(args[0].(*uint64), args[1].(uint64))
	p.Ret(2, ret)
}

// func atomic.SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
func execSwapUintptr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := atomic.SwapUintptr(args[0].(*uintptr), args[1].(uintptr))
	p.Ret(2, ret)
}

// func (*atomic.Value).Load() (x interface{})
func execmValueLoad(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*atomic.Value).Load()
	p.Ret(1, ret)
}

// func (*atomic.Value).Store(x interface{})
func execmValueStore(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*atomic.Value).Store(args[1].(interface{}))
}

// I is a Go package instance.
var I = gop.NewGoPackage("sync/atomic")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*atomic.Value)(nil))),
	)
	I.RegisterFuncs(
		I.Func("AddInt32", atomic.AddInt32, execAddInt32),
		I.Func("AddInt64", atomic.AddInt64, execAddInt64),
		I.Func("AddUint32", atomic.AddUint32, execAddUint32),
		I.Func("AddUint64", atomic.AddUint64, execAddUint64),
		I.Func("AddUintptr", atomic.AddUintptr, execAddUintptr),
		I.Func("CompareAndSwapInt32", atomic.CompareAndSwapInt32, execCompareAndSwapInt32),
		I.Func("CompareAndSwapInt64", atomic.CompareAndSwapInt64, execCompareAndSwapInt64),
		I.Func("CompareAndSwapPointer", atomic.CompareAndSwapPointer, execCompareAndSwapPointer),
		I.Func("CompareAndSwapUint32", atomic.CompareAndSwapUint32, execCompareAndSwapUint32),
		I.Func("CompareAndSwapUint64", atomic.CompareAndSwapUint64, execCompareAndSwapUint64),
		I.Func("CompareAndSwapUintptr", atomic.CompareAndSwapUintptr, execCompareAndSwapUintptr),
		I.Func("LoadInt32", atomic.LoadInt32, execLoadInt32),
		I.Func("LoadInt64", atomic.LoadInt64, execLoadInt64),
		I.Func("LoadPointer", atomic.LoadPointer, execLoadPointer),
		I.Func("LoadUint32", atomic.LoadUint32, execLoadUint32),
		I.Func("LoadUint64", atomic.LoadUint64, execLoadUint64),
		I.Func("LoadUintptr", atomic.LoadUintptr, execLoadUintptr),
		I.Func("StoreInt32", atomic.StoreInt32, execStoreInt32),
		I.Func("StoreInt64", atomic.StoreInt64, execStoreInt64),
		I.Func("StorePointer", atomic.StorePointer, execStorePointer),
		I.Func("StoreUint32", atomic.StoreUint32, execStoreUint32),
		I.Func("StoreUint64", atomic.StoreUint64, execStoreUint64),
		I.Func("StoreUintptr", atomic.StoreUintptr, execStoreUintptr),
		I.Func("SwapInt32", atomic.SwapInt32, execSwapInt32),
		I.Func("SwapInt64", atomic.SwapInt64, execSwapInt64),
		I.Func("SwapPointer", atomic.SwapPointer, execSwapPointer),
		I.Func("SwapUint32", atomic.SwapUint32, execSwapUint32),
		I.Func("SwapUint64", atomic.SwapUint64, execSwapUint64),
		I.Func("SwapUintptr", atomic.SwapUintptr, execSwapUintptr),
		I.Func("(*Value).Load", (*atomic.Value).Load, execmValueLoad),
		I.Func("(*Value).Store", (*atomic.Value).Store, execmValueStore),
	)
}
