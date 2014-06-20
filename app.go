package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/cli"
	"github.com/Masterminds/cookoo/web"
	"os"
	"flag"
	"fmt"
)

func main() {
	// Create a new Cookoo environment.
	registry, router, context := cookoo.Cookoo()

	// An internal route that cannot be accessed directly.
	registry.Route("@render", "Send a message to browser.").
		Does(web.Flush, "out").
		Using("content").From("cxt:message").
		Using("contentType").WithDefault("text/plain")

	// An example root.
	registry.Route("GET /", "Print Hello to something").
		Does(SayHello, "message").
		Using("who").From("query:who").
		Includes("@render")

	// Core commands.
	registry.Route("GET /core", "Example using core commands.").
		Does(cookoo.AddToContext, "_").
		Using("message").WithDefault("This will get logged and rendered.").
		Does(cookoo.LogMessage, "log").
		Using("msg").From("cxt:message").
		Does(cookoo.ForwardTo, "fwd").
		Using("route").WithDefault("@render")

	prestart(registry, router, context)
	web.Serve(registry, router, context)
}

func SayHello(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	// Get the value of the "who" parameter, or use "World" if none is set.
	// We want it to be a string.
	who := p.Get("who", "World").(string)

	return fmt.Sprintf("Hello %s\n", who), nil
}

// prestart parses CLI arguments and does any necessary context setup.
func prestart (registry *cookoo.Registry, router *cookoo.Router, context cookoo.Context) {
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

