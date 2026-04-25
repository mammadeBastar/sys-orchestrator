package main

import (
	"os"

	"sys-orchestrator/internal/sysapp"
)

func main() {
	os.Exit(sysapp.New(sysapp.Options{}).Run(os.Args[1:]))
}
