package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"text/template"
)

type RequestLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}

var RequestLogTemplate = `
Remote address:   {{.RemoteAddr}}
Content-Type:     {{.ContentType}}
HTTP method:      {{.Method}}
path:             {{.Path}}
query string:     {{.Query}}
body:
{{.Body}}
`

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		reqBody := bufbody.String()

		line := RequestLog{
			r.RemoteAddr,
			r.Header.Get("Content-Type"),
			r.URL.Path,
			r.URL.RawQuery,
			r.Method,
			reqBody,
		}
		tmpl, err := template.New("line").Parse(RequestLogTemplate)
		if err != nil {
			panic(err)
		}

		bufline := new(bytes.Buffer)
		err = tmpl.Execute(bufline, line)
		if err != nil {
			panic(err)
		}

		log.Printf(bufline.String())
		handler.ServeHTTP(w, r)
	})
}

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", LoggingMiddleware(http.DefaultServeMux))
}
