package list

import (
	"container/list"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*list.Element).Next() *list.Element
func execmElementNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*list.Element).Next()
	p.Ret(1, ret)
}

// func (*list.Element).Prev() *list.Element
func execmElementPrev(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*list.Element).Prev()
	p.Ret(1, ret)
}

// func (*list.List).Back() *list.Element
func execmListBack(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*list.List).Back()
	p.Ret(1, ret)
}

// func (*list.List).Front() *list.Element
func execmListFront(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*list.List).Front()
	p.Ret(1, ret)
}

// func (*list.List).Init() *list.List
func execmListInit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*list.List).Init()
	p.Ret(1, ret)
}

// func (*list.List).InsertAfter(v interface{}, mark *list.Element) *list.Element
func execmListInsertAfter(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*list.List).InsertAfter(args[1].(interface{}), args[2].(*list.Element))
	p.Ret(3, ret)
}

// func (*list.List).InsertBefore(v interface{}, mark *list.Element) *list.Element
func execmListInsertBefore(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*list.List).InsertBefore(args[1].(interface{}), args[2].(*list.Element))
	p.Ret(3, ret)
}

// func (*list.List).Len() int
func execmListLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*list.List).Len()
	p.Ret(1, ret)
}

// func (*list.List).MoveAfter(e *list.Element, mark *list.Element)
func execmListMoveAfter(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*list.List).MoveAfter(args[1].(*list.Element), args[2].(*list.Element))
}

// func (*list.List).MoveBefore(e *list.Element, mark *list.Element)
func execmListMoveBefore(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*list.List).MoveBefore(args[1].(*list.Element), args[2].(*list.Element))
}

// func (*list.List).MoveToBack(e *list.Element)
func execmListMoveToBack(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*list.List).MoveToBack(args[1].(*list.Element))
}

// func (*list.List).MoveToFront(e *list.Element)
func execmListMoveToFront(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*list.List).MoveToFront(args[1].(*list.Element))
}

// func (*list.List).PushBack(v interface{}) *list.Element
func execmListPushBack(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*list.List).PushBack(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*list.List).PushBackList(other *list.List)
func execmListPushBackList(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*list.List).PushBackList(args[1].(*list.List))
}

// func (*list.List).PushFront(v interface{}) *list.Element
func execmListPushFront(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*list.List).PushFront(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*list.List).PushFrontList(other *list.List)
func execmListPushFrontList(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*list.List).PushFrontList(args[1].(*list.List))
}

// func (*list.List).Remove(e *list.Element) interface{}
func execmListRemove(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*list.List).Remove(args[1].(*list.Element))
	p.Ret(2, ret)
}

// func list.New() *list.List
func execNew(_ int, p *gop.Context) {
	ret := list.New()
	p.Ret(0, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("container/list")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*list.Element)(nil))),
		I.Rtype(reflect.TypeOf((*list.List)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Element).Next", (*list.Element).Next, execmElementNext),
		I.Func("(*Element).Prev", (*list.Element).Prev, execmElementPrev),
		I.Func("(*List).Back", (*list.List).Back, execmListBack),
		I.Func("(*List).Front", (*list.List).Front, execmListFront),
		I.Func("(*List).Init", (*list.List).Init, execmListInit),
		I.Func("(*List).InsertAfter", (*list.List).InsertAfter, execmListInsertAfter),
		I.Func("(*List).InsertBefore", (*list.List).InsertBefore, execmListInsertBefore),
		I.Func("(*List).Len", (*list.List).Len, execmListLen),
		I.Func("(*List).MoveAfter", (*list.List).MoveAfter, execmListMoveAfter),
		I.Func("(*List).MoveBefore", (*list.List).MoveBefore, execmListMoveBefore),
		I.Func("(*List).MoveToBack", (*list.List).MoveToBack, execmListMoveToBack),
		I.Func("(*List).MoveToFront", (*list.List).MoveToFront, execmListMoveToFront),
		I.Func("(*List).PushBack", (*list.List).PushBack, execmListPushBack),
		I.Func("(*List).PushBackList", (*list.List).PushBackList, execmListPushBackList),
		I.Func("(*List).PushFront", (*list.List).PushFront, execmListPushFront),
		I.Func("(*List).PushFrontList", (*list.List).PushFrontList, execmListPushFrontList),
		I.Func("(*List).Remove", (*list.List).Remove, execmListRemove),
		I.Func("New", list.New, execNew),
	)
}
