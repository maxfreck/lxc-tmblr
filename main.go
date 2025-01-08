package main

import (
	"flag"
	"fmt"
	"lxc-tmblr/appflags"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		appflags.ShowUsage()
		return
	}

	flag.Usage = appflags.ShowUsage
	flags := appflags.GetAppFlags()

	fmt.Println(flags)
}
