package main

import (
	"net/http"

	"github.com/TranQuocToan1996/goWebSocket/internal/handlers"
	"github.com/bmizerany/pat"
)

// return New PatternServeMux, see comment below
func newMux() http.Handler {
	mux := pat.New()
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}

// PatternServeMux is an HTTP request multiplexer. It matches the URL of each
// incoming request against a list of registered patterns with their associated
// methods and calls the handler for the pattern that most closely matches the
// URL.
//
// Pattern matching attempts each pattern in the order in which they were
// registered.
//
// Patterns may contain literals or captures. Capture names start with a colon
// and consist of letters A-Z, a-z, _, and 0-9. The rest of the pattern
// matches literally. The portion of the URL matching each name ends with an
// occurrence of the character in the pattern immediately following the name,
// or a /, whichever comes first. It is possible for a name to match the empty
// string.
//
// Example pattern with one capture:
//   /hello/:name
// Will match:
//   /hello/blake
//   /hello/keith
// Will not match:
//   /hello/blake/
//   /hello/blake/foo
//   /foo
//   /foo/bar
//
// Example 2:
//    /hello/:name/
// Will match:
//   /hello/blake/
//   /hello/keith/foo
//   /hello/blake
//   /hello/keith
// Will not match:
//   /foo
//   /foo/bar
//
// A pattern ending with a slash will add an implicit redirect for its non-slash
// version. For example: Get("/foo/", handler) also registers
// Get("/foo", handler) as a redirect. You may override it by registering
// Get("/foo", anotherhandler) before the slash version.
//
// Retrieve the capture from the r.URL.Query().Get(":name") in a handler (note
// the colon). If a capture name appears more than once, the additional values
// are appended to the previous values (see
// http://golang.org/pkg/net/url/#Values)
//
// A trivial example server is:
//
//	package main
//
//	import (
//		"io"
//		"net/http"
//		"github.com/bmizerany/pat"
//		"log"
//	)
//
//	// hello world, the web server
//	func HelloServer(w http.ResponseWriter, req *http.Request) {
//		io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
//	}
//
//	func main() {
//		m := pat.New()
//		m.Get("/hello/:name", http.HandlerFunc(HelloServer))
//
//		// Register this pat with the default serve mux so that other packages
//		// may also be exported. (i.e. /debug/pprof/*)
//		http.Handle("/", m)
//		err := http.ListenAndServe(":12345", nil)
//		if err != nil {
//			log.Fatal("ListenAndServe: ", err)
//		}
//	}
//
// When "Method Not Allowed":
//
// Pat knows what methods are allowed given a pattern and a URI. For
// convenience, PatternServeMux will add the Allow header for requests that
// match a pattern for a method other than the method requested and set the
// Status to "405 Method Not Allowed".
//
// If the NotFound handler is set, then it is used whenever the pattern doesn't
// match the request path for the current method (and the Allow header is not
// altered).
