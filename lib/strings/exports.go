package strings

import (
	"io"
	"reflect"
	"strings"
	"unicode"

	"github.com/qiniu/goplus/gop"
)

// func (*strings.Builder).Cap() int
func execmBuilderCap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strings.Builder).Cap()
	p.Ret(1, ret)
}

// func (*strings.Builder).Grow(n int)
func execmBuilderGrow(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*strings.Builder).Grow(args[1].(int))
}

// func (*strings.Builder).Len() int
func execmBuilderLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strings.Builder).Len()
	p.Ret(1, ret)
}

// func (*strings.Builder).Reset()
func execmBuilderReset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*strings.Builder).Reset()
}

// func (*strings.Builder).String() string
func execmBuilderString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strings.Builder).String()
	p.Ret(1, ret)
}

// func (*strings.Builder).Write(p []byte) (int, error)
func execmBuilderWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*strings.Builder).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*strings.Builder).WriteByte(c byte) error
func execmBuilderWriteByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*strings.Builder).WriteByte(args[1].(byte))
	p.Ret(2, ret)
}

// func (*strings.Builder).WriteRune(r rune) (int, error)
func execmBuilderWriteRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*strings.Builder).WriteRune(args[1].(rune))
	p.Ret(2, ret, ret1)
}

// func (*strings.Builder).WriteString(s string) (int, error)
func execmBuilderWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*strings.Builder).WriteString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func strings.Compare(a string, b string) int
func execCompare(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Compare(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.Contains(s string, substr string) bool
func execContains(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Contains(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.ContainsAny(s string, chars string) bool
func execContainsAny(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.ContainsAny(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.ContainsRune(s string, r rune) bool
func execContainsRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.ContainsRune(args[0].(string), args[1].(rune))
	p.Ret(2, ret)
}

// func strings.Count(s string, substr string) int
func execCount(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Count(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.EqualFold(s string, t string) bool
func execEqualFold(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.EqualFold(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.Fields(s string) []string
func execFields(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strings.Fields(args[0].(string))
	p.Ret(1, ret)
}

// func strings.FieldsFunc(s string, f func(rune) bool) []string
func execFieldsFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.FieldsFunc(args[0].(string), args[1].(func(rune) bool))
	p.Ret(2, ret)
}

// func strings.HasPrefix(s string, prefix string) bool
func execHasPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.HasPrefix(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.HasSuffix(s string, suffix string) bool
func execHasSuffix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.HasSuffix(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.Index(s string, substr string) int
func execIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Index(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.IndexAny(s string, chars string) int
func execIndexAny(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.IndexAny(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.IndexByte(s string, c byte) int
func execIndexByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.IndexByte(args[0].(string), args[1].(byte))
	p.Ret(2, ret)
}

// func strings.IndexFunc(s string, f func(rune) bool) int
func execIndexFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.IndexFunc(args[0].(string), args[1].(func(rune) bool))
	p.Ret(2, ret)
}

// func strings.IndexRune(s string, r rune) int
func execIndexRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.IndexRune(args[0].(string), args[1].(rune))
	p.Ret(2, ret)
}

// func strings.Join(elems []string, sep string) string
func execJoin(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Join(args[0].([]string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.LastIndex(s string, substr string) int
func execLastIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.LastIndex(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.LastIndexAny(s string, chars string) int
func execLastIndexAny(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.LastIndexAny(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.LastIndexByte(s string, c byte) int
func execLastIndexByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.LastIndexByte(args[0].(string), args[1].(byte))
	p.Ret(2, ret)
}

// func strings.LastIndexFunc(s string, f func(rune) bool) int
func execLastIndexFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.LastIndexFunc(args[0].(string), args[1].(func(rune) bool))
	p.Ret(2, ret)
}

// func strings.Map(mapping func(rune) rune, s string) string
func execMap(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Map(args[0].(func(rune) rune), args[1].(string))
	p.Ret(2, ret)
}

// func strings.NewReader(s string) *strings.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strings.NewReader(args[0].(string))
	p.Ret(1, ret)
}

// func strings.NewReplacer(oldnew ..string) *strings.Replacer
func execNewReplacer(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret := strings.NewReplacer(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func (*strings.Reader).Len() int
func execmReaderLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strings.Reader).Len()
	p.Ret(1, ret)
}

// func (*strings.Reader).Read(b []byte) (n int, err error)
func execmReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*strings.Reader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*strings.Reader).ReadAt(b []byte, off int64) (n int, err error)
func execmReaderReadAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*strings.Reader).ReadAt(args[1].([]byte), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func (*strings.Reader).ReadByte() (byte, error)
func execmReaderReadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*strings.Reader).ReadByte()
	p.Ret(1, ret, ret1)
}

// func (*strings.Reader).ReadRune() (ch rune, size int, err error)
func execmReaderReadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*strings.Reader).ReadRune()
	p.Ret(1, ret, ret1, ret2)
}

// func (*strings.Reader).Reset(s string)
func execmReaderReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*strings.Reader).Reset(args[1].(string))
}

// func (*strings.Reader).Seek(offset int64, whence int) (int64, error)
func execmReaderSeek(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*strings.Reader).Seek(args[1].(int64), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func (*strings.Reader).Size() int64
func execmReaderSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strings.Reader).Size()
	p.Ret(1, ret)
}

// func (*strings.Reader).UnreadByte() error
func execmReaderUnreadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strings.Reader).UnreadByte()
	p.Ret(1, ret)
}

// func (*strings.Reader).UnreadRune() error
func execmReaderUnreadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strings.Reader).UnreadRune()
	p.Ret(1, ret)
}

// func (*strings.Reader).WriteTo(w io.Writer) (n int64, err error)
func execmReaderWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*strings.Reader).WriteTo(args[1].(io.Writer))
	p.Ret(2, ret, ret1)
}

// func strings.Repeat(s string, count int) string
func execRepeat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Repeat(args[0].(string), args[1].(int))
	p.Ret(2, ret)
}

// func strings.Replace(s string, old string, new string, n int) string
func execReplace(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := strings.Replace(args[0].(string), args[1].(string), args[2].(string), args[3].(int))
	p.Ret(4, ret)
}

// func strings.ReplaceAll(s string, old string, new string) string
func execReplaceAll(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := strings.ReplaceAll(args[0].(string), args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// func (*strings.Replacer).Replace(s string) string
func execmReplacerReplace(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*strings.Replacer).Replace(args[1].(string))
	p.Ret(2, ret)
}

// func (*strings.Replacer).WriteString(w io.Writer, s string) (n int, err error)
func execmReplacerWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*strings.Replacer).WriteString(args[1].(io.Writer), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func strings.Split(s string, sep string) []string
func execSplit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Split(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.SplitAfter(s string, sep string) []string
func execSplitAfter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.SplitAfter(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.SplitAfterN(s string, sep string, n int) []string
func execSplitAfterN(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := strings.SplitAfterN(args[0].(string), args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func strings.SplitN(s string, sep string, n int) []string
func execSplitN(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := strings.SplitN(args[0].(string), args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func strings.Title(s string) string
func execTitle(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strings.Title(args[0].(string))
	p.Ret(1, ret)
}

// func strings.ToLower(s string) string
func execToLower(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strings.ToLower(args[0].(string))
	p.Ret(1, ret)
}

// func strings.ToLowerSpecial(c unicode.SpecialCase, s string) string
func execToLowerSpecial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.ToLowerSpecial(args[0].(unicode.SpecialCase), args[1].(string))
	p.Ret(2, ret)
}

// func strings.ToTitle(s string) string
func execToTitle(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strings.ToTitle(args[0].(string))
	p.Ret(1, ret)
}

// func strings.ToTitleSpecial(c unicode.SpecialCase, s string) string
func execToTitleSpecial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.ToTitleSpecial(args[0].(unicode.SpecialCase), args[1].(string))
	p.Ret(2, ret)
}

// func strings.ToUpper(s string) string
func execToUpper(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strings.ToUpper(args[0].(string))
	p.Ret(1, ret)
}

// func strings.ToUpperSpecial(c unicode.SpecialCase, s string) string
func execToUpperSpecial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.ToUpperSpecial(args[0].(unicode.SpecialCase), args[1].(string))
	p.Ret(2, ret)
}

// func strings.ToValidUTF8(s string, replacement string) string
func execToValidUTF8(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.ToValidUTF8(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.Trim(s string, cutset string) string
func execTrim(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.Trim(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.TrimFunc(s string, f func(rune) bool) string
func execTrimFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.TrimFunc(args[0].(string), args[1].(func(rune) bool))
	p.Ret(2, ret)
}

// func strings.TrimLeft(s string, cutset string) string
func execTrimLeft(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.TrimLeft(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.TrimLeftFunc(s string, f func(rune) bool) string
func execTrimLeftFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.TrimLeftFunc(args[0].(string), args[1].(func(rune) bool))
	p.Ret(2, ret)
}

// func strings.TrimPrefix(s string, prefix string) string
func execTrimPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.TrimPrefix(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.TrimRight(s string, cutset string) string
func execTrimRight(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.TrimRight(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func strings.TrimRightFunc(s string, f func(rune) bool) string
func execTrimRightFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.TrimRightFunc(args[0].(string), args[1].(func(rune) bool))
	p.Ret(2, ret)
}

// func strings.TrimSpace(s string) string
func execTrimSpace(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strings.TrimSpace(args[0].(string))
	p.Ret(1, ret)
}

// func strings.TrimSuffix(s string, suffix string) string
func execTrimSuffix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strings.TrimSuffix(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("strings")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*strings.Builder)(nil))),
		I.Rtype(reflect.TypeOf((*strings.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*strings.Replacer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Builder).Cap", (*strings.Builder).Cap, execmBuilderCap),
		I.Func("(*Builder).Grow", (*strings.Builder).Grow, execmBuilderGrow),
		I.Func("(*Builder).Len", (*strings.Builder).Len, execmBuilderLen),
		I.Func("(*Builder).Reset", (*strings.Builder).Reset, execmBuilderReset),
		I.Func("(*Builder).String", (*strings.Builder).String, execmBuilderString),
		I.Func("(*Builder).Write", (*strings.Builder).Write, execmBuilderWrite),
		I.Func("(*Builder).WriteByte", (*strings.Builder).WriteByte, execmBuilderWriteByte),
		I.Func("(*Builder).WriteRune", (*strings.Builder).WriteRune, execmBuilderWriteRune),
		I.Func("(*Builder).WriteString", (*strings.Builder).WriteString, execmBuilderWriteString),
		I.Func("Compare", strings.Compare, execCompare),
		I.Func("Contains", strings.Contains, execContains),
		I.Func("ContainsAny", strings.ContainsAny, execContainsAny),
		I.Func("ContainsRune", strings.ContainsRune, execContainsRune),
		I.Func("Count", strings.Count, execCount),
		I.Func("EqualFold", strings.EqualFold, execEqualFold),
		I.Func("Fields", strings.Fields, execFields),
		I.Func("FieldsFunc", strings.FieldsFunc, execFieldsFunc),
		I.Func("HasPrefix", strings.HasPrefix, execHasPrefix),
		I.Func("HasSuffix", strings.HasSuffix, execHasSuffix),
		I.Func("Index", strings.Index, execIndex),
		I.Func("IndexAny", strings.IndexAny, execIndexAny),
		I.Func("IndexByte", strings.IndexByte, execIndexByte),
		I.Func("IndexFunc", strings.IndexFunc, execIndexFunc),
		I.Func("IndexRune", strings.IndexRune, execIndexRune),
		I.Func("Join", strings.Join, execJoin),
		I.Func("LastIndex", strings.LastIndex, execLastIndex),
		I.Func("LastIndexAny", strings.LastIndexAny, execLastIndexAny),
		I.Func("LastIndexByte", strings.LastIndexByte, execLastIndexByte),
		I.Func("LastIndexFunc", strings.LastIndexFunc, execLastIndexFunc),
		I.Func("Map", strings.Map, execMap),
		I.Func("NewReader", strings.NewReader, execNewReader),
		I.Func("(*Reader).Len", (*strings.Reader).Len, execmReaderLen),
		I.Func("(*Reader).Read", (*strings.Reader).Read, execmReaderRead),
		I.Func("(*Reader).ReadAt", (*strings.Reader).ReadAt, execmReaderReadAt),
		I.Func("(*Reader).ReadByte", (*strings.Reader).ReadByte, execmReaderReadByte),
		I.Func("(*Reader).ReadRune", (*strings.Reader).ReadRune, execmReaderReadRune),
		I.Func("(*Reader).Reset", (*strings.Reader).Reset, execmReaderReset),
		I.Func("(*Reader).Seek", (*strings.Reader).Seek, execmReaderSeek),
		I.Func("(*Reader).Size", (*strings.Reader).Size, execmReaderSize),
		I.Func("(*Reader).UnreadByte", (*strings.Reader).UnreadByte, execmReaderUnreadByte),
		I.Func("(*Reader).UnreadRune", (*strings.Reader).UnreadRune, execmReaderUnreadRune),
		I.Func("(*Reader).WriteTo", (*strings.Reader).WriteTo, execmReaderWriteTo),
		I.Func("Repeat", strings.Repeat, execRepeat),
		I.Func("Replace", strings.Replace, execReplace),
		I.Func("ReplaceAll", strings.ReplaceAll, execReplaceAll),
		I.Func("(*Replacer).Replace", (*strings.Replacer).Replace, execmReplacerReplace),
		I.Func("(*Replacer).WriteString", (*strings.Replacer).WriteString, execmReplacerWriteString),
		I.Func("Split", strings.Split, execSplit),
		I.Func("SplitAfter", strings.SplitAfter, execSplitAfter),
		I.Func("SplitAfterN", strings.SplitAfterN, execSplitAfterN),
		I.Func("SplitN", strings.SplitN, execSplitN),
		I.Func("Title", strings.Title, execTitle),
		I.Func("ToLower", strings.ToLower, execToLower),
		I.Func("ToLowerSpecial", strings.ToLowerSpecial, execToLowerSpecial),
		I.Func("ToTitle", strings.ToTitle, execToTitle),
		I.Func("ToTitleSpecial", strings.ToTitleSpecial, execToTitleSpecial),
		I.Func("ToUpper", strings.ToUpper, execToUpper),
		I.Func("ToUpperSpecial", strings.ToUpperSpecial, execToUpperSpecial),
		I.Func("ToValidUTF8", strings.ToValidUTF8, execToValidUTF8),
		I.Func("Trim", strings.Trim, execTrim),
		I.Func("TrimFunc", strings.TrimFunc, execTrimFunc),
		I.Func("TrimLeft", strings.TrimLeft, execTrimLeft),
		I.Func("TrimLeftFunc", strings.TrimLeftFunc, execTrimLeftFunc),
		I.Func("TrimPrefix", strings.TrimPrefix, execTrimPrefix),
		I.Func("TrimRight", strings.TrimRight, execTrimRight),
		I.Func("TrimRightFunc", strings.TrimRightFunc, execTrimRightFunc),
		I.Func("TrimSpace", strings.TrimSpace, execTrimSpace),
		I.Func("TrimSuffix", strings.TrimSuffix, execTrimSuffix),
	)
	I.RegisterFuncvs(
		I.Funcv("NewReplacer", strings.NewReplacer, execNewReplacer),
	)
}
