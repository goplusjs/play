package syslog

import (
	"log/syslog"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func syslog.Dial(network string, raddr string, priority syslog.Priority, tag string) (*syslog.Writer, error)
func execDial(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := syslog.Dial(args[0].(string), args[1].(string), syslog.Priority(args[2].(int)), args[3].(string))
	p.Ret(4, ret, ret1)
}

// func syslog.New(priority syslog.Priority, tag string) (*syslog.Writer, error)
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syslog.New(syslog.Priority(args[0].(int)), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func syslog.NewLogger(p syslog.Priority, logFlag int) (*log.Logger, error)
func execNewLogger(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syslog.NewLogger(syslog.Priority(args[0].(int)), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func (*syslog.Writer).Alert(m string) error
func execmWriterAlert(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Alert(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Close() error
func execmWriterClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syslog.Writer).Close()
	p.Ret(1, ret)
}

// func (*syslog.Writer).Crit(m string) error
func execmWriterCrit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Crit(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Debug(m string) error
func execmWriterDebug(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Debug(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Emerg(m string) error
func execmWriterEmerg(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Emerg(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Err(m string) error
func execmWriterErr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Err(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Info(m string) error
func execmWriterInfo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Info(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Notice(m string) error
func execmWriterNotice(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Notice(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Warning(m string) error
func execmWriterWarning(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*syslog.Writer).Warning(args[1].(string))
	p.Ret(2, ret)
}

// func (*syslog.Writer).Write(b []byte) (int, error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*syslog.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("log/syslog")

func init() {
	I.RegisterConsts(
		I.Const("LOG_ALERT", reflect.Int, syslog.LOG_ALERT),
		I.Const("LOG_AUTH", reflect.Int, syslog.LOG_AUTH),
		I.Const("LOG_AUTHPRIV", reflect.Int, syslog.LOG_AUTHPRIV),
		I.Const("LOG_CRIT", reflect.Int, syslog.LOG_CRIT),
		I.Const("LOG_CRON", reflect.Int, syslog.LOG_CRON),
		I.Const("LOG_DAEMON", reflect.Int, syslog.LOG_DAEMON),
		I.Const("LOG_DEBUG", reflect.Int, syslog.LOG_DEBUG),
		I.Const("LOG_EMERG", reflect.Int, syslog.LOG_EMERG),
		I.Const("LOG_ERR", reflect.Int, syslog.LOG_ERR),
		I.Const("LOG_FTP", reflect.Int, syslog.LOG_FTP),
		I.Const("LOG_INFO", reflect.Int, syslog.LOG_INFO),
		I.Const("LOG_KERN", reflect.Int, syslog.LOG_KERN),
		I.Const("LOG_LOCAL0", reflect.Int, syslog.LOG_LOCAL0),
		I.Const("LOG_LOCAL1", reflect.Int, syslog.LOG_LOCAL1),
		I.Const("LOG_LOCAL2", reflect.Int, syslog.LOG_LOCAL2),
		I.Const("LOG_LOCAL3", reflect.Int, syslog.LOG_LOCAL3),
		I.Const("LOG_LOCAL4", reflect.Int, syslog.LOG_LOCAL4),
		I.Const("LOG_LOCAL5", reflect.Int, syslog.LOG_LOCAL5),
		I.Const("LOG_LOCAL6", reflect.Int, syslog.LOG_LOCAL6),
		I.Const("LOG_LOCAL7", reflect.Int, syslog.LOG_LOCAL7),
		I.Const("LOG_LPR", reflect.Int, syslog.LOG_LPR),
		I.Const("LOG_MAIL", reflect.Int, syslog.LOG_MAIL),
		I.Const("LOG_NEWS", reflect.Int, syslog.LOG_NEWS),
		I.Const("LOG_NOTICE", reflect.Int, syslog.LOG_NOTICE),
		I.Const("LOG_SYSLOG", reflect.Int, syslog.LOG_SYSLOG),
		I.Const("LOG_USER", reflect.Int, syslog.LOG_USER),
		I.Const("LOG_UUCP", reflect.Int, syslog.LOG_UUCP),
		I.Const("LOG_WARNING", reflect.Int, syslog.LOG_WARNING),
	)
	I.RegisterTypes(
		I.Type("Priority", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*syslog.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Dial", syslog.Dial, execDial),
		I.Func("New", syslog.New, execNew),
		I.Func("NewLogger", syslog.NewLogger, execNewLogger),
		I.Func("(*Writer).Alert", (*syslog.Writer).Alert, execmWriterAlert),
		I.Func("(*Writer).Close", (*syslog.Writer).Close, execmWriterClose),
		I.Func("(*Writer).Crit", (*syslog.Writer).Crit, execmWriterCrit),
		I.Func("(*Writer).Debug", (*syslog.Writer).Debug, execmWriterDebug),
		I.Func("(*Writer).Emerg", (*syslog.Writer).Emerg, execmWriterEmerg),
		I.Func("(*Writer).Err", (*syslog.Writer).Err, execmWriterErr),
		I.Func("(*Writer).Info", (*syslog.Writer).Info, execmWriterInfo),
		I.Func("(*Writer).Notice", (*syslog.Writer).Notice, execmWriterNotice),
		I.Func("(*Writer).Warning", (*syslog.Writer).Warning, execmWriterWarning),
		I.Func("(*Writer).Write", (*syslog.Writer).Write, execmWriterWrite),
	)
}
