package cmd

import (
	"github.com/Dorrrke/spacectl/cmd/reboot"
	"github.com/Dorrrke/spacectl/ui/helptext"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "spacectl",
	Short: "CLI для SpaceVDI",
}

func init() {
	RootCmd.SetHelpTemplate(helptext.CustomHelpTemplate)
	RootCmd.SetUsageTemplate(helptext.CustomUsageTemplate)
	RootCmd.CompletionOptions.DisableDefaultCmd = true

	startCmd := reboot.NewRebootCommand()
	RootCmd.AddCommand(startCmd)
}
