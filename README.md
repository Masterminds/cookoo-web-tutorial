# 11. Serving Static Files

In the previous chapter we got an overview of several of the commands
that ship with Cookoo. In this chapter we're going to look at serving
static files using the `"cookoo/web".ServeFiles` command.

The `web.ServeFiles` command maps a URI path to a filesystem directory,
and then serves files from that directory.

Take a moment to look at this chapter's source code. We've added a new
directory called `assets/`. This directory contains some static files.
Let's look at the contents of `example.txt`:

```
This is an example text file.
```

Now let's see if we can get Cookoo to serve us this file.
In this chapter's version of `app.go`, we have a route for serving
static files:

```go
	// Map the URI /files to the local path ./assets.
	registry.Route("GET /files/**", "Return static files").
		Does(web.ServeFiles, "file").
		Using("directory").WithDefault("assets/")
```

In a nutshell, the above maps any `GET` request under the path
`files/` to the a path under `assets/`.

Running `go run app.go`, we can then hit the server using curl:

```
$ curl localhost:8080/files/example.txt
Oops... not this one.
```

Wait a minute... why did we get `Opps... not this one`? That's not the
text in the `example.txt` file! If we look at the `assets/` directory,
we see there's a directory in there called `files`, and that that
directory has an `example.txt` as well. That is the file that was served
above. See, `web.ServeFiles` uses the *entire* path portion of the URI.
So it interpreted our request for `/files/example.txt` as a request for
`assets/files/example.txt`.

So how do we serve `assets/example.txt`? We can tell `web.ServeFiles` to
chop off the first part of the path. Here's a route that does that:

```
	registry.Route("GET /docs/**", "Return static files").
		Does(web.ServeFiles, "file").
		Using("directory").WithDefault("assets/").
		Using("removePrefix").WithDefault("/docs")
```

The important addition above is `removePrefix`. This works on the raw
path info handed to Cookoo by the web server. For that reason, we need
the leading slash (but not a trailing slash).

Effectively, we are now matching the contents of `**` (in the path)
against the contents of the `assets/` directory.

In the next chapter we'll get an overview of datasources as we prepare
for using HTML templates.
