package ginger

import (
    "testing"
)

var routestest Routes
var routertest *Router

func init() {
	routertest = NewRouter()
	routertest.AddRoute("GET", "/a", MockCallback)
	routertest.AddRoute("GET", "/b", MockCallback)
	routertest.AddRoute("GET", "/c", MockCallback)
	routertest.AddRoute("GET", "/ca", MockCallback)
	routestest = routertest.GetRoutes()
}

func TestRoutes(t *testing.T) {
	if len(routestest) != 4 { 
		t.Errorf("TestRoutes: Got %d, expected %d", len(routestest), 4)
	}
}

var matchTests = []struct {
	in		string
	out 	string
}{
	{"/c", "/c"},
	{"/c/", "/c"},
	{"/ca", "/c"},
	{"/cd", "/c"},
	{"/c/d", "/c"},
}

func TestRoutesMatch(t *testing.T) {
	for i, tt := range matchTests {
		route, _ := routertest.Match("GET", tt.in)
		if route.Url != tt.out {
			t.Errorf("%d. routes.Match(%q) => %s returned, expected %s", i, tt.in, route.Url, tt.out)
		}
	}
}

func MockCallback(g *Ginger) {
	
}