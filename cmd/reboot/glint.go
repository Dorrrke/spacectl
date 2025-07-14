package reboot

import (
	"fmt"

	"github.com/Dorrrke/spacectl/internal/reboot"
	"github.com/spf13/cobra"
)

func rebootGlintLauncher() *cobra.Command {
	return &cobra.Command{
		Use:   "glint",
		Short: "Перезапустить Glint-launcher",
		Run: func(cmd *cobra.Command, args []string) {
			if err := reboot.RebootGlintLauncher(); err != nil {
				fmt.Println("Ошибка перезапуска Glint-launcher:", err)
			} else {
				fmt.Println("Glint-launcher перезапущен")
			}
		},
	}
}
