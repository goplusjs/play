package testing

import (
	"reflect"
	"testing"

	"github.com/qiniu/goplus/gop"
)

// func testing.AllocsPerRun(runs int, f func()) (avg float64)
func execAllocsPerRun(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := testing.AllocsPerRun(args[0].(int), args[1].(func()))
	p.Ret(2, ret)
}

// func (*testing.B).ReportAllocs()
func execmBReportAllocs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*testing.B).ReportAllocs()
}

// func (*testing.B).ReportMetric(n float64, unit string)
func execmBReportMetric(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*testing.B).ReportMetric(args[1].(float64), args[2].(string))
}

// func (*testing.B).ResetTimer()
func execmBResetTimer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*testing.B).ResetTimer()
}

// func (*testing.B).Run(name string, f func(b *testing.B)) bool
func execmBRun(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*testing.B).Run(args[1].(string), args[2].(func(b *testing.B)))
	p.Ret(3, ret)
}

// func (*testing.B).RunParallel(body func(*testing.PB))
func execmBRunParallel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*testing.B).RunParallel(args[1].(func(*testing.PB)))
}

// func (*testing.B).SetBytes(n int64)
func execmBSetBytes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*testing.B).SetBytes(args[1].(int64))
}

// func (*testing.B).SetParallelism(p int)
func execmBSetParallelism(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*testing.B).SetParallelism(args[1].(int))
}

// func (*testing.B).StartTimer()
func execmBStartTimer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*testing.B).StartTimer()
}

// func (*testing.B).StopTimer()
func execmBStopTimer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*testing.B).StopTimer()
}

// func testing.Benchmark(f func(b *testing.B)) testing.BenchmarkResult
func execBenchmark(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := testing.Benchmark(args[0].(func(b *testing.B)))
	p.Ret(1, ret)
}

// func (testing.BenchmarkResult).AllocedBytesPerOp() int64
func execmBenchmarkResultAllocedBytesPerOp(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(testing.BenchmarkResult).AllocedBytesPerOp()
	p.Ret(1, ret)
}

// func (testing.BenchmarkResult).AllocsPerOp() int64
func execmBenchmarkResultAllocsPerOp(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(testing.BenchmarkResult).AllocsPerOp()
	p.Ret(1, ret)
}

// func (testing.BenchmarkResult).MemString() string
func execmBenchmarkResultMemString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(testing.BenchmarkResult).MemString()
	p.Ret(1, ret)
}

// func (testing.BenchmarkResult).NsPerOp() int64
func execmBenchmarkResultNsPerOp(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(testing.BenchmarkResult).NsPerOp()
	p.Ret(1, ret)
}

// func (testing.BenchmarkResult).String() string
func execmBenchmarkResultString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(testing.BenchmarkResult).String()
	p.Ret(1, ret)
}

// func testing.CoverMode() string
func execCoverMode(_ int, p *gop.Context) {
	ret := testing.CoverMode()
	p.Ret(0, ret)
}

// func testing.Coverage() float64
func execCoverage(_ int, p *gop.Context) {
	ret := testing.Coverage()
	p.Ret(0, ret)
}

// func testing.Init()
func execInit(_ int, p *gop.Context) {
	testing.Init()
}

// func (*testing.M).Run() int
func execmMRun(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*testing.M).Run()
	p.Ret(1, ret)
}

// func testing.Main(matchString func(pat string, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample)
func execMain(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	testing.Main(args[0].(func(pat string, str string) (bool, error)), args[1].([]testing.InternalTest), args[2].([]testing.InternalBenchmark), args[3].([]testing.InternalExample))
}

// func testing.MainStart(deps testing.testDeps, tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample) *testing.M
func execMainStart(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := testing.MainStart(args[0].(testing.testDeps), args[1].([]testing.InternalTest), args[2].([]testing.InternalBenchmark), args[3].([]testing.InternalExample))
	p.Ret(4, ret)
}

// func (*testing.PB).Next() bool
func execmPBNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*testing.PB).Next()
	p.Ret(1, ret)
}

// func testing.RegisterCover(c testing.Cover)
func execRegisterCover(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	testing.RegisterCover(args[0].(testing.Cover))
}

// func testing.RunBenchmarks(matchString func(pat string, str string) (bool, error), benchmarks []testing.InternalBenchmark)
func execRunBenchmarks(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	testing.RunBenchmarks(args[0].(func(pat string, str string) (bool, error)), args[1].([]testing.InternalBenchmark))
}

// func testing.RunExamples(matchString func(pat string, str string) (bool, error), examples []testing.InternalExample) (ok bool)
func execRunExamples(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := testing.RunExamples(args[0].(func(pat string, str string) (bool, error)), args[1].([]testing.InternalExample))
	p.Ret(2, ret)
}

// func testing.RunTests(matchString func(pat string, str string) (bool, error), tests []testing.InternalTest) (ok bool)
func execRunTests(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := testing.RunTests(args[0].(func(pat string, str string) (bool, error)), args[1].([]testing.InternalTest))
	p.Ret(2, ret)
}

// func testing.Short() bool
func execShort(_ int, p *gop.Context) {
	ret := testing.Short()
	p.Ret(0, ret)
}

// func (*testing.T).Parallel()
func execmTParallel(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*testing.T).Parallel()
}

// func (*testing.T).Run(name string, f func(t *testing.T)) bool
func execmTRun(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*testing.T).Run(args[1].(string), args[2].(func(t *testing.T)))
	p.Ret(3, ret)
}

// func testing.Verbose() bool
func execVerbose(_ int, p *gop.Context) {
	ret := testing.Verbose()
	p.Ret(0, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("testing")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*testing.B)(nil))),
		I.Rtype(reflect.TypeOf((*testing.BenchmarkResult)(nil))),
		I.Rtype(reflect.TypeOf((*testing.Cover)(nil))),
		I.Rtype(reflect.TypeOf((*testing.CoverBlock)(nil))),
		I.Rtype(reflect.TypeOf((*testing.InternalBenchmark)(nil))),
		I.Rtype(reflect.TypeOf((*testing.InternalExample)(nil))),
		I.Rtype(reflect.TypeOf((*testing.InternalTest)(nil))),
		I.Rtype(reflect.TypeOf((*testing.M)(nil))),
		I.Rtype(reflect.TypeOf((*testing.PB)(nil))),
		I.Rtype(reflect.TypeOf((*testing.T)(nil))),
	)
	I.RegisterFuncs(
		I.Func("AllocsPerRun", testing.AllocsPerRun, execAllocsPerRun),
		I.Func("(*B).ReportAllocs", (*testing.B).ReportAllocs, execmBReportAllocs),
		I.Func("(*B).ReportMetric", (*testing.B).ReportMetric, execmBReportMetric),
		I.Func("(*B).ResetTimer", (*testing.B).ResetTimer, execmBResetTimer),
		I.Func("(*B).Run", (*testing.B).Run, execmBRun),
		I.Func("(*B).RunParallel", (*testing.B).RunParallel, execmBRunParallel),
		I.Func("(*B).SetBytes", (*testing.B).SetBytes, execmBSetBytes),
		I.Func("(*B).SetParallelism", (*testing.B).SetParallelism, execmBSetParallelism),
		I.Func("(*B).StartTimer", (*testing.B).StartTimer, execmBStartTimer),
		I.Func("(*B).StopTimer", (*testing.B).StopTimer, execmBStopTimer),
		I.Func("Benchmark", testing.Benchmark, execBenchmark),
		I.Func("(BenchmarkResult).AllocedBytesPerOp", (testing.BenchmarkResult).AllocedBytesPerOp, execmBenchmarkResultAllocedBytesPerOp),
		I.Func("(BenchmarkResult).AllocsPerOp", (testing.BenchmarkResult).AllocsPerOp, execmBenchmarkResultAllocsPerOp),
		I.Func("(BenchmarkResult).MemString", (testing.BenchmarkResult).MemString, execmBenchmarkResultMemString),
		I.Func("(BenchmarkResult).NsPerOp", (testing.BenchmarkResult).NsPerOp, execmBenchmarkResultNsPerOp),
		I.Func("(BenchmarkResult).String", (testing.BenchmarkResult).String, execmBenchmarkResultString),
		I.Func("CoverMode", testing.CoverMode, execCoverMode),
		I.Func("Coverage", testing.Coverage, execCoverage),
		I.Func("Init", testing.Init, execInit),
		I.Func("(*M).Run", (*testing.M).Run, execmMRun),
		I.Func("Main", testing.Main, execMain),
		I.Func("MainStart", testing.MainStart, execMainStart),
		I.Func("(*PB).Next", (*testing.PB).Next, execmPBNext),
		I.Func("RegisterCover", testing.RegisterCover, execRegisterCover),
		I.Func("RunBenchmarks", testing.RunBenchmarks, execRunBenchmarks),
		I.Func("RunExamples", testing.RunExamples, execRunExamples),
		I.Func("RunTests", testing.RunTests, execRunTests),
		I.Func("Short", testing.Short, execShort),
		I.Func("(*T).Parallel", (*testing.T).Parallel, execmTParallel),
		I.Func("(*T).Run", (*testing.T).Run, execmTRun),
		I.Func("Verbose", testing.Verbose, execVerbose),
	)
}
