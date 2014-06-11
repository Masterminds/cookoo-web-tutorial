package main

import (
	"github.com/Masterminds/cookoo"
)

func main() {
	// Create a new Cookoo environment.
	registry, router, context := cookoo.Cookoo()

	// Add one route to this app.
	registry.Route("hello", "Print Hello World").
		Does(cookoo.LogMessage, "out").
		Using("msg").WithDefault("Hello World")

	// Run that route.
	router.HandleRequest("hello", context, true)
}
