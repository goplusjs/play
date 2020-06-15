package textproto

import (
	"bufio"
	"io"
	"net/textproto"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func textproto.CanonicalMIMEHeaderKey(s string) string
func execCanonicalMIMEHeaderKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := textproto.CanonicalMIMEHeaderKey(args[0].(string))
	p.Ret(1, ret)
}

// func (*textproto.Conn).Close() error
func execmConnClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*textproto.Conn).Close()
	p.Ret(1, ret)
}

// func (*textproto.Conn).Cmd(format string, args ..interface{}) (id uint, err error)
func execmConnCmd(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*textproto.Conn).Cmd(args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func textproto.Dial(network string, addr string) (*textproto.Conn, error)
func execDial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := textproto.Dial(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*textproto.Error).Error() string
func execmErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*textproto.Error).Error()
	p.Ret(1, ret)
}

// func (textproto.MIMEHeader).Add(key string, value string)
func execmMIMEHeaderAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(textproto.MIMEHeader).Add(args[1].(string), args[2].(string))
}

// func (textproto.MIMEHeader).Del(key string)
func execmMIMEHeaderDel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(textproto.MIMEHeader).Del(args[1].(string))
}

// func (textproto.MIMEHeader).Get(key string) string
func execmMIMEHeaderGet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(textproto.MIMEHeader).Get(args[1].(string))
	p.Ret(2, ret)
}

// func (textproto.MIMEHeader).Set(key string, value string)
func execmMIMEHeaderSet(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(textproto.MIMEHeader).Set(args[1].(string), args[2].(string))
}

// func (textproto.MIMEHeader).Values(key string) []string
func execmMIMEHeaderValues(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(textproto.MIMEHeader).Values(args[1].(string))
	p.Ret(2, ret)
}

// func textproto.NewConn(conn io.ReadWriteCloser) *textproto.Conn
func execNewConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := textproto.NewConn(args[0].(io.ReadWriteCloser))
	p.Ret(1, ret)
}

// func textproto.NewReader(r *bufio.Reader) *textproto.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := textproto.NewReader(args[0].(*bufio.Reader))
	p.Ret(1, ret)
}

// func textproto.NewWriter(w *bufio.Writer) *textproto.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := textproto.NewWriter(args[0].(*bufio.Writer))
	p.Ret(1, ret)
}

// func (*textproto.Pipeline).EndRequest(id uint)
func execmPipelineEndRequest(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*textproto.Pipeline).EndRequest(args[1].(uint))
}

// func (*textproto.Pipeline).EndResponse(id uint)
func execmPipelineEndResponse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*textproto.Pipeline).EndResponse(args[1].(uint))
}

// func (*textproto.Pipeline).Next() uint
func execmPipelineNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*textproto.Pipeline).Next()
	p.Ret(1, ret)
}

// func (*textproto.Pipeline).StartRequest(id uint)
func execmPipelineStartRequest(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*textproto.Pipeline).StartRequest(args[1].(uint))
}

// func (*textproto.Pipeline).StartResponse(id uint)
func execmPipelineStartResponse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*textproto.Pipeline).StartResponse(args[1].(uint))
}

// func (textproto.ProtocolError).Error() string
func execmProtocolErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(textproto.ProtocolError).Error()
	p.Ret(1, ret)
}

// func (*textproto.Reader).DotReader() io.Reader
func execmReaderDotReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*textproto.Reader).DotReader()
	p.Ret(1, ret)
}

// func (*textproto.Reader).ReadCodeLine(expectCode int) (code int, message string, err error)
func execmReaderReadCodeLine(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*textproto.Reader).ReadCodeLine(args[1].(int))
	p.Ret(2, ret, ret1, ret2)
}

// func (*textproto.Reader).ReadContinuedLine() (string, error)
func execmReaderReadContinuedLine(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*textproto.Reader).ReadContinuedLine()
	p.Ret(1, ret, ret1)
}

// func (*textproto.Reader).ReadContinuedLineBytes() ([]byte, error)
func execmReaderReadContinuedLineBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*textproto.Reader).ReadContinuedLineBytes()
	p.Ret(1, ret, ret1)
}

// func (*textproto.Reader).ReadDotBytes() ([]byte, error)
func execmReaderReadDotBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*textproto.Reader).ReadDotBytes()
	p.Ret(1, ret, ret1)
}

// func (*textproto.Reader).ReadDotLines() ([]string, error)
func execmReaderReadDotLines(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*textproto.Reader).ReadDotLines()
	p.Ret(1, ret, ret1)
}

// func (*textproto.Reader).ReadLine() (string, error)
func execmReaderReadLine(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*textproto.Reader).ReadLine()
	p.Ret(1, ret, ret1)
}

// func (*textproto.Reader).ReadLineBytes() ([]byte, error)
func execmReaderReadLineBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*textproto.Reader).ReadLineBytes()
	p.Ret(1, ret, ret1)
}

// func (*textproto.Reader).ReadMIMEHeader() (textproto.MIMEHeader, error)
func execmReaderReadMIMEHeader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*textproto.Reader).ReadMIMEHeader()
	p.Ret(1, ret, ret1)
}

// func (*textproto.Reader).ReadResponse(expectCode int) (code int, message string, err error)
func execmReaderReadResponse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*textproto.Reader).ReadResponse(args[1].(int))
	p.Ret(2, ret, ret1, ret2)
}

// func textproto.TrimBytes(b []byte) []byte
func execTrimBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := textproto.TrimBytes(args[0].([]byte))
	p.Ret(1, ret)
}

// func textproto.TrimString(s string) string
func execTrimString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := textproto.TrimString(args[0].(string))
	p.Ret(1, ret)
}

// func (*textproto.Writer).DotWriter() io.WriteCloser
func execmWriterDotWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*textproto.Writer).DotWriter()
	p.Ret(1, ret)
}

// func (*textproto.Writer).PrintfLine(format string, args ..interface{}) error
func execmWriterPrintfLine(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*textproto.Writer).PrintfLine(args[1].(string), args[2:]...)
	p.Ret(arity, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/textproto")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*textproto.Conn)(nil))),
		I.Rtype(reflect.TypeOf((*textproto.Error)(nil))),
		I.Rtype(reflect.TypeOf((*textproto.Pipeline)(nil))),
		I.Type("ProtocolError", qspec.TyString),
		I.Rtype(reflect.TypeOf((*textproto.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*textproto.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("CanonicalMIMEHeaderKey", textproto.CanonicalMIMEHeaderKey, execCanonicalMIMEHeaderKey),
		I.Func("(*Conn).Close", (*textproto.Conn).Close, execmConnClose),
		I.Func("Dial", textproto.Dial, execDial),
		I.Func("(*Error).Error", (*textproto.Error).Error, execmErrorError),
		I.Func("(MIMEHeader).Add", (textproto.MIMEHeader).Add, execmMIMEHeaderAdd),
		I.Func("(MIMEHeader).Del", (textproto.MIMEHeader).Del, execmMIMEHeaderDel),
		I.Func("(MIMEHeader).Get", (textproto.MIMEHeader).Get, execmMIMEHeaderGet),
		I.Func("(MIMEHeader).Set", (textproto.MIMEHeader).Set, execmMIMEHeaderSet),
		I.Func("(MIMEHeader).Values", (textproto.MIMEHeader).Values, execmMIMEHeaderValues),
		I.Func("NewConn", textproto.NewConn, execNewConn),
		I.Func("NewReader", textproto.NewReader, execNewReader),
		I.Func("NewWriter", textproto.NewWriter, execNewWriter),
		I.Func("(*Pipeline).EndRequest", (*textproto.Pipeline).EndRequest, execmPipelineEndRequest),
		I.Func("(*Pipeline).EndResponse", (*textproto.Pipeline).EndResponse, execmPipelineEndResponse),
		I.Func("(*Pipeline).Next", (*textproto.Pipeline).Next, execmPipelineNext),
		I.Func("(*Pipeline).StartRequest", (*textproto.Pipeline).StartRequest, execmPipelineStartRequest),
		I.Func("(*Pipeline).StartResponse", (*textproto.Pipeline).StartResponse, execmPipelineStartResponse),
		I.Func("(ProtocolError).Error", (textproto.ProtocolError).Error, execmProtocolErrorError),
		I.Func("(*Reader).DotReader", (*textproto.Reader).DotReader, execmReaderDotReader),
		I.Func("(*Reader).ReadCodeLine", (*textproto.Reader).ReadCodeLine, execmReaderReadCodeLine),
		I.Func("(*Reader).ReadContinuedLine", (*textproto.Reader).ReadContinuedLine, execmReaderReadContinuedLine),
		I.Func("(*Reader).ReadContinuedLineBytes", (*textproto.Reader).ReadContinuedLineBytes, execmReaderReadContinuedLineBytes),
		I.Func("(*Reader).ReadDotBytes", (*textproto.Reader).ReadDotBytes, execmReaderReadDotBytes),
		I.Func("(*Reader).ReadDotLines", (*textproto.Reader).ReadDotLines, execmReaderReadDotLines),
		I.Func("(*Reader).ReadLine", (*textproto.Reader).ReadLine, execmReaderReadLine),
		I.Func("(*Reader).ReadLineBytes", (*textproto.Reader).ReadLineBytes, execmReaderReadLineBytes),
		I.Func("(*Reader).ReadMIMEHeader", (*textproto.Reader).ReadMIMEHeader, execmReaderReadMIMEHeader),
		I.Func("(*Reader).ReadResponse", (*textproto.Reader).ReadResponse, execmReaderReadResponse),
		I.Func("TrimBytes", textproto.TrimBytes, execTrimBytes),
		I.Func("TrimString", textproto.TrimString, execTrimString),
		I.Func("(*Writer).DotWriter", (*textproto.Writer).DotWriter, execmWriterDotWriter),
	)
	I.RegisterFuncvs(
		I.Funcv("(*Conn).Cmd", (*textproto.Conn).Cmd, execmConnCmd),
		I.Funcv("(*Writer).PrintfLine", (*textproto.Writer).PrintfLine, execmWriterPrintfLine),
	)
}
