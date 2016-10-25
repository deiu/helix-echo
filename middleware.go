package helix

import (
	"net/http"
)

var (
	emptyHandler = http.Handler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {}))
)

type Middleware struct {
	Wares []func(http.Handler) http.Handler
}

// Return an empty middleware that is ready to use
func NewMiddleware() *Middleware {
	return &Middleware{}
}

// Add a piece of middleware which is an http.Handler generator
// (signature: func(http.Handler)http.Handler) which, somewhere before it
// finishes, is expected to call .ServeHTTP on the handler that is passed to it.
// Failure to call .ServeHTTP within the http.Handler generator will cause part
// of the stack not to be called.
func (mw *Middleware) Use(handler func(http.Handler) http.Handler) {
	mw.Wares = append(mw.Wares, handler)
}

// Add a piece of middleware which is simply any http.Handler
// (signature: http.Handler). Unlike with Use, we will automatically call
// .ServeHTTP to ensure that the rest of the middleware stack is called.
func (mw *Middleware) UseHandler(handler http.Handler) {
	x := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			handler.ServeHTTP(w, req)
			next.ServeHTTP(w, req)
		})
	}

	mw.Use(x)
}

// Handler returns the composed handler of all the wares. Warning: you would need to
// call this again if you change the wares.
func (mw *Middleware) Handler() http.Handler {
	//Initialize with an empty http.Handler
	next := emptyHandler

	//Call the middleware stack in FIFO order
	for i := len(mw.Wares) - 1; i >= 0; i-- {
		next = mw.Wares[i](next)
	}
	return next
}

// Satisfies the net/http Handler interface and calls the middleware stack
func (mw *Middleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if len(mw.Wares) < 1 {
		return
	}

	next := mw.Handler()

	//Finally, serve back up the chain
	next.ServeHTTP(w, req)
}
