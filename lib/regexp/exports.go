package regexp

import (
	"io"
	"reflect"
	"regexp"

	"github.com/qiniu/goplus/gop"
)

// func regexp.Compile(expr string) (*regexp.Regexp, error)
func execCompile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := regexp.Compile(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func regexp.CompilePOSIX(expr string) (*regexp.Regexp, error)
func execCompilePOSIX(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := regexp.CompilePOSIX(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func regexp.Match(pattern string, b []byte) (matched bool, err error)
func execMatch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := regexp.Match(args[0].(string), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func regexp.MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
func execMatchReader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := regexp.MatchReader(args[0].(string), args[1].(io.RuneReader))
	p.Ret(2, ret, ret1)
}

// func regexp.MatchString(pattern string, s string) (matched bool, err error)
func execMatchString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := regexp.MatchString(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func regexp.MustCompile(str string) *regexp.Regexp
func execMustCompile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := regexp.MustCompile(args[0].(string))
	p.Ret(1, ret)
}

// func regexp.MustCompilePOSIX(str string) *regexp.Regexp
func execMustCompilePOSIX(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := regexp.MustCompilePOSIX(args[0].(string))
	p.Ret(1, ret)
}

// func regexp.QuoteMeta(s string) string
func execQuoteMeta(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := regexp.QuoteMeta(args[0].(string))
	p.Ret(1, ret)
}

// func (*regexp.Regexp).Copy() *regexp.Regexp
func execmRegexpCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*regexp.Regexp).Copy()
	p.Ret(1, ret)
}

// func (*regexp.Regexp).Expand(dst []byte, template []byte, src []byte, match []int) []byte
func execmRegexpExpand(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := args[0].(*regexp.Regexp).Expand(args[1].([]byte), args[2].([]byte), args[3].([]byte), args[4].([]int))
	p.Ret(5, ret)
}

// func (*regexp.Regexp).ExpandString(dst []byte, template string, src string, match []int) []byte
func execmRegexpExpandString(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := args[0].(*regexp.Regexp).ExpandString(args[1].([]byte), args[2].(string), args[3].(string), args[4].([]int))
	p.Ret(5, ret)
}

// func (*regexp.Regexp).Find(b []byte) []byte
func execmRegexpFind(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).Find(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindAll(b []byte, n int) [][]byte
func execmRegexpFindAll(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAll(args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindAllIndex(b []byte, n int) [][]int
func execmRegexpFindAllIndex(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAllIndex(args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindAllString(s string, n int) []string
func execmRegexpFindAllString(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAllString(args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindAllStringIndex(s string, n int) [][]int
func execmRegexpFindAllStringIndex(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAllStringIndex(args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindAllStringSubmatch(s string, n int) [][]string
func execmRegexpFindAllStringSubmatch(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAllStringSubmatch(args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindAllStringSubmatchIndex(s string, n int) [][]int
func execmRegexpFindAllStringSubmatchIndex(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAllStringSubmatchIndex(args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindAllSubmatch(b []byte, n int) [][][]byte
func execmRegexpFindAllSubmatch(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAllSubmatch(args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindAllSubmatchIndex(b []byte, n int) [][]int
func execmRegexpFindAllSubmatchIndex(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).FindAllSubmatchIndex(args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).FindIndex(b []byte) (loc []int)
func execmRegexpFindIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindIndex(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindReaderIndex(r io.RuneReader) (loc []int)
func execmRegexpFindReaderIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindReaderIndex(args[1].(io.RuneReader))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindReaderSubmatchIndex(r io.RuneReader) []int
func execmRegexpFindReaderSubmatchIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindReaderSubmatchIndex(args[1].(io.RuneReader))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindString(s string) string
func execmRegexpFindString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindString(args[1].(string))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindStringIndex(s string) (loc []int)
func execmRegexpFindStringIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindStringIndex(args[1].(string))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindStringSubmatch(s string) []string
func execmRegexpFindStringSubmatch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindStringSubmatch(args[1].(string))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindStringSubmatchIndex(s string) []int
func execmRegexpFindStringSubmatchIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindStringSubmatchIndex(args[1].(string))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindSubmatch(b []byte) [][]byte
func execmRegexpFindSubmatch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindSubmatch(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).FindSubmatchIndex(b []byte) []int
func execmRegexpFindSubmatchIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).FindSubmatchIndex(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).LiteralPrefix() (prefix string, complete bool)
func execmRegexpLiteralPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*regexp.Regexp).LiteralPrefix()
	p.Ret(1, ret, ret1)
}

// func (*regexp.Regexp).Longest()
func execmRegexpLongest(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*regexp.Regexp).Longest()
}

// func (*regexp.Regexp).Match(b []byte) bool
func execmRegexpMatch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).Match(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).MatchReader(r io.RuneReader) bool
func execmRegexpMatchReader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).MatchReader(args[1].(io.RuneReader))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).MatchString(s string) bool
func execmRegexpMatchString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*regexp.Regexp).MatchString(args[1].(string))
	p.Ret(2, ret)
}

// func (*regexp.Regexp).NumSubexp() int
func execmRegexpNumSubexp(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*regexp.Regexp).NumSubexp()
	p.Ret(1, ret)
}

// func (*regexp.Regexp).ReplaceAll(src []byte, repl []byte) []byte
func execmRegexpReplaceAll(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).ReplaceAll(args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
func execmRegexpReplaceAllFunc(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).ReplaceAllFunc(args[1].([]byte), args[2].(func([]byte) []byte))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).ReplaceAllLiteral(src []byte, repl []byte) []byte
func execmRegexpReplaceAllLiteral(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).ReplaceAllLiteral(args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).ReplaceAllLiteralString(src string, repl string) string
func execmRegexpReplaceAllLiteralString(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).ReplaceAllLiteralString(args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).ReplaceAllString(src string, repl string) string
func execmRegexpReplaceAllString(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).ReplaceAllString(args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).ReplaceAllStringFunc(src string, repl func(string) string) string
func execmRegexpReplaceAllStringFunc(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).ReplaceAllStringFunc(args[1].(string), args[2].(func(string) string))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).Split(s string, n int) []string
func execmRegexpSplit(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*regexp.Regexp).Split(args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func (*regexp.Regexp).String() string
func execmRegexpString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*regexp.Regexp).String()
	p.Ret(1, ret)
}

// func (*regexp.Regexp).SubexpNames() []string
func execmRegexpSubexpNames(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*regexp.Regexp).SubexpNames()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("regexp")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*regexp.Regexp)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Compile", regexp.Compile, execCompile),
		I.Func("CompilePOSIX", regexp.CompilePOSIX, execCompilePOSIX),
		I.Func("Match", regexp.Match, execMatch),
		I.Func("MatchReader", regexp.MatchReader, execMatchReader),
		I.Func("MatchString", regexp.MatchString, execMatchString),
		I.Func("MustCompile", regexp.MustCompile, execMustCompile),
		I.Func("MustCompilePOSIX", regexp.MustCompilePOSIX, execMustCompilePOSIX),
		I.Func("QuoteMeta", regexp.QuoteMeta, execQuoteMeta),
		I.Func("(*Regexp).Copy", (*regexp.Regexp).Copy, execmRegexpCopy),
		I.Func("(*Regexp).Expand", (*regexp.Regexp).Expand, execmRegexpExpand),
		I.Func("(*Regexp).ExpandString", (*regexp.Regexp).ExpandString, execmRegexpExpandString),
		I.Func("(*Regexp).Find", (*regexp.Regexp).Find, execmRegexpFind),
		I.Func("(*Regexp).FindAll", (*regexp.Regexp).FindAll, execmRegexpFindAll),
		I.Func("(*Regexp).FindAllIndex", (*regexp.Regexp).FindAllIndex, execmRegexpFindAllIndex),
		I.Func("(*Regexp).FindAllString", (*regexp.Regexp).FindAllString, execmRegexpFindAllString),
		I.Func("(*Regexp).FindAllStringIndex", (*regexp.Regexp).FindAllStringIndex, execmRegexpFindAllStringIndex),
		I.Func("(*Regexp).FindAllStringSubmatch", (*regexp.Regexp).FindAllStringSubmatch, execmRegexpFindAllStringSubmatch),
		I.Func("(*Regexp).FindAllStringSubmatchIndex", (*regexp.Regexp).FindAllStringSubmatchIndex, execmRegexpFindAllStringSubmatchIndex),
		I.Func("(*Regexp).FindAllSubmatch", (*regexp.Regexp).FindAllSubmatch, execmRegexpFindAllSubmatch),
		I.Func("(*Regexp).FindAllSubmatchIndex", (*regexp.Regexp).FindAllSubmatchIndex, execmRegexpFindAllSubmatchIndex),
		I.Func("(*Regexp).FindIndex", (*regexp.Regexp).FindIndex, execmRegexpFindIndex),
		I.Func("(*Regexp).FindReaderIndex", (*regexp.Regexp).FindReaderIndex, execmRegexpFindReaderIndex),
		I.Func("(*Regexp).FindReaderSubmatchIndex", (*regexp.Regexp).FindReaderSubmatchIndex, execmRegexpFindReaderSubmatchIndex),
		I.Func("(*Regexp).FindString", (*regexp.Regexp).FindString, execmRegexpFindString),
		I.Func("(*Regexp).FindStringIndex", (*regexp.Regexp).FindStringIndex, execmRegexpFindStringIndex),
		I.Func("(*Regexp).FindStringSubmatch", (*regexp.Regexp).FindStringSubmatch, execmRegexpFindStringSubmatch),
		I.Func("(*Regexp).FindStringSubmatchIndex", (*regexp.Regexp).FindStringSubmatchIndex, execmRegexpFindStringSubmatchIndex),
		I.Func("(*Regexp).FindSubmatch", (*regexp.Regexp).FindSubmatch, execmRegexpFindSubmatch),
		I.Func("(*Regexp).FindSubmatchIndex", (*regexp.Regexp).FindSubmatchIndex, execmRegexpFindSubmatchIndex),
		I.Func("(*Regexp).LiteralPrefix", (*regexp.Regexp).LiteralPrefix, execmRegexpLiteralPrefix),
		I.Func("(*Regexp).Longest", (*regexp.Regexp).Longest, execmRegexpLongest),
		I.Func("(*Regexp).Match", (*regexp.Regexp).Match, execmRegexpMatch),
		I.Func("(*Regexp).MatchReader", (*regexp.Regexp).MatchReader, execmRegexpMatchReader),
		I.Func("(*Regexp).MatchString", (*regexp.Regexp).MatchString, execmRegexpMatchString),
		I.Func("(*Regexp).NumSubexp", (*regexp.Regexp).NumSubexp, execmRegexpNumSubexp),
		I.Func("(*Regexp).ReplaceAll", (*regexp.Regexp).ReplaceAll, execmRegexpReplaceAll),
		I.Func("(*Regexp).ReplaceAllFunc", (*regexp.Regexp).ReplaceAllFunc, execmRegexpReplaceAllFunc),
		I.Func("(*Regexp).ReplaceAllLiteral", (*regexp.Regexp).ReplaceAllLiteral, execmRegexpReplaceAllLiteral),
		I.Func("(*Regexp).ReplaceAllLiteralString", (*regexp.Regexp).ReplaceAllLiteralString, execmRegexpReplaceAllLiteralString),
		I.Func("(*Regexp).ReplaceAllString", (*regexp.Regexp).ReplaceAllString, execmRegexpReplaceAllString),
		I.Func("(*Regexp).ReplaceAllStringFunc", (*regexp.Regexp).ReplaceAllStringFunc, execmRegexpReplaceAllStringFunc),
		I.Func("(*Regexp).Split", (*regexp.Regexp).Split, execmRegexpSplit),
		I.Func("(*Regexp).String", (*regexp.Regexp).String, execmRegexpString),
		I.Func("(*Regexp).SubexpNames", (*regexp.Regexp).SubexpNames, execmRegexpSubexpNames),
	)
}
