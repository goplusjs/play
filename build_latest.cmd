@echo off
go mod edit -require=github.com/goplus/gop@latest
go tidy
gopherjs build -v -m -o goplus-play.js
