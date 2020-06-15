package exec

import (
	"context"
	"os/exec"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*exec.Cmd).CombinedOutput() ([]byte, error)
func execmCmdCombinedOutput(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*exec.Cmd).CombinedOutput()
	p.Ret(1, ret, ret1)
}

// func (*exec.Cmd).Output() ([]byte, error)
func execmCmdOutput(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*exec.Cmd).Output()
	p.Ret(1, ret, ret1)
}

// func (*exec.Cmd).Run() error
func execmCmdRun(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*exec.Cmd).Run()
	p.Ret(1, ret)
}

// func (*exec.Cmd).Start() error
func execmCmdStart(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*exec.Cmd).Start()
	p.Ret(1, ret)
}

// func (*exec.Cmd).StderrPipe() (io.ReadCloser, error)
func execmCmdStderrPipe(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*exec.Cmd).StderrPipe()
	p.Ret(1, ret, ret1)
}

// func (*exec.Cmd).StdinPipe() (io.WriteCloser, error)
func execmCmdStdinPipe(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*exec.Cmd).StdinPipe()
	p.Ret(1, ret, ret1)
}

// func (*exec.Cmd).StdoutPipe() (io.ReadCloser, error)
func execmCmdStdoutPipe(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*exec.Cmd).StdoutPipe()
	p.Ret(1, ret, ret1)
}

// func (*exec.Cmd).String() string
func execmCmdString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*exec.Cmd).String()
	p.Ret(1, ret)
}

// func (*exec.Cmd).Wait() error
func execmCmdWait(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*exec.Cmd).Wait()
	p.Ret(1, ret)
}

// func exec.Command(name string, arg ..string) *exec.Cmd
func execCommand(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret := exec.Command(args[0].(string), conv(args[1:])...)
	p.Ret(arity, ret)
}

// func exec.CommandContext(ctx context.Context, name string, arg ..string) *exec.Cmd
func execCommandContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret := exec.CommandContext(args[0].(context.Context), args[1].(string), conv(args[2:])...)
	p.Ret(arity, ret)
}

// func (*exec.Error).Error() string
func execmErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*exec.Error).Error()
	p.Ret(1, ret)
}

// func (*exec.Error).Unwrap() error
func execmErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*exec.Error).Unwrap()
	p.Ret(1, ret)
}

// func (*exec.ExitError).Error() string
func execmExitErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*exec.ExitError).Error()
	p.Ret(1, ret)
}

// func exec.LookPath(file string) (string, error)
func execLookPath(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := exec.LookPath(args[0].(string))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("os/exec")

func init() {
	I.RegisterVars(
		I.Var("ErrNotFound", &exec.ErrNotFound),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*exec.Cmd)(nil))),
		I.Rtype(reflect.TypeOf((*exec.Error)(nil))),
		I.Rtype(reflect.TypeOf((*exec.ExitError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Cmd).CombinedOutput", (*exec.Cmd).CombinedOutput, execmCmdCombinedOutput),
		I.Func("(*Cmd).Output", (*exec.Cmd).Output, execmCmdOutput),
		I.Func("(*Cmd).Run", (*exec.Cmd).Run, execmCmdRun),
		I.Func("(*Cmd).Start", (*exec.Cmd).Start, execmCmdStart),
		I.Func("(*Cmd).StderrPipe", (*exec.Cmd).StderrPipe, execmCmdStderrPipe),
		I.Func("(*Cmd).StdinPipe", (*exec.Cmd).StdinPipe, execmCmdStdinPipe),
		I.Func("(*Cmd).StdoutPipe", (*exec.Cmd).StdoutPipe, execmCmdStdoutPipe),
		I.Func("(*Cmd).String", (*exec.Cmd).String, execmCmdString),
		I.Func("(*Cmd).Wait", (*exec.Cmd).Wait, execmCmdWait),
		I.Func("(*Error).Error", (*exec.Error).Error, execmErrorError),
		I.Func("(*Error).Unwrap", (*exec.Error).Unwrap, execmErrorUnwrap),
		I.Func("(*ExitError).Error", (*exec.ExitError).Error, execmExitErrorError),
		I.Func("LookPath", exec.LookPath, execLookPath),
	)
	I.RegisterFuncvs(
		I.Funcv("Command", exec.Command, execCommand),
		I.Funcv("CommandContext", exec.CommandContext, execCommandContext),
	)
}
