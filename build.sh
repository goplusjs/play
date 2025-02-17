#!/bin/bash
rm docs/*.js
rm docs/*.js.map
rm docs/*.wasm
cp $GOROOT/misc/wasm/wasm_exec.js docs/
go run make.go