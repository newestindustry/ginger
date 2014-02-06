package ginger

import (
	"net/http"
	"log"
	"fmt"
)

type Ginger struct {
	Request Request
	Response Response
	Options Options
	Routes Routes
}

func Init() (Ginger) {
	ginger := Ginger{}
	ginger.Options = Options{"localhost", 4242} 
	ginger.Routes = Routes{}
	
	return ginger
}

func (g Ginger) Run() {
	http.HandleFunc("/", g.Handle)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", g.Options.Hostname, g.Options.Port), nil)
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
}

func (g Ginger) Handle(res http.ResponseWriter, req *http.Request) {
	currentRoute, err := g.Routes.Match(fmt.Sprintf("%s", req.URL))
	if err != nil {
		
	}
	fmt.Println(currentRoute)
	g.Response = Response{res, 200}
	g.Request.Http = req
}