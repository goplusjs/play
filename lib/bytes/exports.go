package bytes

import (
	"bytes"
	"io"
	"reflect"
	"unicode"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*bytes.Buffer).Bytes() []byte
func execmBufferBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Buffer).Bytes()
	p.Ret(1, ret)
}

// func (*bytes.Buffer).Cap() int
func execmBufferCap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Buffer).Cap()
	p.Ret(1, ret)
}

// func (*bytes.Buffer).Grow(n int)
func execmBufferGrow(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*bytes.Buffer).Grow(args[1].(int))
}

// func (*bytes.Buffer).Len() int
func execmBufferLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Buffer).Len()
	p.Ret(1, ret)
}

// func (*bytes.Buffer).Next(n int) []byte
func execmBufferNext(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*bytes.Buffer).Next(args[1].(int))
	p.Ret(2, ret)
}

// func (*bytes.Buffer).Read(p []byte) (n int, err error)
func execmBufferRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Buffer).ReadByte() (byte, error)
func execmBufferReadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*bytes.Buffer).ReadByte()
	p.Ret(1, ret, ret1)
}

// func (*bytes.Buffer).ReadBytes(delim byte) (line []byte, err error)
func execmBufferReadBytes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).ReadBytes(args[1].(byte))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Buffer).ReadFrom(r io.Reader) (n int64, err error)
func execmBufferReadFrom(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).ReadFrom(args[1].(io.Reader))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Buffer).ReadRune() (r rune, size int, err error)
func execmBufferReadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*bytes.Buffer).ReadRune()
	p.Ret(1, ret, ret1, ret2)
}

// func (*bytes.Buffer).ReadString(delim byte) (line string, err error)
func execmBufferReadString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).ReadString(args[1].(byte))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Buffer).Reset()
func execmBufferReset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*bytes.Buffer).Reset()
}

// func (*bytes.Buffer).String() string
func execmBufferString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Buffer).String()
	p.Ret(1, ret)
}

// func (*bytes.Buffer).Truncate(n int)
func execmBufferTruncate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*bytes.Buffer).Truncate(args[1].(int))
}

// func (*bytes.Buffer).UnreadByte() error
func execmBufferUnreadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Buffer).UnreadByte()
	p.Ret(1, ret)
}

// func (*bytes.Buffer).UnreadRune() error
func execmBufferUnreadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Buffer).UnreadRune()
	p.Ret(1, ret)
}

// func (*bytes.Buffer).Write(p []byte) (n int, err error)
func execmBufferWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Buffer).WriteByte(c byte) error
func execmBufferWriteByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*bytes.Buffer).WriteByte(args[1].(byte))
	p.Ret(2, ret)
}

// func (*bytes.Buffer).WriteRune(r rune) (n int, err error)
func execmBufferWriteRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).WriteRune(args[1].(rune))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Buffer).WriteString(s string) (n int, err error)
func execmBufferWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).WriteString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Buffer).WriteTo(w io.Writer) (n int64, err error)
func execmBufferWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Buffer).WriteTo(args[1].(io.Writer))
	p.Ret(2, ret, ret1)
}

// func bytes.Compare(a []byte, b []byte) int
func execCompare(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Compare(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.Contains(b []byte, subslice []byte) bool
func execContains(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Contains(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.ContainsAny(b []byte, chars string) bool
func execContainsAny(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.ContainsAny(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func bytes.ContainsRune(b []byte, r rune) bool
func execContainsRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.ContainsRune(args[0].([]byte), args[1].(rune))
	p.Ret(2, ret)
}

// func bytes.Count(s []byte, sep []byte) int
func execCount(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Count(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.Equal(a []byte, b []byte) bool
func execEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Equal(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.EqualFold(s []byte, t []byte) bool
func execEqualFold(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.EqualFold(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.Fields(s []byte) [][]byte
func execFields(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.Fields(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.FieldsFunc(s []byte, f func(rune) bool) [][]byte
func execFieldsFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.FieldsFunc(args[0].([]byte), args[1].(func(rune) bool))
	p.Ret(2, ret)
}

// func bytes.HasPrefix(s []byte, prefix []byte) bool
func execHasPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.HasPrefix(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.HasSuffix(s []byte, suffix []byte) bool
func execHasSuffix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.HasSuffix(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.Index(s []byte, sep []byte) int
func execIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Index(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.IndexAny(s []byte, chars string) int
func execIndexAny(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.IndexAny(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func bytes.IndexByte(b []byte, c byte) int
func execIndexByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.IndexByte(args[0].([]byte), args[1].(byte))
	p.Ret(2, ret)
}

// func bytes.IndexFunc(s []byte, f func(r rune) bool) int
func execIndexFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.IndexFunc(args[0].([]byte), args[1].(func(r rune) bool))
	p.Ret(2, ret)
}

// func bytes.IndexRune(s []byte, r rune) int
func execIndexRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.IndexRune(args[0].([]byte), args[1].(rune))
	p.Ret(2, ret)
}

// func bytes.Join(s [][]byte, sep []byte) []byte
func execJoin(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Join(args[0].([][]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.LastIndex(s []byte, sep []byte) int
func execLastIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.LastIndex(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.LastIndexAny(s []byte, chars string) int
func execLastIndexAny(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.LastIndexAny(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func bytes.LastIndexByte(s []byte, c byte) int
func execLastIndexByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.LastIndexByte(args[0].([]byte), args[1].(byte))
	p.Ret(2, ret)
}

// func bytes.LastIndexFunc(s []byte, f func(r rune) bool) int
func execLastIndexFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.LastIndexFunc(args[0].([]byte), args[1].(func(r rune) bool))
	p.Ret(2, ret)
}

// func bytes.Map(mapping func(r rune) rune, s []byte) []byte
func execMap(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Map(args[0].(func(r rune) rune), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.NewBuffer(buf []byte) *bytes.Buffer
func execNewBuffer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.NewBuffer(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.NewBufferString(s string) *bytes.Buffer
func execNewBufferString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.NewBufferString(args[0].(string))
	p.Ret(1, ret)
}

// func bytes.NewReader(b []byte) *bytes.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.NewReader(args[0].([]byte))
	p.Ret(1, ret)
}

// func (*bytes.Reader).Len() int
func execmReaderLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Reader).Len()
	p.Ret(1, ret)
}

// func (*bytes.Reader).Read(b []byte) (n int, err error)
func execmReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Reader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*bytes.Reader).ReadAt(b []byte, off int64) (n int, err error)
func execmReaderReadAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*bytes.Reader).ReadAt(args[1].([]byte), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func (*bytes.Reader).ReadByte() (byte, error)
func execmReaderReadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*bytes.Reader).ReadByte()
	p.Ret(1, ret, ret1)
}

// func (*bytes.Reader).ReadRune() (ch rune, size int, err error)
func execmReaderReadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*bytes.Reader).ReadRune()
	p.Ret(1, ret, ret1, ret2)
}

// func (*bytes.Reader).Reset(b []byte)
func execmReaderReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*bytes.Reader).Reset(args[1].([]byte))
}

// func (*bytes.Reader).Seek(offset int64, whence int) (int64, error)
func execmReaderSeek(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*bytes.Reader).Seek(args[1].(int64), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func (*bytes.Reader).Size() int64
func execmReaderSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Reader).Size()
	p.Ret(1, ret)
}

// func (*bytes.Reader).UnreadByte() error
func execmReaderUnreadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Reader).UnreadByte()
	p.Ret(1, ret)
}

// func (*bytes.Reader).UnreadRune() error
func execmReaderUnreadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bytes.Reader).UnreadRune()
	p.Ret(1, ret)
}

// func (*bytes.Reader).WriteTo(w io.Writer) (n int64, err error)
func execmReaderWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bytes.Reader).WriteTo(args[1].(io.Writer))
	p.Ret(2, ret, ret1)
}

// func bytes.Repeat(b []byte, count int) []byte
func execRepeat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Repeat(args[0].([]byte), args[1].(int))
	p.Ret(2, ret)
}

// func bytes.Replace(s []byte, old []byte, new []byte, n int) []byte
func execReplace(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := bytes.Replace(args[0].([]byte), args[1].([]byte), args[2].([]byte), args[3].(int))
	p.Ret(4, ret)
}

// func bytes.ReplaceAll(s []byte, old []byte, new []byte) []byte
func execReplaceAll(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := bytes.ReplaceAll(args[0].([]byte), args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret)
}

// func bytes.Runes(s []byte) []rune
func execRunes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.Runes(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.Split(s []byte, sep []byte) [][]byte
func execSplit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Split(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.SplitAfter(s []byte, sep []byte) [][]byte
func execSplitAfter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.SplitAfter(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.SplitAfterN(s []byte, sep []byte, n int) [][]byte
func execSplitAfterN(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := bytes.SplitAfterN(args[0].([]byte), args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func bytes.SplitN(s []byte, sep []byte, n int) [][]byte
func execSplitN(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := bytes.SplitN(args[0].([]byte), args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func bytes.Title(s []byte) []byte
func execTitle(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.Title(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.ToLower(s []byte) []byte
func execToLower(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.ToLower(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
func execToLowerSpecial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.ToLowerSpecial(args[0].(unicode.SpecialCase), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.ToTitle(s []byte) []byte
func execToTitle(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.ToTitle(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
func execToTitleSpecial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.ToTitleSpecial(args[0].(unicode.SpecialCase), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.ToUpper(s []byte) []byte
func execToUpper(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.ToUpper(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
func execToUpperSpecial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.ToUpperSpecial(args[0].(unicode.SpecialCase), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.ToValidUTF8(s []byte, replacement []byte) []byte
func execToValidUTF8(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.ToValidUTF8(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.Trim(s []byte, cutset string) []byte
func execTrim(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.Trim(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func bytes.TrimFunc(s []byte, f func(r rune) bool) []byte
func execTrimFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.TrimFunc(args[0].([]byte), args[1].(func(r rune) bool))
	p.Ret(2, ret)
}

// func bytes.TrimLeft(s []byte, cutset string) []byte
func execTrimLeft(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.TrimLeft(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func bytes.TrimLeftFunc(s []byte, f func(r rune) bool) []byte
func execTrimLeftFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.TrimLeftFunc(args[0].([]byte), args[1].(func(r rune) bool))
	p.Ret(2, ret)
}

// func bytes.TrimPrefix(s []byte, prefix []byte) []byte
func execTrimPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.TrimPrefix(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func bytes.TrimRight(s []byte, cutset string) []byte
func execTrimRight(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.TrimRight(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func bytes.TrimRightFunc(s []byte, f func(r rune) bool) []byte
func execTrimRightFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.TrimRightFunc(args[0].([]byte), args[1].(func(r rune) bool))
	p.Ret(2, ret)
}

// func bytes.TrimSpace(s []byte) []byte
func execTrimSpace(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bytes.TrimSpace(args[0].([]byte))
	p.Ret(1, ret)
}

// func bytes.TrimSuffix(s []byte, suffix []byte) []byte
func execTrimSuffix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bytes.TrimSuffix(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("bytes")

func init() {
	I.RegisterConsts(
		I.Const("MinRead", qspec.ConstUnboundInt, bytes.MinRead),
	)
	I.RegisterVars(
		I.Var("ErrTooLarge", &bytes.ErrTooLarge),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*bytes.Buffer)(nil))),
		I.Rtype(reflect.TypeOf((*bytes.Reader)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Buffer).Bytes", (*bytes.Buffer).Bytes, execmBufferBytes),
		I.Func("(*Buffer).Cap", (*bytes.Buffer).Cap, execmBufferCap),
		I.Func("(*Buffer).Grow", (*bytes.Buffer).Grow, execmBufferGrow),
		I.Func("(*Buffer).Len", (*bytes.Buffer).Len, execmBufferLen),
		I.Func("(*Buffer).Next", (*bytes.Buffer).Next, execmBufferNext),
		I.Func("(*Buffer).Read", (*bytes.Buffer).Read, execmBufferRead),
		I.Func("(*Buffer).ReadByte", (*bytes.Buffer).ReadByte, execmBufferReadByte),
		I.Func("(*Buffer).ReadBytes", (*bytes.Buffer).ReadBytes, execmBufferReadBytes),
		I.Func("(*Buffer).ReadFrom", (*bytes.Buffer).ReadFrom, execmBufferReadFrom),
		I.Func("(*Buffer).ReadRune", (*bytes.Buffer).ReadRune, execmBufferReadRune),
		I.Func("(*Buffer).ReadString", (*bytes.Buffer).ReadString, execmBufferReadString),
		I.Func("(*Buffer).Reset", (*bytes.Buffer).Reset, execmBufferReset),
		I.Func("(*Buffer).String", (*bytes.Buffer).String, execmBufferString),
		I.Func("(*Buffer).Truncate", (*bytes.Buffer).Truncate, execmBufferTruncate),
		I.Func("(*Buffer).UnreadByte", (*bytes.Buffer).UnreadByte, execmBufferUnreadByte),
		I.Func("(*Buffer).UnreadRune", (*bytes.Buffer).UnreadRune, execmBufferUnreadRune),
		I.Func("(*Buffer).Write", (*bytes.Buffer).Write, execmBufferWrite),
		I.Func("(*Buffer).WriteByte", (*bytes.Buffer).WriteByte, execmBufferWriteByte),
		I.Func("(*Buffer).WriteRune", (*bytes.Buffer).WriteRune, execmBufferWriteRune),
		I.Func("(*Buffer).WriteString", (*bytes.Buffer).WriteString, execmBufferWriteString),
		I.Func("(*Buffer).WriteTo", (*bytes.Buffer).WriteTo, execmBufferWriteTo),
		I.Func("Compare", bytes.Compare, execCompare),
		I.Func("Contains", bytes.Contains, execContains),
		I.Func("ContainsAny", bytes.ContainsAny, execContainsAny),
		I.Func("ContainsRune", bytes.ContainsRune, execContainsRune),
		I.Func("Count", bytes.Count, execCount),
		I.Func("Equal", bytes.Equal, execEqual),
		I.Func("EqualFold", bytes.EqualFold, execEqualFold),
		I.Func("Fields", bytes.Fields, execFields),
		I.Func("FieldsFunc", bytes.FieldsFunc, execFieldsFunc),
		I.Func("HasPrefix", bytes.HasPrefix, execHasPrefix),
		I.Func("HasSuffix", bytes.HasSuffix, execHasSuffix),
		I.Func("Index", bytes.Index, execIndex),
		I.Func("IndexAny", bytes.IndexAny, execIndexAny),
		I.Func("IndexByte", bytes.IndexByte, execIndexByte),
		I.Func("IndexFunc", bytes.IndexFunc, execIndexFunc),
		I.Func("IndexRune", bytes.IndexRune, execIndexRune),
		I.Func("Join", bytes.Join, execJoin),
		I.Func("LastIndex", bytes.LastIndex, execLastIndex),
		I.Func("LastIndexAny", bytes.LastIndexAny, execLastIndexAny),
		I.Func("LastIndexByte", bytes.LastIndexByte, execLastIndexByte),
		I.Func("LastIndexFunc", bytes.LastIndexFunc, execLastIndexFunc),
		I.Func("Map", bytes.Map, execMap),
		I.Func("NewBuffer", bytes.NewBuffer, execNewBuffer),
		I.Func("NewBufferString", bytes.NewBufferString, execNewBufferString),
		I.Func("NewReader", bytes.NewReader, execNewReader),
		I.Func("(*Reader).Len", (*bytes.Reader).Len, execmReaderLen),
		I.Func("(*Reader).Read", (*bytes.Reader).Read, execmReaderRead),
		I.Func("(*Reader).ReadAt", (*bytes.Reader).ReadAt, execmReaderReadAt),
		I.Func("(*Reader).ReadByte", (*bytes.Reader).ReadByte, execmReaderReadByte),
		I.Func("(*Reader).ReadRune", (*bytes.Reader).ReadRune, execmReaderReadRune),
		I.Func("(*Reader).Reset", (*bytes.Reader).Reset, execmReaderReset),
		I.Func("(*Reader).Seek", (*bytes.Reader).Seek, execmReaderSeek),
		I.Func("(*Reader).Size", (*bytes.Reader).Size, execmReaderSize),
		I.Func("(*Reader).UnreadByte", (*bytes.Reader).UnreadByte, execmReaderUnreadByte),
		I.Func("(*Reader).UnreadRune", (*bytes.Reader).UnreadRune, execmReaderUnreadRune),
		I.Func("(*Reader).WriteTo", (*bytes.Reader).WriteTo, execmReaderWriteTo),
		I.Func("Repeat", bytes.Repeat, execRepeat),
		I.Func("Replace", bytes.Replace, execReplace),
		I.Func("ReplaceAll", bytes.ReplaceAll, execReplaceAll),
		I.Func("Runes", bytes.Runes, execRunes),
		I.Func("Split", bytes.Split, execSplit),
		I.Func("SplitAfter", bytes.SplitAfter, execSplitAfter),
		I.Func("SplitAfterN", bytes.SplitAfterN, execSplitAfterN),
		I.Func("SplitN", bytes.SplitN, execSplitN),
		I.Func("Title", bytes.Title, execTitle),
		I.Func("ToLower", bytes.ToLower, execToLower),
		I.Func("ToLowerSpecial", bytes.ToLowerSpecial, execToLowerSpecial),
		I.Func("ToTitle", bytes.ToTitle, execToTitle),
		I.Func("ToTitleSpecial", bytes.ToTitleSpecial, execToTitleSpecial),
		I.Func("ToUpper", bytes.ToUpper, execToUpper),
		I.Func("ToUpperSpecial", bytes.ToUpperSpecial, execToUpperSpecial),
		I.Func("ToValidUTF8", bytes.ToValidUTF8, execToValidUTF8),
		I.Func("Trim", bytes.Trim, execTrim),
		I.Func("TrimFunc", bytes.TrimFunc, execTrimFunc),
		I.Func("TrimLeft", bytes.TrimLeft, execTrimLeft),
		I.Func("TrimLeftFunc", bytes.TrimLeftFunc, execTrimLeftFunc),
		I.Func("TrimPrefix", bytes.TrimPrefix, execTrimPrefix),
		I.Func("TrimRight", bytes.TrimRight, execTrimRight),
		I.Func("TrimRightFunc", bytes.TrimRightFunc, execTrimRightFunc),
		I.Func("TrimSpace", bytes.TrimSpace, execTrimSpace),
		I.Func("TrimSuffix", bytes.TrimSuffix, execTrimSuffix),
	)
}
