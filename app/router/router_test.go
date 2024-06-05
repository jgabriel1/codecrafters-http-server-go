package router

import (
	"reflect"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
	"github.com/stretchr/testify/assert"
)

type spy[F any] struct {
	fn    F
	calls *int
}

func spyOn[F any](f F) spy[F] {
	calls := 0
	fn := reflect.ValueOf(f)
	wrapper := reflect.MakeFunc(fn.Type(),
		func(args []reflect.Value) (results []reflect.Value) {
			calls++
			return fn.Call(args)
		})
	return spy[F]{
		fn:    wrapper.Interface().(F),
		calls: &calls,
	}
}

func TestMatchesRoute(t *testing.T) {
	spy := spyOn(stubHandler())
	r := New()
	r.AddRoute("/foo", "GET", spy.fn)
	res := r.Handle(stubRequest("GET", "/foo"), stubConfig())
	assert.NotNil(t, res, "response should be not nil")
	assert.Equal(t, 1, *spy.calls, "handler should've been called once")
}

func TestMatchesMultipleRoutes(t *testing.T) {
	spyFooGet := spyOn(stubHandler())
	spyFooPost := spyOn(stubHandler())
	spyBar := spyOn(stubHandler())
	r := New()
	r.AddRoute("/foo", "GET", spyFooGet.fn)
	r.AddRoute("/foo", "POST", spyFooPost.fn)
	r.AddRoute("/bar", "GET", spyBar.fn)
	for range 2 {
		r.Handle(stubRequest("GET", "/foo"), stubConfig())
	}
	for range 3 {
		r.Handle(stubRequest("POST", "/foo"), stubConfig())
	}
	for range 5 {
		r.Handle(stubRequest("GET", "/bar"), stubConfig())
	}
	assert.Equal(t, 2, *spyFooGet.calls)
	assert.Equal(t, 3, *spyFooPost.calls)
	assert.Equal(t, 5, *spyBar.calls)
}

func TestMatchesRoutesWithParameters(t *testing.T) {
	spy := spyOn(stubHandler())
	r := New()
	r.AddRoute("/foo/bar/:param", "GET", spy.fn)
	res := r.Handle(stubRequest("GET", "/foo/bar/abc"), stubConfig())
	assert.NotNil(t, res, "response should be not nil")
	assert.Equal(t, 1, *spy.calls, "handler should've been called once")
}

func TestNotFound(t *testing.T) {
	r := New()
	res := r.Handle(stubRequest("GET", "/foo"), stubConfig())
	assert.EqualValues(t, response.DefaultNotFound(), res)
}

func stubRequest(method string, path string) request.Request {
	return request.Request{Path: path, Method: method}
}

func stubConfig() config.Config {
	return config.Config{}
}

func stubHandler() Handler {
	return func(req request.Request, cfg config.Config) *response.Response {
		return new(response.Response)
	}
}
