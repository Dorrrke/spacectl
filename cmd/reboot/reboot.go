package reboot

import "github.com/spf13/cobra"

func NewRebootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reboot",
		Short: "Перезапуск служб SpaceVDI",
	}

	cmd.AddCommand(rebootAgentLauncher())
	cmd.AddCommand(rebootGlintLauncher())

	return cmd
}
