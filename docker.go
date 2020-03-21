package main

import (
	"context"
	"github.com/docker/docker/api/types"
	client "github.com/docker/docker/client"

)

type Container struct {
	Id  string
}

func ContainerList() []string{
	var containers []string

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	listOptions := types.ContainerListOptions{
				Quiet: false,
	}
	containerList, err := cli.ContainerList(ctx,listOptions )

	for i := 0; i < len(containerList); i++ {
		containerObj := containerList[i]
		containers = append(containers, containerObj.ID)
	}
	return containers
}

func getContainerPid(containerId string) int {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containerInspect, err := cli.ContainerInspect(ctx,containerId)

	return containerInspect.State.Pid
}