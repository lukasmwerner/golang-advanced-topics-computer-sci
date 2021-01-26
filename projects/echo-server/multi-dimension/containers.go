package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	network "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// CreateNewContainer creates a new container on the given network with the given image. returns a container hostname and an error
func CreateNewContainer(image string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", fmt.Errorf("CreateNewContainer: unable to create client: %v", err)
	}
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	networkFilter := filters.NewArgs()
	networkFilter.Add("name", "micro-net")
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{
		Filters: networkFilter,
	})
	if err != nil {
		return "", fmt.Errorf("CreateNewContainer: unable to find network: %v", err)
	}
	networkNameConfig := &network.EndpointSettings{
		NetworkID: networks[0].ID,
	}
	networkConfig.EndpointsConfig["micro-net"] = networkNameConfig
	config := &container.Config{
		Image: image,
	}

	cntr, err := cli.ContainerCreate(context.Background(), config, nil, networkConfig, nil, "")
	if err != nil {
		return "", fmt.Errorf("CreateNewContainer: unable to create container: %v", err)
	}
	cli.ContainerStart(context.Background(), cntr.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", fmt.Errorf("CreateNewContainer: unable to start container: %v", err)
	}
	cont, err := cli.ContainerInspect(context.Background(), cntr.ID)
	return cont.Config.Hostname, nil

}

// GetContainers get the available containers
func GetContainers(quiet bool, size bool, all bool) ([]types.Container, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("GetContainers: unable to create client: %v", err)
	}
	return cli.ContainerList(context.Background(), types.ContainerListOptions{
		Quiet: quiet,
		Size:  size,
		All:   all,
	})
}
