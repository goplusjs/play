#!/bin/bash
rm docs/*.js
rm docs/*.wasm
rm docs/static/playground_*.js
go run make.go -llgo