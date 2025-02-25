package pkg

//go:generate qexp -outdir . github.com/goplus/gop/ast
//go:generate qexp -outdir . github.com/goplus/gop/parser
//go:generate qexp -outdir . github.com/goplus/gop/printer
//go:generate qexp -outdir . github.com/goplus/gop/scanner
//go:generate qexp -outdir . github.com/goplus/gop/token

import (
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/ast"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/parser"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/printer"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/scanner"
	_ "github.com/goplusjs/play/pkg/github.com/goplus/gop/token"
)
