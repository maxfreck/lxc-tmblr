package lxd

import (
	"fmt"
	"lxc-tmblr/appflags"
	"lxc-tmblr/config"

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
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	processor.connection = connection

	return processor
}

func (l LxdProcessor) Process() {
	instances := l.getInstances()

	fmt.Println(instances)
}

func (l LxdProcessor) getInstances() map[string]api.Instance {
	instancesMap := make(map[string]api.Instance)

	instances, err := l.connection.GetInstancesAllProjects("")
	if err != nil {
		panic(fmt.Errorf("unable to get instances: %w", err))
	}

	for _, v := range instances {
		instancesMap[v.Name] = v
	}

	return instancesMap
}
