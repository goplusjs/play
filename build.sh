#!/bin/bash
#rm docs/*.js
#rm docs/*.js.map
#rm docs/*.wasm
#rm docs/static/playground_*.js
WASM_EXEC="$(go env GOROOT)/lib/wasm/wasm_exec.js"
if [ ! -f "$WASM_EXEC" ]; then
  WASM_EXEC="$(go env GOROOT)/misc/wasm/wasm_exec.js"
fi
cp "$WASM_EXEC" docs/wasm_exec.js
go run make.go -buildonly=true