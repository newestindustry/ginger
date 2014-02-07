package ginger

import (
	"net/http"
//	"net/url"
)

// Ginger's request object
type Request struct {
	Http *http.Request
	Headers http.Header
	Method string
	Data Parameters
	Filter Parameters
}

// Do stuff with request headers
func (g *Ginger) ParseRequestHeaders() {
	accept := g.Request.Headers["Accept"]
	
	if len(accept) > 0 {
		switch accept[0] {
			default:
				g.Response.Type = "json"
				break;
				
			case "application/xml":
				g.Response.Type = "xml"
				break;
		}
	}
}
