package bufio

import (
	"bufio"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func bufio.NewReadWriter(r *bufio.Reader, w *bufio.Writer) *bufio.ReadWriter
func execNewReadWriter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bufio.NewReadWriter(args[0].(*bufio.Reader), args[1].(*bufio.Writer))
	p.Ret(2, ret)
}

// func bufio.NewReader(rd io.Reader) *bufio.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bufio.NewReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func bufio.NewReaderSize(rd io.Reader, size int) *bufio.Reader
func execNewReaderSize(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bufio.NewReaderSize(args[0].(io.Reader), args[1].(int))
	p.Ret(2, ret)
}

// func bufio.NewScanner(r io.Reader) *bufio.Scanner
func execNewScanner(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bufio.NewScanner(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func bufio.NewWriter(w io.Writer) *bufio.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bufio.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func bufio.NewWriterSize(w io.Writer, size int) *bufio.Writer
func execNewWriterSize(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bufio.NewWriterSize(args[0].(io.Writer), args[1].(int))
	p.Ret(2, ret)
}

// func (*bufio.Reader).Buffered() int
func execmReaderBuffered(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Reader).Buffered()
	p.Ret(1, ret)
}

// func (*bufio.Reader).Discard(n int) (discarded int, err error)
func execmReaderDiscard(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Reader).Discard(args[1].(int))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Reader).Peek(n int) ([]byte, error)
func execmReaderPeek(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Reader).Peek(args[1].(int))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Reader).Read(p []byte) (n int, err error)
func execmReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Reader).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Reader).ReadByte() (byte, error)
func execmReaderReadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*bufio.Reader).ReadByte()
	p.Ret(1, ret, ret1)
}

// func (*bufio.Reader).ReadBytes(delim byte) ([]byte, error)
func execmReaderReadBytes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Reader).ReadBytes(args[1].(byte))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Reader).ReadLine() (line []byte, isPrefix bool, err error)
func execmReaderReadLine(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*bufio.Reader).ReadLine()
	p.Ret(1, ret, ret1, ret2)
}

// func (*bufio.Reader).ReadRune() (r rune, size int, err error)
func execmReaderReadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*bufio.Reader).ReadRune()
	p.Ret(1, ret, ret1, ret2)
}

// func (*bufio.Reader).ReadSlice(delim byte) (line []byte, err error)
func execmReaderReadSlice(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Reader).ReadSlice(args[1].(byte))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Reader).ReadString(delim byte) (string, error)
func execmReaderReadString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Reader).ReadString(args[1].(byte))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Reader).Reset(r io.Reader)
func execmReaderReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*bufio.Reader).Reset(args[1].(io.Reader))
}

// func (*bufio.Reader).Size() int
func execmReaderSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Reader).Size()
	p.Ret(1, ret)
}

// func (*bufio.Reader).UnreadByte() error
func execmReaderUnreadByte(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Reader).UnreadByte()
	p.Ret(1, ret)
}

// func (*bufio.Reader).UnreadRune() error
func execmReaderUnreadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Reader).UnreadRune()
	p.Ret(1, ret)
}

// func (*bufio.Reader).WriteTo(w io.Writer) (n int64, err error)
func execmReaderWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Reader).WriteTo(args[1].(io.Writer))
	p.Ret(2, ret, ret1)
}

// func bufio.ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
func execScanBytes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := bufio.ScanBytes(args[0].([]byte), args[1].(bool))
	p.Ret(2, ret, ret1, ret2)
}

// func bufio.ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
func execScanLines(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := bufio.ScanLines(args[0].([]byte), args[1].(bool))
	p.Ret(2, ret, ret1, ret2)
}

// func bufio.ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
func execScanRunes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := bufio.ScanRunes(args[0].([]byte), args[1].(bool))
	p.Ret(2, ret, ret1, ret2)
}

// func bufio.ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
func execScanWords(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := bufio.ScanWords(args[0].([]byte), args[1].(bool))
	p.Ret(2, ret, ret1, ret2)
}

// func (*bufio.Scanner).Buffer(buf []byte, max int)
func execmScannerBuffer(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*bufio.Scanner).Buffer(args[1].([]byte), args[2].(int))
}

// func (*bufio.Scanner).Bytes() []byte
func execmScannerBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Scanner).Bytes()
	p.Ret(1, ret)
}

// func (*bufio.Scanner).Err() error
func execmScannerErr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Scanner).Err()
	p.Ret(1, ret)
}

// func (*bufio.Scanner).Scan() bool
func execmScannerScan(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Scanner).Scan()
	p.Ret(1, ret)
}

// func (*bufio.Scanner).Split(split bufio.SplitFunc)
func execmScannerSplit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*bufio.Scanner).Split(args[1].(bufio.SplitFunc))
}

// func (*bufio.Scanner).Text() string
func execmScannerText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Scanner).Text()
	p.Ret(1, ret)
}

// func (*bufio.Writer).Available() int
func execmWriterAvailable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Writer).Available()
	p.Ret(1, ret)
}

// func (*bufio.Writer).Buffered() int
func execmWriterBuffered(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Writer).Buffered()
	p.Ret(1, ret)
}

// func (*bufio.Writer).Flush() error
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Writer).Flush()
	p.Ret(1, ret)
}

// func (*bufio.Writer).ReadFrom(r io.Reader) (n int64, err error)
func execmWriterReadFrom(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Writer).ReadFrom(args[1].(io.Reader))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Writer).Reset(w io.Writer)
func execmWriterReset(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*bufio.Writer).Reset(args[1].(io.Writer))
}

// func (*bufio.Writer).Size() int
func execmWriterSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*bufio.Writer).Size()
	p.Ret(1, ret)
}

// func (*bufio.Writer).Write(p []byte) (nn int, err error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Writer).WriteByte(c byte) error
func execmWriterWriteByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*bufio.Writer).WriteByte(args[1].(byte))
	p.Ret(2, ret)
}

// func (*bufio.Writer).WriteRune(r rune) (size int, err error)
func execmWriterWriteRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Writer).WriteRune(args[1].(rune))
	p.Ret(2, ret, ret1)
}

// func (*bufio.Writer).WriteString(s string) (int, error)
func execmWriterWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*bufio.Writer).WriteString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("bufio")

func init() {
	I.RegisterConsts(
		I.Const("MaxScanTokenSize", qspec.ConstUnboundInt, bufio.MaxScanTokenSize),
	)
	I.RegisterVars(
		I.Var("ErrAdvanceTooFar", &bufio.ErrAdvanceTooFar),
		I.Var("ErrBufferFull", &bufio.ErrBufferFull),
		I.Var("ErrFinalToken", &bufio.ErrFinalToken),
		I.Var("ErrInvalidUnreadByte", &bufio.ErrInvalidUnreadByte),
		I.Var("ErrInvalidUnreadRune", &bufio.ErrInvalidUnreadRune),
		I.Var("ErrNegativeAdvance", &bufio.ErrNegativeAdvance),
		I.Var("ErrNegativeCount", &bufio.ErrNegativeCount),
		I.Var("ErrTooLong", &bufio.ErrTooLong),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*bufio.ReadWriter)(nil))),
		I.Rtype(reflect.TypeOf((*bufio.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*bufio.Scanner)(nil))),
		I.Rtype(reflect.TypeOf((*bufio.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewReadWriter", bufio.NewReadWriter, execNewReadWriter),
		I.Func("NewReader", bufio.NewReader, execNewReader),
		I.Func("NewReaderSize", bufio.NewReaderSize, execNewReaderSize),
		I.Func("NewScanner", bufio.NewScanner, execNewScanner),
		I.Func("NewWriter", bufio.NewWriter, execNewWriter),
		I.Func("NewWriterSize", bufio.NewWriterSize, execNewWriterSize),
		I.Func("(*Reader).Buffered", (*bufio.Reader).Buffered, execmReaderBuffered),
		I.Func("(*Reader).Discard", (*bufio.Reader).Discard, execmReaderDiscard),
		I.Func("(*Reader).Peek", (*bufio.Reader).Peek, execmReaderPeek),
		I.Func("(*Reader).Read", (*bufio.Reader).Read, execmReaderRead),
		I.Func("(*Reader).ReadByte", (*bufio.Reader).ReadByte, execmReaderReadByte),
		I.Func("(*Reader).ReadBytes", (*bufio.Reader).ReadBytes, execmReaderReadBytes),
		I.Func("(*Reader).ReadLine", (*bufio.Reader).ReadLine, execmReaderReadLine),
		I.Func("(*Reader).ReadRune", (*bufio.Reader).ReadRune, execmReaderReadRune),
		I.Func("(*Reader).ReadSlice", (*bufio.Reader).ReadSlice, execmReaderReadSlice),
		I.Func("(*Reader).ReadString", (*bufio.Reader).ReadString, execmReaderReadString),
		I.Func("(*Reader).Reset", (*bufio.Reader).Reset, execmReaderReset),
		I.Func("(*Reader).Size", (*bufio.Reader).Size, execmReaderSize),
		I.Func("(*Reader).UnreadByte", (*bufio.Reader).UnreadByte, execmReaderUnreadByte),
		I.Func("(*Reader).UnreadRune", (*bufio.Reader).UnreadRune, execmReaderUnreadRune),
		I.Func("(*Reader).WriteTo", (*bufio.Reader).WriteTo, execmReaderWriteTo),
		I.Func("ScanBytes", bufio.ScanBytes, execScanBytes),
		I.Func("ScanLines", bufio.ScanLines, execScanLines),
		I.Func("ScanRunes", bufio.ScanRunes, execScanRunes),
		I.Func("ScanWords", bufio.ScanWords, execScanWords),
		I.Func("(*Scanner).Buffer", (*bufio.Scanner).Buffer, execmScannerBuffer),
		I.Func("(*Scanner).Bytes", (*bufio.Scanner).Bytes, execmScannerBytes),
		I.Func("(*Scanner).Err", (*bufio.Scanner).Err, execmScannerErr),
		I.Func("(*Scanner).Scan", (*bufio.Scanner).Scan, execmScannerScan),
		I.Func("(*Scanner).Split", (*bufio.Scanner).Split, execmScannerSplit),
		I.Func("(*Scanner).Text", (*bufio.Scanner).Text, execmScannerText),
		I.Func("(*Writer).Available", (*bufio.Writer).Available, execmWriterAvailable),
		I.Func("(*Writer).Buffered", (*bufio.Writer).Buffered, execmWriterBuffered),
		I.Func("(*Writer).Flush", (*bufio.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).ReadFrom", (*bufio.Writer).ReadFrom, execmWriterReadFrom),
		I.Func("(*Writer).Reset", (*bufio.Writer).Reset, execmWriterReset),
		I.Func("(*Writer).Size", (*bufio.Writer).Size, execmWriterSize),
		I.Func("(*Writer).Write", (*bufio.Writer).Write, execmWriterWrite),
		I.Func("(*Writer).WriteByte", (*bufio.Writer).WriteByte, execmWriterWriteByte),
		I.Func("(*Writer).WriteRune", (*bufio.Writer).WriteRune, execmWriterWriteRune),
		I.Func("(*Writer).WriteString", (*bufio.Writer).WriteString, execmWriterWriteString),
	)
}
