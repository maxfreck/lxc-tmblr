package main

import (
	"flag"
	"fmt"
	"os"
)

type arrayFlags []string

type programFlags struct {
	start              arrayFlags
	stop               arrayFlags
	ignoreDependencies bool
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	flag.Usage = printUsage
	flags := parseFlags()

	fmt.Println(flags)
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

// Prints program usage info
func printUsage() {
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

// Parses command line args
func parseFlags() programFlags {
	var flags programFlags

	flag.Var(&flags.start, "start", "Start the container with all the dependencies")
	flag.Var(&flags.stop, "stop", "Stop the container with all the dependencies")
	flag.BoolVar(&flags.ignoreDependencies, "i", flags.ignoreDependencies, "Ignore dependencies and work only with specified containers")
	flag.BoolVar(&flags.ignoreDependencies, "no-dependencies", flags.ignoreDependencies, "Ignore dependencies and work only with specified containers")

	flag.Parse()

	return flags
}
