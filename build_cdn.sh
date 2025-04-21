#!/bin/bash
rm docs/*.js
rm docs/*.js.map
rm docs/*.wasm
rm docs/static/playground_*.js
#cp $GOROOT/misc/wasm/wasm_exec.js docs/
cp wasm_exec_rt.js docs/
go run make.go -domain https://play-static.gopluscdn.com