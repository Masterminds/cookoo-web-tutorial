package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/web"
	"fmt"
)

func main() {
	// Create a new Cookoo environment.
	registry, router, context := cookoo.Cookoo()

	// An internal route that cannot be accessed directly.
	registry.Route("render", "Send a message to browser.").
		Does(web.Flush, "out").
		Using("content").From("cxt:message").
		Using("contentType").WithDefault("text/plain")
	registry.Route("@render", "Send a message to browser.").
		Does(web.Flush, "out").
		Using("content").From("cxt:message").
		Using("contentType").WithDefault("text/plain")

	// An example root.
	registry.Route("GET /", "Print Hello to something").
		Does(SayHello, "message").
		Using("who").From("query:who").
		Includes("@render")

	// Example of ForwardTo.
	registry.Route("GET /hello", "Print Hello World").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("Hello World").
		Does(cookoo.ForwardTo, "fwd").
		Using("route").WithDefault("@render")

	// Dynamic foward
	registry.Route("GET /fwd", "Show a dynamic forward").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("Hello World").
		Using("destination").WithDefault("@render").
		Does(cookoo.ForwardTo, "fwd").
		Using("route").From("cxt:destination")

	// Use ForwardToRender to forward.
	registry.Route("GET /custom", "Show custom ForwardToRender").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("Hello World").
		Does(ForwardToRender, "fwd")


	web.Serve(registry, router, context)
}

func SayHello(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	// Get the value of the "who" parameter, or use "World" if none is set.
	// We want it to be a string.
	who := p.Get("who", "World").(string)

	return fmt.Sprintf("Hello %s\n", who), nil
}

func ForwardToRender(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	return nil, &cookoo.Reroute{"@render"}
}
