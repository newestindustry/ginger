package ginger

import (
	"net/http"
	"log"
	"fmt"
)

type Handle func(*Ginger) ResponseData

type Ginger struct {
	Request Request
	Response Response
	Options *Options
	Router *Router
}

func NewGinger() (*Ginger) {
	return &Ginger{}
}

func Init() (*Ginger) {
	g := NewGinger()
	g.Options = &Options{"localhost", 4242}
	g.Router = NewRouter()
	return g
}

func (g *Ginger) Run() {
	http.HandleFunc("/", g.Handle)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", g.Options.Hostname, g.Options.Port), nil)
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
}

func (g *Ginger) Handle(res http.ResponseWriter, req *http.Request) {
	g.Request.Headers = req.Header
	var function Handle
	
	current, err := g.Router.Match(fmt.Sprintf("%s", req.URL))
	if err != nil {
		function = NotFoundHandler
		log.Printf("No route found in %s", req.URL) 
	} else {
		function = current.Handler
	}

	g.Response = Response{res, nil, 200, "json"}
	g.Request.Http = req
	g.ParseRequestHeaders()
	g.Response.Data = function(g) 
	g.sendResponse()
}

// Send response
func (g *Ginger) sendResponse() {
	g.Response.Writer.Header().Set("Server", "Ginger")
	data := g.setResponse(g.Response.Type)
	g.Response.Writer.WriteHeader(g.Response.Status)
	g.Response.Writer.Write(data)
}

// Default 404 handler
func NotFoundHandler(g *Ginger) (resp ResponseData) {
	g.Response.Status = 404
	return resp
}

// Set response header and formats data
func (g *Ginger) setResponse(accept string) (data []byte) {
	switch accept {
		default: 
			g.Response.Writer.Header().Set("Content-Type", "application/json")
			data = ToJson(g.Response.Data)
			break;
			
		case "xml":
			g.Response.Writer.Header().Set("Content-Type", "application/xml")
			data = ToXml(g.Response.Data)
			break;
	}
	
	return data
}
