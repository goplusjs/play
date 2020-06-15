package sync

import (
	"reflect"
	"sync"

	"github.com/qiniu/goplus/gop"
)

// func (*sync.Cond).Broadcast()
func execmCondBroadcast(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.Cond).Broadcast()
}

// func (*sync.Cond).Signal()
func execmCondSignal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.Cond).Signal()
}

// func (*sync.Cond).Wait()
func execmCondWait(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.Cond).Wait()
}

// func (*sync.Map).Delete(key interface{})
func execmMapDelete(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sync.Map).Delete(args[1].(interface{}))
}

// func (*sync.Map).Load(key interface{}) (value interface{}, ok bool)
func execmMapLoad(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*sync.Map).Load(args[1].(interface{}))
	p.Ret(2, ret, ret1)
}

// func (*sync.Map).LoadOrStore(key interface{}, value interface{}) (actual interface{}, loaded bool)
func execmMapLoadOrStore(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*sync.Map).LoadOrStore(args[1].(interface{}), args[2].(interface{}))
	p.Ret(3, ret, ret1)
}

// func (*sync.Map).Range(f func(key interface{}, value interface{}) bool)
func execmMapRange(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sync.Map).Range(args[1].(func(key interface{}, value interface{}) bool))
}

// func (*sync.Map).Store(key interface{}, value interface{})
func execmMapStore(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*sync.Map).Store(args[1].(interface{}), args[2].(interface{}))
}

// func (*sync.Mutex).Lock()
func execmMutexLock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.Mutex).Lock()
}

// func (*sync.Mutex).Unlock()
func execmMutexUnlock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.Mutex).Unlock()
}

// func sync.NewCond(l sync.Locker) *sync.Cond
func execNewCond(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sync.NewCond(args[0].(sync.Locker))
	p.Ret(1, ret)
}

// func (*sync.Once).Do(f func())
func execmOnceDo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sync.Once).Do(args[1].(func()))
}

// func (*sync.Pool).Get() interface{}
func execmPoolGet(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sync.Pool).Get()
	p.Ret(1, ret)
}

// func (*sync.Pool).Put(x interface{})
func execmPoolPut(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sync.Pool).Put(args[1].(interface{}))
}

// func (*sync.RWMutex).Lock()
func execmRWMutexLock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.RWMutex).Lock()
}

// func (*sync.RWMutex).RLock()
func execmRWMutexRLock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.RWMutex).RLock()
}

// func (*sync.RWMutex).RLocker() sync.Locker
func execmRWMutexRLocker(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sync.RWMutex).RLocker()
	p.Ret(1, ret)
}

// func (*sync.RWMutex).RUnlock()
func execmRWMutexRUnlock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.RWMutex).RUnlock()
}

// func (*sync.RWMutex).Unlock()
func execmRWMutexUnlock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.RWMutex).Unlock()
}

// func (*sync.WaitGroup).Add(delta int)
func execmWaitGroupAdd(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sync.WaitGroup).Add(args[1].(int))
}

// func (*sync.WaitGroup).Done()
func execmWaitGroupDone(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.WaitGroup).Done()
}

// func (*sync.WaitGroup).Wait()
func execmWaitGroupWait(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*sync.WaitGroup).Wait()
}

// I is a Go package instance.
var I = gop.NewGoPackage("sync")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*sync.Cond)(nil))),
		I.Rtype(reflect.TypeOf((*sync.Map)(nil))),
		I.Rtype(reflect.TypeOf((*sync.Mutex)(nil))),
		I.Rtype(reflect.TypeOf((*sync.Once)(nil))),
		I.Rtype(reflect.TypeOf((*sync.Pool)(nil))),
		I.Rtype(reflect.TypeOf((*sync.RWMutex)(nil))),
		I.Rtype(reflect.TypeOf((*sync.WaitGroup)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Cond).Broadcast", (*sync.Cond).Broadcast, execmCondBroadcast),
		I.Func("(*Cond).Signal", (*sync.Cond).Signal, execmCondSignal),
		I.Func("(*Cond).Wait", (*sync.Cond).Wait, execmCondWait),
		I.Func("(*Map).Delete", (*sync.Map).Delete, execmMapDelete),
		I.Func("(*Map).Load", (*sync.Map).Load, execmMapLoad),
		I.Func("(*Map).LoadOrStore", (*sync.Map).LoadOrStore, execmMapLoadOrStore),
		I.Func("(*Map).Range", (*sync.Map).Range, execmMapRange),
		I.Func("(*Map).Store", (*sync.Map).Store, execmMapStore),
		I.Func("(*Mutex).Lock", (*sync.Mutex).Lock, execmMutexLock),
		I.Func("(*Mutex).Unlock", (*sync.Mutex).Unlock, execmMutexUnlock),
		I.Func("NewCond", sync.NewCond, execNewCond),
		I.Func("(*Once).Do", (*sync.Once).Do, execmOnceDo),
		I.Func("(*Pool).Get", (*sync.Pool).Get, execmPoolGet),
		I.Func("(*Pool).Put", (*sync.Pool).Put, execmPoolPut),
		I.Func("(*RWMutex).Lock", (*sync.RWMutex).Lock, execmRWMutexLock),
		I.Func("(*RWMutex).RLock", (*sync.RWMutex).RLock, execmRWMutexRLock),
		I.Func("(*RWMutex).RLocker", (*sync.RWMutex).RLocker, execmRWMutexRLocker),
		I.Func("(*RWMutex).RUnlock", (*sync.RWMutex).RUnlock, execmRWMutexRUnlock),
		I.Func("(*RWMutex).Unlock", (*sync.RWMutex).Unlock, execmRWMutexUnlock),
		I.Func("(*WaitGroup).Add", (*sync.WaitGroup).Add, execmWaitGroupAdd),
		I.Func("(*WaitGroup).Done", (*sync.WaitGroup).Done, execmWaitGroupDone),
		I.Func("(*WaitGroup).Wait", (*sync.WaitGroup).Wait, execmWaitGroupWait),
	)
}
