Ginger Rest Framework for GO
=========

The Ginger Rest Framework written in Go!

The idea is the same, only useable in Google's Go this time.

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

This is still in early dev phase.