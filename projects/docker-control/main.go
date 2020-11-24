package main

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func startContainer(context *client.Client, image string, name string, port string) error {
	natPort, err := nat.NewPort("tcp", port)
	if err != nil {
		return err
	}
	
}

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("docker-control: %v", err)
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
	}, nil, nil, nil, "")
	if err != nil {
		log.Fatalf("docker-control: %v", err)
	}
	if err := cti.ContainerStart(ctx, resp.ID, )
}
