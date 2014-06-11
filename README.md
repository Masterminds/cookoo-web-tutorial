# 3. Hello World

Now for that glorious chapter we've all been waiting for! It's time to
build a *hello world* server!

This is the only non-Web app we are going to see. It's here just to
introduce Cookoo.

We can start out by running `app.go`, and then work backward from there.

```
$ go run app.go
info2014/06/11 00:54:14 Hello World
```

We just ran a program that logs the message "Hello World".

Now let's take a look at this program's `main` function:

```go
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
```

**Note:** We rarely walk through all of the code in a branch, so it's
often a good idea to take a quick look at the source to see the example
in context.

The first thing `main` does is get all the pieces necessary for a Cookoo
app. As a joke, the authors of Cookoo made the main app creation
function `cookoo.Cookoo()` (like a cookoo clock). That's how we get a
handle on the three major pieces of any cookoo app:

* The registry: This is where we register routes and tell Cookoo what
  each route does.
* The router: This is responsible for executing routes.
* The context: This is the data container that stores information as a
  route executes.

We'll look at these in detail as we go along. In the example above we
can see two things.

First, we build a route to execute:

```go
registry.Route("hello", "Print Hello World").
  Does(cookoo.LogMessage, "out").
		Using("msg").WithDefault("Hello World")
```

By line, here's what the above does:

1. Create a new `Route` called `hello`. Cookoo can be configured to be
   self-documenting, so every route requires a short description (`Print
   Hello World`).
2. Add a step to this route. This step executes the command
   `cookoo.LogMessage`. A command has a name, which other parts of the
   route may later. In our example, it is named `out`.
3. On the third line we pass some information into the `LogMessage`
   command. Specifically, we tell it the message to log (`msg`) is just
   `Hello World`. Later we'll see some cool ways of passing more complex
   data into a `Using` clause.

So the above says that when the app runs the `hello` route, a log
message should be written with the message `Hello World`. And that's
exactly what we saw in the opening example.

## Running a Route

The last line of our `main` function runs the route:

```go
router.HandleRequest("hello", context, true)
```

The above tells the router to handle a request named `hello` (the route
we created above) using the given base `context`. The final param
indicates whether the route is "tainted". If this is set to `true`, the
router won't let certain requests be executed. In our example, though,
it makes absolutely no difference in functionality.

So now we have our first simple Cookoo example. Let's turn it into a web
server.
