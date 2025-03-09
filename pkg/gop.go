package pkg

//go:generate qexp -outdir . github.com/goplus/gop/ast
//go:generate qexp -outdir . github.com/goplus/gop/parser
//go:generate qexp -outdir . github.com/goplus/gop/printer
//go:generate qexp -outdir . github.com/goplus/gop/scanner
//go:generate qexp -outdir . github.com/goplus/gop/token
//go:generate qexp -outdir . github.com/goplus/gop/tpl/...

import (
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/ast"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/parser"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/printer"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/scanner"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/token"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/ast"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/cl"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/encoding/csv"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/encoding/json"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/encoding/regexp"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/encoding/regexposix"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/encoding/xml"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/matcher"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/parser"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/token"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/tpl/types"
)
