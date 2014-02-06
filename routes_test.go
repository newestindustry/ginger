package ginger

import (
    "testing"
)

var routes Routes = Routes{
		Route{"/comments", TestHandle},
		Route{"/comments/mary", TestHandle},
		Route{"/ding/flof", TestHandle},
		Route{"/ding", TestHandle},
	}

func TestRoutes(t *testing.T) {
	if len(routes) != 4 {
		t.Errorf("Routes: Amount is %d, should be %d", len(routes), 4)
	}
}

var matchTests = []struct {
        in		string
        out 	string
}{
        {"/comments", "/comments"},
        {"/comments/", "/comments"},
        {"/comments/mary", "/comments"},
        {"/ding", "/ding"},
        {"/ding/", "/ding"},
        {"/ding/flof", "/ding/flof"},
}

func TestRoutesMatch(t *testing.T) {
	for i, tt := range matchTests {
		route, _ := routes.Match(tt.in)
		if route.Url != tt.out {
			t.Errorf("%d. routes.Match(%q) => %s returned, expected %s", i, tt.in, route.Url, tt.out)
		}
	}
}

func TestHandle(t *testing.T) {

}