package main

import (
	"flag"
	"fmt"
	"lxc-tmblr/appflags"
	"lxc-tmblr/config"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		appflags.ShowUsage()
		return
	}

	flag.Usage = appflags.ShowUsage
	flags := appflags.GetAppFlags()
	config := config.GetAppConfig()

	fmt.Println(flags)
	fmt.Println(config.Containers["php7-maria"])
}
