package main

import (
	"os"
    "github.com/turamant/goyava/http/route"
    "goyave.dev/goyave/v4"
)

func main() {
    if err := goyave.Start(route.Register); err != nil {
      os.Exit(err.(*goyave.Error).ExitCode)
    }
}