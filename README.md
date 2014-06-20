# 10. Built-in Commands

Over the last few chapters we have seen how to construct routes from
commands. Here we'll take a quick breather and survey the main commands
that come built-in to Cookoo.

There are three "packages" of commands in Cookoo:

- Core: Commands that are directly in the `cookoo` package. These are
  general-purpose commands.
- CLI: These commands are in `cookoo/cli`, and are usually intended for
  command-line applications. However, when you're building web servers
  they can come in useful for handling application startup.
- Web: Package `cookoo/web` has several useful and general-purpose web
  oriented commands.

Here we'll walk through the commands by name. The `app.go` version for
this chapter shows how the core and CLI commands are used. The web
commands we will see again in future chapters.

**Note:** The Cookoo source often bundles up several commands into a
file called `commands.go`. So if you're hunting for the source, look
there first. Big commands sometimes get their own files, though.

## Core Commands

The following are all core commands:

* `cookoo.AddToContext`: Add an arbitary name/value pair to the context.
* `cookoo.LogMessage`: Log a message to the logging system. This is a
  command wrapper around the `cookoo.Context.Log()` function.
* `cookoo.ForwardTo`: This is used to forward from the current command
  to another command. [The last chapter](https://github.com/Masterminds/cookoo-web-tutorial/tree/9_Forwarding)
  covers this in more detail.

Here's a simple example that uses all three commands:

```go
	// Core commands.
	registry.Route("GET /core", "Example using core commands.").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("This will get logged and rendered.").
		Does(cookoo.LogMessage, "log").
		Using("msg").From("cxt:message").
		Does(cookoo.ForwardTo, "fwd").
		Using("route").WithDefault("@render")
```

It puts the name/value pair "message"/"This will get logged and
rendered" into the context. Then it logs that message, and then it
re-routes to the `@render` route that we defined in the previous
chapter.

**Note:** Cookoo has a thin logging system that piggy-backs atop the
core Go logger. By default, log messages go to the standard output of
the console that started the process. Later we'll look at logging in
more detail.

## CLI Commands

The `cookoo/cli` package has the following commands:

* `cli.ParseArgs` parses commaindline args and extracts flags. This is
  closely related to the Go built-in `flag` package.
* `cli.ShowHelp` is a convenience function for displaying help info when
  the flag `-h` or `--help` is used.
* `cli.RunSubcommand` is a utility for building a CLI with subcommand
  support, like `git commit` or `go run`. It supports embedding flags
  between commands, as Git does: `mycli -global-flag mysub -sub-flag`.
* `cli.ShiftArgs` is a helper to work with `RunSubcommand` to support
  multiple commands. It makes it easy to shift subcommands off of the
  arguments list.

Web applications often use the CLI library to add command line arguments
to the main server command. Here's an example function, called
`prestart`, that illustrates how the CLI library can be used to add
command line flags to the server.

```go
// prestart parses CLI arguments and does any necessary context setup.
func prestart (registry *cookoo.Registry, router *cookoo.Router, context cookoo.Context) {
	// CLI commands.

	// Create flags.
	flags := flag.NewFlagSet("global", flag.PanicOnError)
	flags.Bool("h", false, "Print help text")

	// Put the args into the context.
	context.Put("os.Args", os.Args)

	// Define a pre-start route.
	registry.Route("prestart", "Do some stuff before starting the webserver.").
		// Shift the name off of the front of os.Args.
		Does(cli.ShiftArgs, "_").Using("n").WithDefault(1).

		// Parse the CLI arguments.
		Does(cli.ParseArgs, "Parse CLI arguments").
		Using("flagset").WithDefault(flags).
		Using("args").From("cxt:os.Args").

		// Show help if -h (cxt:h) was set.
		Does(cli.ShowHelp, "showHelp").
		Using("show").From("cxt:h").
		Using("summary").WithDefault("Run a demo web app server.").
		Using("usage").WithDefault("go run app.go").
		Using("flags").WithDefault(flags).

		// Log a startup message.
		Does(cookoo.LogMessage, "starting").
		Using("msg").WithDefault("Starting...")

	// Run the prestart route.
	router.HandleRequest("prestart", context, false)

	// If help mode (-h) was on, we should stop now.
	if context.Get("showHelp", false).(bool) {
		os.Exit(0)
	}
}
```

There is quite a bit going on above. Here's the summary:

1. Create a `flag.FlagSet` and add the `-h` flag to turn on help text.
2. Add `os.Args` to the context so that Cookoo can access them easily.
3. Create the `prestart` route, which does the following:
  * Shift off the first value from `os.Args` (which is always the
    command name).
  * Parse `os.Args`, given the `flag.FlagSet` we already defined.
  * If the `-h` flag was true, display the help text and stop.
  * (If there is no `-h` flag) log a message: `Starting...`
4. Finally, run the `prestart` route that we just defined.

Note that we can manually run routes (`router.HandleRequest`) and then
later hand the router over to `web.Serve()` to serve HTTP requests.
Being able to manually run routes like this is good for start and stop
handlers, launching auxilliary goroutines, and also running automated
tests.

**Note:** One common use for using a `prestart` setup like this is to
pass configuration data (like a JSON or YAML file) into Cookoo.

## Web Commands

Finally, we'll take a look at some of the web-specific commands in the
package `cookoo/web`.

* `web.Flush` sends the content back to the web client (or to another
   `io.Writer` if specified).
* `web.RenderHTML` takes a template and renders it through the HTML
  renderer. We will see this in a future chapter.
* `web.ServerInfo` is for debugging. It dumps both the request and the
  response objects to the client.
* `web.ServeFiles` takes a directory and serves the files in that
  directory. This is an easy way to add support for serving images or
  other pieces of static content.

Additionally, the package `cookoo/web/auth` has commands that implement
HTTP Basic Authentication so you can provide users with basic login and
password protection.

In the next chapter we will look at serving static files. This will pave
the way for a future chapter on using `web.RenderHTML` and Go templates.
