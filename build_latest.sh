#!/bin/bash
go mod edit -require=github.com/goplus/gop@latest
go mod tidy
gopherjs build -a -v -m -o goplus-play.js
