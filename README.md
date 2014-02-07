Ginger Rest Framework for GO
=========

The Ginger Rest API Framework written in Go!

Package documentation can be found at https://godoc.org/github.com/mvmaasakkers/ginger

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
