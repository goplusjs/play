package errors

import (
	"errors"

	"github.com/qiniu/goplus/gop"
)

// func errors.As(err error, target interface{}) bool
func execAs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := errors.As(args[0].(error), args[1].(interface{}))
	p.Ret(2, ret)
}

// func errors.Is(err error, target error) bool
func execIs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := errors.Is(args[0].(error), args[1].(error))
	p.Ret(2, ret)
}

// func errors.New(text string) error
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := errors.New(args[0].(string))
	p.Ret(1, ret)
}

// func errors.Unwrap(err error) error
func execUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := errors.Unwrap(args[0].(error))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("errors")

func init() {
	I.RegisterFuncs(
		I.Func("As", errors.As, execAs),
		I.Func("Is", errors.Is, execIs),
		I.Func("New", errors.New, execNew),
		I.Func("Unwrap", errors.Unwrap, execUnwrap),
	)
}
