package fmt

import (
	"fmt"
	"io"

	"github.com/qiniu/goplus/gop"
)

// func fmt.Errorf(format string, a ..interface{}) error
func execErrorf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := fmt.Errorf(args[0].(string), args[1:]...)
	p.Ret(arity, ret)
}

// func fmt.Fprint(w io.Writer, a ..interface{}) (n int, err error)
func execFprint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Fprint(args[0].(io.Writer), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Fprintf(w io.Writer, format string, a ..interface{}) (n int, err error)
func execFprintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Fprintf(args[0].(io.Writer), args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Fprintln(w io.Writer, a ..interface{}) (n int, err error)
func execFprintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Fprintln(args[0].(io.Writer), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Fscan(r io.Reader, a ..interface{}) (n int, err error)
func execFscan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Fscan(args[0].(io.Reader), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Fscanf(r io.Reader, format string, a ..interface{}) (n int, err error)
func execFscanf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Fscanf(args[0].(io.Reader), args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Fscanln(r io.Reader, a ..interface{}) (n int, err error)
func execFscanln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Fscanln(args[0].(io.Reader), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Print(a ..interface{}) (n int, err error)
func execPrint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Print(args[0:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Printf(format string, a ..interface{}) (n int, err error)
func execPrintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Printf(args[0].(string), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Println(a ..interface{}) (n int, err error)
func execPrintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Println(args[0:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Scan(a ..interface{}) (n int, err error)
func execScan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Scan(args[0:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Scanf(format string, a ..interface{}) (n int, err error)
func execScanf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Scanf(args[0].(string), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Scanln(a ..interface{}) (n int, err error)
func execScanln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Scanln(args[0:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Sprint(a ..interface{}) string
func execSprint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := fmt.Sprint(args[0:]...)
	p.Ret(arity, ret)
}

// func fmt.Sprintf(format string, a ..interface{}) string
func execSprintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := fmt.Sprintf(args[0].(string), args[1:]...)
	p.Ret(arity, ret)
}

// func fmt.Sprintln(a ..interface{}) string
func execSprintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := fmt.Sprintln(args[0:]...)
	p.Ret(arity, ret)
}

// func fmt.Sscan(str string, a ..interface{}) (n int, err error)
func execSscan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Sscan(args[0].(string), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Sscanf(str string, format string, a ..interface{}) (n int, err error)
func execSscanf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Sscanf(args[0].(string), args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func fmt.Sscanln(str string, a ..interface{}) (n int, err error)
func execSscanln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := fmt.Sscanln(args[0].(string), args[1:]...)
	p.Ret(arity, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("fmt")

func init() {
	I.RegisterFuncvs(
		I.Funcv("Errorf", fmt.Errorf, execErrorf),
		I.Funcv("Fprint", fmt.Fprint, execFprint),
		I.Funcv("Fprintf", fmt.Fprintf, execFprintf),
		I.Funcv("Fprintln", fmt.Fprintln, execFprintln),
		I.Funcv("Fscan", fmt.Fscan, execFscan),
		I.Funcv("Fscanf", fmt.Fscanf, execFscanf),
		I.Funcv("Fscanln", fmt.Fscanln, execFscanln),
		I.Funcv("Print", fmt.Print, execPrint),
		I.Funcv("Printf", fmt.Printf, execPrintf),
		I.Funcv("Println", fmt.Println, execPrintln),
		I.Funcv("Scan", fmt.Scan, execScan),
		I.Funcv("Scanf", fmt.Scanf, execScanf),
		I.Funcv("Scanln", fmt.Scanln, execScanln),
		I.Funcv("Sprint", fmt.Sprint, execSprint),
		I.Funcv("Sprintf", fmt.Sprintf, execSprintf),
		I.Funcv("Sprintln", fmt.Sprintln, execSprintln),
		I.Funcv("Sscan", fmt.Sscan, execSscan),
		I.Funcv("Sscanf", fmt.Sscanf, execSscanf),
		I.Funcv("Sscanln", fmt.Sscanln, execSscanln),
	)
}
