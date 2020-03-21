package main

import (
	"context"
	"fmt"
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

	for _, container := range containerList {
		list := append(containers, container.ID)
		fmt.Println(list)
	}
	return containers
}

func getContainerPid(container Container) int {

	objectId := container.Id

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containerInspect, err := cli.ContainerInspect(ctx,objectId)

	return containerInspect.State.Pid
}