package log

import (
	"io"
	"log"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func log.Fatal(v ..interface{})
func execFatal(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Fatal(args[0:]...)
}

// func log.Fatalf(format string, v ..interface{})
func execFatalf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Fatalf(args[0].(string), args[1:]...)
}

// func log.Fatalln(v ..interface{})
func execFatalln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Fatalln(args[0:]...)
}

// func log.Flags() int
func execFlags(_ int, p *gop.Context) {
	ret := log.Flags()
	p.Ret(0, ret)
}

// func (*log.Logger).Fatal(v ..interface{})
func execmLoggerFatal(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Fatal(args[1:]...)
}

// func (*log.Logger).Fatalf(format string, v ..interface{})
func execmLoggerFatalf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Fatalf(args[1].(string), args[2:]...)
}

// func (*log.Logger).Fatalln(v ..interface{})
func execmLoggerFatalln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Fatalln(args[1:]...)
}

// func (*log.Logger).Flags() int
func execmLoggerFlags(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*log.Logger).Flags()
	p.Ret(1, ret)
}

// func (*log.Logger).Output(calldepth int, s string) error
func execmLoggerOutput(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*log.Logger).Output(args[1].(int), args[2].(string))
	p.Ret(3, ret)
}

// func (*log.Logger).Panic(v ..interface{})
func execmLoggerPanic(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Panic(args[1:]...)
}

// func (*log.Logger).Panicf(format string, v ..interface{})
func execmLoggerPanicf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Panicf(args[1].(string), args[2:]...)
}

// func (*log.Logger).Panicln(v ..interface{})
func execmLoggerPanicln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Panicln(args[1:]...)
}

// func (*log.Logger).Prefix() string
func execmLoggerPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*log.Logger).Prefix()
	p.Ret(1, ret)
}

// func (*log.Logger).Print(v ..interface{})
func execmLoggerPrint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Print(args[1:]...)
}

// func (*log.Logger).Printf(format string, v ..interface{})
func execmLoggerPrintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Printf(args[1].(string), args[2:]...)
}

// func (*log.Logger).Println(v ..interface{})
func execmLoggerPrintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	args[0].(*log.Logger).Println(args[1:]...)
}

// func (*log.Logger).SetFlags(flag int)
func execmLoggerSetFlags(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*log.Logger).SetFlags(args[1].(int))
}

// func (*log.Logger).SetOutput(w io.Writer)
func execmLoggerSetOutput(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*log.Logger).SetOutput(args[1].(io.Writer))
}

// func (*log.Logger).SetPrefix(prefix string)
func execmLoggerSetPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*log.Logger).SetPrefix(args[1].(string))
}

// func (*log.Logger).Writer() io.Writer
func execmLoggerWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*log.Logger).Writer()
	p.Ret(1, ret)
}

// func log.New(out io.Writer, prefix string, flag int) *log.Logger
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := log.New(args[0].(io.Writer), args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func log.Output(calldepth int, s string) error
func execOutput(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := log.Output(args[0].(int), args[1].(string))
	p.Ret(2, ret)
}

// func log.Panic(v ..interface{})
func execPanic(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Panic(args[0:]...)
}

// func log.Panicf(format string, v ..interface{})
func execPanicf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Panicf(args[0].(string), args[1:]...)
}

// func log.Panicln(v ..interface{})
func execPanicln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Panicln(args[0:]...)
}

// func log.Prefix() string
func execPrefix(_ int, p *gop.Context) {
	ret := log.Prefix()
	p.Ret(0, ret)
}

// func log.Print(v ..interface{})
func execPrint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Print(args[0:]...)
}

// func log.Printf(format string, v ..interface{})
func execPrintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Printf(args[0].(string), args[1:]...)
}

// func log.Println(v ..interface{})
func execPrintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	log.Println(args[0:]...)
}

// func log.SetFlags(flag int)
func execSetFlags(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	log.SetFlags(args[0].(int))
}

// func log.SetOutput(w io.Writer)
func execSetOutput(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	log.SetOutput(args[0].(io.Writer))
}

// func log.SetPrefix(prefix string)
func execSetPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	log.SetPrefix(args[0].(string))
}

// func log.Writer() io.Writer
func execWriter(_ int, p *gop.Context) {
	ret := log.Writer()
	p.Ret(0, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("log")

func init() {
	I.RegisterConsts(
		I.Const("LUTC", qspec.ConstUnboundInt, log.LUTC),
		I.Const("Ldate", qspec.ConstUnboundInt, log.Ldate),
		I.Const("Llongfile", qspec.ConstUnboundInt, log.Llongfile),
		I.Const("Lmicroseconds", qspec.ConstUnboundInt, log.Lmicroseconds),
		I.Const("Lmsgprefix", qspec.ConstUnboundInt, log.Lmsgprefix),
		I.Const("Lshortfile", qspec.ConstUnboundInt, log.Lshortfile),
		I.Const("LstdFlags", qspec.ConstUnboundInt, log.LstdFlags),
		I.Const("Ltime", qspec.ConstUnboundInt, log.Ltime),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*log.Logger)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Flags", log.Flags, execFlags),
		I.Func("(*Logger).Flags", (*log.Logger).Flags, execmLoggerFlags),
		I.Func("(*Logger).Output", (*log.Logger).Output, execmLoggerOutput),
		I.Func("(*Logger).Prefix", (*log.Logger).Prefix, execmLoggerPrefix),
		I.Func("(*Logger).SetFlags", (*log.Logger).SetFlags, execmLoggerSetFlags),
		I.Func("(*Logger).SetOutput", (*log.Logger).SetOutput, execmLoggerSetOutput),
		I.Func("(*Logger).SetPrefix", (*log.Logger).SetPrefix, execmLoggerSetPrefix),
		I.Func("(*Logger).Writer", (*log.Logger).Writer, execmLoggerWriter),
		I.Func("New", log.New, execNew),
		I.Func("Output", log.Output, execOutput),
		I.Func("Prefix", log.Prefix, execPrefix),
		I.Func("SetFlags", log.SetFlags, execSetFlags),
		I.Func("SetOutput", log.SetOutput, execSetOutput),
		I.Func("SetPrefix", log.SetPrefix, execSetPrefix),
		I.Func("Writer", log.Writer, execWriter),
	)
	I.RegisterFuncvs(
		I.Funcv("Fatal", log.Fatal, execFatal),
		I.Funcv("Fatalf", log.Fatalf, execFatalf),
		I.Funcv("Fatalln", log.Fatalln, execFatalln),
		I.Funcv("(*Logger).Fatal", (*log.Logger).Fatal, execmLoggerFatal),
		I.Funcv("(*Logger).Fatalf", (*log.Logger).Fatalf, execmLoggerFatalf),
		I.Funcv("(*Logger).Fatalln", (*log.Logger).Fatalln, execmLoggerFatalln),
		I.Funcv("(*Logger).Panic", (*log.Logger).Panic, execmLoggerPanic),
		I.Funcv("(*Logger).Panicf", (*log.Logger).Panicf, execmLoggerPanicf),
		I.Funcv("(*Logger).Panicln", (*log.Logger).Panicln, execmLoggerPanicln),
		I.Funcv("(*Logger).Print", (*log.Logger).Print, execmLoggerPrint),
		I.Funcv("(*Logger).Printf", (*log.Logger).Printf, execmLoggerPrintf),
		I.Funcv("(*Logger).Println", (*log.Logger).Println, execmLoggerPrintln),
		I.Funcv("Panic", log.Panic, execPanic),
		I.Funcv("Panicf", log.Panicf, execPanicf),
		I.Funcv("Panicln", log.Panicln, execPanicln),
		I.Funcv("Print", log.Print, execPrint),
		I.Funcv("Printf", log.Printf, execPrintf),
		I.Funcv("Println", log.Println, execPrintln),
	)
}
