package daemon_test

import (
	"cli/internal/cmdrunner"
	"cli/internal/daemon"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaemon_DockerComposeUp(t *testing.T) {
	d := daemon.NewDaemon(nil, cmdrunner.NewDockerComposeRunner())
	middlewares := []string{"middleware-1", "middleware-2"}
	for _, m := range middlewares {
		err := d.StartMiddleware(m)
		assert.NoError(t, err)
	}
}

func TestDaemon_DockerComposeStop(t *testing.T) {
	d := daemon.NewDaemon(nil, cmdrunner.NewDockerComposeRunner())
	middlewares := []string{"middleware-1", "middleware-2"}
	for _, m := range middlewares {
		err := d.StopMiddleware(m)
		assert.NoError(t, err)
	}
}
