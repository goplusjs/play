package template

import (
	"io"
	"reflect"
	"text/template"
	"text/template/parse"

	"github.com/qiniu/goplus/gop"
)

// func (template.ExecError).Error() string
func execmExecErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(template.ExecError).Error()
	p.Ret(1, ret)
}

// func (template.ExecError).Unwrap() error
func execmExecErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(template.ExecError).Unwrap()
	p.Ret(1, ret)
}

// func template.HTMLEscape(w io.Writer, b []byte)
func execHTMLEscape(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	template.HTMLEscape(args[0].(io.Writer), args[1].([]byte))
}

// func template.HTMLEscapeString(s string) string
func execHTMLEscapeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := template.HTMLEscapeString(args[0].(string))
	p.Ret(1, ret)
}

// func template.HTMLEscaper(args ..interface{}) string
func execHTMLEscaper(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := template.HTMLEscaper(args[0:]...)
	p.Ret(arity, ret)
}

// func template.IsTrue(val interface{}) (truth bool, ok bool)
func execIsTrue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := template.IsTrue(args[0].(interface{}))
	p.Ret(1, ret, ret1)
}

// func template.JSEscape(w io.Writer, b []byte)
func execJSEscape(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	template.JSEscape(args[0].(io.Writer), args[1].([]byte))
}

// func template.JSEscapeString(s string) string
func execJSEscapeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := template.JSEscapeString(args[0].(string))
	p.Ret(1, ret)
}

// func template.JSEscaper(args ..interface{}) string
func execJSEscaper(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := template.JSEscaper(args[0:]...)
	p.Ret(arity, ret)
}

// func template.Must(t *template.Template, err error) *template.Template
func execMust(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := template.Must(args[0].(*template.Template), args[1].(error))
	p.Ret(2, ret)
}

// func template.New(name string) *template.Template
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := template.New(args[0].(string))
	p.Ret(1, ret)
}

// func template.ParseFiles(filenames ..string) (*template.Template, error)
func execParseFiles(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret, ret1 := template.ParseFiles(conv(args[0:])...)
	p.Ret(arity, ret, ret1)
}

// func template.ParseGlob(pattern string) (*template.Template, error)
func execParseGlob(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := template.ParseGlob(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*template.Template).AddParseTree(name string, tree *parse.Tree) (*template.Template, error)
func execmTemplateAddParseTree(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*template.Template).AddParseTree(args[1].(string), args[2].(*parse.Tree))
	p.Ret(3, ret, ret1)
}

// func (*template.Template).Clone() (*template.Template, error)
func execmTemplateClone(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*template.Template).Clone()
	p.Ret(1, ret, ret1)
}

// func (*template.Template).DefinedTemplates() string
func execmTemplateDefinedTemplates(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*template.Template).DefinedTemplates()
	p.Ret(1, ret)
}

// func (*template.Template).Delims(left string, right string) *template.Template
func execmTemplateDelims(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*template.Template).Delims(args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// func (*template.Template).Execute(wr io.Writer, data interface{}) error
func execmTemplateExecute(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*template.Template).Execute(args[1].(io.Writer), args[2].(interface{}))
	p.Ret(3, ret)
}

// func (*template.Template).ExecuteTemplate(wr io.Writer, name string, data interface{}) error
func execmTemplateExecuteTemplate(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*template.Template).ExecuteTemplate(args[1].(io.Writer), args[2].(string), args[3].(interface{}))
	p.Ret(4, ret)
}

// func (*template.Template).Funcs(funcMap template.FuncMap) *template.Template
func execmTemplateFuncs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*template.Template).Funcs(args[1].(template.FuncMap))
	p.Ret(2, ret)
}

// func (*template.Template).Lookup(name string) *template.Template
func execmTemplateLookup(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*template.Template).Lookup(args[1].(string))
	p.Ret(2, ret)
}

// func (*template.Template).Name() string
func execmTemplateName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*template.Template).Name()
	p.Ret(1, ret)
}

// func (*template.Template).New(name string) *template.Template
func execmTemplateNew(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*template.Template).New(args[1].(string))
	p.Ret(2, ret)
}

// func (*template.Template).Option(opt ..string) *template.Template
func execmTemplateOption(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret := args[0].(*template.Template).Option(conv(args[1:])...)
	p.Ret(arity, ret)
}

// func (*template.Template).Parse(text string) (*template.Template, error)
func execmTemplateParse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*template.Template).Parse(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*template.Template).ParseFiles(filenames ..string) (*template.Template, error)
func execmTemplateParseFiles(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret, ret1 := args[0].(*template.Template).ParseFiles(conv(args[1:])...)
	p.Ret(arity, ret, ret1)
}

// func (*template.Template).ParseGlob(pattern string) (*template.Template, error)
func execmTemplateParseGlob(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*template.Template).ParseGlob(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*template.Template).Templates() []*template.Template
func execmTemplateTemplates(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*template.Template).Templates()
	p.Ret(1, ret)
}

// func template.URLQueryEscaper(args ..interface{}) string
func execURLQueryEscaper(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := template.URLQueryEscaper(args[0:]...)
	p.Ret(arity, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("text/template")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*template.ExecError)(nil))),
		I.Rtype(reflect.TypeOf((*template.Template)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(ExecError).Error", (template.ExecError).Error, execmExecErrorError),
		I.Func("(ExecError).Unwrap", (template.ExecError).Unwrap, execmExecErrorUnwrap),
		I.Func("HTMLEscape", template.HTMLEscape, execHTMLEscape),
		I.Func("HTMLEscapeString", template.HTMLEscapeString, execHTMLEscapeString),
		I.Func("IsTrue", template.IsTrue, execIsTrue),
		I.Func("JSEscape", template.JSEscape, execJSEscape),
		I.Func("JSEscapeString", template.JSEscapeString, execJSEscapeString),
		I.Func("Must", template.Must, execMust),
		I.Func("New", template.New, execNew),
		I.Func("ParseGlob", template.ParseGlob, execParseGlob),
		I.Func("(*Template).AddParseTree", (*template.Template).AddParseTree, execmTemplateAddParseTree),
		I.Func("(*Template).Clone", (*template.Template).Clone, execmTemplateClone),
		I.Func("(*Template).DefinedTemplates", (*template.Template).DefinedTemplates, execmTemplateDefinedTemplates),
		I.Func("(*Template).Delims", (*template.Template).Delims, execmTemplateDelims),
		I.Func("(*Template).Execute", (*template.Template).Execute, execmTemplateExecute),
		I.Func("(*Template).ExecuteTemplate", (*template.Template).ExecuteTemplate, execmTemplateExecuteTemplate),
		I.Func("(*Template).Funcs", (*template.Template).Funcs, execmTemplateFuncs),
		I.Func("(*Template).Lookup", (*template.Template).Lookup, execmTemplateLookup),
		I.Func("(*Template).Name", (*template.Template).Name, execmTemplateName),
		I.Func("(*Template).New", (*template.Template).New, execmTemplateNew),
		I.Func("(*Template).Parse", (*template.Template).Parse, execmTemplateParse),
		I.Func("(*Template).ParseGlob", (*template.Template).ParseGlob, execmTemplateParseGlob),
		I.Func("(*Template).Templates", (*template.Template).Templates, execmTemplateTemplates),
	)
	I.RegisterFuncvs(
		I.Funcv("HTMLEscaper", template.HTMLEscaper, execHTMLEscaper),
		I.Funcv("JSEscaper", template.JSEscaper, execJSEscaper),
		I.Funcv("ParseFiles", template.ParseFiles, execParseFiles),
		I.Funcv("(*Template).Option", (*template.Template).Option, execmTemplateOption),
		I.Funcv("(*Template).ParseFiles", (*template.Template).ParseFiles, execmTemplateParseFiles),
		I.Funcv("URLQueryEscaper", template.URLQueryEscaper, execURLQueryEscaper),
	)
}
