#!/bin/bash
go mod edit -require=github.com/goplus/gop@master
go list --tags="js repl"
gopherjs build -v -m -o goplus-play.js
gopherjs build -v -m --tags repl -o repl.js repl.go