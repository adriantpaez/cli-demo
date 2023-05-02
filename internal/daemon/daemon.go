package daemon

import (
	"cli/internal/cmdrunner"
	"context"
	"fmt"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Daemon interface {
	StartMiddleware(middleware string) error
	StopMiddleware(middleware string) error
	CommitContainer(containerID string) error
	DownMiddleware(middleware string) error
}

type daemon struct {
	dockerClient client.APIClient
	cmdCompose   cmdrunner.CmdDockerCompose
}

func NewDaemon(dockerClient client.APIClient, cmdCompose cmdrunner.CmdDockerCompose) Daemon {
	return &daemon{
		dockerClient: dockerClient,
		cmdCompose:   cmdCompose,
	}
}

func (d *daemon) StartMiddleware(middleware string) error {
	_, err := d.cmdCompose.Up(
		[]string{},
		cmdrunner.DockerComposeUpOpts{
			Detach: true,
		},
		cmdrunner.DockerComposeOpts{
			File: []string{filepath.Join("data", middleware, "docker-compose.yml")},
		},
	)
	return err
}

func (d *daemon) StopMiddleware(middleware string) error {
	_, err := d.cmdCompose.Stop(
		[]string{},
		cmdrunner.DockerComposeOpts{
			File: []string{filepath.Join("data", middleware, "docker-compose.yml")},
		},
	)
	return err
}

func (d *daemon) DownMiddleware(middleware string) error {
	_, err := d.cmdCompose.Down(
		cmdrunner.DockerComposeOpts{
			File: []string{filepath.Join("data", middleware, "docker-compose.yml")},
		},
	)
	return err
}

func (d *daemon) CommitContainer(containerID string) error {
	_, err := d.dockerClient.ContainerCommit(
		context.Background(),
		containerID,
		types.ContainerCommitOptions{
			Reference: fmt.Sprintf("el/backup:%s", containerID),
		},
	)
	return err
}
