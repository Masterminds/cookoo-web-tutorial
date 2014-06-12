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

	registry.Route("GET /hello", "Print Hello").
		Does(web.Flush, "out").
		Using("content").WithDefault("Hello")

	registry.Route("POST /hello", "Print Hello POST").
		Does(web.Flush, "out").
		Using("content").WithDefault("Hello POST")

	registry.Route("* /goodbye", "Print Goodbye Web").
		Does(web.Flush, "out").
		Using("content").WithDefault("Goodbye Web")

	registry.Route("GET /goodbye/*", "Print Goodbye").
		Does(web.Flush, "out").
		Using("content").WithDefault("Goodbye")

	registry.Route("GET /goodbye/**", "Print Goodbye Everybody").
		Does(web.Flush, "out").
		Using("content").WithDefault("Goodbye Everybody")



	// Synchronize access to the context.
	context = cookoo.SyncContext(context)

	// Start a server
	web.Serve(registry, router, context)
}
