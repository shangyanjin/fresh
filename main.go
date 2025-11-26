/*
Fresh is a command line tool that builds and (re)starts your web application everytime you save a go or template file.

If the web framework you are using supports the Fresh runner, it will show build errors on your browser.
*/
package main

import (
	"flag"
	"fmt"

	"fresh/runner"
)

const Version = "20251126"

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		fmt.Printf("fresh version %s, passing args: %v\n", Version, args)
	} else {
		fmt.Printf("fresh version %s\n", Version)
	}

	runner.Start(args)
}
