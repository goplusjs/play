#!/bin/bash
git checkout master
go mod edit -require=github.com/qiniu/goplus@master
go mod tidy
gopherjs build -v -m