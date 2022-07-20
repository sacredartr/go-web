package context

import "net/http"

type H map[string]interface{}

type Context struct {
	writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
