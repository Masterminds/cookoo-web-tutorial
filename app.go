package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/web"
)

func main() {
	// Create a new Cookoo environment.
	registry, router, context := cookoo.Cookoo()

	// Add one route to this app.
	registry.Route("GET /", "Print Hello Web").
		Does(web.Flush, "out").
		Using("content").WithDefault("Hello Web")

	// Synchronize access to the context.
	context = cookoo.SyncContext(context)

	// Start a server
	web.Serve(registry, router, context)
}
