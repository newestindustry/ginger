package ginger

import (
	"net/url"
	"strings"
	"math"
	"reflect"
	"strconv"
	"fmt"
)

type Parameters map[string]interface{}

func ParseDataParameters(data url.Values) (Parameters) {
	params := Parameters{}
	
	for key, val := range data {
		if len(val) > 0 {
			params[key] = CastParameterValue(val[0])
		}
//		params[key] = CastParameterValue(val[0])
//		fmt.Println(key, val[0])
	}
	
	fmt.Println(params)
	
	return params
}

func (r *Route) ParseFilterParameters(url string) (Parameters) {
	route := r.Url

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