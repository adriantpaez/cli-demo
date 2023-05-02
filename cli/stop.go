package cli

import (
	"cli/internal/daemon"
	"github.com/spf13/cobra"
)

func StopCmd(d daemon.Daemon) *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "stop [flags] [name]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name = args[0]
			return d.StopMiddleware(name)
		},
	}
	return cmd
}
