# 7. A First Command

In the past few chapters we have worked only with one or two built-in
Cookoo commands -- mostly just `web.Flush`. In this section, we'll
create our own command.

Commands are just functions that meet the definition of a
`cookoo.Command`.

```go
type Command func(cxt Context, params *Params) (interface{}, Interrupt)
```

Commands get two arguments and return two results.

**Arguments**

* `cxt`: The cookoo.Context that gets passed along a request. You can
  use this to get or store data.
* `params`: The `cookoo.Params` object, which has the data passed in
  from the `Using().WithDefault().From()` chain.

** Return Values**

* `interface{}`: Whatever return value you want this function to return.
* `cookoo.Interrupt`: Any exceptional condition, including an `error`.
  We'll take a look at this in more depth later.

Let's look at the example in this branch.

## The SayHello Command

Here's our simple starter command.

It takes one parameter (`who`) and returns the string `Hello %s`, where
`%s` is replaced with the value of `who`.

```go
func SayHello(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	// Get the value of the "who" parameter, or use "World" if none is set.
	// We want it to be a string.
	who := p.Get("who", "World").(string)

	return fmt.Sprintf("Hello %s\n", who), nil
}
```

Let's look at each of the two lines in the function's body:

```go
who := p.Get("who", "World").(string)
```

This sets the value of `who` to whatever value gets passed in from
`Using().From().WithDefault()`, and it makes sure that this data is
typed to a `string`.

`p.Get()` takes two parameters. The first is the name of the param that
we want to get. The second is a default value. (If this method doesn't
tickle your fance, check out `p.Has()`, which does not return a default
value.)

The second line formats a string and returns it:

```go
return fmt.Sprintf("Hello %s\n", who), nil
```

Since we have no errors, we always return `string, nil`.

## Wiring Up Our New Command

Now we can use our command from within a route. Here's our old "hello
web" example re-tooled to use our command:

```go
	registry.Route("GET /", "Print Hello to something").

		Does(SayHello, "message").
		Using("who").From("query:who").

		Does(web.Flush, "out").
		Using("content").From("cxt:message").
		Using("contentType").WithDefault("text/plain")

```

I've added some empty lines to make it easier to visualize. Right now,
the route `GET /` executes two commands in sequence:

1. `SayHello`
2. `web.Flush`

Let's look at the spec for the first:

```go
		Does(SayHello, "message").
		Using("who").From("query:who").
```

By now, this should be pretty straightforward. The `SayHello` command
will be executed. The `who` parameter will
get set to the value of the query param `?who=XXX`.

But now something that was not important before is **very** important.
*Now the name of the command matters.* By default, Cookoo will store the
return value of a command inside of the `cookoo.Context`. And it's name
will be the name of the command.

So when `SayHello` executes, the return value will be put in
`cxt:message` (since `SayHello`'s name is `message`).

Now the second command will execute:

```go
		Does(web.Flush, "out").
		Using("content").From("cxt:message").
		Using("contentType").WithDefault("text/plain")
```

Notice the second line? We get `content` from `cxt:message`. So
basically we are feeding data from the previous `SayHello` command into
the `web.Flush` command.

Let's run it and see what happens.

```
$ curl localhost:8080/?who=You
Hello You
$ curl localhost:8080/
Hello World
```

Now we're starting to get into what makes Cookoo powerful: We can
compose routes by chaining together commands. And just like we've used
`web.Flush` over and over again, if we write our commands well we will
be able to reuse them.

The second nicety of the chain of commands style is that we can see at a
glance what each route does without having to dive any deeper than the
route declarations.

In the next section, we'll look at how we can re-use route definitions
themselves to simplify our chains.

[git checkout 8_Includes](https://github.com/Masterminds/cookoo-web-tutorial/tree/8_Includes)
