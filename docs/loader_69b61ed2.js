// Set the default global log for use by wasm_exec.js
go_log = console.log;

if (!window.__IXGO_LOADER_STARTED__) {
window.__IXGO_LOADER_STARTED__ = true;

var useWasm = true; //location.href.includes("?wasm");

let isWasmLoaded = false;
let currentGoInstance = null;

let wasmOverflowCallback;

// Track wasm_exec.js loading state.
let wasmExecRtLoaded = false;
let wasmExecRtLoadPromise = null;

function assetURL(path) {
  var base = window.__ASSET_BASE__ || ".";
  if (path.charAt(0) !== "/") {
    path = "/" + path;
  }
  return base + path;
}

async function fetchAsset(path) {
  var primary = assetURL(path);
  try {
    var resp = await fetch(primary);
    if (resp.ok) {
      return resp;
    }
  } catch (e) {}
  if ((window.__ASSET_BASE__ || ".") !== ".") {
    return fetch("./" + path);
  }
  return fetch(primary);
}

window.setIxgoOverflowCallback = function (callback) {
  wasmOverflowCallback = callback;
};

window.isIxgoLoaded = function () {
  return isWasmLoaded;
};

window.goWriteSync = function (text) {
  console.log(text);
};

// Install fs before wasm_exec.js so Go and LLGo both route output here.
(() => {
  const decoder = new TextDecoder("utf-8");
  const enosys = () => {
    const err = new Error("not implemented");
    err.code = "ENOSYS";
    return err;
  };
  if (!globalThis.fs) {
    globalThis.fs = {
      constants: { O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1, O_DIRECTORY: -1 },
      writeSync(fd, buf) {
        const text = decoder.decode(buf);
        if (typeof globalThis.goWriteSync === "function") {
          globalThis.goWriteSync(text);
        } else {
          console.log(text);
        }
        return buf.length;
      },
      write(fd, buf, offset, length, position, callback) {
        if (offset !== 0 || length !== buf.length || position !== null) {
          callback(enosys());
          return;
        }
        callback(null, this.writeSync(fd, buf));
      },
      fsync(fd, callback) { callback(null); },
    };
    for (const name of ["chmod", "chown", "close", "fchmod", "fchown", "fstat", "ftruncate", "lchown", "link", "lstat", "mkdir", "open", "read", "readdir", "readlink", "rename", "rmdir", "stat", "symlink", "truncate", "unlink", "utimes"]) {
      globalThis.fs[name] = (...args) => args[args.length - 1](enosys());
    }
  }
})();

function loadRuntimeScript(src, fallbackSrc) {
  return new Promise(function (resolve, reject) {
    var script = document.createElement("script");
    script.onload = resolve;
    script.onerror = function () {
      if (!fallbackSrc || fallbackSrc === src) {
        reject(new Error("Failed to load wasm_exec.js"));
        return;
      }
      var retry = document.createElement("script");
      retry.onload = resolve;
      retry.onerror = function () {
        reject(new Error("Failed to load wasm_exec.js"));
      };
      retry.src = fallbackSrc;
      document.head.appendChild(retry);
    };
    script.src = src;
    document.head.appendChild(script);
  });
}

if (useWasm) {
  wasmExecRtLoadPromise = loadRuntimeScript(
    assetURL("wasm_exec.js"),
    "./wasm_exec.js",
  ).then(function () {
    // polyfill
    if (!WebAssembly.instantiateStreaming) {
      WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
      };
    }
    wasmExecRtLoaded = true;
    loadWasm();
  });
} else {
  var script = document.createElement("script");
  script.src = "ixgo_69b61ed2.js";
  document.head.appendChild(script);
}

function handleGlobalError(event) {
  if (
    event.error instanceof RangeError &&
    event.message.includes("Maximum call stack size exceeded")
  ) {
    event.preventDefault();
    console.error("Stack overflow detected, reload WASM module...");
    if (typeof wasmOverflowCallback === "function") {
      wasmOverflowCallback(event);
    }
    if (isWasmLoaded) {
      reloadWasm();
    }
  }
}

window.addEventListener("error", handleGlobalError);

async function loadWasm() {
  // Ensure wasm_exec.js is loaded before proceeding.
  if (!wasmExecRtLoaded) {
      await wasmExecRtLoadPromise;
  }

  const go = new Go();
  currentGoInstance = go;
  let mod, inst;
  WebAssembly.instantiateStreaming(fetchAsset("ixgo_69b61ed2.wasm"), go.importObject).then(
    (result) => {
      mod = result.module;
      inst = result.instance;
      isWasmLoaded = true;
      console.log("Load WASM module.");
      run();
    },
  );

  async function run() {
    await go.run(inst);
    inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
  }
}

async function reloadWasm() {
  isWasmLoaded = false;

  if (currentGoInstance) {
    currentGoInstance.exit(0);
    currentGoInstance = null;
  }

  await loadWasm();
}

}
