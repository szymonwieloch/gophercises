package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
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
				stack = transformCallStack(stack)
				fmt.Fprintf(w, "<pre>%s</pre>", stack)
			}
		}()
		brw := bufferedResponseWriter{w: w}
		h.ServeHTTP(&brw, r)
		// didn't panic
		brw.flush()
	}
}

// Converts the "normal" call stack into a html-based callstack with links
func transformCallStack(cs string) string {
	var buf strings.Builder
	lines := strings.Split(cs, "\n")
	for _, line := range lines {
		if !strings.HasPrefix(line, "\t") {
			buf.WriteString(line)
			buf.WriteString("\n")
			continue
		}
		colIdx := strings.Index(line, ":")
		fileName := strings.TrimSpace(line[:colIdx])
		lineNum := strings.Split(line[colIdx+1:], " ")[0]
		buf.WriteString(line[:(colIdx - len(fileName))]) // whitespace prefix
		urlValues := url.Values{}
		urlValues.Add("path", fileName)
		urlValues.Add("line", lineNum)
		buf.WriteString("<a href=\"/source?")
		buf.WriteString(urlValues.Encode())
		buf.WriteString("\" >")
		buf.WriteString(fileName)
		buf.WriteString(":")
		buf.WriteString(lineNum)
		buf.WriteString("</a>")
		buf.WriteString(line[colIdx+1+len(lineNum):])
		buf.WriteString("\n")
	}
	return buf.String()
}
