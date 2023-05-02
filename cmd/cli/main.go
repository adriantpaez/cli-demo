package main

import (
	"cli/cli"
	"cli/internal/cmdrunner"
	"cli/internal/daemon"

	"github.com/docker/docker/client"
)

func main() {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer dockerClient.Close()
	daemon := daemon.NewDaemon(dockerClient, cmdrunner.NewDockerComposeRunner())
	if err := cli.RootCmd(daemon).Execute(); err != nil {
		panic(err)
	}
}
