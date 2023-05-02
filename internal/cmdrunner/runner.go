package cmdrunner

func NewDockerComposeRunner() CmdDockerCompose {
	return &dockerCompose{}
}
