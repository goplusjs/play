# xgo-play

Playground of the XGo language.

- https://play.xgo.dev/
- https://goplusjs.github.io/play/

## Build

Go 1.26+ with wasm support.

```
git clone https://github.com/goplusjs/play
cd play
./build.sh
```

Other toolchains:

```
./build_llgo.sh         # LLGo
./build_ifacepatch.sh   # GOFLAGS='-tags=goplus.ifacefuncval'
```

Output is written to `docs/`. Asset names include a hash of the sources and the toolchain (`go`, `llgo`, or `GOFLAGS`), so the three builds do not overwrite each other. An LLGo build also shows `LLGo` in the About version line.

## Publish

Push `docs/` once. The page loads scripts and wasm from `https://play-static.gopluscdn.com` first, then falls back to the GitHub Pages copy until the CDN catches up.

Local preview (`localhost`, `127.0.0.1`) and `?local` skip the CDN probe and use files from `docs/` directly.
