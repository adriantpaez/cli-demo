package cli

import (
	"cli/internal/daemon"

	"github.com/spf13/cobra"
)

func RootCmd(d daemon.Daemon) *cobra.Command {
	cmd := &cobra.Command{
		Use: "tool",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(
		StartCmd(d),
		StopCmd(d),
		DownCmd(d),
		CommitCmd(d),
	)
	return cmd
}
