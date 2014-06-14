# 8. Includes

In the last chapter we saw how to create a route with more than one
command. In this chapter we will look at another aspect of building
chains of commands: including routes.

## A Basic Include Route

In the `app.go` for this chapter, the first route looks like this:

```go
// An internal route that cannot be accessed directly.
registry.Route("@render", "Send a message to browser.").
  Does(web.Flush, "out").
  Using("content").From("cxt:message").
  Using("contentType").WithDefault("text/plain")
```

Notice that the name of this route is not a URI-like path. It begins
with an `@`, and has no path. This is an internal route. Internal routes
cannot be accessed directly. They can only be the target of includes and
redirects.

Typically, an internal route performs a few commands, but doesn't form a
sequence that would normally be a complete chain of commands. In our
simple case, the internal route just encapsulates a call to `web.Flush`.

Note that it assumes that the context has a `message` value. Internal
routes are often dependent on regular routes to supply information (in
the form of context variables).

Other than these things, internal routes are the same as any other
route.

Now let's see how they are used.

## Using "Includes"

One common way to use an internal route is to include it into another
route using the `Includes()` method. In this chapter's code there are
two routes that do this:

```go
	registry.Route("GET /", "Print Hello to something").
		Does(SayHello, "message").
		Using("who").From("query:who").
		Includes("@render")

	// Another example route.
	registry.Route("GET /hello", "Print Hello World").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("Hello World").
		Includes("@render")
```

The first route is just a slightly modified version of the route we
build in the previous chapter. But where we used to call the `web.Flush`
command, we now call `Includes("@render")`. Effectively, this inlines
the `@render` route into this chain of commands.

The second route ends the same way, by including `@render` into its
chain of commands. This illustrates how internal routes can be re-used.

**Note:** This second route also makes use of the built-in
`cookoo.AddToContext` command, which inserts data directly into the
context as name/value pairs. Every `Using().From().WithDefault()` clause
attached to this command will result in a new pair being added to the
context. It's a great way to populate a context with standard data.

Go ahead and run the exampels above. At this point the outcome should be
predictable.

Later we will see how our `@render` route can be used to encapsulate
more complex logic -- in our case, rendering data through HTML
templates.

But first we will take a look at another concept similar to `Includes`,
and that is *route forwarding*.
