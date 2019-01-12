package main

import "fmt"

const unset = "unset"

var ( // build info
	version = unset
	date    = unset
	commit  = unset
)

func printVersion() {
	fmt.Printf("Version info: %s:%s\n", version, date)
	fmt.Printf("commit: %s\n", commit)
	fmt.Println()
}
