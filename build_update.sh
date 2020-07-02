#!/bin/bash
go mod edit -require=github.com/qiniu/goplus@master
go list --tags js
gopherjs build -v -m