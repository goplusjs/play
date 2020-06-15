package sort

import (
	"sort"

	"github.com/qiniu/goplus/gop"
)

// func (sort.Float64Slice).Len() int
func execmFloat64SliceLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(sort.Float64Slice).Len()
	p.Ret(1, ret)
}

// func (sort.Float64Slice).Less(i int, j int) bool
func execmFloat64SliceLess(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(sort.Float64Slice).Less(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (sort.Float64Slice).Search(x float64) int
func execmFloat64SliceSearch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(sort.Float64Slice).Search(args[1].(float64))
	p.Ret(2, ret)
}

// func (sort.Float64Slice).Sort()
func execmFloat64SliceSort(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(sort.Float64Slice).Sort()
}

// func (sort.Float64Slice).Swap(i int, j int)
func execmFloat64SliceSwap(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(sort.Float64Slice).Swap(args[1].(int), args[2].(int))
}

// func sort.Float64s(a []float64)
func execFloat64s(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	sort.Float64s(args[0].([]float64))
}

// func sort.Float64sAreSorted(a []float64) bool
func execFloat64sAreSorted(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sort.Float64sAreSorted(args[0].([]float64))
	p.Ret(1, ret)
}

// func (sort.IntSlice).Len() int
func execmIntSliceLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(sort.IntSlice).Len()
	p.Ret(1, ret)
}

// func (sort.IntSlice).Less(i int, j int) bool
func execmIntSliceLess(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(sort.IntSlice).Less(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (sort.IntSlice).Search(x int) int
func execmIntSliceSearch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(sort.IntSlice).Search(args[1].(int))
	p.Ret(2, ret)
}

// func (sort.IntSlice).Sort()
func execmIntSliceSort(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(sort.IntSlice).Sort()
}

// func (sort.IntSlice).Swap(i int, j int)
func execmIntSliceSwap(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(sort.IntSlice).Swap(args[1].(int), args[2].(int))
}

// func sort.Ints(a []int)
func execInts(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	sort.Ints(args[0].([]int))
}

// func sort.IntsAreSorted(a []int) bool
func execIntsAreSorted(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sort.IntsAreSorted(args[0].([]int))
	p.Ret(1, ret)
}

// func sort.IsSorted(data sort.Interface) bool
func execIsSorted(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sort.IsSorted(args[0].(sort.Interface))
	p.Ret(1, ret)
}

// func sort.Reverse(data sort.Interface) sort.Interface
func execReverse(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sort.Reverse(args[0].(sort.Interface))
	p.Ret(1, ret)
}

// func sort.Search(n int, f func(int) bool) int
func execSearch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := sort.Search(args[0].(int), args[1].(func(int) bool))
	p.Ret(2, ret)
}

// func sort.SearchFloat64s(a []float64, x float64) int
func execSearchFloat64s(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := sort.SearchFloat64s(args[0].([]float64), args[1].(float64))
	p.Ret(2, ret)
}

// func sort.SearchInts(a []int, x int) int
func execSearchInts(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := sort.SearchInts(args[0].([]int), args[1].(int))
	p.Ret(2, ret)
}

// func sort.SearchStrings(a []string, x string) int
func execSearchStrings(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := sort.SearchStrings(args[0].([]string), args[1].(string))
	p.Ret(2, ret)
}

// func sort.Slice(slice interface{}, less func(i int, j int) bool)
func execSlice(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	sort.Slice(args[0].(interface{}), args[1].(func(i int, j int) bool))
}

// func sort.SliceIsSorted(slice interface{}, less func(i int, j int) bool) bool
func execSliceIsSorted(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := sort.SliceIsSorted(args[0].(interface{}), args[1].(func(i int, j int) bool))
	p.Ret(2, ret)
}

// func sort.SliceStable(slice interface{}, less func(i int, j int) bool)
func execSliceStable(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	sort.SliceStable(args[0].(interface{}), args[1].(func(i int, j int) bool))
}

// func sort.Sort(data sort.Interface)
func execSort(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	sort.Sort(args[0].(sort.Interface))
}

// func sort.Stable(data sort.Interface)
func execStable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	sort.Stable(args[0].(sort.Interface))
}

// func (sort.StringSlice).Len() int
func execmStringSliceLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(sort.StringSlice).Len()
	p.Ret(1, ret)
}

// func (sort.StringSlice).Less(i int, j int) bool
func execmStringSliceLess(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(sort.StringSlice).Less(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (sort.StringSlice).Search(x string) int
func execmStringSliceSearch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(sort.StringSlice).Search(args[1].(string))
	p.Ret(2, ret)
}

// func (sort.StringSlice).Sort()
func execmStringSliceSort(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(sort.StringSlice).Sort()
}

// func (sort.StringSlice).Swap(i int, j int)
func execmStringSliceSwap(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(sort.StringSlice).Swap(args[1].(int), args[2].(int))
}

// func sort.Strings(a []string)
func execStrings(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	sort.Strings(args[0].([]string))
}

// func sort.StringsAreSorted(a []string) bool
func execStringsAreSorted(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sort.StringsAreSorted(args[0].([]string))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("sort")

func init() {
	I.RegisterFuncs(
		I.Func("(Float64Slice).Len", (sort.Float64Slice).Len, execmFloat64SliceLen),
		I.Func("(Float64Slice).Less", (sort.Float64Slice).Less, execmFloat64SliceLess),
		I.Func("(Float64Slice).Search", (sort.Float64Slice).Search, execmFloat64SliceSearch),
		I.Func("(Float64Slice).Sort", (sort.Float64Slice).Sort, execmFloat64SliceSort),
		I.Func("(Float64Slice).Swap", (sort.Float64Slice).Swap, execmFloat64SliceSwap),
		I.Func("Float64s", sort.Float64s, execFloat64s),
		I.Func("Float64sAreSorted", sort.Float64sAreSorted, execFloat64sAreSorted),
		I.Func("(IntSlice).Len", (sort.IntSlice).Len, execmIntSliceLen),
		I.Func("(IntSlice).Less", (sort.IntSlice).Less, execmIntSliceLess),
		I.Func("(IntSlice).Search", (sort.IntSlice).Search, execmIntSliceSearch),
		I.Func("(IntSlice).Sort", (sort.IntSlice).Sort, execmIntSliceSort),
		I.Func("(IntSlice).Swap", (sort.IntSlice).Swap, execmIntSliceSwap),
		I.Func("Ints", sort.Ints, execInts),
		I.Func("IntsAreSorted", sort.IntsAreSorted, execIntsAreSorted),
		I.Func("IsSorted", sort.IsSorted, execIsSorted),
		I.Func("Reverse", sort.Reverse, execReverse),
		I.Func("Search", sort.Search, execSearch),
		I.Func("SearchFloat64s", sort.SearchFloat64s, execSearchFloat64s),
		I.Func("SearchInts", sort.SearchInts, execSearchInts),
		I.Func("SearchStrings", sort.SearchStrings, execSearchStrings),
		I.Func("Slice", sort.Slice, execSlice),
		I.Func("SliceIsSorted", sort.SliceIsSorted, execSliceIsSorted),
		I.Func("SliceStable", sort.SliceStable, execSliceStable),
		I.Func("Sort", sort.Sort, execSort),
		I.Func("Stable", sort.Stable, execStable),
		I.Func("(StringSlice).Len", (sort.StringSlice).Len, execmStringSliceLen),
		I.Func("(StringSlice).Less", (sort.StringSlice).Less, execmStringSliceLess),
		I.Func("(StringSlice).Search", (sort.StringSlice).Search, execmStringSliceSearch),
		I.Func("(StringSlice).Sort", (sort.StringSlice).Sort, execmStringSliceSort),
		I.Func("(StringSlice).Swap", (sort.StringSlice).Swap, execmStringSliceSwap),
		I.Func("Strings", sort.Strings, execStrings),
		I.Func("StringsAreSorted", sort.StringsAreSorted, execStringsAreSorted),
	)
}
