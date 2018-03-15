package main

import (
	"os"
	"runtime"

	"admission-controller/cmd"
)

func main() {

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	if err := cmd.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}