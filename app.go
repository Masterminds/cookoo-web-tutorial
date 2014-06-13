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
		Using("content").WithDefault("Hello Web").From("query:msg").
		Using("contentType").WithDefault("text/plain")

	registry.Route("GET /hello/*", "Print Hello Web").
		Does(web.Flush, "out").
		Using("content").From("cxt:msg query:msg path:1")

	web.Serve(registry, router, context)
}
