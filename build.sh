#!/bin/bash
rm docs/*.js
rm docs/*.js.map
rm docs/*.wasm
#cp $GOROOT/misc/wasm/wasm_exec.js docs/
cp wasm_exec_rt.js docs/
go run make.go