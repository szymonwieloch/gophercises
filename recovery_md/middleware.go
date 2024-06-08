package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
)

type bufferedResponseWriter struct {
	w      http.ResponseWriter
	status int
	buf    bytes.Buffer
}

func (brw *bufferedResponseWriter) WriteHeader(statusCode int) {
	brw.status = statusCode
}

func (brw *bufferedResponseWriter) Header() http.Header {
	return brw.w.Header()
}

func (brw *bufferedResponseWriter) Write(b []byte) (int, error) {
	return brw.buf.Write(b)
}

func (brw *bufferedResponseWriter) flush() error {
	if brw.status != 0 {
		brw.w.WriteHeader(brw.status)
	}
	_, err := brw.w.Write(brw.buf.Bytes())
	return err
}

func (brw *bufferedResponseWriter) Flush() {
	f, ok := brw.w.(http.Flusher)
	if ok {
		f.Flush()
	}
}

func (brw *bufferedResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := brw.w.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijacking is not supported")
	}
	return h.Hijack()
}

var _ http.ResponseWriter = (*bufferedResponseWriter)(nil)
var _ http.Flusher = (*bufferedResponseWriter)(nil)
var _ http.Hijacker = (*bufferedResponseWriter)(nil)

func recoveryMw(h http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rec := recover()
			if rec == nil {
				return
			}
			log.Println("Panic: ", rec)
			stack := string(debug.Stack())
			log.Println("Stack: ", stack)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "<h1>Something went wrong!</h1>")
			if dev {
				fmt.Fprintf(w, "<pre>%s</pre>", stack)
			}
		}()
		brw := bufferedResponseWriter{w: w}
		h.ServeHTTP(&brw, r)
		// didn't panic
		brw.flush()
	}
}
