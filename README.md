# 4. Hello Web

Now that we have seen a basic Cookoo app we can start building a web
application. Cookoo comes with the tools needed to quickly set up a web
server.

At a glance, the `main` function in our new webserver will not look much
different from the `Hello World` version we saw in the last chapter.

```go
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
```

The shape of the app is the same, but we handle a few things
differently. Let's start by looking at our new route declaration.

## A First Web Route

We've made a few changes to the route.

First, there is the route name. In the previous example, we named the
route `hello`. And when we executed it, we used
`router.HandleRequest('hello', context, true)`.

When writing web apps, we use the route name to indicate the HTTP
request. In this case, we've named it `GET /`, which means "any get
request to the webserver's root directory."

Second, where we formerly wrote `Hello World` as a log message, now we're
using the `coookoo/web.Flush` command. This is a general-purpose command
for sending data to an HTTP client.

## Synchronizing the Context

One new addition we made is, strictly speaking, not necessary. But it is
a good idea. We are synchronizing access to the `context`.

We do this so that if a command executes a goroutine, the context will
be safe from race conditions.

Honestly, our simple `Hello Web` server will not ever have race
conditions, so we could skip synchronization. However, since
synchronizing the context is considered the best practice, I've left
this step in.

## Starting the Server

The last step in creating our simple Cookoo web app is creating the HTTP
server:

```go
web.Serve(registry, router, context)
```

This piece is straightforward: We pass the registry, router, and context
into the a server and let it do its thing.

The `context` that is passed into `web.Serve` is actually not the
context that individual routes will access. Instead, it is treated as
the base context. Each time a new request comes in, this context will be
copied and that copy will serve as the request-specific context.

## Running the Server

We can start the server with `go run app.go`. And then we can run a
simple `curl` request to see what our new server does:

```
$ curl -v localhost:8080/
> GET / HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:8080
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=utf-8
< Date: Thu, 12 Jun 2014 04:45:03 GMT
< Content-Length: 9
<
Hello Web
```

Our app behaves as expected: It sends a simple message back to `curl`:
`Hello Web`. In a later chapter we will see how we can use templates to
render HTML content and then send that back.

Now we have a simple web server. You can stop the server by hitting
`CTRL-C`.

In the next chapter we will look at building different kinds of paths,
including using other verbs like `POST` and using wildcards.

[git checkout 5_Routes](https://github.com/Masterminds/cookoo-web-tutorial/tree/5_Routes)
