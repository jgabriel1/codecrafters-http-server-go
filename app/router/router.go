package router

import (
	"fmt"
	"regexp"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
)

type Handler func(req request.Request, cfg config.Config) *response.Response

type Router struct {
	routes          []string
	handlers        map[string]map[string]Handler
	defaultNotFound *response.Response
}

func New() *Router {
	return &Router{
		routes:          []string{},
		handlers:        make(map[string]map[string]Handler, 0),
		defaultNotFound: response.DefaultNotFound(),
	}
}

func (r *Router) AddRoute(path string, method string, h Handler) {
	r.routes = append(r.routes, path)
	if pathHandlers, ok := r.handlers[path]; ok {
		pathHandlers[method] = h
		return
	}
	r.handlers[path] = map[string]Handler{method: h}
}

func (r *Router) Handle(req request.Request, cfg config.Config) *response.Response {
	route, ok := r.matchRoute(req)
	if !ok {
		return r.defaultNotFound
	}
	h := r.getHandler(req, route)
	if h == nil {
		return r.defaultNotFound
	}
	return h(req, cfg)
}

func (r *Router) matchRoute(req request.Request) (string, bool) {
	for _, route := range r.routes {
		regex := fmt.Sprintf("^%s$",
			regexp.MustCompile(`:[^/]+`).ReplaceAllString(route, `[^/]+`))
		matched, _ := regexp.MatchString(regex, req.Path)
		if matched {
			return route, true
		}
	}
	return "", false
}

func (r *Router) getHandler(req request.Request, matchedRoute string) Handler {
	methodsMap, ok := r.handlers[matchedRoute]
	if !ok {
		return nil
	}
	h, ok := methodsMap[req.Method]
	if !ok {
		return nil
	}
	return h
}
