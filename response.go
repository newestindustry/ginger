package ginger

import (
	"encoding/xml"
	"net/http"
	"encoding/json"
	"reflect"
	"log"
)

// Response is the main response object
type Response struct {
	// Contains the http Response Writer so we can easily send a response
	Writer http.ResponseWriter
	// This contains any data type, struct or variable you want to return
	Data interface{}
	// The HTTP status code to return
	Status int
	// Response type
	Type string
}

// StringResponse makes sure that if the Response.Data structure is just a string
// We cast it in a string response before sending it
// so you can get a nice structure
type StringResponse struct  {
	XMLName   xml.Name `json:"-" xml:"response"`
	Response string `json:"response" xml:"response"`
}

func FixResponseData(data interface{}) interface{} {
	if data != nil {
		if reflect.TypeOf(data).String() == "string" {
			data = StringResponse{Response: data.(string)}
		}
	}
	
	return data
}

// ToJSON is a helper function for casting a data structure into
// a json byte slice
func ToJSON(data interface{}) (b []byte){
	data = FixResponseData(data)
	
	b, err := json.Marshal(data)
	if err != nil {
		return b
    }
	return b
}

// ToXML is a helper function for casting a data structure into
// a xml byte slice
func ToXML(data interface{}) (b []byte){
	data = FixResponseData(data)
	
	b, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(err)
		return b
    }
	return b
}
