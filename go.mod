module github.com/goplus/play

go 1.14

require (
	github.com/goplus/gop v0.7.19
	github.com/goplusjs/gopherjs v1.1.2
)

replace (
	github.com/goplus/gop => github.com/visualfc/goplus v0.0.0-20210609034246-546361c19112
	github.com/goplus/reflectx => github.com/goplusjs/reflectx v0.5.2-goplus
)
