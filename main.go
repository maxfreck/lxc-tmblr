package main

import (
	"flag"
	"lxc-tmblr/appflags"
	"lxc-tmblr/config"
	"lxc-tmblr/lxd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		appflags.ShowUsage()
		return
	}

	flag.Usage = appflags.ShowUsage

	lxd.NewLxdProcessor(appflags.GetAppFlags(), config.GetAppConfig()).Process()
}
