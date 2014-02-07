package ginger

import (
	"strings"
	"errors"
)

// RouteTypeGet is the constant for the GET method
const RouteTypeGet string = "GET"
// RouteTypePost is the constant for the POST method
const RouteTypePost string = "POST"
// RouteTypePut is the constant for the POST method
const RouteTypePut string = "PUT"
// RouteTypeDelete is the constant for the POST method
const RouteTypeDelete string = "DELETE"
// RouteTypeHead is the constant for the POST method
const RouteTypeHead string = "HEAD"
// RouteTypeOptions is the constant for the POST method
const RouteTypeOptions string = "OPTIONS"


// The Route object is where the routing match url and handler is stored 
type Route struct {
	Method string
	Url string
	Handler Handle
	Parameters *Parameters
}

// Routes is the slice of routes
type Routes []Route

// Router holds the routes table and the current route (if matched)
type Router struct {
	Routes Routes
	Current Route
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{Routes: Routes{}}
}

// Get the routes table
func (r *Router) GetRoutes() (Routes) {
	return r.Routes
}

// AddRoute adds a new route to the routes table
func (r *Router) AddRoute(method string, url string, handle func(*Ginger)) (*Router) {
	route := NewRoute(method, url, handle)
	r.Routes = append(r.Routes, route)

	return r
}

// Match the request url against a route to see if there is one found
func (r *Router) Match(method string, url string) (Route, error) {
	for _, route := range r.Routes {
		if strings.HasPrefix(url, route.Url) == true && route.Method == method {
			r.Current = route
			return route, nil
		} 
	}
	
	return Route{}, errors.New("Couldn't match route")
}

// GetCurrent get the current matched route
func (r *Router) GetCurrent() (Route) {
	return r.Current
}

// NewRoute creates a new Route
func NewRoute(method string, url string, handle func(*Ginger)) (r Route) {
	r = Route{Method: method, Url: url, Handler: handle}

	return r
}
