@echo off
go mod edit -require=github.com/goplus/gop@latest
go list --tags js
set GOOS=darwin
gopherjs build -v -m -o goplus-play.js