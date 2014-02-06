package ginger

import (
	"strings"
	"errors"
)

// The Route object is where the routing match url and handler is stored 
type Route struct {
	Method string
	Url string
	Handler Handle
	Parameters *Parameters
}

// Routes is the slice of routes
type Routes []Route

// The router object holds the routes table and the current route
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

// Add a new route to the routes table
func (r *Router) AddRoute(url string, handle func(*Ginger) ResponseData) (*Router) {
	route := NewRoute(url, handle)
	r.Routes = append(r.Routes, route)

	return r
}

// Match the request url against a route to see if there is one found
func (r *Router) Match(url string) (Route, error) {
	for _, route := range r.Routes {
		if strings.HasPrefix(url, route.Url) == true{
			r.Current = route
			return route, nil
		} 
	}
	
	return Route{}, errors.New("Couldn't match route")
}

// Get the current matched route
func (r *Router) GetCurrent() (Route) {
	return r.Current
}

// Create a new Route
func NewRoute(url string, handle func(*Ginger) ResponseData) (r Route) {
	r = Route{Url: url, Handler: handle}

	return r
}
