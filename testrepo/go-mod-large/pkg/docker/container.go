package docker

import (
	"github.com/docker/cli/cli/command/container"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Containers(options types.ContainerListOptions) ([]types.Container, error) {
	ctx := context.Background()
	return apiClient.ContainerList(ctx, options)
}

func ContainerExist(ref string) (bool, error) {
	if _, err := ContainerInspect(ref); err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func ContainerInspect(ref string) (types.ContainerJSON, error) {
	ctx := context.Background()
	return apiClient.ContainerInspect(ctx, ref)
}

func ContainerCommit(ref string, commitOptions types.ContainerCommitOptions) (string, error) {
	ctx := context.Background()
	response, err := apiClient.ContainerCommit(ctx, ref, commitOptions)
	if err != nil {
		return "", err
	}

	return response.ID, nil
}

func ContainerRemove(ref string, options types.ContainerRemoveOptions) error {
	ctx := context.Background()
	err := apiClient.ContainerRemove(ctx, ref, options)
	if err != nil {
		return err
	}

	return nil
}

func CliCreate(args ...string) error {
	cmd := container.NewCreateCommand(cli)
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.SetArgs(args)

	err := cmd.Execute()
	if err != nil {
		return err
	}

	return nil
}

func CliRun(args ...string) error {
	cmd := container.NewRunCommand(cli)
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.SetArgs(args)

	err := cmd.Execute()
	if err != nil {
		return err
	}

	return nil
}

func CliRm(args ...string) error {
	cmd := container.NewRmCommand(cli)
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.SetArgs(args)

	err := cmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
