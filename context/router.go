package context

import "log"

type router struct {
	handlers map[string]Handlerfunc
}

func newRouter() *router {
	return &router{handlers: map[string]Handlerfunc{}}
}

func (r *router) addRoute(method string, pattern string, handler Handlerfunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		log.Printf("404 NOT FOUND: %s\n", c.Path)
	}
}
