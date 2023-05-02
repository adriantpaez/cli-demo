package cli

import (
	"cli/internal/daemon"

	"github.com/spf13/cobra"
)

func CommitCmd(d daemon.Daemon) *cobra.Command {
	var containerID string
	return &cobra.Command{
		Use:  "commit [flags] [containerID]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			containerID = args[0]
			return d.CommitContainer(containerID)
		},
	}
}
