package main

import (
	"github.com/kyl2016/tusd/cmd/tusd/cli"
)

func main() {
	cli.ParseFlags()
	cli.PrepareGreeting()

	// Print version and other information and exit if the -version flag has been
	// passed else we will start the HTTP server
	if cli.Flags.ShowVersion {
		cli.ShowVersion()
	} else {
		cli.CreateComposer()
		cli.Serve()
	}
}
