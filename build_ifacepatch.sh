#!/bin/bash
cd "$(dirname "$0")"
GOFLAGS='-tags=goplus.ifacefuncval' exec ./build.sh
