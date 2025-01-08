package appflags

import (
	"flag"
	"fmt"
	"os"
)

type arrayFlags []string

type AppFlags struct {
	Start              arrayFlags
	Stop               arrayFlags
	IgnoreDependencies bool
}

// String is an implementation of the flag.Value interface
func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

// Set is an implementation of the flag.Value interface
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// Parses command line args
func GetAppFlags() *AppFlags {
	var flags AppFlags

	flag.Var(&flags.Start, "start", "Start the container with all the dependencies")
	flag.Var(&flags.Stop, "stop", "Stop the container with all the dependencies")
	flag.BoolVar(&flags.IgnoreDependencies, "i", flags.IgnoreDependencies, "Ignore dependencies and work only with specified containers")
	flag.BoolVar(&flags.IgnoreDependencies, "no-dependencies", flags.IgnoreDependencies, "Ignore dependencies and work only with specified containers")

	flag.Parse()

	return &flags
}

// Prints program usage info
func ShowUsage() {
	fmt.Println("Description:")
	fmt.Println("  A simple LXC orchestrator that allows you to start or stop containers based on their dependencies")
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Printf("  %s [command]", os.Args[0])
	fmt.Println()
	fmt.Println()

	fmt.Println("Available Commands:")
	fmt.Println("  -start    Start the container with all the dependencies")
	fmt.Println("  -stop     Stop the container with all the dependencies")
	fmt.Println()

	fmt.Println("Flags:")
	fmt.Println("  -n, -no-dependencies    Ignore dependencies and work only with specified containers")
	fmt.Println()
}
