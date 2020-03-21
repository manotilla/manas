package main

import (
	"context"
	"fmt"
	client "github.com/docker/docker/client"
)

type Container struct {
	Id  string
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