package token

import (
	"go/token"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*token.File).AddLine(offset int)
func execmFileAddLine(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*token.File).AddLine(args[1].(int))
}

// func (*token.File).AddLineColumnInfo(offset int, filename string, line int, column int)
func execmFileAddLineColumnInfo(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*token.File).AddLineColumnInfo(args[1].(int), args[2].(string), args[3].(int), args[4].(int))
}

// func (*token.File).AddLineInfo(offset int, filename string, line int)
func execmFileAddLineInfo(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*token.File).AddLineInfo(args[1].(int), args[2].(string), args[3].(int))
}

// func (*token.File).Base() int
func execmFileBase(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*token.File).Base()
	p.Ret(1, ret)
}

// func (*token.File).Line(p token.Pos) int
func execmFileLine(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.File).Line(token.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*token.File).LineCount() int
func execmFileLineCount(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*token.File).LineCount()
	p.Ret(1, ret)
}

// func (*token.File).LineStart(line int) token.Pos
func execmFileLineStart(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.File).LineStart(args[1].(int))
	p.Ret(2, ret)
}

// func (*token.File).MergeLine(line int)
func execmFileMergeLine(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*token.File).MergeLine(args[1].(int))
}

// func (*token.File).Name() string
func execmFileName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*token.File).Name()
	p.Ret(1, ret)
}

// func (*token.File).Offset(p token.Pos) int
func execmFileOffset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.File).Offset(token.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*token.File).Pos(offset int) token.Pos
func execmFilePos(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.File).Pos(args[1].(int))
	p.Ret(2, ret)
}

// func (*token.File).Position(p token.Pos) (pos token.Position)
func execmFilePosition(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.File).Position(token.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*token.File).PositionFor(p token.Pos, adjusted bool) (pos token.Position)
func execmFilePositionFor(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*token.File).PositionFor(token.Pos(args[1].(int)), args[2].(bool))
	p.Ret(3, ret)
}

// func (*token.File).SetLines(lines []int) bool
func execmFileSetLines(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.File).SetLines(args[1].([]int))
	p.Ret(2, ret)
}

// func (*token.File).SetLinesForContent(content []byte)
func execmFileSetLinesForContent(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*token.File).SetLinesForContent(args[1].([]byte))
}

// func (*token.File).Size() int
func execmFileSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*token.File).Size()
	p.Ret(1, ret)
}

// func (*token.FileSet).AddFile(filename string, base int, size int) *token.File
func execmFileSetAddFile(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*token.FileSet).AddFile(args[1].(string), args[2].(int), args[3].(int))
	p.Ret(4, ret)
}

// func (*token.FileSet).Base() int
func execmFileSetBase(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*token.FileSet).Base()
	p.Ret(1, ret)
}

// func (*token.FileSet).File(p token.Pos) (f *token.File)
func execmFileSetFile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.FileSet).File(token.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*token.FileSet).Iterate(f func(*token.File) bool)
func execmFileSetIterate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*token.FileSet).Iterate(args[1].(func(*token.File) bool))
}

// func (*token.FileSet).Position(p token.Pos) (pos token.Position)
func execmFileSetPosition(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.FileSet).Position(token.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*token.FileSet).PositionFor(p token.Pos, adjusted bool) (pos token.Position)
func execmFileSetPositionFor(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*token.FileSet).PositionFor(token.Pos(args[1].(int)), args[2].(bool))
	p.Ret(3, ret)
}

// func (*token.FileSet).Read(decode func(interface{}) error) error
func execmFileSetRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.FileSet).Read(args[1].(func(interface{}) error))
	p.Ret(2, ret)
}

// func (*token.FileSet).Write(encode func(interface{}) error) error
func execmFileSetWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*token.FileSet).Write(args[1].(func(interface{}) error))
	p.Ret(2, ret)
}

// func token.IsExported(name string) bool
func execIsExported(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := token.IsExported(args[0].(string))
	p.Ret(1, ret)
}

// func token.IsIdentifier(name string) bool
func execIsIdentifier(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := token.IsIdentifier(args[0].(string))
	p.Ret(1, ret)
}

// func token.IsKeyword(name string) bool
func execIsKeyword(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := token.IsKeyword(args[0].(string))
	p.Ret(1, ret)
}

// func token.Lookup(ident string) token.Token
func execLookup(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := token.Lookup(args[0].(string))
	p.Ret(1, ret)
}

// func token.NewFileSet() *token.FileSet
func execNewFileSet(_ int, p *gop.Context) {
	ret := token.NewFileSet()
	p.Ret(0, ret)
}

// func (token.Pos).IsValid() bool
func execmPosIsValid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(token.Pos).IsValid()
	p.Ret(1, ret)
}

// func (*token.Position).IsValid() bool
func execmPositionIsValid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*token.Position).IsValid()
	p.Ret(1, ret)
}

// func (token.Position).String() string
func execmPositionString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(token.Position).String()
	p.Ret(1, ret)
}

// func (token.Token).IsKeyword() bool
func execmTokenIsKeyword(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(token.Token).IsKeyword()
	p.Ret(1, ret)
}

// func (token.Token).IsLiteral() bool
func execmTokenIsLiteral(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(token.Token).IsLiteral()
	p.Ret(1, ret)
}

// func (token.Token).IsOperator() bool
func execmTokenIsOperator(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(token.Token).IsOperator()
	p.Ret(1, ret)
}

// func (token.Token).Precedence() int
func execmTokenPrecedence(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(token.Token).Precedence()
	p.Ret(1, ret)
}

// func (token.Token).String() string
func execmTokenString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(token.Token).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/token")

func init() {
	I.RegisterConsts(
		I.Const("ADD", reflect.Int, token.ADD),
		I.Const("ADD_ASSIGN", reflect.Int, token.ADD_ASSIGN),
		I.Const("AND", reflect.Int, token.AND),
		I.Const("AND_ASSIGN", reflect.Int, token.AND_ASSIGN),
		I.Const("AND_NOT", reflect.Int, token.AND_NOT),
		I.Const("AND_NOT_ASSIGN", reflect.Int, token.AND_NOT_ASSIGN),
		I.Const("ARROW", reflect.Int, token.ARROW),
		I.Const("ASSIGN", reflect.Int, token.ASSIGN),
		I.Const("BREAK", reflect.Int, token.BREAK),
		I.Const("CASE", reflect.Int, token.CASE),
		I.Const("CHAN", reflect.Int, token.CHAN),
		I.Const("CHAR", reflect.Int, token.CHAR),
		I.Const("COLON", reflect.Int, token.COLON),
		I.Const("COMMA", reflect.Int, token.COMMA),
		I.Const("COMMENT", reflect.Int, token.COMMENT),
		I.Const("CONST", reflect.Int, token.CONST),
		I.Const("CONTINUE", reflect.Int, token.CONTINUE),
		I.Const("DEC", reflect.Int, token.DEC),
		I.Const("DEFAULT", reflect.Int, token.DEFAULT),
		I.Const("DEFER", reflect.Int, token.DEFER),
		I.Const("DEFINE", reflect.Int, token.DEFINE),
		I.Const("ELLIPSIS", reflect.Int, token.ELLIPSIS),
		I.Const("ELSE", reflect.Int, token.ELSE),
		I.Const("EOF", reflect.Int, token.EOF),
		I.Const("EQL", reflect.Int, token.EQL),
		I.Const("FALLTHROUGH", reflect.Int, token.FALLTHROUGH),
		I.Const("FLOAT", reflect.Int, token.FLOAT),
		I.Const("FOR", reflect.Int, token.FOR),
		I.Const("FUNC", reflect.Int, token.FUNC),
		I.Const("GEQ", reflect.Int, token.GEQ),
		I.Const("GO", reflect.Int, token.GO),
		I.Const("GOTO", reflect.Int, token.GOTO),
		I.Const("GTR", reflect.Int, token.GTR),
		I.Const("HighestPrec", qspec.ConstUnboundInt, token.HighestPrec),
		I.Const("IDENT", reflect.Int, token.IDENT),
		I.Const("IF", reflect.Int, token.IF),
		I.Const("ILLEGAL", reflect.Int, token.ILLEGAL),
		I.Const("IMAG", reflect.Int, token.IMAG),
		I.Const("IMPORT", reflect.Int, token.IMPORT),
		I.Const("INC", reflect.Int, token.INC),
		I.Const("INT", reflect.Int, token.INT),
		I.Const("INTERFACE", reflect.Int, token.INTERFACE),
		I.Const("LAND", reflect.Int, token.LAND),
		I.Const("LBRACE", reflect.Int, token.LBRACE),
		I.Const("LBRACK", reflect.Int, token.LBRACK),
		I.Const("LEQ", reflect.Int, token.LEQ),
		I.Const("LOR", reflect.Int, token.LOR),
		I.Const("LPAREN", reflect.Int, token.LPAREN),
		I.Const("LSS", reflect.Int, token.LSS),
		I.Const("LowestPrec", qspec.ConstUnboundInt, token.LowestPrec),
		I.Const("MAP", reflect.Int, token.MAP),
		I.Const("MUL", reflect.Int, token.MUL),
		I.Const("MUL_ASSIGN", reflect.Int, token.MUL_ASSIGN),
		I.Const("NEQ", reflect.Int, token.NEQ),
		I.Const("NOT", reflect.Int, token.NOT),
		I.Const("NoPos", reflect.Int, token.NoPos),
		I.Const("OR", reflect.Int, token.OR),
		I.Const("OR_ASSIGN", reflect.Int, token.OR_ASSIGN),
		I.Const("PACKAGE", reflect.Int, token.PACKAGE),
		I.Const("PERIOD", reflect.Int, token.PERIOD),
		I.Const("QUO", reflect.Int, token.QUO),
		I.Const("QUO_ASSIGN", reflect.Int, token.QUO_ASSIGN),
		I.Const("RANGE", reflect.Int, token.RANGE),
		I.Const("RBRACE", reflect.Int, token.RBRACE),
		I.Const("RBRACK", reflect.Int, token.RBRACK),
		I.Const("REM", reflect.Int, token.REM),
		I.Const("REM_ASSIGN", reflect.Int, token.REM_ASSIGN),
		I.Const("RETURN", reflect.Int, token.RETURN),
		I.Const("RPAREN", reflect.Int, token.RPAREN),
		I.Const("SELECT", reflect.Int, token.SELECT),
		I.Const("SEMICOLON", reflect.Int, token.SEMICOLON),
		I.Const("SHL", reflect.Int, token.SHL),
		I.Const("SHL_ASSIGN", reflect.Int, token.SHL_ASSIGN),
		I.Const("SHR", reflect.Int, token.SHR),
		I.Const("SHR_ASSIGN", reflect.Int, token.SHR_ASSIGN),
		I.Const("STRING", reflect.Int, token.STRING),
		I.Const("STRUCT", reflect.Int, token.STRUCT),
		I.Const("SUB", reflect.Int, token.SUB),
		I.Const("SUB_ASSIGN", reflect.Int, token.SUB_ASSIGN),
		I.Const("SWITCH", reflect.Int, token.SWITCH),
		I.Const("TYPE", reflect.Int, token.TYPE),
		I.Const("UnaryPrec", qspec.ConstUnboundInt, token.UnaryPrec),
		I.Const("VAR", reflect.Int, token.VAR),
		I.Const("XOR", reflect.Int, token.XOR),
		I.Const("XOR_ASSIGN", reflect.Int, token.XOR_ASSIGN),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*token.File)(nil))),
		I.Rtype(reflect.TypeOf((*token.FileSet)(nil))),
		I.Type("Pos", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*token.Position)(nil))),
		I.Type("Token", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("(*File).AddLine", (*token.File).AddLine, execmFileAddLine),
		I.Func("(*File).AddLineColumnInfo", (*token.File).AddLineColumnInfo, execmFileAddLineColumnInfo),
		I.Func("(*File).AddLineInfo", (*token.File).AddLineInfo, execmFileAddLineInfo),
		I.Func("(*File).Base", (*token.File).Base, execmFileBase),
		I.Func("(*File).Line", (*token.File).Line, execmFileLine),
		I.Func("(*File).LineCount", (*token.File).LineCount, execmFileLineCount),
		I.Func("(*File).LineStart", (*token.File).LineStart, execmFileLineStart),
		I.Func("(*File).MergeLine", (*token.File).MergeLine, execmFileMergeLine),
		I.Func("(*File).Name", (*token.File).Name, execmFileName),
		I.Func("(*File).Offset", (*token.File).Offset, execmFileOffset),
		I.Func("(*File).Pos", (*token.File).Pos, execmFilePos),
		I.Func("(*File).Position", (*token.File).Position, execmFilePosition),
		I.Func("(*File).PositionFor", (*token.File).PositionFor, execmFilePositionFor),
		I.Func("(*File).SetLines", (*token.File).SetLines, execmFileSetLines),
		I.Func("(*File).SetLinesForContent", (*token.File).SetLinesForContent, execmFileSetLinesForContent),
		I.Func("(*File).Size", (*token.File).Size, execmFileSize),
		I.Func("(*FileSet).AddFile", (*token.FileSet).AddFile, execmFileSetAddFile),
		I.Func("(*FileSet).Base", (*token.FileSet).Base, execmFileSetBase),
		I.Func("(*FileSet).File", (*token.FileSet).File, execmFileSetFile),
		I.Func("(*FileSet).Iterate", (*token.FileSet).Iterate, execmFileSetIterate),
		I.Func("(*FileSet).Position", (*token.FileSet).Position, execmFileSetPosition),
		I.Func("(*FileSet).PositionFor", (*token.FileSet).PositionFor, execmFileSetPositionFor),
		I.Func("(*FileSet).Read", (*token.FileSet).Read, execmFileSetRead),
		I.Func("(*FileSet).Write", (*token.FileSet).Write, execmFileSetWrite),
		I.Func("IsExported", token.IsExported, execIsExported),
		I.Func("IsIdentifier", token.IsIdentifier, execIsIdentifier),
		I.Func("IsKeyword", token.IsKeyword, execIsKeyword),
		I.Func("Lookup", token.Lookup, execLookup),
		I.Func("NewFileSet", token.NewFileSet, execNewFileSet),
		I.Func("(Pos).IsValid", (token.Pos).IsValid, execmPosIsValid),
		I.Func("(*Position).IsValid", (*token.Position).IsValid, execmPositionIsValid),
		I.Func("(Position).String", (token.Position).String, execmPositionString),
		I.Func("(Token).IsKeyword", (token.Token).IsKeyword, execmTokenIsKeyword),
		I.Func("(Token).IsLiteral", (token.Token).IsLiteral, execmTokenIsLiteral),
		I.Func("(Token).IsOperator", (token.Token).IsOperator, execmTokenIsOperator),
		I.Func("(Token).Precedence", (token.Token).Precedence, execmTokenPrecedence),
		I.Func("(Token).String", (token.Token).String, execmTokenString),
	)
}
