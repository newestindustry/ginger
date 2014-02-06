package ginger

import (
    "testing"
)

var routestest Routes
var routertest *Router

func init() {
	routertest = NewRouter()
	routertest.AddRoute("/a", MockCallback)
	routertest.AddRoute("/b", MockCallback)
	routertest.AddRoute("/c", MockCallback)
	routertest.AddRoute("/ca", MockCallback)
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
		route, _ := routertest.Match(tt.in)
		if route.Url != tt.out {
			t.Errorf("%d. routes.Match(%q) => %s returned, expected %s", i, tt.in, route.Url, tt.out)
		}
	}
}

func MockCallback(g *Ginger) {
	
}