#!/bin/bash
go mod edit -require=github.com/goplus/gop@latest
go list --tags js
gopherjs build -v -m -o goplus-play.js