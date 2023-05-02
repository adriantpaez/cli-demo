package cmdrunner

type CmdDockerCompose interface {
	Up(services []string, upOpts DockerComposeUpOpts, globalOpts DockerComposeOpts) ([]byte, error)
	Stop(services []string, globalOpts DockerComposeOpts) ([]byte, error)
	Down(globalOpts DockerComposeOpts) ([]byte, error)
}

type dockerCompose struct{}

type DockerComposeOpts struct {
	// Compose configuration files.
	File []string
}

func (d *DockerComposeOpts) Build() []string {
	if len(d.File) == 0 {
		return []string{}
	}
	build := make([]string, 0, 2*len(d.File))
	for _, f := range d.File {
		build = append(build, "-f", f)
	}
	return build
}

type DockerComposeUpOpts struct {
	// Detached mode: Run containers in the background.
	Detach bool
}

func (d *DockerComposeUpOpts) Build() []string {
	if !d.Detach {
		return []string{}
	}
	return []string{"-d"}
}

func (d *dockerCompose) Up(services []string, upOpts DockerComposeUpOpts, globalOpts DockerComposeOpts) ([]byte, error) {
	name := "docker"
	args := []string{"compose"}
	args = append(args, globalOpts.Build()...)
	args = append(args, "up")
	args = append(args, upOpts.Build()...)
	args = append(args, services...)
	return runCmd(name, args...)
}

func (d *dockerCompose) Stop(services []string, globalOpts DockerComposeOpts) ([]byte, error) {
	name := "docker"
	args := []string{"compose"}
	args = append(args, globalOpts.Build()...)
	args = append(args, "stop")
	args = append(args, services...)
	return runCmd(name, args...)
}

func (d *dockerCompose) Down(globalOpts DockerComposeOpts) ([]byte, error) {
	name := "docker"
	args := []string{"compose"}
	args = append(args, globalOpts.Build()...)
	args = append(args, "down")
	return runCmd(name, args...)
}
