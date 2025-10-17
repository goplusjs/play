// Set the default global log for use by wasm_exec.js
go_log = console.log;

var useWasm = true; //location.href.includes("?wasm");

let isWasmLoaded = false;
let currentGoInstance = null;

let wasmOverflowCallback;

// Track wasm_exec_rt.js loading state.
let wasmExecRtLoaded = false;
let wasmExecRtLoadPromise = null;

window.setIgopOverflowCallback = function (callback) {
  wasmOverflowCallback = callback;
};

window.isIgopLoaded = function () {
  return isWasmLoaded;
};

window.goWriteSync = function (text) {
  console.log(text);
};

var script = document.createElement("script");
if (useWasm) {
  script.src = "https://play-static.gopluscdn.com/wasm_exec_rt.js";

  // Create a promise to track wasm_exec_rt.js loading.
  wasmExecRtLoadPromise = new Promise((resolve, reject) => {
    script.onload = function () {
      // polyfill
      if (!WebAssembly.instantiateStreaming) {
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
          const source = await (await resp).arrayBuffer();
          return await WebAssembly.instantiate(source, importObject);
        };
      }
      wasmExecRtLoaded = true;
      resolve();
      loadWasm();
      //        const go = new Go();
      //        currentGoInstance = go;
      //        let mod, inst;
      //        WebAssembly.instantiateStreaming(fetch("igop_1dd7d1c3.wasm"), go.importObject).then((result) => {
      //            mod = result.module;
      //            inst = result.instance;
      //            isWasmLoaded = true;
      //            run();
      //        })

      //        async function run() {
      //            await go.run(inst);
      //            inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
      //        }
    };

    script.onerror = function () {
      reject(new Error("Failed to load wasm_exec_rt.js"));
    };
  });
} else {
  script.src = "ixgo_9e852bd5.js";
}
document.head.appendChild(script);

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
  // Ensure wasm_exec_rt.js is loaded before proceeding.
  if (!wasmExecRtLoaded) {
      await wasmExecRtLoadPromise;
  }

  const go = new Go();
  currentGoInstance = go;
  let mod, inst;
  WebAssembly.instantiateStreaming(fetch("https://play-static.gopluscdn.com/ixgo_9e852bd5.wasm"), go.importObject).then(
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
