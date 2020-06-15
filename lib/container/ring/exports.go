package ring

import (
	"container/ring"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func ring.New(n int) *ring.Ring
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ring.New(args[0].(int))
	p.Ret(1, ret)
}

// func (*ring.Ring).Do(f func(interface{}))
func execmRingDo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*ring.Ring).Do(args[1].(func(interface{})))
}

// func (*ring.Ring).Len() int
func execmRingLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ring.Ring).Len()
	p.Ret(1, ret)
}

// func (*ring.Ring).Link(s *ring.Ring) *ring.Ring
func execmRingLink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*ring.Ring).Link(args[1].(*ring.Ring))
	p.Ret(2, ret)
}

// func (*ring.Ring).Move(n int) *ring.Ring
func execmRingMove(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*ring.Ring).Move(args[1].(int))
	p.Ret(2, ret)
}

// func (*ring.Ring).Next() *ring.Ring
func execmRingNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ring.Ring).Next()
	p.Ret(1, ret)
}

// func (*ring.Ring).Prev() *ring.Ring
func execmRingPrev(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ring.Ring).Prev()
	p.Ret(1, ret)
}

// func (*ring.Ring).Unlink(n int) *ring.Ring
func execmRingUnlink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*ring.Ring).Unlink(args[1].(int))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("container/ring")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*ring.Ring)(nil))),
	)
	I.RegisterFuncs(
		I.Func("New", ring.New, execNew),
		I.Func("(*Ring).Do", (*ring.Ring).Do, execmRingDo),
		I.Func("(*Ring).Len", (*ring.Ring).Len, execmRingLen),
		I.Func("(*Ring).Link", (*ring.Ring).Link, execmRingLink),
		I.Func("(*Ring).Move", (*ring.Ring).Move, execmRingMove),
		I.Func("(*Ring).Next", (*ring.Ring).Next, execmRingNext),
		I.Func("(*Ring).Prev", (*ring.Ring).Prev, execmRingPrev),
		I.Func("(*Ring).Unlink", (*ring.Ring).Unlink, execmRingUnlink),
	)
}
