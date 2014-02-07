package ginger

import (
	"net/url"
	"strings"
	"math"
	"reflect"
	"strconv"
)


// Parameters container
type Parameters map[string]interface{}

// ParseDataParameters parses the "data" parameters. These are coming from 
// url.Values (coming from http.ParseForm) and returns
// The casted parameters
func ParseDataParameters(data url.Values) (Parameters) {
	params := Parameters{}
	
	for key, val := range data {
		if len(val) > 0 {
			params[key] = CastParameterValue(val[0])
		}
	}

	return params
}

// ParseFilterParameters parses the filter parameters from the current url based
// on the current route
func (r *Route) ParseFilterParameters(url string) (Parameters) {
	route := r.URL

	params := Parameters{}
	bare := strings.Replace(url, route, "", -1)
	if len(bare) > 1 {
		if bare[:1] == "/" {
			bare = bare[1:]
		}
		if bare[len(bare)-1:] == "/" {
			bare = bare[:len(bare)-1]
		}
		
		parts := strings.Split(bare, "/")
		var curkey string
		for key, val := range parts {
			modulo := math.Mod(float64(key), 2)
			
			if modulo == 0 {
				val = strings.ToLower(val)
				params[val] = ""
				curkey = val
			} else {
				params[curkey] = CastParameterValue(val)
			}
		}
	}
	
	return params
}

// CastParameterValue casts a value to it's correct value part for easy usage
func CastParameterValue(value interface{}) interface{} {
	if reflect.TypeOf(value).String() == "string" {
		stringvalue := value.(string)
		if strings.ToLower(stringvalue) == "true" || strings.ToLower(stringvalue) == "false" {
			thisbool, err := strconv.ParseBool(strings.ToLower(stringvalue))
			if err == nil {
				return thisbool
			}
		}
		
		thisint, err := strconv.ParseInt(stringvalue, 0, 64)
		if err == nil {
			return int(thisint)
		}
		
		thisfloat, err := strconv.ParseFloat(stringvalue, 64)
		if err == nil {
			return thisfloat
		}
		
		return stringvalue
	}
	
	return value
}