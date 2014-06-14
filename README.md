# 9. Forwarding

In addition to including routes, it is also possible to forward from one
route to another. In this chapter we will look at two ways of
accomplishing it. The first is using the `cookoo.ForwardTo` command, and
the second is creating a command that returns a `*cookoo.Reroute`.

Forwarding is slightly different than including. Forwards are evaluated
at the time that the route is run. This is slightly more dangerous (less
pre-checking), it also gives us the possibility to embed route logic in
an appropriate way.

The easiest way to forward a route is to use the `cookoo.ForwardTo` command:

```go
	registry.Route("GET /hello", "Print Hello World").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("Hello World").
		Does(cookoo.ForwardTo, "fwd").
		Using("route").WithDefault("@render")
```

Functionally speaking, this produces the same result as doing an
`Includes()` (though it doesn't really inline the route). But we can see
the different between a `cookoo.ForwardTo` and an `Includes` with
another example:

```go
	// Dynamic foward
	registry.Route("GET /fwd", "Show a dynamic forward").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("Hello World").
		Using("destination").WithDefault("@render").
		Does(cookoo.ForwardTo, "fwd").
		Using("route").From("cxt:destination")
```

Note that in this case the destination is determined at runtime, not at
build time. And with a little cleverness (and the `ignoreRoutes` param
to `ForwardTo`) you can actually conditionally re-route.

**Note:** Cookoo 1.1.0 and earlier did not allow forwarding to internal
`@`-routes, though by omitting the `@` you could still forward to a
route that was not exposed to the web server.

## Writing Our Own Forwarding

What `ForwardTo` does is actually trivial in Cookoo, and we can do it
with a very simple command:

```go
func ForwardToRender(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	return nil, &cookoo.Reroute{"@render"}
}
```

The above always Forwards to the `@render` command. It uses a feature we
haven't explored yet. Cookoo commands return a `cookoo.Interrupt`. And
based on the type of interrupt, Cookoo may perform different actions.
Here are the main interrupt types:

* `error`: Processing of the current chain stops and the error is
  reported. For the Cookoo web server, a `500` error is returned.
* `cookoo.FatalError`: Produces the same result as above, but with a
  different error message. If you can, you should return `FatalError`
  when the cause of the error is known.
* `cookoo.RecoverableError`: Cookoo logs the error and then continues to
  the next command on the chain. It does not interrupt processing.
* `cookoo.Stop`: Cookoo stops processing the chain of commands, but does
  not report an error. For the web server, no result is returned.
* `cookoo.Reroute`: Cookoo stops processing the current chain and
  attempts to jump to the route given to `Reroute`. This is what we're
  doing here.

So when our example command returns a `*cookoo.Reroute`, it is
essentially telling Cookoo to jump to a different route and run it.

During a reroute, the context is passed along, too. So any variables we
put into the context will be available on the route we jump to.

Now we can put our `ForwardToReroute` command to work:

```go
	registry.Route("GET /custom", "Show custom ForwardToRender").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("Hello World").
		Does(ForwardToRender, "fwd")
```

The result of this route is the same as the other two.



