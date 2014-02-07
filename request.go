package ginger

import (
	"net/http"
)

// Request object for Ginger Framework
type Request struct {
	HTTP *http.Request
	Headers http.Header
	Method string
	Data Parameters
	Filter Parameters
}

// ParseRequestHeaders parses the request headers
// like accept or authentication
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
