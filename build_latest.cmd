@echo off
go mod edit -require=github.com/goplus/gop@latest
go list --tags="js repl"
set GOOS=darwin
gopherjs build -v -m -o goplus-play.js
gopherjs build -v -m --tags repl -o repl.js repl.go