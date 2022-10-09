module github.com/goplusjs/play

go 1.16

require (
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/goplus/gop v1.1.3
	github.com/goplus/igop v0.9.2
	github.com/goplus/reflectx v0.9.6
)

replace (
	// github.com/goplus/reflectx => github.com/goplusjs/reflectx v0.5.6
	github.com/goplus/reflectx => ../demo/reflectx
	github.com/petermattis/goid => github.com/visualfc/goid v0.1.1
	github.com/visualfc/xtype => github.com/visualfc/xtype_js v0.1.1
// golang.org/x/mod => golang.org/x/mod v0.5.1
// golang.org/x/tools => golang.org/x/tools v0.1.12
// golang.org/x/xerrors => golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898
)
