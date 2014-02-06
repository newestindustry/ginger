package ginger

import (
	"net/http"
)

type Response struct {
	Writer http.ResponseWriter
	
	Status int
	
}

func (r Response) Send() {
	
}