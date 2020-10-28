package common

import (
	"context"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/spf13/afero"
)

type MockClient struct {
	MockClose func() error

	MockContainerCreate func(
		ctx context.Context,
		config *container.Config,
		hostConfig *container.HostConfig,
		networkingConfig *network.NetworkingConfig,
		containerName string,
	) (container.ContainerCreateCreatedBody, error)

	MockContainerStart func(
		ctx context.Context,
		containerID string,
		options types.ContainerStartOptions,
	) error

	MockContainerStop func(
		ctx context.Context,
		containerID string,
		timeout *time.Duration,
	) error

	MockContainerInspect func(
		ctx context.Context,
		containerID string,
	) (types.ContainerJSON, error)

	MockContainerRemove func(
		ctx context.Context,
		containerID string,
		options types.ContainerRemoveOptions,
	) error

	MockNetworkCreate func(
		ctx context.Context,
		name string,
		options types.NetworkCreate,
	) (types.NetworkCreateResponse, error)

	MockNetworkRemove func(
		ctx context.Context,
		name string,
	) error

	MockNetworkList func(
		ctx context.Context,
		options types.NetworkListOptions,
	) ([]types.NetworkResource, error)

	MockNetworkInspect func(
		ctx context.Context,
		networkID string,
	) (types.NetworkResource, error)

	MockIsErrNotFound func(err error) bool
}

func (c *MockClient) Close() error {
	if c.MockClose != nil {
		return c.MockClose()
	}
	return nil
}

func (c *MockClient) ContainerCreate(
	ctx context.Context,
	config *container.Config,
	hostConfig *container.HostConfig,
	networkingConfig *network.NetworkingConfig,
	containerName string,
) (container.ContainerCreateCreatedBody, error) {
	if c.MockContainerCreate != nil {
		return c.MockContainerCreate(
			ctx,
			config,
			hostConfig,
			networkingConfig,
			containerName,
		)
	}
	return container.ContainerCreateCreatedBody{}, nil
}

func (c *MockClient) ContainerStart(
	ctx context.Context, containerID string, options types.ContainerStartOptions,
) error {
	if c.MockContainerStart != nil {
		return c.MockContainerStart(
			ctx, containerID, options,
		)
	}
	return nil
}

func (c *MockClient) ContainerStop(
	ctx context.Context, containerID string, timeout *time.Duration,
) error {
	if c.MockContainerStop != nil {
		return c.MockContainerStop(ctx, containerID, timeout)
	}
	return nil
}

func (c *MockClient) ContainerInspect(
	ctx context.Context, containerID string,
) (types.ContainerJSON, error) {
	if c.MockContainerInspect != nil {
		return c.MockContainerInspect(ctx, containerID)
	}
	return types.ContainerJSON{}, nil
}

func (c *MockClient) ContainerRemove(
	ctx context.Context, containerID string, options types.ContainerRemoveOptions,
) error {
	if c.MockContainerRemove != nil {
		return c.MockContainerRemove(ctx, containerID, options)
	}
	return nil
}

func (c *MockClient) NetworkCreate(
	ctx context.Context, name string, options types.NetworkCreate,
) (types.NetworkCreateResponse, error) {
	if c.MockNetworkCreate != nil {
		return c.MockNetworkCreate(ctx, name, options)
	}
	return types.NetworkCreateResponse{}, nil
}

func (c *MockClient) NetworkRemove(ctx context.Context, name string) error {
	if c.MockNetworkRemove != nil {
		return c.MockNetworkRemove(ctx, name)
	}
	return nil
}

func (c *MockClient) NetworkList(
	ctx context.Context, options types.NetworkListOptions,
) ([]types.NetworkResource, error) {
	if c.MockNetworkList != nil {
		return c.MockNetworkList(ctx, options)
	}
	return []types.NetworkResource{}, nil
}

func (c *MockClient) NetworkInspect(
	ctx context.Context, networkID string,
) (types.NetworkResource, error) {
	if c.MockNetworkInspect != nil {
		return c.MockNetworkInspect(ctx, networkID)
	}
	return types.NetworkResource{}, nil
}

func (c *MockClient) IsErrNotFound(err error) bool {
	if c.MockIsErrNotFound != nil {
		return c.MockIsErrNotFound(err)
	}
	return false
}

func CheckFiles(
	fs afero.Fs, st *testing.T, shouldExist bool, paths ...string,
) (bool, error) {
	for _, p := range paths {
		exists, err := afero.Exists(fs, p)
		if err != nil {
			return false, err
		}
		if exists != shouldExist {
			return false, nil
		}
	}
	return true, nil
}