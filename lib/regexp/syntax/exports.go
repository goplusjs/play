package syntax

import (
	"reflect"
	"regexp/syntax"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func syntax.Compile(re *syntax.Regexp) (*syntax.Prog, error)
func execCompile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syntax.Compile(args[0].(*syntax.Regexp))
	p.Ret(1, ret, ret1)
}

// func syntax.EmptyOpContext(r1 rune, r2 rune) syntax.EmptyOp
func execEmptyOpContext(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syntax.EmptyOpContext(args[0].(rune), args[1].(rune))
	p.Ret(2, ret)
}

// func (*syntax.Error).Error() string
func execmErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Error).Error()
	p.Ret(1, ret)
}

// func (syntax.ErrorCode).String() string
func execmErrorCodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syntax.ErrorCode).String()
	p.Ret(1, ret)
}

// func (*syntax.Inst).MatchEmptyWidth(before rune, after rune) bool
func execmInstMatchEmptyWidth(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*syntax.Inst).MatchEmptyWidth(args[1].(rune), args[2].(rune))
	p.Ret(3, ret)
}

// func (*syntax.Inst).MatchRune(r rune) bool
func execmInstMatchRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syntax.Inst).MatchRune(args[1].(rune))
	p.Ret(2, ret)
}

// func (*syntax.Inst).MatchRunePos(r rune) int
func execmInstMatchRunePos(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syntax.Inst).MatchRunePos(args[1].(rune))
	p.Ret(2, ret)
}

// func (*syntax.Inst).String() string
func execmInstString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Inst).String()
	p.Ret(1, ret)
}

// func (syntax.InstOp).String() string
func execmInstOpString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syntax.InstOp).String()
	p.Ret(1, ret)
}

// func syntax.IsWordChar(r rune) bool
func execIsWordChar(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syntax.IsWordChar(args[0].(rune))
	p.Ret(1, ret)
}

// func (syntax.Op).String() string
func execmOpString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syntax.Op).String()
	p.Ret(1, ret)
}

// func syntax.Parse(s string, flags syntax.Flags) (*syntax.Regexp, error)
func execParse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syntax.Parse(args[0].(string), syntax.Flags(args[1].(uint16)))
	p.Ret(2, ret, ret1)
}

// func (*syntax.Prog).Prefix() (prefix string, complete bool)
func execmProgPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*syntax.Prog).Prefix()
	p.Ret(1, ret, ret1)
}

// func (*syntax.Prog).StartCond() syntax.EmptyOp
func execmProgStartCond(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Prog).StartCond()
	p.Ret(1, ret)
}

// func (*syntax.Prog).String() string
func execmProgString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Prog).String()
	p.Ret(1, ret)
}

// func (*syntax.Regexp).CapNames() []string
func execmRegexpCapNames(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Regexp).CapNames()
	p.Ret(1, ret)
}

// func (*syntax.Regexp).Equal(y *syntax.Regexp) bool
func execmRegexpEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syntax.Regexp).Equal(args[1].(*syntax.Regexp))
	p.Ret(2, ret)
}

// func (*syntax.Regexp).MaxCap() int
func execmRegexpMaxCap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Regexp).MaxCap()
	p.Ret(1, ret)
}

// func (*syntax.Regexp).Simplify() *syntax.Regexp
func execmRegexpSimplify(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Regexp).Simplify()
	p.Ret(1, ret)
}

// func (*syntax.Regexp).String() string
func execmRegexpString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syntax.Regexp).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("regexp/syntax")

func init() {
	I.RegisterConsts(
		I.Const("ClassNL", reflect.Uint16, syntax.ClassNL),
		I.Const("DotNL", reflect.Uint16, syntax.DotNL),
		I.Const("EmptyBeginLine", reflect.Uint8, syntax.EmptyBeginLine),
		I.Const("EmptyBeginText", reflect.Uint8, syntax.EmptyBeginText),
		I.Const("EmptyEndLine", reflect.Uint8, syntax.EmptyEndLine),
		I.Const("EmptyEndText", reflect.Uint8, syntax.EmptyEndText),
		I.Const("EmptyNoWordBoundary", reflect.Uint8, syntax.EmptyNoWordBoundary),
		I.Const("EmptyWordBoundary", reflect.Uint8, syntax.EmptyWordBoundary),
		I.Const("ErrInternalError", reflect.String, syntax.ErrInternalError),
		I.Const("ErrInvalidCharClass", reflect.String, syntax.ErrInvalidCharClass),
		I.Const("ErrInvalidCharRange", reflect.String, syntax.ErrInvalidCharRange),
		I.Const("ErrInvalidEscape", reflect.String, syntax.ErrInvalidEscape),
		I.Const("ErrInvalidNamedCapture", reflect.String, syntax.ErrInvalidNamedCapture),
		I.Const("ErrInvalidPerlOp", reflect.String, syntax.ErrInvalidPerlOp),
		I.Const("ErrInvalidRepeatOp", reflect.String, syntax.ErrInvalidRepeatOp),
		I.Const("ErrInvalidRepeatSize", reflect.String, syntax.ErrInvalidRepeatSize),
		I.Const("ErrInvalidUTF8", reflect.String, syntax.ErrInvalidUTF8),
		I.Const("ErrMissingBracket", reflect.String, syntax.ErrMissingBracket),
		I.Const("ErrMissingParen", reflect.String, syntax.ErrMissingParen),
		I.Const("ErrMissingRepeatArgument", reflect.String, syntax.ErrMissingRepeatArgument),
		I.Const("ErrTrailingBackslash", reflect.String, syntax.ErrTrailingBackslash),
		I.Const("ErrUnexpectedParen", reflect.String, syntax.ErrUnexpectedParen),
		I.Const("FoldCase", reflect.Uint16, syntax.FoldCase),
		I.Const("InstAlt", reflect.Uint8, syntax.InstAlt),
		I.Const("InstAltMatch", reflect.Uint8, syntax.InstAltMatch),
		I.Const("InstCapture", reflect.Uint8, syntax.InstCapture),
		I.Const("InstEmptyWidth", reflect.Uint8, syntax.InstEmptyWidth),
		I.Const("InstFail", reflect.Uint8, syntax.InstFail),
		I.Const("InstMatch", reflect.Uint8, syntax.InstMatch),
		I.Const("InstNop", reflect.Uint8, syntax.InstNop),
		I.Const("InstRune", reflect.Uint8, syntax.InstRune),
		I.Const("InstRune1", reflect.Uint8, syntax.InstRune1),
		I.Const("InstRuneAny", reflect.Uint8, syntax.InstRuneAny),
		I.Const("InstRuneAnyNotNL", reflect.Uint8, syntax.InstRuneAnyNotNL),
		I.Const("Literal", reflect.Uint16, syntax.Literal),
		I.Const("MatchNL", reflect.Uint16, syntax.MatchNL),
		I.Const("NonGreedy", reflect.Uint16, syntax.NonGreedy),
		I.Const("OneLine", reflect.Uint16, syntax.OneLine),
		I.Const("OpAlternate", reflect.Uint8, syntax.OpAlternate),
		I.Const("OpAnyChar", reflect.Uint8, syntax.OpAnyChar),
		I.Const("OpAnyCharNotNL", reflect.Uint8, syntax.OpAnyCharNotNL),
		I.Const("OpBeginLine", reflect.Uint8, syntax.OpBeginLine),
		I.Const("OpBeginText", reflect.Uint8, syntax.OpBeginText),
		I.Const("OpCapture", reflect.Uint8, syntax.OpCapture),
		I.Const("OpCharClass", reflect.Uint8, syntax.OpCharClass),
		I.Const("OpConcat", reflect.Uint8, syntax.OpConcat),
		I.Const("OpEmptyMatch", reflect.Uint8, syntax.OpEmptyMatch),
		I.Const("OpEndLine", reflect.Uint8, syntax.OpEndLine),
		I.Const("OpEndText", reflect.Uint8, syntax.OpEndText),
		I.Const("OpLiteral", reflect.Uint8, syntax.OpLiteral),
		I.Const("OpNoMatch", reflect.Uint8, syntax.OpNoMatch),
		I.Const("OpNoWordBoundary", reflect.Uint8, syntax.OpNoWordBoundary),
		I.Const("OpPlus", reflect.Uint8, syntax.OpPlus),
		I.Const("OpQuest", reflect.Uint8, syntax.OpQuest),
		I.Const("OpRepeat", reflect.Uint8, syntax.OpRepeat),
		I.Const("OpStar", reflect.Uint8, syntax.OpStar),
		I.Const("OpWordBoundary", reflect.Uint8, syntax.OpWordBoundary),
		I.Const("POSIX", reflect.Uint16, syntax.POSIX),
		I.Const("Perl", reflect.Uint16, syntax.Perl),
		I.Const("PerlX", reflect.Uint16, syntax.PerlX),
		I.Const("Simple", reflect.Uint16, syntax.Simple),
		I.Const("UnicodeGroups", reflect.Uint16, syntax.UnicodeGroups),
		I.Const("WasDollar", reflect.Uint16, syntax.WasDollar),
	)
	I.RegisterTypes(
		I.Type("EmptyOp", qspec.TyUint8),
		I.Rtype(reflect.TypeOf((*syntax.Error)(nil))),
		I.Type("ErrorCode", qspec.TyString),
		I.Type("Flags", qspec.TyUint16),
		I.Rtype(reflect.TypeOf((*syntax.Inst)(nil))),
		I.Type("InstOp", qspec.TyUint8),
		I.Type("Op", qspec.TyUint8),
		I.Rtype(reflect.TypeOf((*syntax.Prog)(nil))),
		I.Rtype(reflect.TypeOf((*syntax.Regexp)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Compile", syntax.Compile, execCompile),
		I.Func("EmptyOpContext", syntax.EmptyOpContext, execEmptyOpContext),
		I.Func("(*Error).Error", (*syntax.Error).Error, execmErrorError),
		I.Func("(ErrorCode).String", (syntax.ErrorCode).String, execmErrorCodeString),
		I.Func("(*Inst).MatchEmptyWidth", (*syntax.Inst).MatchEmptyWidth, execmInstMatchEmptyWidth),
		I.Func("(*Inst).MatchRune", (*syntax.Inst).MatchRune, execmInstMatchRune),
		I.Func("(*Inst).MatchRunePos", (*syntax.Inst).MatchRunePos, execmInstMatchRunePos),
		I.Func("(*Inst).String", (*syntax.Inst).String, execmInstString),
		I.Func("(InstOp).String", (syntax.InstOp).String, execmInstOpString),
		I.Func("IsWordChar", syntax.IsWordChar, execIsWordChar),
		I.Func("(Op).String", (syntax.Op).String, execmOpString),
		I.Func("Parse", syntax.Parse, execParse),
		I.Func("(*Prog).Prefix", (*syntax.Prog).Prefix, execmProgPrefix),
		I.Func("(*Prog).StartCond", (*syntax.Prog).StartCond, execmProgStartCond),
		I.Func("(*Prog).String", (*syntax.Prog).String, execmProgString),
		I.Func("(*Regexp).CapNames", (*syntax.Regexp).CapNames, execmRegexpCapNames),
		I.Func("(*Regexp).Equal", (*syntax.Regexp).Equal, execmRegexpEqual),
		I.Func("(*Regexp).MaxCap", (*syntax.Regexp).MaxCap, execmRegexpMaxCap),
		I.Func("(*Regexp).Simplify", (*syntax.Regexp).Simplify, execmRegexpSimplify),
		I.Func("(*Regexp).String", (*syntax.Regexp).String, execmRegexpString),
	)
}
