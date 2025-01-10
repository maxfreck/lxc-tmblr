package lxd

import (
	"fmt"
	"lxc-tmblr/appflags"
	"lxc-tmblr/config"
	"slices"

	incus "github.com/lxc/incus/client"
	"github.com/lxc/incus/shared/api"
)

type LxdProcessor struct {
	connection incus.InstanceServer
	flags      *appflags.AppFlags
	config     *config.AppConfig
}

func NewLxdProcessor(flags *appflags.AppFlags, config *config.AppConfig) *LxdProcessor {
	processor := new(LxdProcessor)
	processor.flags = flags
	processor.config = config

	connection, err := incus.ConnectIncusUnix(config.Socket, nil)
	if err != nil {
		panic(fmt.Errorf("unable to connect to incus: %w", err))
	}

	processor.connection = connection

	return processor
}

func (l LxdProcessor) Process() {
	startInstances := []string{}

	for _, startItem := range l.flags.Start {
		instanceName := l.nameToRoot(startItem)
		if instanceName != "" {
			startInstances = append(startInstances, instanceName)
		}

		if !l.flags.IgnoreDependencies {
			startInstances = append(startInstances, l.getDependencies(startItem, nil)...)
		}
	}

	stopInstances := []string{}

	for _, stopItem := range l.flags.Stop {
		instanceName := l.nameToRoot(stopItem)
		if instanceName != "" && !slices.Contains(startInstances, instanceName) {
			stopInstances = append(stopInstances, instanceName)
		}

		if !l.flags.IgnoreDependencies {
			stopInstances = append(stopInstances, l.getDependencies(stopItem, startInstances)...)
		}
	}

	if len(stopInstances) > 0 {
		l.stop(stopInstances)
	}

	if len(startInstances) > 0 {
		l.start(startInstances)
	}
}

func (l LxdProcessor) getDependencies(name string, exclude []string) []string {
	val, ok := l.config.Containers[name]

	if !ok {
		return nil
	}

	dependencies := []string{}

	for _, dep := range val.Dependencies {
		instanceName := l.nameToRoot(dep)
		if instanceName != "" && !slices.Contains(exclude, instanceName) {
			dependencies = append(dependencies, instanceName)
		}

		dependencies = append(dependencies, l.getDependencies(dep, exclude)...)
	}

	return dependencies
}

func (l LxdProcessor) nameToRoot(name string) string {
	//Check if record exists in the config
	val, ok := l.config.Containers[name]
	if ok {
		return val.Root
	}

	return name
}

func (l LxdProcessor) stop(instances []string) {
	for _, name := range instances {
		inst, etag, err := l.connection.GetInstance(name)
		if err != nil {
			fmt.Println(fmt.Errorf("unable to get instance %s (%w)", name, err))
			continue
		}

		if !inst.IsActive() {
			fmt.Printf("instance %s is already inactive\r\n", name)
			continue
		}

		fmt.Printf("stopping %s: ", name)

		op, err := l.connection.UpdateInstanceState(name, api.InstanceStatePut{Action: "stop"}, etag)
		if err != nil {
			fmt.Println(fmt.Errorf("%w", err))
		}

		op.Wait()
		fmt.Println("OK")
	}
}

func (l LxdProcessor) start(instances []string) {
	for _, name := range instances {
		inst, etag, err := l.connection.GetInstance(name)
		if err != nil {
			fmt.Println(fmt.Errorf("unable to get instance %s (%w)", name, err))
			continue
		}

		if inst.IsActive() {
			fmt.Printf("instance %s is already active\r\n", name)
			continue
		}

		fmt.Printf("starting %s: ", name)

		op, err := l.connection.UpdateInstanceState(name, api.InstanceStatePut{Action: "start"}, etag)
		if err != nil {
			fmt.Println(fmt.Errorf("%w", err))
		}

		op.Wait()
		fmt.Println("OK")
	}
}
