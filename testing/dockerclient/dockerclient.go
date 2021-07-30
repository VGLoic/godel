package dockerclient

import (
	"context"
	"errors"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type DockerCli struct {
	*client.Client
}

func NewDockerCli() (*DockerCli, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &DockerCli{cli}, nil
}

func (cli *DockerCli) listContainers(ctx context.Context) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return []types.Container{}, err
	}
	return containers, nil
}

func (cli *DockerCli) pullImage(ctx context.Context, image string) error {
	out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := ioutil.ReadAll(out); err != nil {
		return err
	}
	return nil
}

func (cli *DockerCli) CreateContainer(ctx context.Context, containerConfig *container.Config, hostConfig *container.HostConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	image := containerConfig.Image
	if image == "" {
		return container.ContainerCreateCreatedBody{}, errors.New("Image has not been defined when creating container")
	}
	if err := cli.pullImage(ctx, image); err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	cont, err := cli.ContainerCreate(
		context.Background(),
		containerConfig,
		hostConfig,
		&network.NetworkingConfig{},
		nil,
		containerName,
	)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	err = cli.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	return cont, nil
}

func (cli *DockerCli) StopContainers(ctx context.Context) error {
	containers, err := cli.listContainers(ctx)
	if err != nil {
		return err
	}
	for _, container := range containers {
		err = cli.ContainerStop(ctx, container.ID, nil)
		if err != nil {
			return err
		}
		err = cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
