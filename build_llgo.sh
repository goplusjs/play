#!/bin/bash
rm docs/*.js
rm docs/*.wasm
rm docs/static/playground_*.js
rm docs/*.json
go run make.go -llgo
rm docs/*.json

