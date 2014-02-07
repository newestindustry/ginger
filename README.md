Ginger Rest Framework for GO
=========

The Ginger Rest API Framework written in Go!

Package documentation can be found at https://godoc.org/github.com/newestindustry/ginger

Just like the PHP Framework key features include:

- key/value pair url parsing (into filters)
- automatic value casting for data and filter parameters
- easy creation of routes and handles

## Filter parameters
Filter parameters are key/value parameters passed on to the URI and value casted as described below.
Say we have a route GET /thing/ and a GET request to /thing/id/212/public/true is sent, the values will be parsed in to a Paramers object (g.Request.Filter) looking like:

```go
g.Request.Filter["id"] = 212
g.Request.Filter["public"] = true
```

These values can be used any way you'd like. What we usually do is pass them on to our DB layer as the SELECT parameters.

## Data parameters
Data parameters are postfield parameters and value casted as described below.
Say we have a route POST /thing/ and a POST request with postdata name=Thing&public=false&sortorder=42 is sent, the values will be parsed in to a Paramers object (g.Request.Data) looking like:

```go
g.Request.Data["name"] = "Thing"
g.Request.Data["public"] = false
g.Request.Data["sortorder"] = 42
```

These values can be used any way you'd like. What we usually do is pass them on to our DB layer as the INSERT values.


## Value casting
We try to cast all incoming values to their correct datatype for ease of use. The datatypes currently supported are:

| Type | Possible value | Description |
| ------ | ---- | ------ |
| integer | 42 | Strings which are numeric are cast to an integer |
| float | 42.24 | Strings which are floating point decimals are cast to a floating point decimal |
| boolean | true | If a parameter is “true” or “false” it is cast as a boolean |
| default | Just a bunch of characters | Default behavior is to cast as string. |


## Implementation


Implementation is simple:

```Go
package main

import (
	"ni.nl/ginger"
)

func main() {
	g := ginger.Init()
	
	g.Router.AddRoute(ginger.ROUTE_TYPE_GET, "/thing", func(g *ginger.Ginger) {
							
							g.SendResponse(g.Request.Filter)
					  })
					  
	g.Router.AddRoute(ginger.ROUTE_TYPE_POST, "/thing", func(g *ginger.Ginger) {
							
							g.SendResponse(g.Request.Data)
					  })				  
	g.Run()

}

```

When you go to http://localhost:4242/thing the result will be:

```js
{}
```

Go to http://localhost:4242/thing/bool/true/a/b/id/123/number/12.42 and the result will be:
```js
{
    "a": "b",
    "bool": true,
    "id": 123,
    "number": 12.42
}
```

Post a=b&float=12.42&int=12&bool=true to http://localhost:4242/thing/ and the result will be:
```js
{
    "a": "b",
    "bool": true,
    "float": 12.42,
    "int": 12
}
```


This is still in early dev phase.
