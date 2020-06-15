package flag

import (
	"flag"
	"io"
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func flag.Arg(i int) string
func execArg(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := flag.Arg(args[0].(int))
	p.Ret(1, ret)
}

// func flag.Args() []string
func execArgs(_ int, p *gop.Context) {
	ret := flag.Args()
	p.Ret(0, ret)
}

// func flag.Bool(name string, value bool, usage string) *bool
func execBool(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.Bool(args[0].(string), args[1].(bool), args[2].(string))
	p.Ret(3, ret)
}

// func flag.BoolVar(p *bool, name string, value bool, usage string)
func execBoolVar(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.BoolVar(args[0].(*bool), args[1].(string), args[2].(bool), args[3].(string))
}

// func flag.Duration(name string, value time.Duration, usage string) *time.Duration
func execDuration(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.Duration(args[0].(string), time.Duration(args[1].(int64)), args[2].(string))
	p.Ret(3, ret)
}

// func flag.DurationVar(p *time.Duration, name string, value time.Duration, usage string)
func execDurationVar(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.DurationVar(args[0].(*time.Duration), args[1].(string), time.Duration(args[2].(int64)), args[3].(string))
}

// func (*flag.FlagSet).Arg(i int) string
func execmFlagSetArg(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*flag.FlagSet).Arg(args[1].(int))
	p.Ret(2, ret)
}

// func (*flag.FlagSet).Args() []string
func execmFlagSetArgs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flag.FlagSet).Args()
	p.Ret(1, ret)
}

// func (*flag.FlagSet).Bool(name string, value bool, usage string) *bool
func execmFlagSetBool(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).Bool(args[1].(string), args[2].(bool), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).BoolVar(p *bool, name string, value bool, usage string)
func execmFlagSetBoolVar(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).BoolVar(args[1].(*bool), args[2].(string), args[3].(bool), args[4].(string))
}

// func (*flag.FlagSet).Duration(name string, value time.Duration, usage string) *time.Duration
func execmFlagSetDuration(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).Duration(args[1].(string), time.Duration(args[2].(int64)), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).DurationVar(p *time.Duration, name string, value time.Duration, usage string)
func execmFlagSetDurationVar(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).DurationVar(args[1].(*time.Duration), args[2].(string), time.Duration(args[3].(int64)), args[4].(string))
}

// func (*flag.FlagSet).ErrorHandling() flag.ErrorHandling
func execmFlagSetErrorHandling(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flag.FlagSet).ErrorHandling()
	p.Ret(1, ret)
}

// func (*flag.FlagSet).Float64(name string, value float64, usage string) *float64
func execmFlagSetFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).Float64(args[1].(string), args[2].(float64), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).Float64Var(p *float64, name string, value float64, usage string)
func execmFlagSetFloat64Var(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).Float64Var(args[1].(*float64), args[2].(string), args[3].(float64), args[4].(string))
}

// func (*flag.FlagSet).Init(name string, errorHandling flag.ErrorHandling)
func execmFlagSetInit(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*flag.FlagSet).Init(args[1].(string), flag.ErrorHandling(args[2].(int)))
}

// func (*flag.FlagSet).Int(name string, value int, usage string) *int
func execmFlagSetInt(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).Int(args[1].(string), args[2].(int), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).Int64(name string, value int64, usage string) *int64
func execmFlagSetInt64(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).Int64(args[1].(string), args[2].(int64), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).Int64Var(p *int64, name string, value int64, usage string)
func execmFlagSetInt64Var(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).Int64Var(args[1].(*int64), args[2].(string), args[3].(int64), args[4].(string))
}

// func (*flag.FlagSet).IntVar(p *int, name string, value int, usage string)
func execmFlagSetIntVar(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).IntVar(args[1].(*int), args[2].(string), args[3].(int), args[4].(string))
}

// func (*flag.FlagSet).Lookup(name string) *flag.Flag
func execmFlagSetLookup(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*flag.FlagSet).Lookup(args[1].(string))
	p.Ret(2, ret)
}

// func (*flag.FlagSet).NArg() int
func execmFlagSetNArg(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flag.FlagSet).NArg()
	p.Ret(1, ret)
}

// func (*flag.FlagSet).NFlag() int
func execmFlagSetNFlag(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flag.FlagSet).NFlag()
	p.Ret(1, ret)
}

// func (*flag.FlagSet).Name() string
func execmFlagSetName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flag.FlagSet).Name()
	p.Ret(1, ret)
}

// func (*flag.FlagSet).Output() io.Writer
func execmFlagSetOutput(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flag.FlagSet).Output()
	p.Ret(1, ret)
}

// func (*flag.FlagSet).Parse(arguments []string) error
func execmFlagSetParse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*flag.FlagSet).Parse(args[1].([]string))
	p.Ret(2, ret)
}

// func (*flag.FlagSet).Parsed() bool
func execmFlagSetParsed(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*flag.FlagSet).Parsed()
	p.Ret(1, ret)
}

// func (*flag.FlagSet).PrintDefaults()
func execmFlagSetPrintDefaults(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*flag.FlagSet).PrintDefaults()
}

// func (*flag.FlagSet).Set(name string, value string) error
func execmFlagSetSet(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*flag.FlagSet).Set(args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// func (*flag.FlagSet).SetOutput(output io.Writer)
func execmFlagSetSetOutput(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*flag.FlagSet).SetOutput(args[1].(io.Writer))
}

// func (*flag.FlagSet).String(name string, value string, usage string) *string
func execmFlagSetString(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).String(args[1].(string), args[2].(string), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).StringVar(p *string, name string, value string, usage string)
func execmFlagSetStringVar(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).StringVar(args[1].(*string), args[2].(string), args[3].(string), args[4].(string))
}

// func (*flag.FlagSet).Uint(name string, value uint, usage string) *uint
func execmFlagSetUint(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).Uint(args[1].(string), args[2].(uint), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).Uint64(name string, value uint64, usage string) *uint64
func execmFlagSetUint64(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*flag.FlagSet).Uint64(args[1].(string), args[2].(uint64), args[3].(string))
	p.Ret(4, ret)
}

// func (*flag.FlagSet).Uint64Var(p *uint64, name string, value uint64, usage string)
func execmFlagSetUint64Var(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).Uint64Var(args[1].(*uint64), args[2].(string), args[3].(uint64), args[4].(string))
}

// func (*flag.FlagSet).UintVar(p *uint, name string, value uint, usage string)
func execmFlagSetUintVar(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*flag.FlagSet).UintVar(args[1].(*uint), args[2].(string), args[3].(uint), args[4].(string))
}

// func (*flag.FlagSet).Var(value flag.Value, name string, usage string)
func execmFlagSetVar(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*flag.FlagSet).Var(args[1].(flag.Value), args[2].(string), args[3].(string))
}

// func (*flag.FlagSet).Visit(fn func(*flag.Flag))
func execmFlagSetVisit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*flag.FlagSet).Visit(args[1].(func(*flag.Flag)))
}

// func (*flag.FlagSet).VisitAll(fn func(*flag.Flag))
func execmFlagSetVisitAll(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*flag.FlagSet).VisitAll(args[1].(func(*flag.Flag)))
}

// func flag.Float64(name string, value float64, usage string) *float64
func execFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.Float64(args[0].(string), args[1].(float64), args[2].(string))
	p.Ret(3, ret)
}

// func flag.Float64Var(p *float64, name string, value float64, usage string)
func execFloat64Var(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.Float64Var(args[0].(*float64), args[1].(string), args[2].(float64), args[3].(string))
}

// func flag.Int(name string, value int, usage string) *int
func execInt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.Int(args[0].(string), args[1].(int), args[2].(string))
	p.Ret(3, ret)
}

// func flag.Int64(name string, value int64, usage string) *int64
func execInt64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.Int64(args[0].(string), args[1].(int64), args[2].(string))
	p.Ret(3, ret)
}

// func flag.Int64Var(p *int64, name string, value int64, usage string)
func execInt64Var(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.Int64Var(args[0].(*int64), args[1].(string), args[2].(int64), args[3].(string))
}

// func flag.IntVar(p *int, name string, value int, usage string)
func execIntVar(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.IntVar(args[0].(*int), args[1].(string), args[2].(int), args[3].(string))
}

// func flag.Lookup(name string) *flag.Flag
func execLookup(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := flag.Lookup(args[0].(string))
	p.Ret(1, ret)
}

// func flag.NArg() int
func execNArg(_ int, p *gop.Context) {
	ret := flag.NArg()
	p.Ret(0, ret)
}

// func flag.NFlag() int
func execNFlag(_ int, p *gop.Context) {
	ret := flag.NFlag()
	p.Ret(0, ret)
}

// func flag.NewFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet
func execNewFlagSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := flag.NewFlagSet(args[0].(string), flag.ErrorHandling(args[1].(int)))
	p.Ret(2, ret)
}

// func flag.Parse()
func execParse(_ int, p *gop.Context) {
	flag.Parse()
}

// func flag.Parsed() bool
func execParsed(_ int, p *gop.Context) {
	ret := flag.Parsed()
	p.Ret(0, ret)
}

// func flag.PrintDefaults()
func execPrintDefaults(_ int, p *gop.Context) {
	flag.PrintDefaults()
}

// func flag.Set(name string, value string) error
func execSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := flag.Set(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func flag.String(name string, value string, usage string) *string
func execString(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.String(args[0].(string), args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// func flag.StringVar(p *string, name string, value string, usage string)
func execStringVar(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.StringVar(args[0].(*string), args[1].(string), args[2].(string), args[3].(string))
}

// func flag.Uint(name string, value uint, usage string) *uint
func execUint(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.Uint(args[0].(string), args[1].(uint), args[2].(string))
	p.Ret(3, ret)
}

// func flag.Uint64(name string, value uint64, usage string) *uint64
func execUint64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := flag.Uint64(args[0].(string), args[1].(uint64), args[2].(string))
	p.Ret(3, ret)
}

// func flag.Uint64Var(p *uint64, name string, value uint64, usage string)
func execUint64Var(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.Uint64Var(args[0].(*uint64), args[1].(string), args[2].(uint64), args[3].(string))
}

// func flag.UintVar(p *uint, name string, value uint, usage string)
func execUintVar(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	flag.UintVar(args[0].(*uint), args[1].(string), args[2].(uint), args[3].(string))
}

// func flag.UnquoteUsage(flag *flag.Flag) (name string, usage string)
func execUnquoteUsage(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := flag.UnquoteUsage(args[0].(*flag.Flag))
	p.Ret(1, ret, ret1)
}

// func flag.Var(value flag.Value, name string, usage string)
func execVar(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	flag.Var(args[0].(flag.Value), args[1].(string), args[2].(string))
}

// func flag.Visit(fn func(*flag.Flag))
func execVisit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	flag.Visit(args[0].(func(*flag.Flag)))
}

// func flag.VisitAll(fn func(*flag.Flag))
func execVisitAll(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	flag.VisitAll(args[0].(func(*flag.Flag)))
}

// I is a Go package instance.
var I = gop.NewGoPackage("flag")

func init() {
	I.RegisterConsts(
		I.Const("ContinueOnError", reflect.Int, flag.ContinueOnError),
		I.Const("ExitOnError", reflect.Int, flag.ExitOnError),
		I.Const("PanicOnError", reflect.Int, flag.PanicOnError),
	)
	I.RegisterVars(
		I.Var("CommandLine", &flag.CommandLine),
		I.Var("ErrHelp", &flag.ErrHelp),
		I.Var("Usage", &flag.Usage),
	)
	I.RegisterTypes(
		I.Type("ErrorHandling", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*flag.Flag)(nil))),
		I.Rtype(reflect.TypeOf((*flag.FlagSet)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Arg", flag.Arg, execArg),
		I.Func("Args", flag.Args, execArgs),
		I.Func("Bool", flag.Bool, execBool),
		I.Func("BoolVar", flag.BoolVar, execBoolVar),
		I.Func("Duration", flag.Duration, execDuration),
		I.Func("DurationVar", flag.DurationVar, execDurationVar),
		I.Func("(*FlagSet).Arg", (*flag.FlagSet).Arg, execmFlagSetArg),
		I.Func("(*FlagSet).Args", (*flag.FlagSet).Args, execmFlagSetArgs),
		I.Func("(*FlagSet).Bool", (*flag.FlagSet).Bool, execmFlagSetBool),
		I.Func("(*FlagSet).BoolVar", (*flag.FlagSet).BoolVar, execmFlagSetBoolVar),
		I.Func("(*FlagSet).Duration", (*flag.FlagSet).Duration, execmFlagSetDuration),
		I.Func("(*FlagSet).DurationVar", (*flag.FlagSet).DurationVar, execmFlagSetDurationVar),
		I.Func("(*FlagSet).ErrorHandling", (*flag.FlagSet).ErrorHandling, execmFlagSetErrorHandling),
		I.Func("(*FlagSet).Float64", (*flag.FlagSet).Float64, execmFlagSetFloat64),
		I.Func("(*FlagSet).Float64Var", (*flag.FlagSet).Float64Var, execmFlagSetFloat64Var),
		I.Func("(*FlagSet).Init", (*flag.FlagSet).Init, execmFlagSetInit),
		I.Func("(*FlagSet).Int", (*flag.FlagSet).Int, execmFlagSetInt),
		I.Func("(*FlagSet).Int64", (*flag.FlagSet).Int64, execmFlagSetInt64),
		I.Func("(*FlagSet).Int64Var", (*flag.FlagSet).Int64Var, execmFlagSetInt64Var),
		I.Func("(*FlagSet).IntVar", (*flag.FlagSet).IntVar, execmFlagSetIntVar),
		I.Func("(*FlagSet).Lookup", (*flag.FlagSet).Lookup, execmFlagSetLookup),
		I.Func("(*FlagSet).NArg", (*flag.FlagSet).NArg, execmFlagSetNArg),
		I.Func("(*FlagSet).NFlag", (*flag.FlagSet).NFlag, execmFlagSetNFlag),
		I.Func("(*FlagSet).Name", (*flag.FlagSet).Name, execmFlagSetName),
		I.Func("(*FlagSet).Output", (*flag.FlagSet).Output, execmFlagSetOutput),
		I.Func("(*FlagSet).Parse", (*flag.FlagSet).Parse, execmFlagSetParse),
		I.Func("(*FlagSet).Parsed", (*flag.FlagSet).Parsed, execmFlagSetParsed),
		I.Func("(*FlagSet).PrintDefaults", (*flag.FlagSet).PrintDefaults, execmFlagSetPrintDefaults),
		I.Func("(*FlagSet).Set", (*flag.FlagSet).Set, execmFlagSetSet),
		I.Func("(*FlagSet).SetOutput", (*flag.FlagSet).SetOutput, execmFlagSetSetOutput),
		I.Func("(*FlagSet).String", (*flag.FlagSet).String, execmFlagSetString),
		I.Func("(*FlagSet).StringVar", (*flag.FlagSet).StringVar, execmFlagSetStringVar),
		I.Func("(*FlagSet).Uint", (*flag.FlagSet).Uint, execmFlagSetUint),
		I.Func("(*FlagSet).Uint64", (*flag.FlagSet).Uint64, execmFlagSetUint64),
		I.Func("(*FlagSet).Uint64Var", (*flag.FlagSet).Uint64Var, execmFlagSetUint64Var),
		I.Func("(*FlagSet).UintVar", (*flag.FlagSet).UintVar, execmFlagSetUintVar),
		I.Func("(*FlagSet).Var", (*flag.FlagSet).Var, execmFlagSetVar),
		I.Func("(*FlagSet).Visit", (*flag.FlagSet).Visit, execmFlagSetVisit),
		I.Func("(*FlagSet).VisitAll", (*flag.FlagSet).VisitAll, execmFlagSetVisitAll),
		I.Func("Float64", flag.Float64, execFloat64),
		I.Func("Float64Var", flag.Float64Var, execFloat64Var),
		I.Func("Int", flag.Int, execInt),
		I.Func("Int64", flag.Int64, execInt64),
		I.Func("Int64Var", flag.Int64Var, execInt64Var),
		I.Func("IntVar", flag.IntVar, execIntVar),
		I.Func("Lookup", flag.Lookup, execLookup),
		I.Func("NArg", flag.NArg, execNArg),
		I.Func("NFlag", flag.NFlag, execNFlag),
		I.Func("NewFlagSet", flag.NewFlagSet, execNewFlagSet),
		I.Func("Parse", flag.Parse, execParse),
		I.Func("Parsed", flag.Parsed, execParsed),
		I.Func("PrintDefaults", flag.PrintDefaults, execPrintDefaults),
		I.Func("Set", flag.Set, execSet),
		I.Func("String", flag.String, execString),
		I.Func("StringVar", flag.StringVar, execStringVar),
		I.Func("Uint", flag.Uint, execUint),
		I.Func("Uint64", flag.Uint64, execUint64),
		I.Func("Uint64Var", flag.Uint64Var, execUint64Var),
		I.Func("UintVar", flag.UintVar, execUintVar),
		I.Func("UnquoteUsage", flag.UnquoteUsage, execUnquoteUsage),
		I.Func("Var", flag.Var, execVar),
		I.Func("Visit", flag.Visit, execVisit),
		I.Func("VisitAll", flag.VisitAll, execVisitAll),
	)
}
