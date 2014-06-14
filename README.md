# 6. Request Data

In this chapter we'll look at how to get data out of an HTTP request and
pass it on to commands.

To get there, we need to take another look at the way Cookoo passes data
into commands.

## "Does", "Using", and "WithDefault"

Our working example thus far has been something like this:

```go
	registry.Route("GET /", "Print Hello Web").
		Does(web.Flush, "out").
		Using("content").WithDefault("Hello Web")
```

We've seen already that...

* `Does()` specifies which command to run
* `Using()` maps a parameter name, and...
* `WithDefault()` gives that parameter a default value.

So in the case of the above, we're calling Cookoo's `web.Flush()`
command, which takes any of the following parameters:

* `content`: The content to write to the web browser
* `contentType`: The MIME type (`text/html`, `application/json`, etc.).
* `responseCode`: The HTTP status code, as captured in Go's `net/http`
  library.
* `headers`: a `map[string]string` of HTTP headers
* `writer`: A Go `io.Writer` if we want to send data somewhere besides
  the web browser.

In all of our examples above, we've had only one `Using()` call per
command. However, if we wanted to use more than one, we could do so:

```go
	registry.Route("GET /", "Print Hello Web").
		Does(web.Flush, "out").
		Using("content").WithDefault("Hello Web").
		Using("contentType").WithDefault("text/plain")
```

Now the `web.Flush` command will use our values for both of those
parameters.

But in the example above, we're usign `WithDefault()` to send our
specified default values. So `content` will *always* be `Hello Web`.
What if instead we wanted to pass in some data generated elsewhere?

## Using "From"

Here's `Using()` can take values not only from `WithDefault()`, but also
from `From()`.

Technically speaking, `From()` formalizes a way to extract values from a
`cookoo.KeyValueDatasource`. In practice, what that means is that we can
retrieve values that were computed by other parts of Cookoo.

Alright... enough talk. Let's see an example:

```go
	registry.Route("GET /", "Print Hello Web").
		Does(web.Flush, "out").
		Using("content").WithDefault("Hello Web").From("query:msg").
		Using("contentType").WithDefault("text/plain")
```

The main change above is the addition of `.From("query:msg")`. In the
Cookoo web library, `query` is the datasource that stores query
parameters (a.k.a. GET params) in the URL. So the above tells Cookoo to
get the `msg` value out of the query string.

We can see this in action by running the server and then using Curl to
send this request:

```
$ curl localhost:8080/?msg=Hi
```

(You might need to add a backslash before `?` for shell escaping.)

The output of the above is: `Hi`.

**Note:** The code above is *definitely not* secure. We're not filtering
the content of `msg`.

Behind the scenes, what's going on above is that Cookoo is handling the
web request, putting all of the query parameters into the `query`
datasource, and then running `web.Flush()`. It then takes the value of
`From("query:msg")` and sets it as the value to `content`.

Or, in short, it creates a name value pair from `content` to
whatever is passed in `msg`.

### And the default value...

So what if we don't provide a `msg` in the query string?

```
$ curl localhost:8080/
Hello Web
```

Without the query param, Cookoo calls back to the value specified in
`WithDefault()`.

## The Places You Can Get Things

So we've seen how to get query parameters through `From()`. What else
can we get that way?

First and foremost, there's `cxt`: `From("cxt:foo")` will look in the
context for a value. This is the most common way of using Cookoo's
`From()` clause to pass data from one command to another, and we'll see
it often from here on.

Here are some of the others:

* `url`: Grab various components out of the URL. For example, `url:Host`
  gets the domain name. If you ever need raw query parameters, you can
also use `url:RawQuery`.
* `post`: Grab name/value pairs from POST data. This works just like
  query, only it looks in decoded POST data.
* `path`: Grab (by index) a part of a URL. For example, with route
  `/foo/bar`, using `path:0` will fetch `foo` and `path:1` will get
`bar`.

## Using More Than One Source

You can also pass multiple sources in a `From()` statement:

```go
registry.Route("GET /hello/*", "Print Hello Web").
		Does(web.Flush, "out").
		Using("content").From("cxt:msg query:msg path:1")
```

Here our `From()` has three sources. It will use the first non-Nil
source that it finds.

`cxt:msg` will always be empty (because we don't ever set it).
`query:msg` will be set if we put it in the context, and `path:1` will
be set from the second item in the path.

```
$ curl localhost:8080/hello/world?msg=Hi
Hi
$ curl localhost:8080/hello/world
world
```

That's how Cookoo web applications can get data from their environment.

**Note:** For a Cookoo web app, the `Context` object has some built-in
values that you can get from `From()` calls:

* `From("cxt:http.Request")`: The `net/http.Request` object.
* `From("cxt:http.ResponseWriter")`: The `net/http.ResponseWriter` object.
* `From("cxt:server.Address")`: The address of the currently running
  server.

In the next chapter we'll create a custom command and see how commands
can be chained together on a route.

[git checkout 7_First_Command](https://github.com/Masterminds/cookoo-web-tutorial/tree/7_First_Command)
