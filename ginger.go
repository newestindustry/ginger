package ginger

import (
	"net/http"
	"log"
	"fmt"
)

type Handle func(*Ginger)

type Ginger struct {
	Request Request
	Response Response
	Options *Options
	Router *Router
}


func NewGinger() (*Ginger) {
	return &Ginger{}
}

// Initialize the ginger framework
// Set default options an router object
func Init() (*Ginger) {
	g := NewGinger()
	g.Options = &Options{"localhost", 4242}
	g.Router = NewRouter()
	return g
}

// Start the http server on the port set in options
func (g *Ginger) Run() {
	http.HandleFunc("/", g.Handle)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", g.Options.Hostname, g.Options.Port), nil)
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
}

// Read headers, check the current route, parse the parameters
// and call the Handle function with the Ginger object
func (g *Ginger) Handle(res http.ResponseWriter, req *http.Request) {
	g.Request.Headers = req.Header
	var function Handle
	
	current, err := g.Router.Match(req.Method, fmt.Sprintf("%s", req.URL))
	if err != nil {
		function = NotFoundHandler
		log.Printf("No route found in %s", req.URL) 
	} else {
		function = current.Handler
	}

	g.Response = Response{res, nil, 200, "json"}
	g.Request.Http = req
	g.Request.Method = req.Method
	g.Request.Filter = current.ParseFilterParameters(fmt.Sprintf("%s", req.URL))
	if req.Method == "POST" || req.Method == "PUT" {
		req.ParseForm()
		g.Request.Data = ParseDataParameters(req.Form)
	}
	g.ParseRequestHeaders()
	
	function(g)
}

// Send data d back to the client
// Set the response headers as well
func (g *Ginger) SendResponse(d interface{}) {
	g.Response.Writer.Header().Set("Server", "Ginger")
	data := g.setResponseData(g.Response.Type, d)
	g.Response.Writer.WriteHeader(g.Response.Status)
	g.Response.Writer.Write(data)
}

// Default 404 handler
func NotFoundHandler(g *Ginger) {
	g.Response.Status = 404
	g.SendResponse("Not found")
}

// Set response header & data and formats data
func (g *Ginger) setResponseData(accept string, d interface{}) (data []byte) {
	switch accept {
		default: 
			g.Response.Writer.Header().Set("Content-Type", "application/json")
			data = ToJson(d)
			break;
			
		case "xml":
			g.Response.Writer.Header().Set("Content-Type", "application/xml")
			data = ToXml(d)
			break;
	}
	
	return data
}
