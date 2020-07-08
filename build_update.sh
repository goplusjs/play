#!/bin/bash
go mod edit -require=github.com/goplus/gop@master
go list --tags js
gopherjs build -v -m