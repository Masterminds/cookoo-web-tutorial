package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/web"
	"fmt"
)

func main() {
	// Create a new Cookoo environment.
	registry, router, context := cookoo.Cookoo()

	// Add one route to this app.
	registry.Route("GET /", "Print Hello to something").

		Does(SayHello, "message").
		Using("who").From("query:who").

		Does(web.Flush, "out").
		Using("content").From("cxt:message").
		Using("contentType").WithDefault("text/plain")


	web.Serve(registry, router, context)
}

func SayHello(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	// Get the value of the "who" parameter, or use "World" if none is set.
	// We want it to be a string.
	who := p.Get("who", "World").(string)

	return fmt.Sprintf("Hello %s\n", who), nil
}
