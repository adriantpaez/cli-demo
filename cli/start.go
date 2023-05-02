package cli

import (
	"cli/internal/daemon"
	"github.com/spf13/cobra"
)

func StartCmd(d daemon.Daemon) *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "start [flags] [name]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name = args[0]
			return d.StartMiddleware(name)
		},
	}
	return cmd
}
