// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
In the absence of any formal way to specify interfaces in JavaScript,
here's a skeleton implementation of a playground transport.

        function Transport() {
                // Set up any transport state (eg, make a websocket connection).
                return {
                        Run: function(body, output, options) {
                                // Compile and run the program 'body' with 'options'.
				// Call the 'output' callback to display program output.
                                return {
                                        Kill: function() {
                                                // Kill the running program.
                                        }
                                };
                        }
                };
        }

	// The output callback is called multiple times, and each time it is
	// passed an object of this form.
        var write = {
                Kind: 'string', // 'start', 'stdout', 'stderr', 'end'
                Body: 'string'  // content of write or end status message
        }

	// The first call must be of Kind 'start' with no body.
	// Subsequent calls may be of Kind 'stdout' or 'stderr'
	// and must have a non-null Body string.
	// The final call should be of Kind 'end' with an optional
	// Body string, signifying a failure ("killed", for example).

	// The output callback must be of this form.
	// See PlaygroundOutput (below) for an implementation.
        function outputCallback(write) {
        }
*/

const host = "https://play.goplus.org";

var bIgopInit = false;

function hasIgop() {
    if (bIgopInit && isIgopLoaded()) { 
        return true;
    }
    if (typeof gop_ajax === "function") {
        bIgopInit = true;
    }
    return bIgopInit;
}

function js_ajax(url,options) {
    if (hasIgop()) {
       gop_ajax(url,options);
    } else {
    	console.log("pass on play.goplus.org")
        $.ajax(host+url,options);
    }
}

// HTTPTransport is the default transport.
// enableVet enables running vet if a program was compiled and ran successfully.
// If vet returned any errors, display them before the output of a program.
function HTTPTransport(enableVet, fnIsGop) {
	'use strict';

	function playback(output, data) {
		// Backwards compatibility: default values do not affect the output.
		var events = data.Events || [];
		var errors = data.Errors || "";
		var status = data.Status || 0;
		var isTest = data.IsTest || false;
		var testsFailed = data.TestsFailed || 0;
        var isKeep = data.IsKeep || false;

		var timeout;
        if (!isKeep) {
		    output({Kind: 'start'});
        }
		function next() {
			if (!events || events.length === 0) {
				if (isTest) {
					if (testsFailed > 0) {
						output({Kind: 'system', Body: '\n'+testsFailed+' test'+(testsFailed>1?'s':'')+' failed.'});
					} else {
						output({Kind: 'system', Body: '\nAll tests passed.'});
					}
				} else {
					if (status > 0) {
						output({Kind: 'end', Body: 'status ' + status + '.'});
					} else {
						if (errors !== "") {
							// errors are displayed only in the case of timeout.
							output({Kind: 'end', Body: errors + '.'});
						} else {
							output({Kind: 'end'});
						}
					}
				}
				return;
			}
			var e = events.shift();
			if (e.Delay === 0) {
				output({Kind: e.Kind, Body: e.Message});
				next();
				return;
			}
			timeout = setTimeout(function() {
				output({Kind: e.Kind, Body: e.Message});
				next();
			}, e.Delay / 1000000);
		}
		next();
		return {
			Stop: function() {
				clearTimeout(timeout);
			}
		};
	}

	function error(output, msg) {
		output({Kind: 'start'});
		output({Kind: 'stderr', Body: msg});
		output({Kind: 'end'});
	}

	function buildFailed(output, msg) {
		output({Kind: 'start'});
		output({Kind: 'stderr', Body: msg});
		output({Kind: 'system', Body: '\nGo+ build failed.'});
	}

	var seq = 0;
	return {
		Run: function(body, output, options) {
			seq++;
			var cur = seq;
			var playing;
			js_ajax('/compile', {
				type: 'POST',
				data: {'version': 2, 'body': body, 'withVet': enableVet, 'goplus': fnIsGop()},
				dataType: 'json',
				success: function(data) {
					if (seq != cur) return;
					if (!data) return;
					if (playing != null) playing.Stop();
					if (data.Errors) {
						if (data.Errors === 'process took too long') {
							// Playback the output that was captured before the timeout.
							playing = playback(output, data);
						} else {
							buildFailed(output, data.Errors);
						}
						return;
					}
					if (!data.Events) {
						data.Events = [];
					}
					if (data.VetErrors) {
						// Inject errors from the vet as the first events in the output.
						data.Events.unshift({Message: 'Go vet exited.\n\n', Kind: 'system', Delay: 0});
						data.Events.unshift({Message: data.VetErrors, Kind: 'stderr', Delay: 0});
					}

//					if (!enableVet || data.VetOK || data.VetErrors) {
						playing = playback(output, data);
						return;
//					}

					// In case the server support doesn't support
					// compile+vet in same request signaled by the
					// 'withVet' parameter above, also try the old way.
					// TODO: remove this when it falls out of use.
					// It is 2019-05-13 now.
//					js_ajax("/vet", {
//						data: {"body": body},
//						type: "POST",
//						dataType: "json",
//						success: function(dataVet) {
//							if (dataVet.Errors) {
//								// inject errors from the vet as the first events in the output
//								data.Events.unshift({Message: 'Go vet exited.\n\n', Kind: 'system', Delay: 0});
//								data.Events.unshift({Message: dataVet.Errors, Kind: 'stderr', Delay: 0});
//							}
//							playing = playback(output, data);
//						},
//						error: function() {
//							playing = playback(output, data);
//						}
//					});
				},
				error: function() {
					error(output, 'Error communicating with remote server.');
				}
			});
			return {
				Kill: function() {
					if (playing != null) playing.Stop();
					output({Kind: 'end', Body: 'killed'});
				}
			};
		}
	};
}

function SocketTransport() {
	'use strict';

	var id = 0;
	var outputs = {};
	var started = {};
	var websocket;
	if (window.location.protocol == "http:") {
		websocket = new WebSocket('ws://' + window.location.host + '/socket');
	} else if (window.location.protocol == "https:") {
		websocket = new WebSocket('wss://' + window.location.host + '/socket');
	}

	websocket.onclose = function() {
		console.log('websocket connection closed');
	};

	websocket.onmessage = function(e) {
		var m = JSON.parse(e.data);
		var output = outputs[m.Id];
		if (output === null)
			return;
		if (!started[m.Id]) {
			output({Kind: 'start'});
			started[m.Id] = true;
		}
		output({Kind: m.Kind, Body: m.Body});
	};

	function send(m) {
		websocket.send(JSON.stringify(m));
	}

	return {
		Run: function(body, output, options) {
			var thisID = id+'';
			id++;
			outputs[thisID] = output;
			send({Id: thisID, Kind: 'run', Body: body, Options: options});
			return {
				Kill: function() {
					send({Id: thisID, Kind: 'kill'});
				}
			};
		}
	};
}

function PlaygroundOutput(el) {
	'use strict';

	return function(write) {
		if (write.Kind == 'start') {
			el.innerHTML = '';
			return;
		}

		var cl = 'system';
		if (write.Kind == 'stdout' || write.Kind == 'stderr')
			cl = write.Kind;

		var m = write.Body;
		if (write.Kind == 'end') {
			m = '\nProgram exited' + (m?(': '+m):'.');
		}

		if (m.indexOf('IMAGE:') === 0) {
			// TODO(adg): buffer all writes before creating image
			var url = 'data:image/png;base64,' + m.substr(6);
			var img = document.createElement('img');
			img.src = url;
			el.appendChild(img);
			return;
		}

		// ^L clears the screen.
		var s = m.split('\x0c');
		if (s.length > 1) {
			el.innerHTML = '';
			m = s.pop();
		}

		m = m.replace(/&/g, '&amp;');
		m = m.replace(/</g, '&lt;');
		m = m.replace(/>/g, '&gt;');

		var needScroll = (el.scrollTop + el.offsetHeight) == el.scrollHeight;

		var span = document.createElement('span');
		span.className = cl;
		span.innerHTML = m;
		el.appendChild(span);

		if (needScroll)
			el.scrollTop = el.scrollHeight - el.offsetHeight;
	};
}

(function() {
  function lineHighlight(error) {
    var regex = /main.gop?:([0-9]+)/g;
    var r = regex.exec(error);
    while (r) {
      $(".lines div").eq(r[1]-1).addClass("lineerror");
      r = regex.exec(error);
    }
  }
  function highlightOutput(wrappedOutput) {
    return function(write) {
      if (write.Body) lineHighlight(write.Body);
      wrappedOutput(write);
    };
  }
  function lineClear() {
    $(".lineerror").removeClass("lineerror");
  }

  // opts is an object with these keys
  //  codeEl - code editor element
  //  outputEl - program output element
  //  runEl - run button element
  //  fmtEl - fmt button element (optional)
  //  fmtImportEl - fmt "imports" checkbox element (optional)
  //  shareEl - share button element (optional)
  //  shareURLEl - share URL text input element (optional)
  //  shareRedirect - base URL to redirect to on share (optional)
  //  toysEl - toys select element (optional)
  //  enableHistory - enable using HTML5 history API (optional)
  //  transport - playground transport to use (default is HTTPTransport)
  //  enableShortcuts - whether to enable shortcuts (Ctrl+S/Cmd+S to save) (default is false)
  //  enableVet - enable running vet and displaying its errors
  function playground(opts) {
    var code = $(opts.codeEl);
    function isGoplus() {
    	return $(opts.enableGoplus).is(":checked")
    }
    var transport = opts['transport'] || new HTTPTransport(opts['enableVet'],isGoplus);
    var running;
    
    const search = location.search.substring(1); 
    const params = new URLSearchParams(search);
    if (params.has('p')) {
      let snippetUrl = host+"/p/"+params.get('p')+".gop"
      $.ajax(snippetUrl, {
        method: 'GET',
        dataType: 'text',
        success: function(response) {
           setBody(response);
        },
//        error: function(xhr, status, error) {
//           console.error(status, error);
//        }
      })
    } 

    // autoindent helpers.
    function insertTabs(n) {
      // find the selection start and end
      var start = code[0].selectionStart;
      var end   = code[0].selectionEnd;
      // split the textarea content into two, and insert n tabs
      var v = code[0].value;
      var u = v.substr(0, start);
      for (var i=0; i<n; i++) {
        u += "\t";
      }
      u += v.substr(end);
      // set revised content
      code[0].value = u;
      // reset caret position after inserted tabs
      code[0].selectionStart = start+n;
      code[0].selectionEnd = start+n;
    }
    function autoindent(el) {
      var curpos = el.selectionStart;
      var tabs = 0;
      while (curpos > 0) {
        curpos--;
        if (el.value[curpos] == "\t") {
          tabs++;
        } else if (tabs > 0 || el.value[curpos] == "\n") {
          break;
        }
      }
      setTimeout(function() {
        insertTabs(tabs);
      }, 1);
    }

    // NOTE(cbro): e is a jQuery event, not a DOM event.
    function handleSaveShortcut(e) {
      if (e.isDefaultPrevented()) return false;
      if (!e.metaKey && !e.ctrlKey) return false;
      if (e.key != "S" && e.key != "s") return false;

      e.preventDefault();

      // Share and save
      share(function(url) {
        window.location.href = url + ".go?download=true";
      });

      return true;
    }

    function keyHandler(e) {
      if (opts.enableShortcuts && handleSaveShortcut(e)) return;

      if (e.keyCode == 9 && !e.ctrlKey) { // tab (but not ctrl-tab)
        insertTabs(1);
        e.preventDefault();
        return false;
      }
      if (e.keyCode == 13) { // enter
        if (e.shiftKey) { // +shift
          run();
          e.preventDefault();
          return false;
        } if (e.ctrlKey) { // +control
          fmt();
          e.preventDefault();
        } else {
          autoindent(e.target);
        }
      }
      return true;
    }
    code.unbind('keydown').bind('keydown', keyHandler);
    var outdiv = $(opts.outputEl).empty();
    var output = $('<pre/>').appendTo(outdiv);
   
    function body() {
      return $(opts.codeEl).val();
    }
    function setBody(text) {
      $(opts.codeEl).val(text);
    }
    function origin(href) {
      return (""+href).split("/").slice(0, 3).join("/");
    }

    var pushedEmpty = (window.location.pathname == "/");
    function inputChanged() {
      if (pushedEmpty) {
        return;
      }
      pushedEmpty = true;
      $(opts.shareURLEl).hide();
      window.history.pushState(null, "", "/");
    }
    function popState(e) {
      if (e === null) {
        return;
      }
      if (e && e.state && e.state.code) {
        setBody(e.state.code);
      }
    }
    var rewriteHistory = false;
    if (window.history && window.history.pushState && window.addEventListener && opts.enableHistory) {
      rewriteHistory = true;
      code[0].addEventListener('input', inputChanged);
      window.addEventListener('popstate', popState);
    }

    function setError(error) {
      if (running) running.Kill();
      lineClear();
      lineHighlight(error);
      output.empty().addClass("error").text(error);
    }
    let wait = false;
    function loading() {
      wait = true;
      lineClear();
      if (running) running.Kill();
      if (!hasIgop()) {
         output.removeClass("error").text('Waiting for compilation...');
      } else {
         output.removeClass("error").text('');        
      }
    }
    let playout = highlightOutput(PlaygroundOutput(output[0]));
    setIgopOverflowCallback(function(event) {
        gop_ajax = undefined;
        if (running) running.Kill();
        let s = "\n[Stack overflow] "+event.message+"\n\nPlease check your code.";
        playout({Kind: "stderr", Body: s});
    })
    window.goWriteSync = function(s) {
        if (wait && output.text() === 'Waiting for compilation...') {
            wait = false;
            output.text('');
        }
        playout({Kind: "stdout", Body: s});
    }

    function run() {
        loading();
        //if (!hasIgop()) {
        running = transport.Run(body(), highlightOutput(PlaygroundOutput(output[0])));
//        } else {
//	       var data = {"body": body(), "goplus": $(opts.enableGoplus).is(":checked") };
//	       js_ajax("/compile", {
//	           data: data,
//	           type: "POST",
//	           dataType: "json",
//	           success: function(data) {
//	               setOutput("");
//	               if (data.Error) {
//	                   setError(data.Error);
//	               } else {
//	                   data.Kind = "stdout"
//	                   highlightOutput(PlaygroundOutput(output[0]))(data);
//	                   data.Kind = "end";
//	                   data.Body = data.Code;
//	                   highlightOutput(PlaygroundOutput(output[0]))(data);
//	               }
//	           }
//	       });
	    //}
    }

    function fmt() {
      loading();
      var data = {"body": body(), "goplus": $(opts.enableGoplus).is(":checked")};
      if ($(opts.fmtImportEl).is(":checked")) {
        data["imports"] = "true";
      }
      js_ajax("/fmt", {
        data: data,
        type: "POST",
        dataType: "json",
        success: function(data) {
          if (data.Error) {
            setError(data.Error);
          } else {
            setBody(data.Body);
            setError("");
          }
        }
      });
    }

    var shareURL; // jQuery element to show the shared URL.
    var sharing = false; // true if there is a pending request.
    var shareCallbacks = [];
    function share(opt_callback) {
      if (opt_callback) shareCallbacks.push(opt_callback);

      if (sharing) return;
      sharing = true;

      var sharingData = body();
      $.ajax("https://play.goplus.org/share", {
        processData: false,
        data: sharingData,
        type: "POST",
        contentType: "text/plain; charset=utf-8",
        complete: function(xhr) {
          sharing = false;
          if (xhr.status != 200) {
            alert("Server error; try again.");
            return;
          }
          if (opts.shareRedirect) {
            window.location = opts.shareRedirect + xhr.responseText;
          }
          var path = "/p/" + xhr.responseText;
          //var url = origin(window.location) + path;
          var url = "https://play.goplus.org" + path;

          for (var i = 0; i < shareCallbacks.length; i++) {
            shareCallbacks[i](url);
          }
          shareCallbacks = [];

          if (shareURL) {
            shareURL.show().val(url).focus().select();

            if (rewriteHistory) {
            var historyData = {"code": sharingData};
            window.history.pushState(historyData, "", "");
            pushedEmpty = false;
            }
          }
        }
      });
    }

    $(opts.runEl).click(run);
    $(opts.fmtEl).click(fmt);

    if (opts.shareEl !== null && (opts.shareURLEl !== null || opts.shareRedirect !== null)) {
      if (opts.shareURLEl) {
        shareURL = $(opts.shareURLEl).hide();
      }
      $(opts.shareEl).click(function() {
        share();
      });
    }

    if (opts.toysEl !== null) {
      $(opts.toysEl).bind('change', function() {
        var toy = $(this).val();
        switch (toy) {
        case "hello.txt":
            setBody(`// You can edit this code!
// Click here and start typing.

echo "Hello, 世界"
`)
		break;
        case "hellogop.txt":
        		setBody(`fields := [
	"engineering",
	"STEM education",
	"and data science",
]

echo "The Go+ language for", fields.join(", ")
`)
		break;
		case "basic.txt":
			setBody(`echo "Hello, Go+"

echo 1r<<129
echo 1/3r+2/7r*2

arr := [1, 3, 5, 7, 11, 13, 17, 19]
echo arr
echo [x*x for x <- arr if x > 3]

m := {"Hi": 1, "Go+": 2}
echo m
echo {v: k for k, v <- m}
echo [k for k, _ <- m]
echo [v for v <- m]
`)
		break;
		case "range.txt":
			setBody(`a := [1, 3, 5, 7, 11]
b := [x*x for x <- a if x > 3]
echo b // output: [25 49 121]

mapData := {"Hi": 1, "Hello": 2, "Go+": 3}
reversedMap := {v: k for k, v <- mapData}
echo reversedMap // output: map[1:Hi 2:Hello 3:Go+]

sum := 0
for x <- [1, 3, 5, 7, 11, 13, 17] if x > 3 {
	sum += x
}
echo sum
`)
		break;
		case "rational.txt":
			setBody(`a := 1r << 65 // bigint, large than int64
b := 4/5r     // bigrat
c := b - 1/3r + 3*1/2r
echo a, b, c
`)
		break;
		case "slice.txt":
			setBody(`// in Go we do:
a := []float64{1, 2, 3.4}
println(a)

// in GoPlus we do:
b := [1, 2, 3.4]
echo b

//slice literal
c := [1, 3.4] // []float64
printf "%#v %T\\n", c, c

d := [1] // []int
printf "%#v %T\\n", d, d

e := [1+2i, "xsw"] // []interface{}
printf "%#v %T\\n", e, e

f := [1, 3.4, 3+4i] // []complex128
printf "%#v %T\\n", f, f

g := [5+6i] // []complex128
printf "%#v %T\\n", g, g

h := ["xsw", 3] // []interface{}
printf "%#v %T\\n", h, h

empty := [] // []interface{}
printf "%#v %T\\n", empty, empty
`)
		break;
		case "listmap.txt":
		setBody(`a := [x*x for x <- [1, 3, 5, 7, 11]]
echo a
b := [x*x for x <- [1, 3, 5, 7, 11] if x > 3]
echo b
c := [i+v for i, v <- [1, 3, 5, 7, 11] if i%2 == 1]
echo c
d := [k+","+s for k, s <- {"Hello": "xsw", "Hi": "Go+"}]
echo d

arr := [1, 2, 3, 4, 5, 6]
e := [[a, b] for a <- arr if a < b for b <- arr if b > 2]
echo e

x := {x: i for i, x <- [1, 3, 5, 7, 11]}
echo x
y := {x: i for i, x <- [1, 3, 5, 7, 11] if i%2 == 1}
echo y
z := {v: k for k, v <- {1: "Hello", 3: "Hi", 5: "xsw", 7: "Go+"} if k > 3}
echo z
`)
		break;
		case "error.txt":
		setBody(`import (
	"strconv"
)

func add(x, y string) (int, error) {
	return strconv.atoi(x)? + strconv.atoi(y)?, nil
}

func addSafe(x, y string) int {
	return strconv.atoi(x)?:0 + strconv.atoi(y)?:0
}

// Error handling
// We reinvent the error handling specification in Go+. We call them ErrWrap expressions:

// expr! // panic if err
// expr? // return if err
// expr?:defval // use defval if err

echo add("100", "23")!

sum, err := add("10", "abc")
echo sum, err

echo addSafe("10", "abc")
`)
		break;
		case "canvas.txt":
		setBody(`import "syscall/js"

document := js.global.get("document")
canvas := document.call("getElementById", "canvas")
canvas.set "width", 200
canvas.set "height", 200
ctx := canvas.call("getContext", "2d")
ctx.set "fillStyle", "rgba(128,0,0,0.3)"
ctx.call "fillRect", 10, 10, 100, 100
`)
		break;
        case "generic.txt":
        setBody(`
package main

import (
	"fmt"
)

// The index function returns the index of the first occurrence of v in s,
// or -1 if not present.
func index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func main() {
	s := []int{1, 3, 5, 2, 4}
	fmt.Println(index(s, 3))
	fmt.Println(index(s, 6))
}
`)
		break;
}
//        js_ajax("/doc/play/"+toy, {
//          processData: false,
//          type: "GET",
//          complete: function(xhr) {
//            if (xhr.status != 200) {
//              alert("Server error; try again.");
//              return;
//            }
//            setBody(xhr.responseText);
//          }
//        });
      });
    }
  }

  window.playground = playground;
})();