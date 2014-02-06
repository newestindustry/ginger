package ginger

import (
	"strings"
	"errors"
)

type Route struct {
	Url string
	Handler interface{}
}

type Routes []Route

func (routes Routes) Add(route Route) (Routes) {
	
	return routes
}

func (routes Routes) Match(url string) (Route, error) {
	for _, route := range routes {
		if strings.HasPrefix(url, route.Url) == true{
			return route, nil
		} 
	}
	
	return Route{}, errors.New("Couldn't match route")
}