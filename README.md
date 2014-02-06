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
	
	g.Router.AddRoute("/thing", func(g *ginger.Ginger) {
							
							g.SendResponse("asdf")
					  })
	g.Run()

}

```

When you go to http://localhost:4242/thing the result will be:

```js
{
	response: "asdf"
}
```

This is still in early dev phase.
