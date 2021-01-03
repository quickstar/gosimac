package main

import (
	"fmt"

	"github.com/quickstar/wally/cmd"
)

// nolint: gocheckglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("wally %s, commit %s, built at %s\n", version, commit, date)

	cmd.Execute()
}
